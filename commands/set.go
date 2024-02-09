package commands

import (
	"fmt"
	"os"
	"procfile-util/procfile"

	"github.com/josegonzalez/cli-skeleton/command"
	"github.com/posener/complete"
	flag "github.com/spf13/pflag"
)

type SetCommand struct {
	command.Meta
	GlobalFlagCommand

	processType string
	command     string
	stdout      bool
	writePath   string
}

func (c *SetCommand) Name() string {
	return "set"
}

func (c *SetCommand) Synopsis() string {
	return "Eats one or more lollipops"
}

func (c *SetCommand) Help() string {
	return command.CommandHelp(c)
}

func (c *SetCommand) Examples() map[string]string {
	appName := os.Getenv("CLI_APP_NAME")
	return map[string]string{
		"Command": fmt.Sprintf("%s %s", appName, c.Name()),
	}
}

func (c *SetCommand) Arguments() []command.Argument {
	args := []command.Argument{}
	return args
}

func (c *SetCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *SetCommand) ParsedArguments(args []string) (map[string]command.Argument, error) {
	return command.ParseArguments(args, c.Arguments())
}

func (c *SetCommand) FlagSet() *flag.FlagSet {
	f := c.Meta.FlagSet(c.Name(), command.FlagSetClient)
	// Required
	f.StringVarP(&c.processType, "process-type", "p", "", "name of process to set")
	// Required
	f.StringVarP(&c.command, "command", "c", "", "command to set")
	f.BoolVarP(&c.stdout, "stdout", "s", false, "write output to stdout")
	f.StringVarP(&c.writePath, "write-path", "w", "", "path to Procfile to write to")

	c.GlobalFlags(f)
	return f
}

func (c *SetCommand) AutocompleteFlags() complete.Flags {
	return command.MergeAutocompleteFlags(
		c.Meta.AutocompleteFlags(command.FlagSetClient),
		c.AutocompleteGlobalFlags(),
		complete.Flags{
			"--process-type": complete.PredictAnything,
			"--command":      complete.PredictAnything,
			"--sdout":        complete.PredictNothing,
			"--write-path":   complete.PredictAnything,
		},
	)
}

func (c *SetCommand) Run(args []string) int {
	flags := c.FlagSet()
	flags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := flags.Parse(args); err != nil {
		c.Ui.Error(err.Error())
		c.Ui.Error(command.CommandErrorText(c))
		return 1
	}

	_, err := c.ParsedArguments(flags.Args())
	if err != nil {
		c.Ui.Error(err.Error())
		c.Ui.Error(command.CommandErrorText(c))
		return 1
	}

	entries, err := parseProcfile(c.procfile, c.delimiter, c.strict)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	if len(entries) == 0 {
		c.Ui.Error("No processes defined")
		return 1
	}

	var validEntries []procfile.ProcfileEntry
	validEntries = append(validEntries, procfile.ProcfileEntry{
		Name:    c.processType,
		Command: c.command,
	})
	for _, entry := range entries {
		if c.processType == entry.Name {
			continue
		}
		validEntries = append(validEntries, entry)
	}

	if err := procfile.OutputProcfile(c.procfile, c.writePath, c.delimiter, c.stdout, validEntries); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	return 0
}
