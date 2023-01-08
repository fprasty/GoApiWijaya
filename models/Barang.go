package models

type UserBarang struct {
	Id     uint   `json:"id"`
	Title  string `json:"barang_title"`
	Desc   string `json:"barang_deskripsi"`
	Price  int64  `json:"barang_price"`
	Image  string `json:"barang_image"`
	Likes  int64  `json:"barang_likes"`
	UserID string `json:"userid"`
	User   User   `gorm:"foreignkey:UserID"`
}
