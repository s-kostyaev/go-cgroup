package lxc

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Container struct {
	Name  string
	State string
	IP    string
}

func GetContainers() (containers []Container) {
	cmd := exec.Command("heaver", "-L")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	rawcontainers := strings.Split(strings.Trim(out.String(), "\n"), "\n")
	for _, rawcontainer := range rawcontainers {
		var con Container
		cont_str := strings.Split(strings.Trim(rawcontainer, " "), " ")
		con.Name = strings.Trim(cont_str[0], ":")
		con.State = strings.Trim(cont_str[1], ",")
		con.IP = strings.Split(cont_str[3], "/")[0]
		containers = append(containers, con)
	}
	return
}

func (container Container) String() string {
	return fmt.Sprintf("%+v", container)
}
