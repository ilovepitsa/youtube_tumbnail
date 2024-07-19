package app

import (
	"fmt"
	"net/http"
)

func Run(configPath string) error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		return err
	}

	return nil
}
