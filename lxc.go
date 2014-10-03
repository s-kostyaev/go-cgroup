package lxc

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	memoryPathPrefix    = "/sys/fs/cgroup/memory/lxc"
	memoryLimitFilename = "memory.limit_in_bytes"
	memoryUsageFilename = "memory.usage_in_bytes"
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
		contStr := strings.Fields(rawcontainer)
		con.Name = strings.Trim(contStr[0], ":")
		con.State = strings.Trim(contStr[1], ",")
		con.IP = strings.Split(contStr[3], "/")[0]
		result = append(result, con)
	}
	return result, nil
}

func (container Container) String() string {
	return fmt.Sprintf("%+v", container)
}

func GetParam(cgroup, containerName, paramName string) (string, error) {
	switch cgroup {
	case "memory":
		return memoryGetParam(containerName, paramName)
	default:
		return "", errors.New(fmt.Sprintf("Undefined cgroup %s",
			cgroup))
	}
}

func GetParamInt(cgroup, containerName, paramName string) (int, error) {
	res, err := GetParam(cgroup, containerName, paramName)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func memoryGetParam(containerName, paramName string) (string, error) {
	file := ""
	switch paramName {
	case "limit":
		file = memoryLimitFilename
	case "usage":
		file = memoryUsageFilename
	default:
		return "", errors.New("Undefined parameter name")
	}
	paramPath := filepath.Join(memoryPathPrefix, containerName, file)
	if _, err := os.Stat(paramPath); err != nil {
		paramPath = filepath.Join(memoryPathPrefix,
			fmt.Sprintf("%s-1", containerName), file)
	}
	content, err := ioutil.ReadFile(paramPath)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(content), "\n"), nil
}
