package http

import "net/http"

func ServeDocs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "docs/index.html")
}
