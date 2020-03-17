package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

func exportLaunchd(app string, entries []procfileEntry, formations map[string]formationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	service, err := Asset("templates/launchd/launchd.plist.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return false
	}

	s, err := template.New("service").Parse(string(service))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template: %s\n", err)
		return false
	}

	if _, err := os.Stat(location); os.IsNotExist(err) {
		os.MkdirAll(location, os.ModePerm)
	}

	for i, entry := range entries {
		count := processCount(entry, formations)

		num := 1
		for num <= count {
			processName := fmt.Sprintf("%s-%d", entry.Name, num)
			fmt.Println("writing:", app+"-"+processName+".plist")

			config := vars
			config["command"] = entry.Command
			config["command_args"] = entry.commandArgs()
			config["num"] = num
			config["port"] = strconv.Itoa(portFor(i, num, defaultPort))
			config["process_name"] = processName
			config["process_type"] = entry.Name
			if config["description"] == "" {
				config["description"] = fmt.Sprintf("%s process for %s", processName, app)
			}

			f, err := os.Create(location + "/" + app + "-" + processName + ".plist")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error creating file: %s\n", err)
				return false
			}
			defer f.Close()

			if err = s.Execute(f, config); err != nil {
				fmt.Fprintf(os.Stderr, "error writing output: %s\n", err)
				return false
			}

			num += 1
		}
	}

	return true
}

func exportRunit(app string, entries []procfileEntry, formations map[string]formationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	run, err := Asset("templates/runit/run.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return false
	}

	log, err := Asset("templates/runit/log/run.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return false
	}

	r, err := template.New("run").Parse(string(run))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template: %s\n", err)
		return false
	}

	l, err := template.New("log").Parse(string(log))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template: %s\n", err)
		return false
	}

	if _, err := os.Stat(location); os.IsNotExist(err) {
		os.MkdirAll(location, os.ModePerm)
	}

	for i, entry := range entries {
		count := processCount(entry, formations)

		num := 1
		for num <= count {
			processDirectory := fmt.Sprintf("%s-%s-%d", app, entry.Name, num)
			folderPath := location + "/" + processDirectory
			processName := fmt.Sprintf("%s-%d", entry.Name, num)

			fmt.Println("creating:", app+"-"+processName)
			os.MkdirAll(folderPath, os.ModePerm)

			fmt.Println("creating:", app+"-"+processName+"/env")
			os.MkdirAll(folderPath+"/env", os.ModePerm)

			fmt.Println("creating:", app+"-"+processName+"/log")
			os.MkdirAll(folderPath+"/log", os.ModePerm)

			port := portFor(i, num, defaultPort)

			config := vars
			config["command"] = entry.Command
			config["command_args"] = entry.commandArgs()
			config["num"] = num
			config["port"] = port
			config["process_name"] = processName
			config["process_type"] = entry.Name
			if config["description"] == "" {
				config["description"] = fmt.Sprintf("%s process for %s", processName, app)
			}

			fmt.Println("writing:", app+"-"+processName+"/run")
			f, err := os.Create(folderPath + "/run")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error creating file: %s\n", err)
				return false
			}
			defer f.Close()

			if err = r.Execute(f, config); err != nil {
				fmt.Fprintf(os.Stderr, "error writing output: %s\n", err)
				return false
			}

			fmt.Println("setting", app+"-"+processName+"/run", "to mode 755")
			if err := os.Chmod(folderPath+"/run", 0755); err != nil {
				fmt.Fprintf(os.Stderr, "error setting mode: %s\n", err)
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
				fmt.Println("writing:", app+"-"+processName+"/env/"+key)
				f, err = os.Create(folderPath + "/env/" + key)
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

			fmt.Println("writing:", app+"-"+processName+"/log/run")
			f, err = os.Create(folderPath + "/log/run")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error creating file: %s\n", err)
				return false
			}
			defer f.Close()

			if err = l.Execute(f, config); err != nil {
				fmt.Fprintf(os.Stderr, "error writing output: %s\n", err)
				return false
			}

			fmt.Println("setting", app+"-"+processName+"/log/run", "to mode 755")
			if err := os.Chmod(folderPath+"/log/run", 0755); err != nil {
				fmt.Fprintf(os.Stderr, "error setting mode: %s\n", err)
				return false
			}

			num += 1
		}
	}

	return true
}

