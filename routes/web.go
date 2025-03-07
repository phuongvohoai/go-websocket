package routes

import "net/http"

func serveHomepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

func MapWebRoutes() {
	http.HandleFunc("/", serveHomepage)
}
