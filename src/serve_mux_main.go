package main

import (
	"fmt"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
)

type HomeHandler struct {
}

func (hh *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", "HomeHandler")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", "PostHandler")
}

func RegisterHandler(mux *http.ServeMux, pattern string, handler http.Handler) {
	mux.Handle(pattern, handler)
}

func main() {
	mx := http.NewServeMux()
	RegisterHandler(mx, "/", &HomeHandler{})
	RegisterHandler(mx, "/post", http.HandlerFunc(PostHandler))
	err := http.ListenAndServe(":8888", mx)
	/*
			r := mux.NewRouter()
			r.HandleFunc("/mux1/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%s %v\n", "MuxRouterHandler", mux.Vars(r))
			})
			f := func(w http.ResponseWriter, r *http.Request) {
				vars := mux.Vars(r)
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "got vars %v\n", vars)
			}
			r.HandleFunc("/mux2/{key}", f)
		err := http.ListenAndServe(":8888", r)
	*/
	log.Printf("%v\n", err)
}
