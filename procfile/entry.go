package procfile

import (
	"strings"

	"gopkg.in/alessio/shellescape.v1"
)

// ProcfileEntry is a struct containing a process type and the corresponding command
type ProcfileEntry struct {
	Name    string
	Command string
}

func (p *ProcfileEntry) CommandList() []string {
	return strings.Fields(p.Command)
}

func (p *ProcfileEntry) Program() string {
	return strings.Fields(p.Command)[0]
}

func (p *ProcfileEntry) Args() string {
	return strings.Join(strings.Fields(p.Command)[1:], " ")
}

func (p *ProcfileEntry) ArgsEscaped() string {
	return shellescape.Quote(p.Args())
}

// FormationEntry is a struct containing a process type and the corresponding count
type FormationEntry struct {
	Name  string
	Count int
}
