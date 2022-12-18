package models

type BarangComment struct {
	Id       uint       `json:"id"`
	Rating   string     `json:rating`
	Comment  string     `json:comment`
	ImgComnt string     `json:image`
	Star     string     `json:star`
	BarangID string     `json:barangid`
	Barang   UserBarang `json:"userbarangs";gorm:"foreignkey:BarangID"`
	UserID   string     `json:userid`
	User     User       `json:"user";gorm:"foreignkey:UserID"`
}
