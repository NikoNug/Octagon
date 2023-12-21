package controller

import (
	"octagon/db"
	"octagon/models"

	"github.com/gofiber/fiber/v2"
)

func GetPersons(c *fiber.Ctx) error {
	db := db.ConnectDB()

	// rows, err := db.Query("SELECT id, title, body FROM posts order by id")
	rows, err := db.Query("select FirstName, LastName, Address, City from persons")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	defer db.Close()
	result := models.Persons{}

	for rows.Next() {
		person := models.Person{}
		if err := rows.Scan(&person.Firstname, &person.Lastname, &person.City, &person.Address); err != nil {
			return err
		}

		// Append Employee to Employees
		result.Persons = append(result.Persons, person)
	}
	// Return Employees in JSON format
	return c.JSON(result)
}
