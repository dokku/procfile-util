package commands

import (
	"fmt"
	"os"

	"github.com/josegonzalez/cli-skeleton/command"
	"github.com/posener/complete"
	flag "github.com/spf13/pflag"
)

type ExistsCommand struct {
	command.Meta
	GlobalFlagCommand

	processType string
}

func (c *ExistsCommand) Name() string {
	return "exists"
}

func (c *ExistsCommand) Synopsis() string {
	return "Eats one or more lollipops"
}

func (c *ExistsCommand) Help() string {
	return command.CommandHelp(c)
}

func (c *ExistsCommand) Examples() map[string]string {
	appName := os.Getenv("CLI_APP_NAME")
	return map[string]string{
		"Command": fmt.Sprintf("%s %s", appName, c.Name()),
	}
}

func (c *ExistsCommand) Arguments() []command.Argument {
	args := []command.Argument{}
	return args
}

func (c *ExistsCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ExistsCommand) ParsedArguments(args []string) (map[string]command.Argument, error) {
	return command.ParseArguments(args, c.Arguments())
}

func (c *ExistsCommand) FlagSet() *flag.FlagSet {
	f := c.Meta.FlagSet(c.Name(), command.FlagSetClient)
	// required?
	f.StringVarP(&c.processType, "process-type", "p", "", "name of process to delete")
	c.GlobalFlags(f)
	return f
}

func (c *ExistsCommand) AutocompleteFlags() complete.Flags {
	return command.MergeAutocompleteFlags(
		c.Meta.AutocompleteFlags(command.FlagSetClient),
		c.AutocompleteGlobalFlags(),
		complete.Flags{
			"--process-type": complete.PredictAnything,
		},
	)
}

func (c *ExistsCommand) Run(args []string) int {
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

	for _, entry := range entries {
		if c.processType == entry.Name {
			return 0
		}
	}

	c.Ui.Error("No matching process entry found")

	return 1
}
