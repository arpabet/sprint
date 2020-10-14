/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package util

import (
	"github.com/arpabet/sprint/pkg/app"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
	"path/filepath"
	"strconv"
	"runtime"
)

func StartServer(masterKey string) error {

	args := []string { "run", "-" + app.DAEMON_FLAG_KEY }
	exeArgs := app.GetArgs()
	if exeArgs != nil {
		args = append(args, exeArgs...)
	}

	executable, _ := os.Executable()
	dir := filepath.Dir(executable)
	fileName := filepath.Base(executable)

	serverFileName := app.ExecutableRotate(fileName)
	serverFilePath := filepath.Join(dir, serverFileName)

	distrFilePath := filepath.Join(dir, app.ExecutableName)
	distrFilePathAlt := distrFilePath +  "_" + runtime.GOOS
	if _, err := os.Stat(distrFilePath); err == nil {
		if err := CopyFile(distrFilePath, serverFilePath, app.ExeFilePerm); err != nil {
			return err
		}
		args = append(args, "-" + app.DISTR_FLAG_KEY, distrFilePath)
	} else if _, err := os.Stat(distrFilePathAlt); err == nil {
		if err := CopyFile(distrFilePathAlt, serverFilePath, app.ExeFilePerm); err != nil {
			return err
		}
		args = append(args, "-" + app.DISTR_FLAG_KEY, distrFilePathAlt)
	} else {
		serverFilePath = executable
	}

	cmd := exec.Command(serverFilePath, args...)
	fmt.Printf("Run cmd: %v\n", cmd)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer stdin.Close()

	io.WriteString(stdin, masterKey + "\n")

	if err := cmd.Start(); err != nil {
		return err
	}

	fmt.Println("Daemon process ID is : ", cmd.Process.Pid)

	executablePid := app.ExecutablePid()

	content := fmt.Sprintf("%d", cmd.Process.Pid)
	ioutil.WriteFile(executablePid, []byte(content), 0660)

	return nil
}


func KillServer() error {

	executablePid := app.ExecutablePid()

	blob, err := ioutil.ReadFile(executablePid)
	if err != nil {
		return err
	}

	pid := string(blob)

	if _, err := strconv.Atoi(pid); err != nil {
		return errors.Errorf("Invalid pid %s, %v", pid, err)
	}

	cmd := exec.Command("kill", "-2", pid)
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := os.Remove(executablePid); err != nil {
		return errors.Errorf("Can not remove file %s, %v", executablePid, err)
	}

	return cmd.Wait()

}

