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

func CreateComment(c *fiber.Ctx) error {
	var barcomen models.BarangComment
	if err := c.BodyParser(&barcomen); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&barcomen).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Komen berhasil di unggah",
	})

}

func AllComment(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getcomment []models.BarangComment
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getcomment)
	database.DB.Model(&models.BarangComment{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getcomment,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func DetailComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var BarangComment models.BarangComment
	database.DB.Where("id=?", id).Preload("User").First(&BarangComment)
	return c.JSON(fiber.Map{
		"data": BarangComment,
	})

}

func UpdateComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	barang := models.BarangComment{
		Id: uint(id),
	}

	if err := c.BodyParser(&barang); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&barang).Updates(barang)
	return c.JSON(fiber.Map{
		"message": barang,
	})

}

func UniqueComment(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	var comment []models.BarangComment
	database.DB.Model(&comment).Where("user_id=?", id).Preload("User").Find(&comment)

	return c.JSON(comment)

}
func DeleteComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	comment := models.BarangComment{
		Id: uint(id),
	}
	deleteQuery := database.DB.Delete(&comment)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Opps!, record Not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Komen deleted Succesfully",
	})

}
