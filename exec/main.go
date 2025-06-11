package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	var sizeLog string
	var size string
	var used string
	var avail string
	var use string

	if stdout, _, err := sh("journalctl --disk-usage"); err == nil {
		columns := strings.Split(squeezeSpace(stdout), " ")
		for index, column := range columns {
			if column == "up" {
				sizeLog = columns[index+1]

				break
			}
		}
	}
	fmt.Println(sizeLog)

	if stdout, _, err := sh("df -h"); err == nil {
		rows := strings.Split(stdout, "\n")
		for _, row := range rows {
			columns := strings.Split(squeezeSpace(row), " ")
			if len(columns) == 0 {
				continue
			}

			if columns[len(columns)-1] == "/" {
				fmt.Println(columns)
				size = columns[1]
				used = columns[2]
				avail = columns[3]
				use = columns[4]

				break
			}
		}
	}

	fmt.Println(size, used, avail, use)
}

func sh(command string) (stdout string, stderr string, err error) {
	var bufferStdout, bufferStderr bytes.Buffer

	parts := strings.Split(command, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = &bufferStdout
	cmd.Stderr = &bufferStderr

	err = cmd.Run()
	stdout, stderr = string(bufferStdout.Bytes()), string(bufferStderr.Bytes())

	return stdout, stderr, err
}

func squeezeSpace(str string) string {
	reg := regexp.MustCompile("\\s+")

	return reg.ReplaceAllString(str, " ")
}
