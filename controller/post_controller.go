package controller

import (
	"database/sql"
	"log"
	"octagon/db"
	"octagon/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetHelloWorld(c *fiber.Ctx) error {
	return c.JSON("Welcome to the Octagon !")
}

func GetPosts(c *fiber.Ctx) error {
	rows, err := db.DB.Query("select ID, UserID, Title, Body, ImageURL from posts order by CreatedAt desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	result := models.Posts{}

	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.ImageURL); err != nil {
			return err
		}

		user := models.User{}

		userQuery := "SELECT UserID, Firstname, Lastname, Username, Email, Password, Address, City FROM users WHERE UserID = ?"
		if err := db.DB.QueryRow(userQuery, post.UserID).Scan(
			&user.UserID, &user.Firstname, &user.Lastname, &user.Username, &user.Email, &user.Password, &user.Address, &user.City,
		); err != nil {
			return c.Status(500).SendString(err.Error())
		} else {
			post.UserInfo = user // Menambahkan detail user ke post
		}

		// Append Employee to Employees
		result.Posts = append(result.Posts, post)
	}

	// Return Employees in JSON format
	return c.JSON(result)
}

// func GetPost(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var post models.Post

// 	rows, err := db.DB.Query("SELECT ID, UserID, Title, Body, ImageURL FROM posts WHERE ID=?", id)
// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.ImageURL)
// 	}

// 	return c.JSON(post)
// }

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")

	// Validasi ID
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	// Ambil data post berdasarkan ID
	post := models.Post{}
	postQuery := "SELECT ID, UserID, Title, Body, ImageURL FROM posts WHERE ID = ?"
	if err := db.DB.QueryRow(postQuery, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.ImageURL); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"message": "Post not found",
			})
		}
		return c.Status(500).SendString(err.Error())
	}

	// Ambil data user berdasarkan UserID
	user := models.User{}
	userQuery := "SELECT UserID, Firstname, Lastname, Username, Email, Password, Address, City FROM users WHERE UserID = ?"
	if err := db.DB.QueryRow(userQuery, post.UserID).Scan(
		&user.UserID, &user.Firstname, &user.Lastname, &user.Username, &user.Email, &user.Password, &user.Address, &user.City,
	); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"message": "User not found for the post",
			})
		}
		return c.Status(500).SendString(err.Error())
	}

	// Tambahkan detail user ke post
	post.UserInfo = user

	// Return data dalam format JSON
	return c.JSON(post)
}

func AddPost(c *fiber.Ctx) error {
	// New Employee struct
	u := new(models.Post)

	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Insert Employee into database
	// rows, err := db.DB.Query("INSERT INTO posts (UserID, body) VALUES (?,?)", u.Title, u.Body)
	// insert into posts (UserID, Title, Body, ImageURL) values (347978433, "Berbicara", "Bagaimana Caranya?", "test.jpg")
	rows, err := db.DB.Query("insert into posts (UserID, Title, Body, ImageURL) values (?,?,?,?)", u.UserID, u.Title, u.Body, u.ImageURL)
	if err != nil {
		return err
	}
	rows.Close()

	// Print result
	log.Println(rows)

	// Return Employee in JSON format
	return c.JSON(u)
}

func DeletePost(c *fiber.Ctx) error {
	u := new(models.Post)

	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Delete Employee from database
	res, err := db.DB.Query("DELETE FROM posts WHERE title =?", u.Title)
	if err != nil {
		return err
	}

	// Print result
	log.Println(res)

	// Return Employee in JSON format
	return c.JSON("Data Deleted!")
}
