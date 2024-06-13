package models

import "mime/multipart"

type Produk struct {
	ID               string                `json:"id" bson:"_id,omitempty"`
	Nama_Produk      string                `json:"nama_produk" form:"nama_produk"`
	Harga_Produk     string                `json:"harga_produk" form:"harga_produk"`
	Kategori_Produk  string                `json:"kategori_produk" form:"kategori_produk"`
	Stok_Produk      string                `json:"stok_produk" form:"stok_produk"`
	Gambar_Produk    *multipart.FileHeader `json:"gambar_produk" form:"gambar_produk"`
	Deskripsi_Produk string                `json:"deskripsi_produk" form:"deskripsi_produk"`
	Status_Produk    string                `json:"status_produk" form:"status_produk"`
}
