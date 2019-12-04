package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/russross/blackfriday"
)

var data = map[string]string{
	"author": "Andrew Beacock",
}

func aboutHtmlFunc(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("about").Parse(aboutHtml)
	t.Execute(w, data)
}

func aboutMarkdownFunc(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("about").Parse(aboutMarkdown)
	var buffer bytes.Buffer
	t.Execute(&buffer, data)
	w.Write(blackfriday.MarkdownBasic(buffer.Bytes()))
}

func main() {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	name := localAddr.IP.String()

	port := "8080"

	url := "http://" + name + ":" + port + "/"
	fmt.Println(url)

	http.HandleFunc("/about", aboutMarkdownFunc)
	http.HandleFunc("/aboutHtml", aboutHtmlFunc)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	http.ListenAndServe(":"+port, nil)
}

const aboutHtml string = `
<h1>About</h1>
<p>The author of serve.go is {{.author}}</p>
`

const aboutMarkdown string = `
# About
The author of serve.go is {{.author}}
`
