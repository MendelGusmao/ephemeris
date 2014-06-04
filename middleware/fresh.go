package middleware

import (
	"github.com/pilu/fresh/runner/runnerutils"
	"net/http"
)

func Fresh(w http.ResponseWriter, r *http.Request) {
	if runnerutils.HasErrors() {
		runnerutils.RenderError(w)
	}
}
