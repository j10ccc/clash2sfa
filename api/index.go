package handler

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/xmdhs/clash2sfa/provide"
)

var app = provide.NewApp(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
