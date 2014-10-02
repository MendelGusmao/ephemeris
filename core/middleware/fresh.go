package middleware

import (
	"net/http"

	"github.com/pilu/fresh/runner/runnerutils"
)

func Fresh(w http.ResponseWriter, r *http.Request) {
	if runnerutils.HasErrors() {
		runnerutils.RenderError(w)
	}
}
