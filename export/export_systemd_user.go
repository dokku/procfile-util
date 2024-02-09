package export

import (
	"fmt"
	"os"

	"procfile-util/procfile"

	"github.com/mitchellh/cli"
)

func ExportSystemdUser(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}, ui cli.Ui) bool {
	s, err := loadTemplate("service", "templates/systemd-user/default/program.service.tmpl")
	if err != nil {
		ui.Error(err.Error())
		return false
	}

	path := vars["home"].(string) + "/.config/systemd/user/"
	if _, err := os.Stat(location + path); os.IsNotExist(err) {
		os.MkdirAll(location+path, os.ModePerm)
	}

	for i, entry := range entries {
		num := 1
		count := processCount(entry, formations)

		for num <= count {
			processName := fmt.Sprintf("%s-%d", entry.Name, num)
			port := portFor(i, num, defaultPort)
			config := templateVars(app, entry, processName, num, port, vars)
			if err := writeOutput(s, fmt.Sprintf("%s%s%s-%s.service", location, path, app, processName), config); err != nil {
				ui.Error(err.Error())
				return false
			}

			num += 1
		}
	}

	ui.Output("You will want to run 'systemctl --user daemon-reload' to activate the service on the target host")
	return true
}
