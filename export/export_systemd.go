package export

import (
	"fmt"
	"os"

	"github.com/josegonzalez/go-procfile-util/procfile"
)

func ExportSystemd(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	t, err := loadTemplate("target", "templates/systemd/default/control.target.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	s, err := loadTemplate("service", "templates/systemd/default/program.service.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	if _, err := os.Stat(location + "/etc/systemd/system/"); os.IsNotExist(err) {
		os.MkdirAll(location+"/etc/systemd/system/", os.ModePerm)
	}

	processes := []string{}
	for i, entry := range entries {
		num := 1
		count := processCount(entry, formations)

		for num <= count {
			processName := fmt.Sprintf("%s-%d", entry.Name, num)
			fileName := fmt.Sprintf("%s.%d", entry.Name, num)
			processes = append(processes, fmt.Sprintf(app+"-%s.service", fileName))

			port := portFor(i, num, defaultPort)
			config := templateVars(app, entry, processName, num, port, vars)
			if !writeOutput(s, fmt.Sprintf("%s/etc/systemd/system/%s-%s.service", location, app, fileName), config) {
				return false
			}

			num += 1
		}
	}

	config := vars
	config["processes"] = processes
	if writeOutput(t, fmt.Sprintf("%s/etc/systemd/system/%s.target", location, app), config) {
		fmt.Println("You will want to run 'systemctl --system daemon-reload' to activate the service on the target host")
		return true
	}

	return true
}
