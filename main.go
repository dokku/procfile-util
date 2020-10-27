package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/akamensky/argparse"
	"github.com/josegonzalez/go-procfile-util/procfile"
	"github.com/josegonzalez/go-procfile-util/commands"
)

// Version contains the procfile-util version
var Version string

func parseProcfile(path string, delimiter string, strict bool) ([]procfile.ProcfileEntry, error) {
	var entries []procfile.ProcfileEntry
	text, err := procfile.GetProcfileContent(path)
	if err != nil {
		return entries, err
	}

	return procfile.ParseProcfile(text, delimiter, strict)
}

func main() {
	parser := argparse.NewParser("procfile-util", "A procfile parsing tool")
	procfileFlag := parser.String("P", "procfile", &argparse.Options{Default: "Procfile", Help: "path to a procfile"})
	delimiterFlag := parser.String("D", "delimiter", &argparse.Options{Default: ":", Help: "delimiter in use within procfile"})
	defaultPortFlag := parser.String("d", "default-port", &argparse.Options{Default: "5000", Help: "default port to use"})
	strictFlag := parser.Flag("S", "strict", &argparse.Options{Help: "strictly parse the Procfile"})
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
	formatExportFlag := exportCmd.String("", "format", &argparse.Options{Default: "systemd", Help: "format to export"})
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
	workingDirectoryPathExportFlag := exportCmd.String("", "working-directory-path", &argparse.Options{Default: "/", Help: "working directory path for app"})
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

	if err := parser.Parse(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", parser.Usage(err))
		os.Exit(1)
		return
	}

	if *versionFlag {
		fmt.Printf("procfile-util %v\n", Version)
		os.Exit(0)
		return
	}

	entries, err := parseProcfile(*procfileFlag, *delimiterFlag, *strictFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		return
	}

	defaultPort := 5000
	if *defaultPortFlag != "" {
		i, err := strconv.Atoi(*defaultPortFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid default port value: %v\n", err)
			os.Exit(1)
			return
		}
		defaultPort = i
	}

	success := false
	if checkCmd.Happened() {
		success = commands.CheckCommand(entries)
	} else if deleteCmd.Happened() {
		success = commands.DeleteCommand(entries, *processTypeDeleteFlag, *writePathDeleteFlag, *stdoutDeleteFlag, *delimiterFlag, *procfileFlag)
	} else if existsCmd.Happened() {
		success = commands.ExistsCommand(entries, *processTypeExistsFlag)
	} else if expandCmd.Happened() {
		success = commands.ExpandCommand(entries, *envPathExpandFlag, *allowGetenvExpandFlag, *processTypeExpandFlag, defaultPort, *delimiterFlag)
	} else if exportCmd.Happened() {
		success = commands.ExportCommand(entries, *appExportFlag, *descriptionExportFlag, *envPathExportFlag, *formatExportFlag, *formationExportFlag, *groupExportFlag, *homeExportFlag, *limitCoredumpExportFlag, *limitCputimeExportFlag, *limitDataExportFlag, *limitFileSizeExportFlag, *limitLockedMemoryExportFlag, *limitOpenFilesExportFlag, *limitUserProcessesExportFlag, *limitPhysicalMemoryExportFlag, *limitStackSizeExportFlag, *locationExportFlag, *logPathExportFlag, *niceExportFlag, *prestartExportFlag, *workingDirectoryPathExportFlag, *runExportFlag, *timeoutExportFlag, *userExportFlag, defaultPort)
	} else if listCmd.Happened() {
		success = commands.ListCommand(entries)
	} else if setCmd.Happened() {
		success = commands.SetCommand(entries, *processTypeSetFlag, *commandSetFlag, *writePathSetFlag, *stdoutSetFlag, *delimiterFlag, *procfileFlag)
	} else if showCmd.Happened() {
		success = commands.ShowCommand(entries, *envPathShowFlag, *allowGetenvShowFlag, *processTypeShowFlag, defaultPort)
	} else {
		fmt.Print(parser.Usage(err))
	}

	if !success {
		os.Exit(1)
	}
}
