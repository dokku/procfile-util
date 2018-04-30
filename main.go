package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/joho/godotenv"
)

type procfileEntry struct {
	Name    string
	Command string
}

func parseProcfile(path string) ([]procfileEntry, error) {
	var entries []procfileEntry
	re, _ := regexp.Compile(`^([A-Za-z0-9_]+):\s*(.+)$`)

	f, err := os.Open(path)
	if err != nil {
		return entries, err
	}
	defer f.Close()

	names := make(map[string]bool)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}

		params := re.FindStringSubmatch(scanner.Text())
		if len(params) != 3 {
			continue
		}

		name, cmd := params[1], params[2]

		if names[name] {
			return entries, fmt.Errorf("process names must be unique")
		}
		names[name] = true

		entries = append(entries, procfileEntry{name, cmd})
	}

	if scanner.Err() != nil {
		return entries, scanner.Err()
	}

	if len(entries) == 0 {
		return entries, fmt.Errorf("no entries found in Procfile")
	}

	return entries, nil
}

func expandEnv(s string, envPath string, allowEnv bool) (string, error) {
	expandFunc := func(key string) string {
		return ""
	}

	if allowEnv {
		expandFunc = os.Getenv
	}

	if envPath != "" {
		b, err := ioutil.ReadFile(envPath) // just pass the file name
		if err != nil {
			return "", err
		}

		content := string(b)
		env, err := godotenv.Unmarshal(content)
		if err != nil {
			return "", err
		}

		expandFunc = func(key string) string {
			if val, ok := env[key]; ok {
				return val
			}
			if allowEnv {
				return os.Getenv(key)
			}
			return ""
		}
	}

	os.Setenv("EXPENV_PARENTHESIS", "$(")
	s = strings.Replace(s, "$(", "${EXPENV_PARENTHESIS}", -1)
	return os.Expand(s, expandFunc), nil
}

func existsCommand(procfile string, processType string) bool {
	entries, err := parseProcfile(procfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	for _, entry := range entries {
		if processType == entry.Name {
			return true
		}
	}

	fmt.Fprint(os.Stderr, "no matching process entry found\n")
	return false
}

func expandCommand(procfile string, envPath string, allowGetenv bool) bool {
	entries, err := parseProcfile(procfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	hasErrors := false
	commands := make(map[string]string)
	for _, entry := range entries {
		command, err := expandEnv(entry.Command, envPath, allowGetenv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error processing command: %s\n", err)
			hasErrors = true
		}

		commands[entry.Name] = command
	}

	if hasErrors {
		return false
	}

	for k, v := range commands {
		fmt.Printf("%v: %v\n", k, v)
	}
	return true
}

func listCommand(procfile string) bool {
	entries, err := parseProcfile(procfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	for _, entry := range entries {
		fmt.Printf("%v\n", entry.Name)
	}
	return true
}

func showCommand(procfile string, envPath string, allowGetenv bool, processType string) bool {
	entries, err := parseProcfile(procfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	var foundEntry procfileEntry
	for _, entry := range entries {
		if processType == entry.Name {
			foundEntry = entry
			break
		}
	}

	if foundEntry == (procfileEntry{}) {
		fmt.Fprintf(os.Stderr, "no matching process entry found\n")
		return false
	}

	command, err := expandEnv(foundEntry.Command, envPath, allowGetenv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing command: %s\n", err)
		return false
	}

	fmt.Printf("%v\n", command)
	return true
}

func main() {
	parser := argparse.NewParser("procfile-parser", "A procfile parsing tool")
	procfileFlag := parser.String("P", "procfile", &argparse.Options{Default: "Procfile", Help: "path to a procfile"})

	existsCmd := parser.NewCommand("exists", "check if a process type exists")
	processTypeExistsFlag := existsCmd.String("p", "process-type", &argparse.Options{Help: "name of process to retrieve"})

	expandCmd := parser.NewCommand("expand", "expands a procfile against a specific environment")
	allowGetenvExpandFlag := expandCmd.Flag("a", "allow-getenv", &argparse.Options{Help: "allow the use of the existing env when expanding commands"})
	envPathExpandFlag := expandCmd.String("e", "env-file", &argparse.Options{Help: "path to a dotenv file"})

	listCmd := parser.NewCommand("list", "list all process types in a procfile")

	showCmd := parser.NewCommand("show", "show the command for a specific process type")
	allowGetenvShowFlag := showCmd.Flag("a", "allow-getenv", &argparse.Options{Help: "allow the use of the existing env when expanding commands"})
	envPathShowFlag := showCmd.String("e", "env-file", &argparse.Options{Help: "path to a dotenv file"})
	processTypeShowFlag := showCmd.String("p", "process-type", &argparse.Options{Help: "name of process to retrieve"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", parser.Usage(err))
		os.Exit(1)
		return
	}

	success := false
	if existsCmd.Happened() {
		success = existsCommand(*procfileFlag, *processTypeExistsFlag)
	} else if expandCmd.Happened() {
		success = expandCommand(*procfileFlag, *envPathExpandFlag, *allowGetenvExpandFlag)
	} else if listCmd.Happened() {
		success = listCommand(*procfileFlag)
	} else if showCmd.Happened() {
		success = showCommand(*procfileFlag, *envPathShowFlag, *allowGetenvShowFlag, *processTypeShowFlag)
	} else {
		fmt.Print(parser.Usage(err))
	}

	if !success {
		os.Exit(1)
	}
}
