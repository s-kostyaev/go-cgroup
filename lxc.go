package lxc

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Container struct {
	Name  string
	State string
	IP    string
}

func GetContainers() ([]Container, error) {
	result := []Container{}
	cmd := exec.Command("heaver", "-L")
	cmd.Stdout = &bytes.Buffer{}
	err := cmd.Run()
	if err != nil {
		return result, err
	}
	rawcontainers := strings.Split(strings.Trim(
		cmd.Stdout.(*bytes.Buffer).String(), "\n"), "\n")
	for _, rawcontainer := range rawcontainers {
		con := Container{}
		cont_str := strings.Fields(rawcontainer)
		con.Name = strings.Trim(cont_str[0], ":")
		con.State = strings.Trim(cont_str[1], ",")
		con.IP = strings.Split(cont_str[3], "/")[0]
		result = append(result, con)
	}
	return result, nil
}

func (container Container) String() string {
	return fmt.Sprintf("%+v", container)
}
