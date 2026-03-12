package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"myapp/assets"
	"myapp/ui/pages"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	datastar "github.com/starfederation/datastar-go/datastar"
	"github.com/templui/templui/utils"
)

func main() {
	initDotEnv()

	mux := http.NewServeMux()
	setupAssetsRoutes(mux)
	mux.Handle("GET /", templ.Handler(pages.Landing()))
	mux.Handle("GET /repro", templ.Handler(pages.Repro()))
	mux.HandleFunc("GET /repro/fragment", handleReproFragment)

	fmt.Println("Server is running on http://localhost:8090")
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		panic(err)
	}
}

func initDotEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func setupAssetsRoutes(mux *http.ServeMux) {
	isDevelopment := os.Getenv("GO_ENV") != "production"

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))

	utils.SetupScriptRoutes(mux, isDevelopment)
}

func handleReproFragment(w http.ResponseWriter, r *http.Request) {
	iteration := parseIteration(r.URL.Query().Get("iteration"))
	sse := newSSE(w, r)
	if err := sse.PatchElementTempl(pages.ReproFragment(iteration)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func parseIteration(raw string) int {
	if raw == "" {
		return 1
	}

	n, err := strconv.Atoi(raw)
	if err != nil || n < 1 {
		return 1
	}

	return n
}

func newSSE(w http.ResponseWriter, r *http.Request) *datastar.ServerSentEventGenerator {
	if rc := http.NewResponseController(w); rc != nil {
		_ = rc.SetWriteDeadline(time.Time{})
	}

	w.Header().Set("Content-Type", "text/event-stream")
	if r.ProtoMajor == 1 {
		w.Header().Set("Connection", "keep-alive")
	}
	w.WriteHeader(http.StatusOK)

	return datastar.NewSSE(w, r)
}
