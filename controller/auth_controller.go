package controller

import (
	"database/sql"
	"log"
	"octagon/db"
	"octagon/dtos"
	"octagon/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Insert Employee into database
	user.UserID = int(uuid.New().ID())
	rows, err := db.DB.Query("INSERT INTO users (UserID, FirstName, LastName, Username, Password, Email, Address, City) VALUES (?,?,?,?,?,?,?,?)", user.UserID, user.Firstname, user.Lastname, user.Username, hashedPassword, user.Email, user.Address, user.City)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	rows.Close()

	// Return Employee in JSON format
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "register success",
	})
}

func Login(c *fiber.Ctx) error {
	userInput := new(models.User)
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Cek credentials in db
	user := new(models.User)
	err := db.DB.QueryRow(`SELECT Email, Password FROM users WHERE Email = ?`, userInput.Email).Scan(&user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(fiber.Map{
				"error": "Email or Password is Incorrect",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Create JWT Token
	expTime := time.Now().Add(time.Minute * 1)
	claims := &dtos.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "octagon",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Algorithm for signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(dtos.JWT_KEY)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HTTPOnly: true,
	})

	return c.Status(200).JSON(fiber.Map{
		"message": "Login Success!",
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout Success",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	user := new(models.User)

	// Parse body into struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Delete Employee from database
	rows, err := db.DB.Query("DELETE FROM users WHERE UserID=?", user.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	// Print result
	log.Println(rows)

	// Return Employee in JSON format
	return c.Status(fiber.StatusOK).JSON("User Deleted!")
}
