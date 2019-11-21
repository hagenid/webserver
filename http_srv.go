package main

import (
	"fmt"
	"net/http"
	"github.com/juju/loggo"

)

func main() {
	
	logger := loggo.GetLogger("foo.bar")
	writer, err := loggo.RemoveWriter("default")
	// err is non-nil if and only if the name isn't found.
	loggo.RegisterWriter("default", writer)
	loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":80", nil)



}
