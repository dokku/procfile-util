package commands

import (
	"fmt"

	"github.com/josegonzalez/go-procfile-util/procfile"
)

func ListCommand(entries []procfile.ProcfileEntry) bool {
	for _, entry := range entries {
		fmt.Printf("%v\n", entry.Name)
	}
	return true
}
