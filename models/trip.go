package models

type Trip struct {
	ID          uint   `json:"id"`
	Banner      string `json:"banner_image"`
	Destination string `json:"destination"`
	Vehicles    string `json:"email" gorm:"unique"`
	Hotel       string `json:"hotel"`
}
