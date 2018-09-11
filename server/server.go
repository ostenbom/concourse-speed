package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func HandleHome(templateFolder string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homeTemplate, err := ioutil.ReadFile(filepath.Join(templateFolder, "index.html"))
		if err != nil {
			fmt.Fprintf(w, "<html><p>Error: %s</p></html>", err)
		}
		io.WriteString(w, string(homeTemplate))
	}
}
