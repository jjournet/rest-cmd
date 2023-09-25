package cmds

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type cmd struct {
	Command   string   `json:"command"`
	Verb      string   `json:"verb"`
	Arguments []string `json:"arguments"`
}

type Cmds map[string]cmd

var Mycmds Cmds

func Command_handler(w http.ResponseWriter, r *http.Request) {
	// get path from request without leading slash
	path := r.URL.Path[1:]
	// get command from map
	command := Mycmds[path].Command
	out := []byte{}
	err := error(nil)
	// if arguments are not empty, concatenate them
	if len(Mycmds[path].Arguments) > 0 {
		arguments := strings.Join(Mycmds[path].Arguments, " ")
		out, err = exec.Command(command, arguments).Output()
	} else {
		out, err = exec.Command(command).Output()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	// write output to response
	w.Write(out)
	return
}
