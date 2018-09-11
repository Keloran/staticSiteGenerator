package settings

type Config struct {
	ScreenWidth  int  `json:"screenWidth"`
	ScreenHeight int  `json:"screenHeight"`
	FullScreen   bool `json:"fullScreen"`
}
