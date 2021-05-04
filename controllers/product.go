package controllers

import (
	"crud_fiber_go_gorm/database"
	"crud_fiber_go_gorm/models"

	"github.com/gofiber/fiber/v2"
)

func GetProductsAll(c *fiber.Ctx) error {
	db := database.DBConn
	var product []models.Product
	db.Find(&product)
	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var product []models.Product
	db.Find(&product, id)
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {

	db := database.DBConn

	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return err
	}

	db.Create(&product)

	return c.JSON(product)

}

func UpdateProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	db := database.DBConn

	product := models.Product{}

	error := db.Model(product).Where("id = ?", id).First(&product).Error
	if error != nil {
		return c.Status(500).SendString(error.Error())
	}

	if error = c.BodyParser(&product); error != nil {
		return c.Status(500).SendString(error.Error())
	}

	result := db.Model(product).Where("id = ?", id).Updates(product)

	if result.Error != nil {
		return c.Status(404).SendString("No Product Found")
	}

	return c.Status(200).JSON(map[string]interface{}{"result": "success"})

}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var product models.Product

	db.First(&product, id)
	if product.P_Name == "" {
		return c.Status(500).SendString("No Product Found with ID")

	}

	// Soft  permanently
	db.Delete(&product, id)

	// Delete permanently
	// db.Unscoped().Delete(&product, id)

	return c.SendString("Product Successfully deleted")

}
