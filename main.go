package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"log/slog"
	"net/http"
)

type LoginData struct{}
type ArticleData struct {
	Id   string
	Name string
}

func HiThereHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("response received", "uri", r.RequestURI)

	slog.Info("debug", "url", r.URL, "query", r.URL.Query())

	params := r.URL.Query()
	q := params.Get("q")
	_, err := fmt.Fprintf(w, "Hi there, I love %s!", q)

	if err != nil {
		slog.Error("failed to serve", "uri", r.RequestURI, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("response received", "uri", r.RequestURI)

	filenames := []string{"views/login.htmx", "views/head.htmx"}
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		slog.Error("failed to read file", "uri", r.RequestURI, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, LoginData{})
	if err != nil {
		slog.Error("failed to send response", "uri", r.RequestURI, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("response received", "uri", r.RequestURI)

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "admin" {
		headers := w.Header()
		headers.Add("HX-Redirect", "/articles/view/")
		w.WriteHeader(200)
		_, err := fmt.Fprintf(w, "You successfully hacked in........!!!!!")
		if err != nil {
			slog.Error("failed to send response", "uri", r.RequestURI, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	_, err := fmt.Fprintf(w, "Access denied...........................")
	if err != nil {
		slog.Error("failed to send response", "uri", r.RequestURI, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func ArticlesViewHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("response received", "uri", r.RequestURI)

	filenames := []string{"views/articles/many.htmx", "views/head.htmx"}
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		slog.Error("failed to read file", "uri", r.RequestURI, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, LoginData{})
	if err != nil {
		slog.Error("failed to send response", "uri", r.RequestURI, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ArticleCreateHandler(w http.ResponseWriter, r *http.Request) {

}

func ArticlePatchHandler(w http.ResponseWriter, r *http.Request) {

}

func ArticleDeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func ArticleGetHandler(w http.ResponseWriter, r *http.Request) {

}

func ArticleViewHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("response received", "uri", r.RequestURI)

	filenames := []string{"views/articles/one.htmx", "views/head.htmx"}
	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		slog.Error("failed to read file", "uri", r.RequestURI, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, ArticleData{
		Id:   "test id 123",
		Name: "Why wombats have cubic poop",
	})
	if err != nil {
		slog.Error("failed to send response", "uri", r.RequestURI, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := chi.NewRouter()

	r.Get("/", HiThereHandler)
	r.Get("/login/", LoginViewHandler)
	r.Post("/login/", LoginHandler)

	r.Get("/articles/view/", ArticlesViewHandler)
	r.Get("/articles/{id}/view/", ArticleViewHandler)

	r.Get("/articles/", ArticlesHandler)
	r.Post("/articles/", ArticleCreateHandler)
	r.Get("/articles/{id}/", ArticleGetHandler)
	r.Patch("/articles/{id}/", ArticlePatchHandler)
	r.Delete("/articles/{id}/", ArticleDeleteHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
