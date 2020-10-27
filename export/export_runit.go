package export

import (
	"fmt"
	"os"
	"strconv"

	"github.com/josegonzalez/go-procfile-util/procfile"
)

func ExportRunit(app string, entries []procfile.ProcfileEntry, formations map[string]procfile.FormationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	r, err := loadTemplate("run", "templates/runit/run.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}
	l, err := loadTemplate("log", "templates/runit/log/run.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
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

			fmt.Println("creating:", folderPath)
			os.MkdirAll(folderPath, os.ModePerm)

			fmt.Println("creating:", folderPath+"/env")
			os.MkdirAll(folderPath+"/env", os.ModePerm)

			fmt.Println("creating:", folderPath+"/log")
			os.MkdirAll(folderPath+"/log", os.ModePerm)

			port := portFor(i, num, defaultPort)
			config := templateVars(app, entry, processName, num, port, vars)

			if !writeOutput(r, fmt.Sprintf("%s/run", folderPath), config) {
				return false
			}

			env, ok := config["env"].(map[string]string)
			if !ok {
				fmt.Fprintf(os.Stderr, "invalid env map\n")
				return false
			}

			env["PORT"] = strconv.Itoa(port)
			env["PS"] = app + "-" + processName

			for key, value := range env {
				fmt.Println("writing:", folderPath+"/env/"+key)
				f, err := os.Create(folderPath + "/env/" + key)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error creating file: %s\n", err)
					return false
				}
				defer f.Close()

				if _, err = f.WriteString(value); err != nil {
					fmt.Fprintf(os.Stderr, "error writing output: %s\n", err)
					return false
				}

				if err = f.Sync(); err != nil {
					fmt.Fprintf(os.Stderr, "error syncing output: %s\n", err)
					return false
				}
			}

			if !writeOutput(l, fmt.Sprintf("%s/log/run", folderPath), config) {
				return false
			}

			num += 1
		}
	}

	return true
}
