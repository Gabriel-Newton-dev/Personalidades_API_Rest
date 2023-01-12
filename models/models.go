package models

type Personalidade struct {
	//gorm.Model
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}