func exportSystemd(app string, entries []procfileEntry, formations map[string]formationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	target, err := Asset("templates/systemd/default/control.target.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return false
	}

	service, err := Asset("templates/systemd/default/program.service.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return false
	}

	t, err := template.New("target").Parse(string(target))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template: %s\n", err)
		return false
	}

	s, err := template.New("service").Parse(string(service))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template: %s\n", err)
		return false
	}

	if _, err := os.Stat(location); os.IsNotExist(err) {
		os.MkdirAll(location, os.ModePerm)
	}

	processes := []string{}
	for i, entry := range entries {
		count := processCount(entry, formations)

		num := 1
		for num <= count {
			processName := fmt.Sprintf("%s-%d", entry.Name, num)
			fileName := fmt.Sprintf("%s.%d", entry.Name, num)
			processes = append(processes, fmt.Sprintf(app+"-%s.service", fileName))
			fmt.Println("writing:", app+"-"+fileName+".service")

			config := vars
			config["command"] = entry.Command
			config["command_args"] = entry.commandArgs()
			config["num"] = num
			config["port"] = portFor(i, num, defaultPort)
			config["process_name"] = processName
			config["process_type"] = entry.Name
			if config["description"] == "" {
				config["description"] = fmt.Sprintf("%s process for %s", processName, app)
			}

			f, err := os.Create(location + "/" + app + "-" + fileName + ".service")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error creating file: %s\n", err)
				return false
			}
			defer f.Close()

			if err = s.Execute(f, config); err != nil {
				fmt.Fprintf(os.Stderr, "error writing output: %s\n", err)
				return false
			}

			num += 1
		}
	}

	config := vars
	config["processes"] = processes
	fmt.Println("writing:", app+".target")
	f, err := os.Create(location + "/" + app + ".target")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating file: %s\n", err)
		return false
	}
	defer f.Close()

	if err = t.Execute(f, config); err != nil {
		fmt.Fprintf(os.Stderr, "error writing output: %s\n", err)
		return false
	}

	return true
}

func exportSystemdUser(app string, entries []procfileEntry, formations map[string]formationEntry, location string, defaultPort int, vars map[string]interface{}) bool {
	service, err := Asset("templates/systemd-user/default/program.service.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return false
	}

	s, err := template.New("service").Parse(string(service))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template: %s\n", err)
		return false
	}

	if _, err := os.Stat(location); os.IsNotExist(err) {
		os.MkdirAll(location, os.ModePerm)
	}

	processes := []string{}
	for i, entry := range entries {
		count := processCount(entry, formations)

		num := 1
		for num <= count {
			processName := fmt.Sprintf("%s-%d", entry.Name, num)
			processes = append(processes, fmt.Sprintf("%s.service", processName))
			fmt.Println("writing:", app+"-"+processName+".service")

			config := vars
			config["command"] = entry.Command
			config["num"] = num
			config["port"] = portFor(i, num, defaultPort)
			config["process_name"] = processName
			config["process_type"] = entry.Name
			if config["description"] == "" {
				config["description"] = fmt.Sprintf("%s process for %s", processName, app)
			}

			f, err := os.Create(location + "/" + app + "-" + processName + ".service")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error creating file: %s\n", err)
				return false
			}
			defer f.Close()

			if err = s.Execute(f, config); err != nil {
				fmt.Fprintf(os.Stderr, "error writing output: %s\n", err)
				return false
			}

			num += 1
		}
	}

	return true
}

func processCount(entry procfileEntry, formations map[string]formationEntry) int {
	count := 0
	if f, ok := formations["all"]; ok {
		count = f.Count
	}
	if f, ok := formations[entry.Name]; ok {
		count = f.Count
	}
	return count
}

func portFor(processIndex int, instance int, base int) int {
	return 5000 + (processIndex * 100) + (instance - 1)
}
