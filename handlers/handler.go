package handlers

import (
	"fmt"
	"net/http"
)

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is a stable(ish) service!")
}
