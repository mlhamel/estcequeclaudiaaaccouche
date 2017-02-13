package accouchement

import (
	"encoding/json"
  "net/http"
)

func BuildResponse(status string) map[string] string {
	return map[string] string {"Status": status}
}

func DisplayStatus(w http.ResponseWriter, r *http.Request) {
	params := BuildResponse(GetStatus())
	renderTemplate(w, "templates/status.html", params)
}

func ApiStatus(w http.ResponseWriter, r *http.Request) {
	params := BuildResponse(GetStatus())

	json.NewEncoder(w).Encode(params)
}

func ToggleStatus(w http.ResponseWriter, r *http.Request) {
	status := GetStatus()
	params := BuildResponse(status)

	if status == no {
		status = EnableStatus()
	} else {
		status = DisableStatus()
	}

	json.NewEncoder(w).Encode(params)
}
