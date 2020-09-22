/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package util

import (
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
	"strconv"
)

func StartServer(masterKey string) error {

	args := []string { "run", "-" + app.DAEMON_FLAG_KEY }
	args = append(args, app.GetArgs()...)

	executable, _ := os.Executable()
	cmd := exec.Command(executable, args...)
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

	content := fmt.Sprintf("%d", cmd.Process.Pid)
	ioutil.WriteFile(app.ExecutablePID, []byte(content), 0660)

	return nil
}


func KillServer() error {

	blob, err := ioutil.ReadFile(app.ExecutablePID)
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

	if err := os.Remove(app.ExecutablePID); err != nil {
		return errors.Errorf("Can not remove file %s, %v", app.ExecutablePID, err)
	}

	return cmd.Wait()

}

