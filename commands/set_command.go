package commands

import (
	"procfile-util/procfile"
)

func SetCommand(entries []procfile.ProcfileEntry, processType string, command string, writePath string, stdout bool, delimiter string, path string) bool {
	var validEntries []procfile.ProcfileEntry
	validEntries = append(validEntries, procfile.ProcfileEntry{processType, command})
	for _, entry := range entries {
		if processType == entry.Name {
			continue
		}
		validEntries = append(validEntries, entry)
	}

	return procfile.OutputProcfile(path, writePath, delimiter, stdout, validEntries)
}