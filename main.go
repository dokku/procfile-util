package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/andrew-d/go-termutil"
	"github.com/joho/godotenv"
)

type procfileEntry struct {
	Name    string
	Command string
}

const portEnvVar = "PORT"

// Version contains the procfile-util version
var Version string

// Loglevel stores the current app log level
var Loglevel = "info"

func logMessage(message string, level string) {
	if level == "info" {
		fmt.Println(message)
		return
	}

	if Loglevel == "debug" {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%v\n", message))
	}
}

func debugMessage(message string) {
	logMessage(message, "debug")
}

func infoMessage(message string) {
	logMessage(message, "info")
}

func getProcfile(path string) (string, error) {
	if !termutil.Isatty(os.Stdin.Fd()) {
		debugMessage("Reading input from stdin")
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}

	debugMessage(fmt.Sprintf("Reading input from file: %v", path))
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return strings.Join(lines, "\n"), err
}

func parseProcfile(path string, delimiter string) ([]procfileEntry, error) {
	var entries []procfileEntry
	re, _ := regexp.Compile(`^([A-Za-z0-9_]+)` + delimiter + `\s*(.+)$`)

	text, err := getProcfile(path)
	if err != nil {
		return entries, err
	}

	names := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		params := re.FindStringSubmatch(line)
		if len(params) != 3 {
			debugMessage(fmt.Sprintf("No matching params in line: %v", line))
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

func expandEnv(e procfileEntry, envPath string, allowEnv bool, defaultPort string) (string, error) {
	baseExpandFunc := func(key string) string {
		if key == "PS" {
			return os.Getenv("PS")
		}
		if key == portEnvVar {
			return defaultPort
		}
		return ""
	}

	expandFunc := func(key string) string {
		return baseExpandFunc(key)
	}

	if allowEnv {
		debugMessage("Allowing getenv variable expansion")
		expandFunc = func(key string) string {
			value := os.Getenv(key)
			if value == "" {
				value = baseExpandFunc(key)
			}
			return value
		}
	}

	if envPath != "" {
		b, err := ioutil.ReadFile(envPath)
		if err != nil {
			return "", err
		}

		content := string(b)
		env, err := godotenv.Unmarshal(content)
		if err != nil {
			return "", err
		}

		debugMessage("Allowing .env variable expansion")
		expandFunc = func(key string) string {
			if val, ok := env[key]; ok {
				return val
			}
			value := ""
			if allowEnv {
				value = os.Getenv(key)
			}
			if value == "" {
				value = baseExpandFunc(key)
			}
			return value
		}
	}

	os.Setenv("PS", e.Name)
	os.Setenv("EXPENV_PARENTHESIS", "$(")
	s := strings.Replace(e.Command, "$(", "${EXPENV_PARENTHESIS}", -1)
	return os.Expand(s, expandFunc), nil
}

func checkCommand(entries []procfileEntry) bool {
	if len(entries) == 0 {
		fmt.Fprintf(os.Stderr, "no processes defined\n")
		return false
	}

	names := []string{}
	for _, entry := range entries {
		names = append(names, entry.Name)
	}

	processNames := strings.Join(names[:], ", ")
	fmt.Printf("valid procfile detected %v\n", processNames)
	return true
}

func existsCommand(entries []procfileEntry, processType string) bool {
	for _, entry := range entries {
		if processType == entry.Name {
			return true
		}
	}

	fmt.Fprint(os.Stderr, "no matching process entry found\n")
	return false
}

func expandCommand(entries []procfileEntry, envPath string, allowGetenv bool, processType string, defaultPort string) bool {
	hasErrors := false
	commands := make(map[string]string)
	for _, entry := range entries {
		command, err := expandEnv(entry, envPath, allowGetenv, defaultPort)
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
		if processType == "" || processType == k {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
	return true
}

func listCommand(entries []procfileEntry) bool {
	for _, entry := range entries {
		fmt.Printf("%v\n", entry.Name)
	}
	return true
}

func showCommand(entries []procfileEntry, envPath string, allowGetenv bool, processType string, defaultPort string) bool {
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

	command, err := expandEnv(foundEntry, envPath, allowGetenv, defaultPort)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing command: %s\n", err)
		return false
	}

	fmt.Printf("%v\n", command)
	return true
}

func main() {
	parser := argparse.NewParser("procfile-util", "A procfile parsing tool")
	loglevelFlag := parser.Selector("l", "loglevel", []string{"info", "debug"}, &argparse.Options{Default: "info", Help: "loglevel to use"})
	procfileFlag := parser.String("P", "procfile", &argparse.Options{Default: "Procfile", Help: "path to a procfile"})
	delimiterFlag := parser.String("D", "delimiter", &argparse.Options{Default: ":", Help: "delimiter in use within procfile"})
	versionFlag := parser.Flag("v", "version", &argparse.Options{Help: "show version"})
	defaultPortFlag := parser.String("d", "default-port", &argparse.Options{Default: "5000", Help: "default port to use"})

	existsCmd := parser.NewCommand("exists", "check if a process type exists")
	processTypeExistsFlag := existsCmd.String("p", "process-type", &argparse.Options{Help: "name of process to retrieve"})

	expandCmd := parser.NewCommand("expand", "expands a procfile against a specific environment")
	allowGetenvExpandFlag := expandCmd.Flag("a", "allow-getenv", &argparse.Options{Help: "allow the use of the existing env when expanding commands"})
	envPathExpandFlag := expandCmd.String("e", "env-file", &argparse.Options{Help: "path to a dotenv file"})
	processTypeExpandFlag := expandCmd.String("p", "process-type", &argparse.Options{Help: "name of process to retrieve"})

	listCmd := parser.NewCommand("list", "list all process types in a procfile")

	checkCmd := parser.NewCommand("check", "check that the specified procfile is valid")

	showCmd := parser.NewCommand("show", "show the command for a specific process type")
	allowGetenvShowFlag := showCmd.Flag("a", "allow-getenv", &argparse.Options{Help: "allow the use of the existing env when expanding commands"})
	envPathShowFlag := showCmd.String("e", "env-file", &argparse.Options{Help: "path to a dotenv file"})
	processTypeShowFlag := showCmd.String("p", "process-type", &argparse.Options{Help: "name of process to retrieve", Required: true})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", parser.Usage(err))
		os.Exit(1)
		return
	}

	if *versionFlag {
		fmt.Printf("procfile-util %v\n", Version)
		os.Exit(0)
		return
	}

	Loglevel = *loglevelFlag

	entries, err := parseProcfile(*procfileFlag, *delimiterFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		return
	}

	success := false
	if checkCmd.Happened() {
		success = checkCommand(entries)
	} else if existsCmd.Happened() {
		success = existsCommand(entries, *processTypeExistsFlag)
	} else if expandCmd.Happened() {
		success = expandCommand(entries, *envPathExpandFlag, *allowGetenvExpandFlag, *processTypeExpandFlag, *defaultPortFlag)
	} else if listCmd.Happened() {
		success = listCommand(entries)
	} else if showCmd.Happened() {
		success = showCommand(entries, *envPathShowFlag, *allowGetenvShowFlag, *processTypeShowFlag, *defaultPortFlag)
	} else {
		fmt.Print(parser.Usage(err))
	}

	if !success {
		os.Exit(1)
	}
}
