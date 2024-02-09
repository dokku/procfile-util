package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"procfile-util/export"
	"procfile-util/procfile"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/josegonzalez/cli-skeleton/command"
	"github.com/posener/complete"
	flag "github.com/spf13/pflag"
)

type ExportCommand struct {
	command.Meta
	GlobalFlagCommand

	app                  string
	description          string
	envPath              string
	format               string
	formation            string
	group                string
	home                 string
	limitCoredump        string
	limitCputime         string
	limitData            string
	limitFileSize        string
	limitLockedMemory    string
	limitOpenFiles       string
	limitUserProcesses   string
	limitPhysicalMemory  string
	limitStackSize       string
	location             string
	logPath              string
	nice                 string
	prestart             string
	workingDirectoryPath string
	run                  string
	timeout              int
	user                 string
}

func (c *ExportCommand) Name() string {
	return "export"
}

func (c *ExportCommand) Synopsis() string {
	return "Eats one or more lollipops"
}

func (c *ExportCommand) Help() string {
	return command.CommandHelp(c)
}

func (c *ExportCommand) Examples() map[string]string {
	appName := os.Getenv("CLI_APP_NAME")
	return map[string]string{
		"Command": fmt.Sprintf("%s %s", appName, c.Name()),
	}
}

func (c *ExportCommand) Arguments() []command.Argument {
	args := []command.Argument{}
	return args
}

func (c *ExportCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ExportCommand) ParsedArguments(args []string) (map[string]command.Argument, error) {
	return command.ParseArguments(args, c.Arguments())
}

func (c *ExportCommand) FlagSet() *flag.FlagSet {
	f := c.Meta.FlagSet(c.Name(), command.FlagSetClient)
	f.StringVar(&c.app, "app", "app", "name of app")
	f.StringVar(&c.description, "description", "", "process description")
	f.StringVarP(&c.envPath, "env-file", "e", "", "path to a dotenv file")
	f.StringVar(&c.format, "format", "systemd", "format to export")
	f.StringVar(&c.formation, "formation", "all=1", "specify what processes will run and how many")
	f.StringVar(&c.group, "group", "", "group to run the command as")
	f.StringVar(&c.home, "home", "", "home directory for program")
	f.StringVar(&c.limitCoredump, "limit-coredump", "", "Largest size (in blocks) of a core file that can be created. (setrlimit RLIMIT_CORE)")
	f.StringVar(&c.limitCputime, "limit-cputime", "", "Maximum amount of cpu time (in seconds) a program may use. (setrlimit RLIMIT_CPU)")
	f.StringVar(&c.limitData, "limit-data", "", "Maximum data segment size (setrlimit RLIMIT_DATA)")
	f.StringVar(&c.limitFileSize, "limit-file-size", "", "Maximum size (in blocks) of a file receiving writes (setrlimit RLIMIT_FSIZE)")
	f.StringVar(&c.limitLockedMemory, "limit-locked-memory", "", "Maximum amount of memory (in bytes) lockable with mlock(2) (setrlimit RLIMIT_MEMLOCK)")
	f.StringVar(&c.limitOpenFiles, "limit-open-files", "", "maximum number of open files, sockets, etc. (setrlimit RLIMIT_NOFILE)")
	f.StringVar(&c.limitUserProcesses, "limit-user-processes", "", "Maximum number of running processes (or threads!) for this user id. Not recommended because this setting applies to the user, not the process group. (setrlimit RLIMIT_NPROC)")
	f.StringVar(&c.limitPhysicalMemory, "limit-physical-memory", "", "Maximum resident set size (in bytes); the amount of physical memory used by a process. (setrlimit RLIMIT_RSS)")
	f.StringVar(&c.limitStackSize, "limit-stack-size", "", "Maximum size (in bytes) of a stack segment (setrlimit RLIMIT_STACK)")
	f.StringVar(&c.location, "location", "", "location to output to")
	f.StringVar(&c.logPath, "log-path", "/var/log", "log directory")
	f.StringVar(&c.nice, "nice", "", "nice level to add to this program before running")
	f.StringVar(&c.prestart, "prestart", "", "A command to execute before starting and restarting. A failure of this command will cause the start/restart to abort. This is useful for health checks, config tests, or similar operations.")
	f.StringVar(&c.workingDirectoryPath, "working-directory-path", "/", "working directory path for app")
	f.StringVar(&c.run, "run", "", "run pid file directory, defaults to /var/run/<app>")
	f.IntVar(&c.timeout, "timeout", 5, "amount of time (in seconds) processes have to shutdown gracefully before receiving a SIGKILL")
	f.StringVar(&c.user, "user", "", "user to run the command as")
	c.GlobalFlags(f)
	return f
}

