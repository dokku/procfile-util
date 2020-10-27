package commands

import (
	"fmt"
	"os"

	"github.com/josegonzalez/go-procfile-util/procfile"
)

func ShowCommand(entries []procfile.ProcfileEntry, envPath string, allowGetenv bool, processType string, defaultPort int) bool {
	var foundEntry procfile.ProcfileEntry
	for _, entry := range entries {
		if processType == entry.Name {
			foundEntry = entry
			break
		}
	}

	if foundEntry == (procfile.ProcfileEntry{}) {
		fmt.Fprintf(os.Stderr, "no matching process entry found\n")
		return false
	}

	command, err := expandEnv(foundEntry, envPath, allowGetenv, defaultPort)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing command: %s\n", err)
		return false
	}

	fmt.Printf("%v\n", command)
	return true
}
