package main

import (
	"github.com/bugfixes/go-bugfixes"
	"github.com/zserge/webview"
	"html/template"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var events chan string
var dir string

var windowWidth = 1024
var windowHeight = 768

func init() {
	events = make(chan string, 1000)

	var err error
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		bugfixes.Error("FilePath", err)
	}
}

func main() {
	prefixChannel := make(chan string)
	go app(prefixChannel)

	prefix := <- prefixChannel

	settings := webview.Settings{
		Debug: true,
		Height: windowHeight,
		Width: windowWidth,
		Title: "Static Site Generator",
		URL: prefix + "/dist/index.html",
		Resizable: true,
	}
	view := webview.New(settings)
	view.Run()
}

func app(prefixChannel chan string) {
	mux := http.NewServeMux()
	mux.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir(dir + "/dist"))))
	mux.Handle("/", http.FileServer(http.Dir(dir + "/dist")))
	mux.HandleFunc("/tester", start)
	mux.HandleFunc("/exit", kill)

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		bugfixes.Error("Listener", err)
	}

	portAddress := listener.Addr().String()
	prefixChannel <- "http://" + portAddress
	listener.Close()

	server := &http.Server{
		Addr: portAddress,
		Handler: mux,
	}
	server.ListenAndServe()
}

func start(response http.ResponseWriter, request *http.Request) {
	temp := template.New("Text")
	temp.Execute(response, "Tester")
}

func kill(response http.ResponseWriter, request *http.Request) {
	os.Exit(0)
}
