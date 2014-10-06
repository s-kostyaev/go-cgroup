package cgroup

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	pathPrefix  = "/sys/fs/cgroup"
	MemoryLimit = "memory.limit_in_bytes"
	MemoryUsage = "memory.usage_in_bytes"
)

func GetParam(cgroup, param string) (string, error) {
	paramFile := filepath.Join(pathPrefix, cgroup, param)
	if _, err := os.Stat(paramFile); err != nil {
		return "", err
	}
	content, err := ioutil.ReadFile(paramFile)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(content), "\n"), nil
}

func GetParamInt(cgroup, paramName string) (int, error) {
	res, err := GetParam(cgroup, paramName)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}
