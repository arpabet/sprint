/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintserver

import (
	"fmt"
	"go.arpabet.com/glue"
	"go.arpabet.com/sprint/sprint"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type implRedirectHttpsPage struct {
	Properties glue.Properties `inject`

	beanName       string
	redirectAddr   string
	redirectSuffix string
}

func RedirectHttpsPage(beanName string) sprint.Router {
	return &implRedirectHttpsPage{
		beanName: beanName,
	}
}

func (t *implRedirectHttpsPage) BeanName() string {
	return t.beanName
}

func (t *implRedirectHttpsPage) PostConstruct() (err error) {
	t.redirectAddr = t.Properties.GetString(fmt.Sprintf("%s.%s", t.beanName, "redirect-address"), "")
	if t.redirectAddr == "" {
		return errors.Errorf("property '%s.redirect-address' is not found in context", t.beanName)
	}

	i := strings.IndexByte(t.redirectAddr, ':')
	if i != -1 {
		t.redirectSuffix = t.redirectAddr[i:]
	} else {
		t.redirectSuffix = ""
	}

	return
}

func (t *implRedirectHttpsPage) Pattern() string {
	return "/"
}

func (t *implRedirectHttpsPage) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			http.Error(w, fmt.Sprintf("%v", r), http.StatusInternalServerError)
		}
	}()

	hostname := strings.Split(req.Host, ":")[0]
	url := fmt.Sprintf("https://%s%s%s", hostname, t.redirectSuffix, req.RequestURI)
	http.Redirect(w, req, url, http.StatusMovedPermanently)
}
