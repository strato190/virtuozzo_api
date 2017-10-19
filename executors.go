package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Prlctl(args ...string) (string, error) {
	var stdout, stderr bytes.Buffer
	var prlctlPath string
	if prlctlPath == "" {
		var err error
		prlctlPath, err = exec.LookPath("prlctl")
		if err != nil {
			return stdout.String(), err
		}
	}

	log.Printf("Executing prlctl: %#v", args)
	cmd := exec.Command(prlctlPath, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	stdoutString := strings.TrimSpace(stdout.String())
	stderrString := strings.TrimSpace(stderr.String())

	if _, ok := err.(*exec.ExitError); ok {
		err = fmt.Errorf("prlctl error: %s", stderrString)
	}

	log.Printf("stdout: %s", stdoutString)
	log.Printf("stderr: %s", stderrString)

	return stdoutString, err
}

func Vzctl(args ...string) (string, error) {
	var stdout, stderr bytes.Buffer
	var prlctlPath string
	if prlctlPath == "" {
		var err error
		prlctlPath, err = exec.LookPath("vzctl")
		if err != nil {
			return stdout.String(), err
		}
	}

	log.Printf("Executing vzctl: %#v", args)
	cmd := exec.Command(prlctlPath, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	stdoutString := strings.TrimSpace(stdout.String())
	stderrString := strings.TrimSpace(stderr.String())

	if _, ok := err.(*exec.ExitError); ok {
		err = fmt.Errorf("vzctl error: %s", stderrString)
	}

	log.Printf("stdout: %s", stdoutString)
	log.Printf("stderr: %s", stderrString)

	return stdoutString, err
}
