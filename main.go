package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please specify the PID: " + os.Args[0] + " PID")
	}

	pid := os.Args[1]
	cmd := exec.Command("lsof", "-a", "-Fn", "-p", pid, "-R", "/")

	fmt.Fprintf(os.Stderr, "About to run command: %s\n", cmd)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "Result: \n\n\n")

	s := bufio.NewScanner(&out)
	for s.Scan() {
		line := s.Text()
		if !strings.HasPrefix(line, "n") {
			continue
		}

		cut, _ := strings.CutPrefix(line, "n")

		fmt.Println(cut)
	}
}
