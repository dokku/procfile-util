package procfile

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/andrew-d/go-termutil"
)

// GetProcfileContent returns the content at a path as a string
func GetProcfileContent(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		if !termutil.Isatty(os.Stdin.Fd()) {
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

func OutputProcfile(path string, writePath string, delimiter string, stdout bool, entries []ProcfileEntry) error {
	if writePath != "" && stdout {
		return errors.New("cannot specify both --stdout and --write-path flags")
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	if stdout {
		for _, entry := range entries {
			fmt.Printf("%v%v %v\n", entry.Name, delimiter, entry.Command)
		}
		return nil
	}

	if writePath != "" {
		path = writePath
	}

	if err := writeProcfile(path, delimiter, entries); err != nil {
		return fmt.Errorf("error writing procfile: %s", err)
	}

	return nil
}

func writeProcfile(path string, delimiter string, entries []ProcfileEntry) error {
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
