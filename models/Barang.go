package models

type UserBarang struct {
	Id     uint   `json:"id"`
	Title  string `json:title`
	Desc   string `json:desc`
	Image  string `json:image`
	Star   string `json:star`
	Price  int    `json:price`
	UserID string `json:userid`
	User   User   `json:"user";gorm:"foreignkey:UserID"`
}
