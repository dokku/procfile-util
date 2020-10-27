package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/josegonzalez/go-procfile-util/export"
	"github.com/josegonzalez/go-procfile-util/procfile"
	"github.com/joho/godotenv"
)

func ExportCommand(entries []procfile.ProcfileEntry, app string, description string, envPath string, format string, formation string, group string, home string, limitCoredump string, limitCputime string, limitData string, limitFileSize string, limitLockedMemory string, limitOpenFiles string, limitUserProcesses string, limitPhysicalMemory string, limitStackSize string, location string, logPath string, nice string, prestart string, workingDirectoryPath string, runPath string, timeout int, processUser string, defaultPort int) bool {
	if format == "" {
		fmt.Fprintf(os.Stderr, "no format specified\n")
		return false
	}
	if location == "" {
		fmt.Fprintf(os.Stderr, "no output location specified\n")
		return false
	}

	formats := map[string]export.ExportFunc{
		"launchd":      export.ExportLaunchd,
		"runit":        export.ExportRunit,
		"systemd":      export.ExportSystemd,
		"systemd-user": export.ExportSystemdUser,
		"sysv":         export.ExportSysv,
		"upstart":      export.ExportUpstart,
	}

	if _, ok := formats[format]; !ok {
		fmt.Fprintf(os.Stderr, "invalid format type: %s\n", format)
		return false
	}

	formations, err := procfile.ParseFormation(formation)
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
