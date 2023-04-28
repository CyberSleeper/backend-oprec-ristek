package controllers

import (
	"time"

	"github.com/CyberSleeper/backend-oprec-ristek/configs"
	"github.com/CyberSleeper/backend-oprec-ristek/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePostHandler(c *fiber.Ctx) error {
	var payload *models.CreatePostSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if errors := models.ValidateStruct(payload); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	now := time.Now()
	newPost := models.Post{
		CreatedAt: now,
		UpdatedAt: now,
		Caption:   payload.Caption,
	}

	result := configs.DB.Create(&newPost)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"post": newPost,
		},
	})
}

func GetPostsHandler(c *fiber.Ctx) error {
	var posts []models.Post

	results := configs.DB.Find(&posts)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": results.Error,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"results": len(posts),
		"posts":   posts,
	})
}

func GetPostByIdHandler(c *fiber.Ctx) error {
	postId := c.Params("postId")
	var post models.Post
	result := configs.DB.Find(&post, "id = ?", postId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "fail",
				"message": "No post with that id exists",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"post": post,
		},
	})
}

func UpdatePostHandler(c *fiber.Ctx) error {
	postId := c.Params("postId")
	var payload *models.UpdatePostSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	var post models.Post
	result := configs.DB.Find(&post, "id = ?", postId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "fail",
				"message": "No post with that id exists",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error,
		})
	}

	updates := make(map[string]interface{})
	updates["updatedAt"] = time.Now()
	if payload.Caption != "" {
		updates["caption"] = payload.Caption
	}

	configs.DB.Model(&post).Updates(updates)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"post": post,
		},
	})
}

func DeletePostHandler(c *fiber.Ctx) error {
	postId := c.Params("postId")

	result := configs.DB.Delete(&models.Post{}, "id = ?", postId)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "No post with that id exists",
		})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
