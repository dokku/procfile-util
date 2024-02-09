package export

import (
	"fmt"
	"os"
	"strconv"
	"text/template"

	"procfile-util/procfile"

	"github.com/mitchellh/cli"
)

type ExportFunc func(string, []procfile.ProcfileEntry, map[string]procfile.FormationEntry, string, int, map[string]interface{}, cli.Ui) bool

func processCount(entry procfile.ProcfileEntry, formations map[string]procfile.FormationEntry) int {
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

func templateVars(app string, entry procfile.ProcfileEntry, processName string, num int, port int, vars map[string]interface{}) map[string]interface{} {
	config := vars
	config["args"] = entry.Args()
	config["args_escaped"] = entry.ArgsEscaped()
	config["command"] = entry.Command
	config["command_list"] = entry.CommandList()
	config["num"] = num
	config["port"] = port
	config["process_name"] = processName
	config["process_type"] = entry.Name
	config["program"] = entry.Program()
	config["ps"] = app + "-" + entry.Name + "." + strconv.Itoa(num)
	if config["description"] == "" {
		config["description"] = fmt.Sprintf("%s.%s process for %s", entry.Name, strconv.Itoa(num), app)
	}

	return config
}

func writeOutput(t *template.Template, outputPath string, variables map[string]interface{}) error {
	fmt.Println("writing:", outputPath)
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer f.Close()

	if err = t.Execute(f, variables); err != nil {
		return fmt.Errorf("error writing output: %w", err)
	}

	if err := os.Chmod(outputPath, 0755); err != nil {
		return fmt.Errorf("error setting mode: %w", err)
	}

	return nil
}

func loadTemplate(name string, filename string) (*template.Template, error) {
	asset, err := Asset(filename)
	if err != nil {
		return nil, err
	}

	t, err := template.New(name).Parse(string(asset))
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %s", err)
	}

	return t, nil
}
