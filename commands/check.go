package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/josegonzalez/cli-skeleton/command"
	"github.com/posener/complete"
	flag "github.com/spf13/pflag"
)

type CheckCommand struct {
	command.Meta
	GlobalFlagCommand
}

func (c *CheckCommand) Name() string {
	return "check"
}

func (c *CheckCommand) Synopsis() string {
	return "Eats one or more lollipops"
}

func (c *CheckCommand) Help() string {
	return command.CommandHelp(c)
}

func (c *CheckCommand) Examples() map[string]string {
	appName := os.Getenv("CLI_APP_NAME")
	return map[string]string{
		"Check if the procfile is valid": fmt.Sprintf("%s %s", appName, c.Name()),
	}
}

func (c *CheckCommand) Arguments() []command.Argument {
	args := []command.Argument{}
	return args
}

func (c *CheckCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *CheckCommand) ParsedArguments(args []string) (map[string]command.Argument, error) {
	return command.ParseArguments(args, c.Arguments())
}

func (c *CheckCommand) FlagSet() *flag.FlagSet {
	f := c.Meta.FlagSet(c.Name(), command.FlagSetClient)
	c.GlobalFlags(f)
	return f
}

func (c *CheckCommand) AutocompleteFlags() complete.Flags {
	return command.MergeAutocompleteFlags(
		c.Meta.AutocompleteFlags(command.FlagSetClient),
		c.AutocompleteGlobalFlags(),
		complete.Flags{},
	)
}

func (c *CheckCommand) Run(args []string) int {
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

	names := []string{}
	for _, entry := range entries {
		names = append(names, entry.Name)
	}

	processNames := strings.Join(names[:], ", ")
	c.Ui.Output(fmt.Sprintf("valid procfile detected %v", processNames))

	return 0
}
