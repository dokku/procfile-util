package commands

import (
	"fmt"
	"os"
	"procfile-util/procfile"

	"github.com/josegonzalez/cli-skeleton/command"
	"github.com/posener/complete"
	flag "github.com/spf13/pflag"
)

type ExpandCommand struct {
	command.Meta
	GlobalFlagCommand

	allowGetenv bool
	envPath     string
	processType string
}

func (c *ExpandCommand) Name() string {
	return "expand"
}

func (c *ExpandCommand) Synopsis() string {
	return "Eats one or more lollipops"
}

func (c *ExpandCommand) Help() string {
	return command.CommandHelp(c)
}

func (c *ExpandCommand) Examples() map[string]string {
	appName := os.Getenv("CLI_APP_NAME")
	return map[string]string{
		"Command": fmt.Sprintf("%s %s", appName, c.Name()),
	}
}

func (c *ExpandCommand) Arguments() []command.Argument {
	args := []command.Argument{}
	return args
}

func (c *ExpandCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ExpandCommand) ParsedArguments(args []string) (map[string]command.Argument, error) {
	return command.ParseArguments(args, c.Arguments())
}

func (c *ExpandCommand) FlagSet() *flag.FlagSet {
	f := c.Meta.FlagSet(c.Name(), command.FlagSetClient)
	f.BoolVarP(&c.allowGetenv, "allow-getenv", "a", false, "allow the use of the existing env when expanding commands")
	f.StringVarP(&c.envPath, "env-file", "e", "", "path to a dotenv file")
	f.StringVarP(&c.processType, "process-type", "p", "", "name of process to expand")

	c.GlobalFlags(f)
	return f
}

func (c *ExpandCommand) AutocompleteFlags() complete.Flags {
	return command.MergeAutocompleteFlags(
		c.Meta.AutocompleteFlags(command.FlagSetClient),
		c.AutocompleteGlobalFlags(),
		complete.Flags{
			"--allow-getenv": complete.PredictNothing,
			"--env-path":     complete.PredictFiles("*"),
			"--process-type": complete.PredictAnything,
		},
	)
}

func (c *ExpandCommand) Run(args []string) int {
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
	if len(entries) == 0 {
		c.Ui.Error("No processes defined")
		return 1
	}

	hasErrors := false
	var expandedEntries []procfile.ProcfileEntry
	for _, entry := range entries {
		command, err := expandEnv(entry, c.envPath, c.allowGetenv, c.defaultPort)
		if err != nil {
			c.Ui.Error(fmt.Sprintf("error processing command: %s", err))
			hasErrors = true
		}

		entry.Command = command
		expandedEntries = append(expandedEntries, entry)
	}

	if hasErrors {
		return 1
	}

	for _, entry := range expandedEntries {
		if c.processType == "" || c.processType == entry.Name {
			c.Ui.Output(fmt.Sprintf("%v%v %v", entry.Name, c.delimiter, entry.Command))
		}
	}

	return 0
}
