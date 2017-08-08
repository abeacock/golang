package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	name, _ := os.Hostname()
	port := "80"

	url := "http://" + name + ":" + port + "/"
	fmt.Println(url)
	clipboard.WriteAll(string(url))

	http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))
}
