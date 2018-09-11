package main

import (
	"github.com/bugfixes/go-bugfixes"
	"github.com/zserge/webview"
	"main/app"
	"main/settings"
	"os"
	"path/filepath"
)

var fileDir string

func init() {
	var err error
	fileDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		bugfixes.Error("FilePath", err)
	}
}

func main() {
	prefixChannel := make(chan string)
	go app.App(prefixChannel, fileDir)

	config := settings.GetConfig(fileDir)

	prefix := <- prefixChannel

	webview.Debug(config)

	webSettings := webview.Settings{
		Debug: true,
		Height: config.ScreenHeight,
		Width: config.ScreenWidth,
		Title: "Static Site Generator",
		URL: prefix + "/dist/index.html",
		Resizable: true,
		ExternalInvokeCallback: app.HandleRPC,
	}
	view := webview.New(webSettings)

	defer view.Exit()

	view.Run()

	if config.FullScreen {
		view.Dispatch(func() {
			view.SetFullscreen(true)
		})
	}
}
