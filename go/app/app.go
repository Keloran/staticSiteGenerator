package app

import (
	"github.com/bugfixes/go-bugfixes"
	"github.com/zserge/webview"
	"main/app/items"
	"main/app/parse"
	"main/settings"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func App(prefixChannel chan string, fileDir string) {
	mux := http.NewServeMux()
	mux.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir(fileDir + "/dist"))))
	mux.Handle("/", http.FileServer(http.Dir(fileDir + "/dist")))
	mux.HandleFunc("/parse", parse.Parse)
	mux.HandleFunc("/list", items.List)

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

func HandleRPC(view webview.WebView, data string) {
	switch data {
	case "close":
		view.Terminate()
	//case "fullscreen":
	//	view.SetFullscreen(true)
	//	settings.SetFullScreen("", true)
	//case "window":
	//	view.SetFullscreen(false)
	//	settings.SetFullScreen("", false)
	//case "info":
	//	view.Dialog(webview.DialogTypeAlert, webview.DialogFlagInfo, "Info Tester", "Tester Content")
	}

	// Change Window Size
	if strings.HasPrefix(data, "windowSize:") {
		var stripped = strings.TrimPrefix(data, "windowSize:")
		var sizes = strings.Split(stripped, "|")
		var width = sizes[0]
		var height = sizes[1]

		widthInt, err := strconv.Atoi(width)
		if err != nil {
			bugfixes.Error("Width", err)
		}

		heightInt, err := strconv.Atoi(height)
		if err != nil {
			bugfixes.Error("Height", err)
		}

		settings.SetWidth("", widthInt)
		settings.SetHeight("", heightInt)
	}

	if strings.HasPrefix(data, "error:") {
		var stripped = strings.TrimPrefix(data, "error:")
		var errorMessage = strings.Split(stripped, "|")

		view.Dialog(webview.DialogTypeAlert, webview.DialogFlagError, errorMessage[0], errorMessage[1])
	}
}