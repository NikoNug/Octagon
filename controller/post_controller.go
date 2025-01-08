package controller

import (
	"log"
	"octagon/db"
	"octagon/models"

	"github.com/gofiber/fiber/v2"
)

func GetHelloWorld(c *fiber.Ctx) error {
	return c.JSON("Welcome to the Octagon !")
}

func GetPosts(c *fiber.Ctx) error {
	rows, err := db.DB.Query("SELECT id, title, body FROM posts order by id")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Posts{}

	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			return err
		}

		// Append Employee to Employees
		result.Posts = append(result.Posts, post)
	}
	// Return Employees in JSON format
	return c.JSON(result)
	// return c.Render("index", fiber.Map{
	// 	"Title":    "Home Page",
	// 	"Subtitle": "Post",
	// 	"Data":     result.Posts,
	// })
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post

	rows, err := db.DB.Query("SELECT * FROM posts WHERE id =?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&post.ID, &post.Title, &post.Body)
	}

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
	rows, err := db.DB.Query("INSERT INTO posts (title, body) VALUES (?,?)", u.Title, u.Body)
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
