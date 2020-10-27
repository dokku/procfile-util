package commands

import (
	"fmt"

	"procfile-util/procfile"
)

func ListCommand(entries []procfile.ProcfileEntry) bool {
	for _, entry := range entries {
		fmt.Printf("%v\n", entry.Name)
	}
	return true
}
