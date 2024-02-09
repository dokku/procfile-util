package export

import (
	"fmt"
	"os"

	"procfile-util/procfile"

	"github.com/mitchellh/cli"
)

func ExportUpstart(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}, ui cli.Ui) bool {
	p, err := loadTemplate("program", "templates/upstart/default/program.conf.tmpl")
	if err != nil {
		ui.Error(err.Error())
		return false
	}

	c, err := loadTemplate("app", "templates/upstart/default/control.conf.tmpl")
	if err != nil {
		ui.Error(err.Error())
		return false
	}

	t, err := loadTemplate("process-type", "templates/upstart/default/process-type.conf.tmpl")
	if err != nil {
		ui.Error(err.Error())
		return false
	}

	if _, err := os.Stat(location + "/etc/init/"); os.IsNotExist(err) {
		os.MkdirAll(location+"/etc/init/", os.ModePerm)
	}

	for i, entry := range entries {
		num := 1
		count := processCount(entry, formations)

		if count > 0 {
			config := vars
			config["process_type"] = entry.Name
			if err := writeOutput(t, fmt.Sprintf("%s/etc/init/%s-%s.conf", location, app, entry.Name), config); err != nil {
				ui.Error(err.Error())
				return false
			}
		}

		for num <= count {
			processName := fmt.Sprintf("%s-%d", entry.Name, num)
			fileName := fmt.Sprintf("%s-%d", entry.Name, num)
			port := portFor(i, num, defaultPort)
			config := templateVars(app, entry, processName, num, port, vars)
			if err := writeOutput(p, fmt.Sprintf("%s/etc/init/%s-%s.conf", location, app, fileName), config); err != nil {
				ui.Error(err.Error())
				return false
			}

			num += 1
		}
	}

	config := vars
	if err := writeOutput(c, fmt.Sprintf("%s/etc/init/%s.conf", location, app), config); err != nil {
		ui.Error(err.Error())
		return false
	}

	return true
}
