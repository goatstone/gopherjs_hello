package gostart

import (
	"net/http"
 )

func init() {
	http.HandleFunc("/admin", HandleTemplate)
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	return
}