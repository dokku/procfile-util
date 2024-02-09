package commands

import (
	"io/ioutil"
	"os"
	"procfile-util/procfile"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const portEnvVar = "PORT"

func parseProcfile(path string, delimiter string, strict bool) ([]procfile.ProcfileEntry, error) {
	var entries []procfile.ProcfileEntry
	text, err := procfile.GetProcfileContent(path)
	if err != nil {
		return entries, err
	}

	return procfile.ParseProcfile(text, delimiter, strict)
}

func expandEnv(e procfile.ProcfileEntry, envPath string, allowEnv bool, defaultPort int) (string, error) {
	baseExpandFunc := func(key string) string {
		if key == "PS" {
			return os.Getenv("PS")
		}
		if key == portEnvVar {
			return strconv.Itoa(defaultPort)
		}
		return ""
	}

	expandFunc := func(key string) string {
		return baseExpandFunc(key)
	}

	if allowEnv {
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
