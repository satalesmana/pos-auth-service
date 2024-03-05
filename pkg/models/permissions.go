package models

type Permissions struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	UserId     string `json:"users_id"`
	ActivityId string `json:"activity_id"`
}
