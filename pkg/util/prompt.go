/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */


package util

import (
	"bufio"
	"github.com/arpabet/templateserv/pkg/app"
	"os"
	"strings"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
)

func Prompt(request string) string {
	reader := bufio.NewReader(os.Stdin)
	print(request)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func PromptPassword(request string) string {
	print(request)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		println()
		password := string(bytePassword)
		return strings.TrimSpace(password)
	} else {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		return strings.TrimSpace(text)
	}
}

func PromptMasterKey() string {
	if app.MasterKey == "" {
		return PromptPassword("Enter master key:")
	} else {
		return app.MasterKey
	}
}
