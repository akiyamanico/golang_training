package models

type User struct {
	Username       string `json:"username"`
	Email          string `json:"email" gorm:"unique"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profile_picture"`
	Nama           string `json:"nama"`
	NoTelp         string `json:"no_telp"`
	Tanggal_Lahir  string `json:"tanggal_lahir"`
	Jenis_Kelamin  string `json:"jenis_kelamin"`
}
