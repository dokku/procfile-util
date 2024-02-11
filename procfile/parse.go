package procfile

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func ParseFormation(formation string) (map[string]FormationEntry, error) {
	entries := make(map[string]FormationEntry)
	for _, formation := range strings.Split(formation, ",") {
		parts := strings.Split(formation, "=")
		if len(parts) != 2 {
			return entries, fmt.Errorf("invalid formation: %s", formation)
		}

		i, err := strconv.Atoi(parts[1])
		if err != nil {
			return entries, fmt.Errorf("invalid formation: %s", err)
		}

		entries[parts[0]] = FormationEntry{
			Name:  parts[0],
			Count: i,
		}
	}

	return entries, nil
}

// ParseProcfile parses text as a procfile and returns a list of procfile entries
func ParseProcfile(text string, delimiter string, strict bool) ([]ProcfileEntry, error) {
	var entries []ProcfileEntry
	reCmd, _ := regexp.Compile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?)\s*` + delimiter + `\s*(.+)$`)
	reOldCmd, _ := regexp.Compile(`^([A-Za-z0-9_-]+)\s*` + delimiter + `\s*(.+)$`)

	reComment, _ := regexp.Compile(`^(.*)\s#.+$`)
	reForwardslashComment, _ := regexp.Compile(`^(.*)\s//.+$`)

	lineNumber := 0
	names := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		oldParams := reOldCmd.FindStringSubmatch(line)
		params := reCmd.FindStringSubmatch(line)
		isCommand := len(params) == 4
		isOldCommand := len(oldParams) == 3
		isComment := strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//")
		if isComment {
			continue
		}

		if !isCommand {
			if !isOldCommand {
				return entries, fmt.Errorf("invalid line in procfile, line %d", lineNumber)
			}

			if strict {
				return entries, fmt.Errorf("process name contains invalid characters, line %d", lineNumber)
			}
		}

		name := ""
		cmd := ""
		if strict {
			name, cmd = params[1], params[3]
		} else {
			name, cmd = oldParams[1], oldParams[2]
		}

		if len(name) > 63 {
			return entries, fmt.Errorf("process name over 63 characters, line %d", lineNumber)
		}

		if names[name] {
			return entries, fmt.Errorf("process names must be unique, line %d", lineNumber)
		}
		names[name] = true

		commentParams := reComment.FindStringSubmatch(cmd)
		reForwardslashCommentParams := reForwardslashComment.FindStringSubmatch(cmd)
		if len(commentParams) == 2 {
			cmd = commentParams[1]
		} else if len(reForwardslashCommentParams) == 2 {
			cmd = reForwardslashCommentParams[1]
		}

		cmd = strings.TrimSpace(cmd)
		if strings.HasPrefix(cmd, "#") || strings.HasPrefix(cmd, "//") {
			return entries, fmt.Errorf("comment specified in place of command, line %d", lineNumber)
		}

		if len(cmd) == 0 {
			return entries, fmt.Errorf("no command specified, line %d", lineNumber)
		}

		entries = append(entries, ProcfileEntry{name, cmd})
	}

	if scanner.Err() != nil {
		return entries, scanner.Err()
	}

	if len(entries) == 0 {
		return entries, fmt.Errorf("no entries found in Procfile")
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	return entries, nil
}
