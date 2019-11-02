package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	fileManip "github.com/jsonvoorhees11/distributed-image/file-manip"
	strHelper "github.com/jsonvoorhees11/distributed-image/string-helpers"
)

//Command consts
const splitImg = "splitImg"
const mergeImg = "mergeImg"
const exit = "exit"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCmd(cmdStr)
	}
}

func runCmd(cmdStr string) error {
	commandStr := strings.TrimSuffix(cmdStr, "\n")
	cmdArr := strings.Fields(commandStr)
	switch strings.ToLower(cmdArr[0]) {
	case strings.ToLower(exit):
		os.Exit(0)
	case strings.ToLower(splitImg):
		filePath, err := strHelper.ParseStrToFlag(cmdArr[1])
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, splitImage(filePath.Value))
		return nil
	}

	cmd := exec.Command(cmdArr[0], cmdArr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func splitImage(imgPath string) string {
	fileManip.SplitFile(imgPath)
	return "OK"
}