func (c *ExportCommand) AutocompleteFlags() complete.Flags {
	return command.MergeAutocompleteFlags(
		c.Meta.AutocompleteFlags(command.FlagSetClient),
		c.AutocompleteGlobalFlags(),
		complete.Flags{
			"--count":                  complete.PredictAnything,
			"--app":                    complete.PredictAnything,
			"--description":            complete.PredictAnything,
			"--env-file":               complete.PredictFiles("*"),
			"--format":                 complete.PredictAnything,
			"--formation":              complete.PredictAnything,
			"--group":                  complete.PredictAnything,
			"--home":                   complete.PredictAnything,
			"--limit-coredump":         complete.PredictAnything,
			"--limit-cputime":          complete.PredictAnything,
			"--limit-data":             complete.PredictAnything,
			"--limit-file-size":        complete.PredictAnything,
			"--limit-locked-memory":    complete.PredictAnything,
			"--limit-open-files":       complete.PredictAnything,
			"--limit-user-processes":   complete.PredictAnything,
			"--limit-physical-memory":  complete.PredictAnything,
			"--limit-stack-size":       complete.PredictAnything,
			"--location":               complete.PredictAnything,
			"--log-path":               complete.PredictAnything,
			"--nice":                   complete.PredictAnything,
			"--prestart":               complete.PredictAnything,
			"--working-directory-path": complete.PredictAnything,
			"--run":                    complete.PredictAnything,
			"--timeout":                complete.PredictAnything,
			"--user":                   complete.PredictAnything,
		},
	)
}

func (c *ExportCommand) Run(args []string) int {
	flags := c.FlagSet()
	flags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := flags.Parse(args); err != nil {
		c.Ui.Error(err.Error())
		c.Ui.Error(command.CommandErrorText(c))
		return 1
	}

	_, err := c.ParsedArguments(flags.Args())
	if err != nil {
		c.Ui.Error(err.Error())
		c.Ui.Error(command.CommandErrorText(c))
		return 1
	}

	entries, err := parseProcfile(c.procfile, c.delimiter, c.strict)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	if len(entries) == 0 {
		c.Ui.Error("No processes defined")
		return 1
	}

	if c.format == "" {
		fmt.Fprintf(os.Stderr, "no format specified\n")
		return 1
	}
	if c.location == "" {
		fmt.Fprintf(os.Stderr, "no output location specified\n")
		return 1
	}

	formats := map[string]export.ExportFunc{
		"launchd":      export.ExportLaunchd,
		"runit":        export.ExportRunit,
		"systemd":      export.ExportSystemd,
		"systemd-user": export.ExportSystemdUser,
		"sysv":         export.ExportSysv,
		"upstart":      export.ExportUpstart,
	}

	if _, ok := formats[c.format]; !ok {
		c.Ui.Error(fmt.Sprintf("Invalid format type: %s", c.format))
		return 1
	}

	formations, err := procfile.ParseFormation(c.formation)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	if c.user == "" {
		c.user = c.app
	}

	if c.group == "" {
		c.group = c.app
	}

	u, err := user.Current()
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	if c.home == "" {
		c.home = "/home/" + u.Username
	}

	env := make(map[string]string)
	if c.envPath != "" {
		b, err := ioutil.ReadFile(c.envPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading env file: %s\n", err)
			return 1
		}

		content := string(b)
		env, err = godotenv.Unmarshal(content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing env file: %s\n", err)
			return 1
		}
	}

	vars := make(map[string]interface{})
	vars["app"] = c.app
	vars["description"] = c.description
	vars["env"] = env
	vars["group"] = c.group
	vars["home"] = c.home
	vars["log"] = c.logPath
	vars["location"] = c.location
	vars["limit_coredump"] = c.limitCoredump
	vars["limit_cputime"] = c.limitCputime
	vars["limit_data"] = c.limitData
	vars["limit_file_size"] = c.limitFileSize
	vars["limit_locked_memory"] = c.limitLockedMemory
	vars["limit_open_files"] = c.limitOpenFiles
	vars["limit_user_processes"] = c.limitUserProcesses
	vars["limit_physical_memory"] = c.limitPhysicalMemory
	vars["limit_stack_size"] = c.limitStackSize
	vars["nice"] = c.nice
	vars["prestart"] = c.prestart
	vars["working_directory"] = c.workingDirectoryPath
	vars["timeout"] = strconv.Itoa(c.timeout)
	vars["ulimit_shell"] = ulimitShell(c.limitCoredump, c.limitCputime, c.limitData, c.limitFileSize, c.limitLockedMemory, c.limitOpenFiles, c.limitUserProcesses, c.limitPhysicalMemory, c.limitStackSize)
	vars["user"] = c.user

	if fn, ok := formats[c.format]; ok {
		if !fn(c.app, entries, formations, c.location, c.defaultPort, vars, c.Ui) {
			return 1
		}
	}

	return 0
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
