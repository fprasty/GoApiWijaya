package models

type BarangComment struct {
	Id       uint   `json:"id"`
	Rating   string `json:rating`
	Comment  string `json:desc`
	ImgComnt string `json:image`
	Star     string `json:star`
	UserID   string `json:userid`
	User     User   `json:"user";gorm:"foreignkey:UserID"`
}
