package processes

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const op = "/internal/lib/processes: "

type Process struct {
	Pid     int
	Command string
	Active  bool
	Out     string
	params  []string
	cmd     *exec.Cmd
}

type Processor interface {
	Spawn(params ...string) error
	Kill() error
}

func New(params ...string) *Process {
	return &Process{
		params: params,
	}
}

func (p *Process) Spawn() error {
	// TODO add logic to set Active to false after process ends

	name := p.params[0]
	var options []string
	if len(p.params) > 1 {
		options = p.params[1:len(p.params)]
	} else {
		options = []string{}
	}
	cmd := exec.Command(name, options...)
	fmt.Println("Starting a process with cmd ", name, options)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(op, "process.Spawn(): ", err)
		return err
	}
	p.cmd = cmd
	p.Out = string(out)
	p.Command = strings.Join(p.params, " ")
	p.Pid = cmd.Process.Pid
	p.Active = true
	fmt.Println("Process ended on pid: ", p.Pid)

	return nil
}

func (p *Process) Kill() error {
	err := p.cmd.Process.Kill()
	if err != nil {
		log.Fatal(op, err)
		return err
	}
	p.cmd = nil
	p.Active = false
	return nil
}
