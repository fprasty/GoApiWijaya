package controllers

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/fprasty/GoApiWijaya/database"
	"github.com/fprasty/GoApiWijaya/models"
	"github.com/fprasty/GoApiWijaya/util"
	"github.com/gofiber/fiber/v2"
)

// Get All User
func AllUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, err := util.Parsejwt(cookie)
	if err != nil {
		return c.JSON(err)
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getuser []models.User
	database.DB.Model(&getuser).Where("id=?", id).Find(&getuser)
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getuser)
	database.DB.Model(&models.User{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getuser,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

// Get User by id
func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	/*if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Unable get token",
			"error":   err,
		})
	}*/
	var user models.User
	if err := database.DB.Model(&user).Where("id=?", id).Find(&user); err == nil {
		c.Status(400)
		return c.JSON("User not found")
	}
	var UserData = make(map[string]interface{})
	UserData["user_id"] = user.Id
	UserData["first_name"] = user.FirstName
	UserData["last_name"] = user.LastName
	UserData["email"] = user.Email
	UserData["phone"] = user.Phone

	return c.JSON(fiber.Map{
		"message": "Sucess get user",
		"user":    UserData,
	})

}

func UpdateUser(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Register>Unable to parse body")
	}
	//Cookie session
	cookie := c.Cookies("jwt")
	cparse, _ := util.Parsejwt(cookie)
	id := c.Params("id")
	user := models.User{
		Id: cparse,
	}

	uparams := models.User{
		Id: id,
	}
	//Params id cek
	if uparams.Id != user.Id {
		c.Status(400)
		return c.JSON("Id not match")
	}

	//Check jika email sudah ada
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	update := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	//UpdateUser
	err := database.DB.Model(&user).Updates(&update)
	if err != nil {
		log.Println(err)
	}

	var UserData = make(map[string]interface{})
	UserData["first_name"] = update.FirstName
	UserData["last_name"] = update.LastName
	UserData["email"] = update.Email
	UserData["phone"] = update.Phone

	c.Status(200)
	return c.JSON(fiber.Map{
		"cookie":  id,
		"userid":  user.Id,
		"user":    UserData,
		"message": "barang updated successfully",
	})
}
