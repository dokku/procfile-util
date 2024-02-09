package export

import (
	"fmt"
	"os"
	"strconv"

	"procfile-util/procfile"

	"github.com/mitchellh/cli"
)

func ExportRunit(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}, ui cli.Ui) bool {
	r, err := loadTemplate("run", "templates/runit/run.tmpl")
	if err != nil {
		ui.Error(err.Error())
		return false
	}
	l, err := loadTemplate("log", "templates/runit/log/run.tmpl")
	if err != nil {
		ui.Error(err.Error())
		return false
	}

	if _, err := os.Stat(location + "/service"); os.IsNotExist(err) {
		os.MkdirAll(location+"/service", os.ModePerm)
	}

	for i, entry := range entries {
		num := 1
		count := processCount(entry, formations)

		for num <= count {
			processDirectory := fmt.Sprintf("%s-%s-%d", app, entry.Name, num)
			folderPath := location + "/service/" + processDirectory
			processName := fmt.Sprintf("%s-%d", entry.Name, num)

			ui.Output(fmt.Sprintf("creating: %s", folderPath))
			os.MkdirAll(folderPath, os.ModePerm)

			ui.Output(fmt.Sprintf("creating: %s/env", folderPath))
			os.MkdirAll(folderPath+"/env", os.ModePerm)

			ui.Output(fmt.Sprintf("creating:  %s/log", folderPath))
			os.MkdirAll(folderPath+"/log", os.ModePerm)

			port := portFor(i, num, defaultPort)
			config := templateVars(app, entry, processName, num, port, vars)

			if err := writeOutput(r, fmt.Sprintf("%s/run", folderPath), config); err != nil {
				ui.Error(err.Error())
				return false
			}

			env, ok := config["env"].(map[string]string)
			if !ok {
				ui.Error("Invalid env map")
				return false
			}

			env["PORT"] = strconv.Itoa(port)
			env["PS"] = app + "-" + processName

			for key, value := range env {
				ui.Output(fmt.Sprintf("writing:  %s/env/%s", folderPath, key))
				f, err := os.Create(folderPath + "/env/" + key)
				if err != nil {
					ui.Error(fmt.Sprintf("Error creating file: %s", err))
					return false
				}
				defer f.Close()

				if _, err = f.WriteString(value); err != nil {
					ui.Error(fmt.Sprintf("Error writing output: %s", err))
					return false
				}

				if err = f.Sync(); err != nil {
					ui.Error(fmt.Sprintf("Error syncing output: %s", err))
					return false
				}
			}

			if err := writeOutput(l, fmt.Sprintf("%s/log/run", folderPath), config); err != nil {
				ui.Error(err.Error())
				return false
			}

			num += 1
		}
	}

	return true
}
