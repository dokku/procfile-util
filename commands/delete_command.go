package commands

import (
	"procfile-util/procfile"
)

func DeleteCommand(entries []procfile.ProcfileEntry, processType string, writePath string, stdout bool, delimiter string, path string) bool {
	var validEntries []procfile.ProcfileEntry
	for _, entry := range entries {
		if processType == entry.Name {
			continue
		}
		validEntries = append(validEntries, entry)
	}

	return procfile.OutputProcfile(path, writePath, delimiter, stdout, validEntries)
}
