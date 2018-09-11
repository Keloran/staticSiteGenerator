package items

type Item struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Parsed   string `json:"parsed"`
	Original string `json:"parsed"`
	Category string `json:"category"`
	Icon 	 string `json:"icon"`
}
