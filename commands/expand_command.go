package commands

import (
	"fmt"
	"os"

	"github.com/josegonzalez/go-procfile-util/procfile"
)

func ExpandCommand(entries []procfile.ProcfileEntry, envPath string, allowGetenv bool, processType string, defaultPort int, delimiter string) bool {
	hasErrors := false
	var expandedEntries []procfile.ProcfileEntry
	for _, entry := range entries {
		command, err := expandEnv(entry, envPath, allowGetenv, defaultPort)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error processing command: %s\n", err)
			hasErrors = true
		}

		entry.Command = command
		expandedEntries = append(expandedEntries, entry)
	}

	if hasErrors {
		return false
	}

	for _, entry := range expandedEntries {
		if processType == "" || processType == entry.Name {
			fmt.Printf("%v%v %v\n", entry.Name, delimiter, entry.Command)
		}
	}

	return true
}
