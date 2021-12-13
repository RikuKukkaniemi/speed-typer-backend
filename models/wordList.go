package models

type WordList struct {
	Language  *string   `json:"language"`
	Words	  *[]string `json:"words"`
}
