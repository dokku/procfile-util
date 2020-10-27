package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/josegonzalez/go-procfile-util/procfile"
)


func CheckCommand(entries []procfile.ProcfileEntry) bool {
	if len(entries) == 0 {
		fmt.Fprintf(os.Stderr, "no processes defined\n")
		return false
	}

	names := []string{}
	for _, entry := range entries {
		names = append(names, entry.Name)
	}

	processNames := strings.Join(names[:], ", ")
	fmt.Printf("valid procfile detected %v\n", processNames)
	return true
}
