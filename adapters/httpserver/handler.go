package httpserver

import (
	"fmt"
	"net/http"

	"github.com/Marcelixoo/go_specs_greet/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}
