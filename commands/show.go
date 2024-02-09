package commands

import (
	"fmt"
	"os"
	"procfile-util/procfile"

	"github.com/josegonzalez/cli-skeleton/command"
	"github.com/posener/complete"
	flag "github.com/spf13/pflag"
)

type ShowCommand struct {
	command.Meta
	GlobalFlagCommand

	allowGetenv bool
	envPath     string
	processType string
}

func (c *ShowCommand) Name() string {
	return "show"
}

func (c *ShowCommand) Synopsis() string {
	return "Eats one or more lollipops"
}

func (c *ShowCommand) Help() string {
	return command.CommandHelp(c)
}

func (c *ShowCommand) Examples() map[string]string {
	appName := os.Getenv("CLI_APP_NAME")
	return map[string]string{
		"Command": fmt.Sprintf("%s %s", appName, c.Name()),
	}
}

func (c *ShowCommand) Arguments() []command.Argument {
	args := []command.Argument{}
	return args
}

func (c *ShowCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ShowCommand) ParsedArguments(args []string) (map[string]command.Argument, error) {
	return command.ParseArguments(args, c.Arguments())
}

func (c *ShowCommand) FlagSet() *flag.FlagSet {
	f := c.Meta.FlagSet(c.Name(), command.FlagSetClient)
	f.BoolVarP(&c.allowGetenv, "allow-getenv", "a", false, "allow the use of the existing env when expanding commands")
	f.StringVarP(&c.envPath, "env-file", "e", "", "path to a dotenv file")
	// required?
	f.StringVarP(&c.processType, "process-type", "p", "", "name of process to show")

	c.GlobalFlags(f)
	return f
}

func (c *ShowCommand) AutocompleteFlags() complete.Flags {
	return command.MergeAutocompleteFlags(
		c.Meta.AutocompleteFlags(command.FlagSetClient),
		c.AutocompleteGlobalFlags(),
		complete.Flags{
			"--allow-getenv": complete.PredictNothing,
			"--env-file":     complete.PredictFiles("*"),
			"--process-type": complete.PredictAnything,
		},
	)
}

func (c *ShowCommand) Run(args []string) int {
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

	var foundEntry procfile.ProcfileEntry
	for _, entry := range entries {
		if c.processType == entry.Name {
			foundEntry = entry
			break
		}
	}

	if foundEntry == (procfile.ProcfileEntry{}) {
		c.Ui.Error("No matching process entry found")
		return 1
	}

	command, err := expandEnv(foundEntry, c.envPath, c.allowGetenv, c.defaultPort)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error processing command: %s", err))
		return 1
	}

	c.Ui.Output(command)

	return 0
}
