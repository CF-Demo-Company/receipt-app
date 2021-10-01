package server

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
)

type Server struct {
	storage *Storage
	content embed.FS
}

func NewServer(s *Storage, content embed.FS) *Server {
	return &Server{storage: s, content: content}
}

// Run a HTTP server with all routes configured
func (s *Server) Run() error {
	r := chi.NewRouter()

	r.Get("/healthz", s.Healthcheck)

	r.Get("/", s.IndexHandler)

	// server static CSS
	authStaticHandler := StaticAssetsHandler(s.content, "/static/dist", "static/dist")

	r.Get("/static/*", authStaticHandler.ServeHTTP)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	return server.ListenAndServe()
}

type IndexTemplateData struct {
	Receipts []Receipt
}

// Healthcheck is a very basic healthcheck endpoint which always returns
// a HTTP 200 OK status.
func (s *Server) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Healthy"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// IndexHandler renders the index HTML page for the application
// which is templated HTML showing a list of receipts
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(s.content, "template/index.html")
	if err != nil {
		writeError(w, err)
		return
	}
	ctx := r.Context()

	receipts, err := s.storage.List(ctx)
	if err != nil {
		writeError(w, err)
		return
	}

	data := IndexTemplateData{
		Receipts: receipts,
	}

	err = t.Execute(w, data)
	if err != nil {
		writeError(w, err)
		return
	}
}

// writeError logs an error and returns a HTTP 500 error to the client
func writeError(w http.ResponseWriter, e error) {
	log.Default().Printf("error: %s\n", e)
	w.WriteHeader(http.StatusInternalServerError)
}
