package export

import (
	"fmt"
	"os"

	"github.com/josegonzalez/go-procfile-util/procfile"
)

func ExportLaunchd(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	l, err := loadTemplate("launchd", "templates/launchd/launchd.plist.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	if _, err := os.Stat(location + "/Library/LaunchDaemons/"); os.IsNotExist(err) {
		os.MkdirAll(location+"/Library/LaunchDaemons/", os.ModePerm)
	}

	for i, entry := range entries {
		num := 1
		count := processCount(entry, formations)

		for num <= count {
			processName := fmt.Sprintf("%s-%d", entry.Name, num)
			port := portFor(i, num, defaultPort)
			config := templateVars(app, entry, processName, num, port, vars)
			if !writeOutput(l, fmt.Sprintf("%s/Library/LaunchDaemons/%s-%s.plist", location, app, processName), config) {
				return false
			}

			num += 1
		}
	}

	return true
}
