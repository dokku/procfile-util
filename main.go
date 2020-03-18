package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/andrew-d/go-termutil"
	"github.com/joho/godotenv"
	"gopkg.in/alessio/shellescape.v1"
)

type procfileEntry struct {
	Name    string
	Command string
}

type formationEntry struct {
	Name  string
	Count int
}

type exportFunc func(string, []procfileEntry, map[string]formationEntry, string, int, map[string]interface{}) bool

func (p *procfileEntry) commandList() []string {
	return strings.Fields(p.Command)
}

func (p *procfileEntry) program() string {
	return strings.Fields(p.Command)[0]
}

func (p *procfileEntry) args() string {
	return strings.Join(strings.Fields(p.Command)[1:], " ")
}

func (p *procfileEntry) argsEscaped() string {
	return shellescape.Quote(p.args())
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
	debugMessage(fmt.Sprintf("Attempting to read input from file: %v", path))
	f, err := os.Open(path)
	if err != nil {
		if !termutil.Isatty(os.Stdin.Fd()) {
			debugMessage("Reading input from stdin")
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		}
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

func outputProcfile(path string, writePath string, delimiter string, stdout bool, entries []procfileEntry) bool {
	if writePath != "" && stdout {
		fmt.Fprintf(os.Stderr, "cannot specify both --stdout and --write-path flags\n")
		return false
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	if stdout {
		for _, entry := range entries {
			fmt.Printf("%v%v %v\n", entry.Name, delimiter, entry.Command)
		}
		return true
	}

	if writePath != "" {
		path = writePath
	}

	if err := writeProcfile(path, delimiter, entries); err != nil {
		fmt.Fprintf(os.Stderr, "error writing procfile: %s\n", err)
		return false
	}

	return true
}

func writeProcfile(path string, delimiter string, entries []procfileEntry) error {
	debugMessage(fmt.Sprintf("Attempting to write output to file: %v", path))
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, entry := range entries {
		fmt.Fprintln(w, fmt.Sprintf("%v%v %v", entry.Name, delimiter, entry.Command))
	}
	return w.Flush()
}

func parseProcfile(path string, delimiter string) ([]procfileEntry, error) {
	var entries []procfileEntry
	reCmd, _ := regexp.Compile(`^([A-Za-z0-9_-]+)` + delimiter + `\s*(.+)$`)
	reComment, _ := regexp.Compile(`^(.*)\s#.+$`)

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

		params := reCmd.FindStringSubmatch(line)
		if len(params) != 3 {
			debugMessage(fmt.Sprintf("No matching params in line: %v", line))
			continue
		}

		name, cmd := params[1], params[2]

		if names[name] {
			return entries, fmt.Errorf("process names must be unique")
		}
		names[name] = true

		commentParams := reComment.FindStringSubmatch(cmd)
		if len(commentParams) == 2 {
			cmd = commentParams[1]
		}

		entries = append(entries, procfileEntry{name, cmd})
	}

	if scanner.Err() != nil {
		return entries, scanner.Err()
	}

	if len(entries) == 0 {
		return entries, fmt.Errorf("no entries found in Procfile")
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	return entries, nil
}

func expandEnv(e procfileEntry, envPath string, allowEnv bool, defaultPort int) (string, error) {
	baseExpandFunc := func(key string) string {
		if key == "PS" {
			return os.Getenv("PS")
		}
		if key == portEnvVar {
			return string(defaultPort)
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

func expandCommand(entries []procfileEntry, envPath string, allowGetenv bool, processType string, defaultPort int, delimiter string) bool {
	hasErrors := false
	var expandedEntries []procfileEntry
	for _, entry := range entries {
		command, err := expandEnv(entry, envPath, allowGetenv, defaultPort)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error processing command: %s\n", err)
			hasErrors = true
		}

		entry.Command = command
		expandedEntries = append(expandedEntries, entry)
	}

	if hasErrors {
		return false
	}

	for _, entry := range expandedEntries {
		if processType == "" || processType == entry.Name {
			fmt.Printf("%v%v %v\n", entry.Name, delimiter, entry.Command)
		}
	}

	return true
}

func exportCommand(entries []procfileEntry, app string, description string, envPath string, format string, formation string, group string, home string, limitCoredump string, limitCputime string, limitData string, limitFileSize string, limitLockedMemory string, limitOpenFiles string, limitUserProcesses string, limitPhysicalMemory string, limitStackSize string, location string, logPath string, nice string, prestart string, workingDirectoryPath string, runPath string, timeout int, processUser string, defaultPort int) bool {
	if format == "" {
		fmt.Fprintf(os.Stderr, "no format specified\n")
		return false
	}
	if location == "" {
		fmt.Fprintf(os.Stderr, "no output location specified\n")
		return false
	}

	formats := map[string]exportFunc{
		"launchd":      exportLaunchd,
		"runt":         exportRunit,
		"systemd":      exportSystemd,
		"systemd-user": exportSystemdUser,
		"sysv":         exportSysv,
		"upstart":      exportUpstart,
	}

	if _, ok := formats[format]; !ok {
		fmt.Fprintf(os.Stderr, "invalid format type: %s\n", format)
		return false
	}

	formations, err := parseFormation(formation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	if processUser == "" {
		processUser = app
	}

	if group == "" {
		group = app
	}

	u, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return false
	}

	if home == "" {
		home = "/home/" + u.Username
	}

	env := make(map[string]string)
	if envPath != "" {
		b, err := ioutil.ReadFile(envPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading env file: %s\n", err)
			return false
		}

		content := string(b)
		env, err = godotenv.Unmarshal(content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing env file: %s\n", err)
			return false
		}
	}

	vars := make(map[string]interface{})
	vars["app"] = app
	vars["description"] = description
	vars["env"] = env
	vars["group"] = group
	vars["home"] = home
	vars["log"] = logPath
	vars["location"] = location
	vars["limit_coredump"] = limitCoredump
	vars["limit_cputime"] = limitCputime
	vars["limit_data"] = limitData
	vars["limit_file_size"] = limitFileSize
	vars["limit_locked_memory"] = limitLockedMemory
	vars["limit_open_files"] = limitOpenFiles
	vars["limit_user_processes"] = limitUserProcesses
	vars["limit_physical_memory"] = limitPhysicalMemory
	vars["limit_stack_size"] = limitStackSize
	vars["nice"] = nice
	vars["prestart"] = prestart
	vars["working_directory"] = workingDirectoryPath
	vars["timeout"] = strconv.Itoa(timeout)
	vars["ulimit_shell"] = ulimitShell(limitCoredump, limitCputime, limitData, limitFileSize, limitLockedMemory, limitOpenFiles, limitUserProcesses, limitPhysicalMemory, limitStackSize)
	vars["user"] = processUser

	if fn, ok := formats[format]; ok {
		return fn(app, entries, formations, location, defaultPort, vars)
	}

	return false
}

func ulimitShell(limitCoredump string, limitCputime string, limitData string, limitFileSize string, limitLockedMemory string, limitOpenFiles string, limitUserProcesses string, limitPhysicalMemory string, limitStackSize string) string {
	s := []string{}
	if limitCoredump != "" {
		s = append(s, "ulimit -c ${limit_coredump}")
	}
	if limitCputime != "" {
		s = append(s, "ulimit -t ${limit_cputime}")
	}
	if limitData != "" {
		s = append(s, "ulimit -d ${limit_data}")
	}
	if limitFileSize != "" {
		s = append(s, "ulimit -f ${limit_file_size}")
	}
	if limitLockedMemory != "" {
		s = append(s, "ulimit -l ${limit_locked_memory}")
	}
	if limitOpenFiles != "" {
		s = append(s, "ulimit -n ${limit_open_files}")
	}
	if limitUserProcesses != "" {
		s = append(s, "ulimit -u ${limit_user_processes}")
	}
	if limitPhysicalMemory != "" {
		s = append(s, "ulimit -m ${limit_physical_memory}")
	}
	if limitStackSize != "" {
		s = append(s, "ulimit -s ${limit_stack_size}")
	}

	return strings.Join(s, "\n")
}

func parseFormation(formation string) (map[string]formationEntry, error) {
	entries := make(map[string]formationEntry)
	for _, formation := range strings.Split(formation, ",") {
		parts := strings.Split(formation, "=")
		if len(parts) != 2 {
			return entries, fmt.Errorf("invalid formation: %s", formation)
		}

		i, err := strconv.Atoi(parts[1])
		if err != nil {
			return entries, fmt.Errorf("invalid formation: %s", err)
		}

		entries[parts[0]] = formationEntry{
			Name:  parts[0],
			Count: i,
		}
	}

	return entries, nil
}

func deleteCommand(entries []procfileEntry, processType string, writePath string, stdout bool, delimiter string, path string) bool {
	var validEntries []procfileEntry
	for _, entry := range entries {
		if processType == entry.Name {
			continue
		}
		validEntries = append(validEntries, entry)
	}

	return outputProcfile(path, writePath, delimiter, stdout, validEntries)
}

func listCommand(entries []procfileEntry) bool {
	for _, entry := range entries {
		fmt.Printf("%v\n", entry.Name)
	}
	return true
}

func setCommand(entries []procfileEntry, processType string, command string, writePath string, stdout bool, delimiter string, path string) bool {
	var validEntries []procfileEntry
	validEntries = append(validEntries, procfileEntry{processType, command})
	for _, entry := range entries {
		if processType == entry.Name {
			continue
		}
		validEntries = append(validEntries, entry)
	}

	return outputProcfile(path, writePath, delimiter, stdout, validEntries)
}

func showCommand(entries []procfileEntry, envPath string, allowGetenv bool, processType string, defaultPort int) bool {
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
	workingDirectoryPath, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	parser := argparse.NewParser("procfile-util", "A procfile parsing tool")
	loglevelFlag := parser.Selector("l", "loglevel", []string{"info", "debug"}, &argparse.Options{Default: "info", Help: "loglevel to use"})
	procfileFlag := parser.String("P", "procfile", &argparse.Options{Default: "Procfile", Help: "path to a procfile"})
	delimiterFlag := parser.String("D", "delimiter", &argparse.Options{Default: ":", Help: "delimiter in use within procfile"})
	defaultPortFlag := parser.Int("d", "default-port", &argparse.Options{Default: 5000, Help: "default port to use"})
	versionFlag := parser.Flag("v", "version", &argparse.Options{Help: "show version"})

	checkCmd := parser.NewCommand("check", "check that the specified procfile is valid")

	deleteCmd := parser.NewCommand("delete", "delete a process type from a file")
	processTypeDeleteFlag := deleteCmd.String("p", "process-type", &argparse.Options{Help: "name of process to delete", Required: true})
	stdoutDeleteFlag := deleteCmd.Flag("s", "stdout", &argparse.Options{Help: "write output to stdout"})
	writePathDeleteFlag := deleteCmd.String("w", "write-path", &argparse.Options{Help: "path to Procfile to write to"})

	existsCmd := parser.NewCommand("exists", "check if a process type exists")
	processTypeExistsFlag := existsCmd.String("p", "process-type", &argparse.Options{Help: "name of process to retrieve"})

	expandCmd := parser.NewCommand("expand", "expands a procfile against a specific environment")
	allowGetenvExpandFlag := expandCmd.Flag("a", "allow-getenv", &argparse.Options{Help: "allow the use of the existing env when expanding commands"})
	envPathExpandFlag := expandCmd.String("e", "env-file", &argparse.Options{Help: "path to a dotenv file"})
	processTypeExpandFlag := expandCmd.String("p", "process-type", &argparse.Options{Help: "name of process to expand"})

	exportCmd := parser.NewCommand("export", "export the application to another process management format")
	appExportFlag := exportCmd.String("", "app", &argparse.Options{Default: "app", Help: "name of app"})
	descriptionExportFlag := exportCmd.String("", "description", &argparse.Options{Help: "process description"})
	envPathExportFlag := exportCmd.String("e", "env-file", &argparse.Options{Help: "path to a dotenv file"})
	formatExportFlag := exportCmd.String("", "format", &argparse.Options{Help: "format to export"})
	formationExportFlag := exportCmd.String("", "formation", &argparse.Options{Default: "all=1", Help: "specify what processes will run and how many"})
	groupExportFlag := exportCmd.String("", "group", &argparse.Options{Help: "group to run the command as"})
	homeExportFlag := exportCmd.String("", "home", &argparse.Options{Help: "home directory for program"})
	limitCoredumpExportFlag := exportCmd.String("", "limit-coredump", &argparse.Options{Help: "Largest size (in blocks) of a core file that can be created. (setrlimit RLIMIT_CORE)"})
	limitCputimeExportFlag := exportCmd.String("", "limit-cputime", &argparse.Options{Help: "Maximum amount of cpu time (in seconds) a program may use. (setrlimit RLIMIT_CPU)"})
	limitDataExportFlag := exportCmd.String("", "limit-data", &argparse.Options{Help: "Maximum data segment size (setrlimit RLIMIT_DATA)"})
	limitFileSizeExportFlag := exportCmd.String("", "limit-file-size", &argparse.Options{Help: "Maximum size (in blocks) of a file receiving writes (setrlimit RLIMIT_FSIZE)"})
	limitLockedMemoryExportFlag := exportCmd.String("", "limit-locked-memory", &argparse.Options{Help: "Maximum amount of memory (in bytes) lockable with mlock(2) (setrlimit RLIMIT_MEMLOCK)"})
	limitOpenFilesExportFlag := exportCmd.String("", "limit-open-files", &argparse.Options{Help: "maximum number of open files, sockets, etc. (setrlimit RLIMIT_NOFILE)"})
	limitUserProcessesExportFlag := exportCmd.String("", "limit-user-processes", &argparse.Options{Help: "Maximum number of running processes (or threads!) for this user id. Not recommended because this setting applies to the user, not the process group. (setrlimit RLIMIT_NPROC)"})
	limitPhysicalMemoryExportFlag := exportCmd.String("", "limit-physical-memory", &argparse.Options{Help: "Maximum resident set size (in bytes); the amount of physical memory used by a process. (setrlimit RLIMIT_RSS)"})
	limitStackSizeExportFlag := exportCmd.String("", "limit-stack-size", &argparse.Options{Help: "Maximum size (in bytes) of a stack segment (setrlimit RLIMIT_STACK)"})
	locationExportFlag := exportCmd.String("", "location", &argparse.Options{Help: "location to output to"})
	logPathExportFlag := exportCmd.String("", "log-path", &argparse.Options{Default: "/var/log", Help: "log directory"})
	niceExportFlag := exportCmd.String("", "nice", &argparse.Options{Help: "nice level to add to this program before running"})
	prestartExportFlag := exportCmd.String("", "prestart", &argparse.Options{Help: "A command to execute before starting and restarting. A failure of this command will cause the start/restart to abort. This is useful for health checks, config tests, or similar operations."})
	workingDirectoryPathExportFlag := exportCmd.String("", "working-directory-path", &argparse.Options{Default: workingDirectoryPath, Help: "working directory path for app"})
	runExportFlag := exportCmd.String("", "run", &argparse.Options{Help: "run pid file directory, defaults to /var/run/<app>"})
	timeoutExportFlag := exportCmd.Int("", "timeout", &argparse.Options{Default: 5, Help: "amount of time (in seconds) processes have to shutdown gracefully before receiving a SIGKILL"})
	userExportFlag := exportCmd.String("", "user", &argparse.Options{Help: "user to run the command as"})

	listCmd := parser.NewCommand("list", "list all process types in a procfile")

	setCmd := parser.NewCommand("set", "set the command for a process type in a file")
	processTypeSetFlag := setCmd.String("p", "process-type", &argparse.Options{Help: "name of process to set", Required: true})
	commandSetFlag := setCmd.String("c", "command", &argparse.Options{Help: "command to set", Required: true})
	stdoutSetFlag := setCmd.Flag("s", "stdout", &argparse.Options{Help: "write output to stdout"})
	writePathSetFlag := setCmd.String("w", "write-path", &argparse.Options{Help: "path to Procfile to write to"})

	showCmd := parser.NewCommand("show", "show the command for a specific process type")
	allowGetenvShowFlag := showCmd.Flag("a", "allow-getenv", &argparse.Options{Help: "allow the use of the existing env when expanding commands"})
	envPathShowFlag := showCmd.String("e", "env-file", &argparse.Options{Help: "path to a dotenv file"})
	processTypeShowFlag := showCmd.String("p", "process-type", &argparse.Options{Help: "name of process to show", Required: true})

	if err = parser.Parse(os.Args); err != nil {
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
	} else if deleteCmd.Happened() {
		success = deleteCommand(entries, *processTypeDeleteFlag, *writePathDeleteFlag, *stdoutDeleteFlag, *delimiterFlag, *procfileFlag)
	} else if existsCmd.Happened() {
		success = existsCommand(entries, *processTypeExistsFlag)
	} else if expandCmd.Happened() {
		success = expandCommand(entries, *envPathExpandFlag, *allowGetenvExpandFlag, *processTypeExpandFlag, *defaultPortFlag, *delimiterFlag)
	} else if exportCmd.Happened() {
		success = exportCommand(entries, *appExportFlag, *descriptionExportFlag, *envPathExportFlag, *formatExportFlag, *formationExportFlag, *groupExportFlag, *homeExportFlag, *limitCoredumpExportFlag, *limitCputimeExportFlag, *limitDataExportFlag, *limitFileSizeExportFlag, *limitLockedMemoryExportFlag, *limitOpenFilesExportFlag, *limitUserProcessesExportFlag, *limitPhysicalMemoryExportFlag, *limitStackSizeExportFlag, *locationExportFlag, *logPathExportFlag, *niceExportFlag, *prestartExportFlag, *workingDirectoryPathExportFlag, *runExportFlag, *timeoutExportFlag, *userExportFlag, *defaultPortFlag)
	} else if listCmd.Happened() {
		success = listCommand(entries)
	} else if setCmd.Happened() {
		success = setCommand(entries, *processTypeSetFlag, *commandSetFlag, *writePathSetFlag, *stdoutSetFlag, *delimiterFlag, *procfileFlag)
	} else if showCmd.Happened() {
		success = showCommand(entries, *envPathShowFlag, *allowGetenvShowFlag, *processTypeShowFlag, *defaultPortFlag)
	} else {
		fmt.Print(parser.Usage(err))
	}

	if !success {
		os.Exit(1)
	}
}
