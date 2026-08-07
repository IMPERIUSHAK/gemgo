package models

type Recipe struct {
	Name        string   `json:"name"`
	TimeMinutes int      `json:"timeMinutes"`
	Calories    int      `json:"calories"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}
