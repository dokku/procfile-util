package export

import (
	"fmt"
	"os"

	"github.com/josegonzalez/go-procfile-util/procfile"
)

func ExportSysv(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	l, err := loadTemplate("launchd", "templates/sysv/default/init.sh.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	if _, err := os.Stat(location + "/etc/init.d/"); os.IsNotExist(err) {
		os.MkdirAll(location+"/etc/init.d/", os.ModePerm)
	}

	for i, entry := range entries {
		num := 1
		count := processCount(entry, formations)

		for num <= count {
			processName := fmt.Sprintf("%s-%d", entry.Name, num)
			port := portFor(i, num, defaultPort)
			config := templateVars(app, entry, processName, num, port, vars)
			if !writeOutput(l, fmt.Sprintf("%s/etc/init.d/%s-%s", location, app, processName), config) {
				return false
			}

			num += 1
		}
	}

	return true
}
