package models

type Activities struct {
	Id      int64  `json:"id" gorm:"primaryKey"`
	Parent  string `json:"parent"`
	Label   string `json:"label"`
	PathUrl string `json:"path_url"`
}
