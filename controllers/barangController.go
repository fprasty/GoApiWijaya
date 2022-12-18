package controllers

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/fprasty/GoApiWijaya/database"
	"github.com/fprasty/GoApiWijaya/models"
	"github.com/fprasty/GoApiWijaya/util"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var barangpost models.UserBarang
	if err := c.BodyParser(&barangpost); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&barangpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Barang berhasil di unggah",
	})

}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.UserBarang
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.UserBarang{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var barangpost models.UserBarang
	database.DB.Where("id=?", id).Preload("User").First(&barangpost)
	return c.JSON(fiber.Map{
		"data": barangpost,
	})

}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	barang := models.UserBarang{
		Id: uint(id),
	}

	if err := c.BodyParser(&barang); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&barang).Updates(barang)
	return c.JSON(fiber.Map{
		"message": "barang updated successfully",
	})

}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	var barang []models.UserBarang
	database.DB.Model(&barang).Where("user_id=?", id).Preload("User").Find(&barang)

	return c.JSON(barang)

}
func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	barang := models.UserBarang{
		Id: uint(id),
	}
	deleteQuery := database.DB.Delete(&barang)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Opps!, record Not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "barang deleted Succesfully",
	})

}
