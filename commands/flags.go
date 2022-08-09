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
	f.StringVarP(&c.procfile, "P", "procfile", "Procfile", "path to a procfile")
	f.StringVarP(&c.delimiter, "D", "delimiter", ":", "delimiter in use within procfile")
	f.IntVarP(&c.defaultPort, "d", "default-port", 5000, "default port to use")
	f.BoolVarP(&c.strict, "S", "strict", false, "strictly parse the Procfile")
}

func (c *GlobalFlagCommand) AutocompleteGlobalFlags() complete.Flags {
	return complete.Flags{
		"--procfile":     complete.PredictAnything,
		"--delimiter":    complete.PredictAnything,
		"--default-port": complete.PredictAnything,
		"--strict":       complete.PredictNothing,
	}
}
