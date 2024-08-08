package handlers

import (
    "gotest/database"
    "gotest/models"
    "path/filepath"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
    "os"
    "time"
)

const staticImgDir = "../static/uploads/images"

func GetUsers(c *fiber.Ctx) error {
    var users []models.User
    database.DB.Find(&users)
    return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    passwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    user.Password = string(passwd)

    // Handle image
    file, err := c.FormFile("image")
    if err == nil {
        timestamp := time.Now().Format("20060102150405")
        newFilename := timestamp + "_" + file.Filename
        imagePath := filepath.Join(staticImgDir, newFilename)
        os.MkdirAll(staticImgDir, os.ModePerm)
        if err := c.SaveFile(file, imagePath); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        user.Image = newFilename
    }

    if err := database.DB.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User

    if err := database.DB.First(&user, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    var updateUser models.User
    if err := c.BodyParser(&updateUser); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Check if the password has been updated
    if updateUser.Password != "" {
        passwd, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        user.Password = string(passwd)
    }

    // Update other fields
    user.Email = updateUser.Email
    user.Name = updateUser.Name

    // Handle image
    file, err := c.FormFile("image")
    if err == nil {
        // Remove old image if exists
        if user.Image != "" {
            oldImagePath := filepath.Join(staticImgDir, filepath.Base(user.Image))
            os.Remove(oldImagePath)
        }

        timestamp := time.Now().Format("20060102150405")
        newFilename := timestamp + "_" + file.Filename
        imagePath := filepath.Join(staticImgDir, newFilename)
        if err := c.SaveFile(file, imagePath); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        user.Image = newFilename
    }

    if err := database.DB.Save(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User

    // Handle image
    if err := database.DB.First(&user, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    if user.Image != "" {
        imagePath := filepath.Join(staticImgDir, filepath.Base(user.Image))
        os.Remove(imagePath)
    }

    // soft delete => database.DB.Delete(&user, id)
    if err := database.DB.Unscoped().Delete(&user, id).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.SendStatus(fiber.StatusOK)
}
