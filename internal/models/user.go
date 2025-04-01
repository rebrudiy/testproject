package models

type RequestResPayload struct {
	User    User    `json:"data"`
	Support Support `json:"support"`
}

type User struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Email     string `gorm:"unique;not null" json:"email"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Support struct {
	URL  string `gorm:"not null" json:"url"`
	Text string `gorm:"not null" json:"text"`
}
