package models

type RequestResPayload struct {
	User    User    `json:"data"`
	Support Support `json:"support"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Support struct {
	URL  string `gorm:"not null" json:"url"`
	Text string `gorm:"not null" json:"text"`
}
