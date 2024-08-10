package handlers

import (
    "gotest/database"
    "gotest/models"
    "gotest/utils"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
	"time"
)

func Register(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return err
    }

    password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

    user := models.User{
        Name: data["name"],
        Email: data["email"],
        Password: string(password),
        Role: "user", // Default role
    }

    // handle UID
	for {
		user.Uid = utils.RandStr()
		if database.DB.Where("uid = ?", user.Uid).First(&models.User{}).Error != nil {
			break
		}
	}

    if err := database.DB.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
    // cookie := c.Cookies()
    // if cookie == "" {
    //     return c.SendStatus(fiber.StatusUnauthorized)
    // }
    // return c.SendStatus(fiber.StatusOK)

    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return err
    }
    var user models.User

    database.DB.Where("email = ?", data["email"]).First(&user)

    if user.ID == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect password"})
    }

    token, err := utils.GenerateJWT(user.ID, user.Role)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

	// use cookie
	c.Cookie(&fiber.Cookie{
        Name:     "jwt",
        Value:    token,
        Expires:  time.Now().Add(72 * time.Hour),
        HTTPOnly: true,
    })

    return c.JSON(fiber.Map{"token": token})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
        Name:     "jwt",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour),
        HTTPOnly: true,
    })
    return c.Redirect("/login")
}
