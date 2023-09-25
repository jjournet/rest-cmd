package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jjournet/rest-go/cmds"

	"github.com/go-chi/chi/v5"
)

var port = flag.Int("port", 3333, "listening port")

func main() {
	flag.Parse()
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, router *http.Request) {
		w.Write([]byte("ok"))
	})

	// read json file into buffer
	jsonFile, err := os.Open("cmds.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	byteValue, _ := io.ReadAll(jsonFile)
	defer jsonFile.Close()
	// unmarshal json file cmds.json

	err2 := json.Unmarshal(byteValue, &cmds.Mycmds)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	for path, cmd := range cmds.Mycmds {
		router.MethodFunc(cmd.Verb, "/"+path, cmds.Command_handler)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
