package run

import (
	"github.com/arpabet/sprint/pkg/app"
	"html/template"
	"net/http"
)

type indexPage struct {
	tpl   *template.Template
}

func (t *indexPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if t.tpl != nil {
		t.tpl.Execute(w, r)
	}
}

func newIndexPage() (http.Handler, error) {
	tpl, err := app.ResourceHtmlTemplate(app.IndexFile)
	if err != nil {
		return nil, err
	}
	return &indexPage{
		tpl: tpl,
	}, nil
}
