package handler

import (
	"io"
	"net/http"

	"github.com/notnmeyer/dngl/internal/envhelper"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	env := r.Context().Value("env").(*envhelper.Env)
	io.WriteString(w, "OK "+env.DNGL_TOKEN)
}
