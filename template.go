package accouchement

import (
	"html/template"
  "net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		t.Execute(w, p)
	}
}
