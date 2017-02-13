package accouchement

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGenerateToggleUrlWithProperFormat(t *testing.T) {
	t.Log("Generate an url with the proper format")

	url := GenerateToggleUrl()

	if !strings.HasPrefix(url, "/") {
		t.Error("%s should starts with a `/`")
	}

	if strings.HasSuffix(url, "/") {
		t.Error("%s should not end with a `/`")
	}

	if utf8.RuneCountInString(url) != 37 {
		t.Error("%s should not end with a `/`")
	}
}

func TestUniqGenerateToggleUrl(t *testing.T) {
	t.Log("Generate a unique url")

	first := GenerateToggleUrl()

	if second := GenerateToggleUrl(); first == second {
		t.Error("%s and %s are identical", first, second)
	}
}

func TestGetToggleUrl(t *testing.T) {

}
