package export

import (
	"fmt"
	"os"

	"procfile-util/procfile"
)

func ExportSystemdUser(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	s, err := loadTemplate("service", "templates/systemd-user/default/program.service.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
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
			if !writeOutput(s, fmt.Sprintf("%s%s%s-%s.service", location, path, app, processName), config) {
				return false
			}

			num += 1
		}
	}

	fmt.Println("You will want to run 'systemctl --user daemon-reload' to activate the service on the target host")
	return true
}
