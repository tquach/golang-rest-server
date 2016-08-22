package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()

	msg := map[string]string{
		"a": "b",
	}

	JSON(msg, http.StatusOK, w)

	if expected, actual := http.StatusOK, w.Code; expected != actual {
		t.Fatalf("expected %d but got %d instead", expected, actual)
	}

	if expected, actual := "application/json; charset=utf-8", w.Header().Get("Content-Type"); expected != actual {
		t.Fatalf("expected %s but got %s instead", expected, actual)
	}
}
