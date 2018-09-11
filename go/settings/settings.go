package settings

import (
	"encoding/json"
	"github.com/bugfixes/go-bugfixes"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getConfigFile(fileDir string) (config Config) {
	_, err := os.Stat(fileDir + "/config.json")
	if err != nil {
		config.FullScreen = false
		config.ScreenHeight = 768
		config.ScreenWidth = 1024

		writeConfigFile(fileDir, config)
	}

	body, err := ioutil.ReadFile(fileDir + "/config.json")
	err = json.Unmarshal([]byte(body), &config)
	if err != nil {
		return config
	}

	return config
}

func writeConfigFile(fileDir string, config Config) {
	setConfig, err := json.Marshal(config)
	if err != nil {
		bugfixes.Fatal("Marshal JSON", err)
	}

	err = ioutil.WriteFile(fileDir + "/config.json", setConfig, 0644)
	if err != nil {
		bugfixes.Fatal("Marshal JSON", err)
	}
}

func GetConfig(fileDir string) (config Config) {
	return getConfigFile(fileDir)
}

func getFileDir() (fileDir string) {
	fileDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		bugfixes.Error("FilePath", err)
	}

	return fileDir
}

func SetWidth(fileDir string, width int) {
	if fileDir == "" {
		fileDir = getFileDir()
	}

	config := getConfigFile(fileDir)
	config.ScreenWidth = width

	writeConfigFile(fileDir, config)
}

func SetHeight(fileDir string, height int) {
	if fileDir == "" {
		fileDir = getFileDir()
	}

	config := getConfigFile(fileDir)
	config.ScreenHeight = height

	writeConfigFile(fileDir, config)
}

func SetFullScreen(fileDir string, fullscreen bool) {
	if fileDir == "" {
		fileDir = getFileDir()
	}

	config := getConfigFile(fileDir)
	config.FullScreen = fullscreen

	writeConfigFile(fileDir, config)
}