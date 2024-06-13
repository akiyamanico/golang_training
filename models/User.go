package models

import "mime/multipart"

type User struct {
	ID              string                `json:"id" bson:"_id,omitempty"`
	Username        string                `json:"username" form:"username"`
	Email           string                `json:"email" form:"email"`
	Password        string                `json:"password" form:"password"`
	Nama            string                `json:"nama" form:"nama"`
	NoTelp          string                `json:"no_telp" form:"no_telp"`
	Tanggal_Lahir   string                `json:"tanggal_lahir" form:"tanggal_lahir"`
	Jenis_Kelamin   string                `json:"jenis_kelamin" form:"jenis_kelamin"`
	Profile_Picture *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
}
