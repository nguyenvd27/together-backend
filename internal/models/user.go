package models

type User struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email" gorm:"unique"`
	Password []byte  `json:"-"`
	Avatar   string  `json:"avatar"`
	Address  int     `json:"address"`
	Events   []Event `json:"events" gorm:"many2many:user_events;"`
}
