package commands

import (
	"github.com/posener/complete"
	flag "github.com/spf13/pflag"
)

type GlobalFlagCommand struct {
	procfile    string
	delimiter   string
	defaultPort int
	strict      bool
}

func (c *GlobalFlagCommand) GlobalFlags(f *flag.FlagSet) {
	f.StringVarP(&c.procfile, "procfile", "P", "Procfile", "path to a procfile")
	f.StringVarP(&c.delimiter, "delimiter", "D", ":", "delimiter in use within procfile")
	f.IntVarP(&c.defaultPort, "default-port", "d", 5000, "default port to use")
	f.BoolVarP(&c.strict, "strict", "S", false, "strictly parse the Procfile")
}

func (c *GlobalFlagCommand) AutocompleteGlobalFlags() complete.Flags {
	return complete.Flags{
		"--procfile":     complete.PredictAnything,
		"--delimiter":    complete.PredictAnything,
		"--default-port": complete.PredictAnything,
		"--strict":       complete.PredictNothing,
	}
}
