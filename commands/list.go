package commands

import (
	"fmt"
	"os"

	"github.com/josegonzalez/cli-skeleton/command"
	"github.com/posener/complete"
	flag "github.com/spf13/pflag"
)

type ListCommand struct {
	command.Meta
	GlobalFlagCommand
}

func (c *ListCommand) Name() string {
	return "list"
}

func (c *ListCommand) Synopsis() string {
	return "Eats one or more lollipops"
}

func (c *ListCommand) Help() string {
	return command.CommandHelp(c)
}

func (c *ListCommand) Examples() map[string]string {
	appName := os.Getenv("CLI_APP_NAME")
	return map[string]string{
		"Command": fmt.Sprintf("%s %s", appName, c.Name()),
	}
}

func (c *ListCommand) Arguments() []command.Argument {
	args := []command.Argument{}
	return args
}

func (c *ListCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ListCommand) ParsedArguments(args []string) (map[string]command.Argument, error) {
	return command.ParseArguments(args, c.Arguments())
}

func (c *ListCommand) FlagSet() *flag.FlagSet {
	f := c.Meta.FlagSet(c.Name(), command.FlagSetClient)
	c.GlobalFlags(f)
	return f
}

func (c *ListCommand) AutocompleteFlags() complete.Flags {
	return command.MergeAutocompleteFlags(
		c.Meta.AutocompleteFlags(command.FlagSetClient),
		c.AutocompleteGlobalFlags(),
		complete.Flags{},
	)
}

func (c *ListCommand) Run(args []string) int {
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
		c.Ui.Output(entry.Name)
	}

	return 0
}
