package commands

import (
	"fmt"
	"os"

	"procfile-util/procfile"
)

func ExistsCommand(entries []procfile.ProcfileEntry, processType string) bool {
	for _, entry := range entries {
		if processType == entry.Name {
			return true
		}
	}

	fmt.Fprint(os.Stderr, "no matching process entry found\n")
	return false
}
