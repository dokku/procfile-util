package export

import (
	"fmt"
	"os"

	"procfile-util/procfile"

	"github.com/mitchellh/cli"
)

func ExportSysv(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}, ui cli.Ui) bool {
	l, err := loadTemplate("launchd", "templates/sysv/default/init.sh.tmpl")
	if err != nil {
		ui.Error(err.Error())
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
			if err := writeOutput(l, fmt.Sprintf("%s/etc/init.d/%s-%s", location, app, processName), config); err != nil {
				ui.Error(err.Error())
				return false
			}

			num += 1
		}
	}

	return true
}
