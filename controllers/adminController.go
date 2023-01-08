package controllers

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/fprasty/GoApiWijaya/database"
	"github.com/fprasty/GoApiWijaya/models"
	"github.com/fprasty/GoApiWijaya/util"
	"github.com/gofiber/fiber/v2"
)

/*------------------Admin Auth Session---------------------------*/
// This for Admin Register
func AdminRegister(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.Admin
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Register>Unable to parse body")
	}

	//is Email correct?
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"massage": "Invalid email address",
		})
	}

	//is Email Exist?
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	user := models.Admin{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	//Hashing Password Here
	user.SetPassword(data["password"].(string))
	//CreateUser
	err := database.DB.Create(&user)
	if err == nil {
		log.Println("Can't create user")
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"massage": "Akun admin berhasil dibuat",
	})
}

// This for Admin Login
func AdminLogin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	var user models.Admin
	//Searching email in database
	database.DB.Where("email=?", data["email"]).First(&user)
	//is Email nill?
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Address doesn't exit",
		})
	}
	//Comparing password
	//is Password Correct?
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	//Generate token jwt
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	//Set Cookie
	cookie := fiber.Cookie{
		Name:     "Admin-jwt", //Cookie's name
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "you have successfully login",
		"user":    user,
	})

}

/*------------------End Admin Auth Session---------------------------*/

/*------------------Admin Controll Session---------------------------*/

// Get All User in Page
func AllUser(c *fiber.Ctx) error {
	cookie := c.Cookies("Admin-jwt")
	_, err := util.Parsejwt(cookie)
	if err != nil {
		return c.JSON(err)
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getuser []models.User
	database.DB.Model(&getuser).Where("id=?", "id").Find(&getuser)
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

// Admin Get Admin by uuid
func AdminGetme(c *fiber.Ctx) error {
	cookie := c.Cookies("Admin-jwt")
	uuid, err := util.Parsejwt(cookie)
	//Validate Token
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Unable get token",
			"error":   err,
		})
	}

	var user models.Admin
	if err := database.DB.Model(&user).Where("id=?", uuid).Find(&user); err == nil {
		c.Status(400)
		return c.JSON("User not found")
	}
	var UserData = make(map[string]interface{})
	UserData["user_id"] = user.Id
	UserData["uuid"] = user.UUID
	UserData["first_name"] = user.FirstName
	UserData["last_name"] = user.LastName
	UserData["email"] = user.Email
	UserData["phone"] = user.Phone

	return c.JSON(fiber.Map{
		"message": "Sucess get user",
		"user":    UserData,
	})

}

// Admin Get User by id
func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("Admin-jwt")
	//Validate Token
	cparams, _ := util.Parsejwt(cookie)
	//Convert token from string to int
	token, _ := strconv.Atoi(string(cparams))
	var user = models.User{
		Id: uint(token),
	}
	id, _ := strconv.Atoi(c.Params("id"))
	var uparams = models.User{
		Id: uint(id),
	}

	database.DB.Model(&user).Where("id=?", id).Find(&user)
	if user.Id != uparams.Id {
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

// Admin Get User by uuid
func GetUserUUID(c *fiber.Ctx) error {
	cookie := c.Cookies("Admin-jwt")
	//Validate Token
	_, err := util.Parsejwt(cookie)
	if err != nil {
		c.Status(400)
		return c.JSON("Token not found")
	}

	id := c.Params("uuid")
	var user = models.User{
		UUID: id,
	}
	database.DB.Model(&user).Where("uuid=?", id).Find(&user)
	if user.Id == 0 {
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
