package controllers

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/MartinAbelRR/blogbackend/database"
	"github.com/MartinAbelRR/blogbackend/models"
	"github.com/MartinAbelRR/blogbackend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

	func CreatePost(c *fiber.Ctx) error {
		var blogpost models.Blog
		if err := c.BodyParser(&blogpost); err != nil {
			fmt.Println("Unable to parse body")
		}

		if err := database.DB.Create(&blogpost).Error; err != nil {
			c.Status(400)
			c.JSON(fiber.Map{
				"message": "Invalid payload",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Contratulation!, Your post is live",		
		})
	}

func AllPost(c *fiber.Ctx) error {
	// Gets the page number of the URL, If not provided, it is set to 1 by default.
	page , _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page-1) * limit
	var total int64
	var getblogs []models.Blog
	
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblogs)
	database.DB.Model(&models.Blog{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": getblogs,
		"meta": fiber.Map{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64((int(total)/limit))),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePost(c * fiber.Ctx) error {
	id , _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&blog).Updates(blog)

	return c.JSON(fiber.Map{
		"message": "Post update successfully",

	})	
}

func UniquePost(c *fiber.Ctx) error {
	cookies := c.Cookies("jwt")
	id , _ := utils.Parsejwt(cookies)
	var blog []models.Blog
	
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
	
	return c.JSON(blog)
}

func DeletePost(c *fiber.Ctx) error {
	id , _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound){
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Opps!, record not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}