package controllers

import (
	"backend/database"
	m "backend/models"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HelloTest1(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + "----" + p.Pass
	fmt.Println("asdsss")
	return c.SendString(str)
}

func HelloTest2(c *fiber.Ctx) error {
	c.Params("name") // "fenny"
	str := "Hello " + c.Params("name")
	return c.SendString(str)
	// ...
}

func HelloTest3(c *fiber.Ctx) error {
	validate := validator.New()
	//Connect to database
	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func AddCalculate(c *fiber.Ctx) error {
	id_id := c.Params("id")
	id, errs := strconv.Atoi(id_id)
	var calculate = factorial(id)
	str := strconv.Itoa(calculate)

	fmt.Println(" ", errs)
	return c.SendString(str)
}

func Register(c *fiber.Ctx) error {
	db := database.DBConn
	validate := validator.New()
	register := new(m.Register)

	if err := c.BodyParser(&register); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := validate.Struct(register)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	db.Create(&register)
	return c.Status(201).JSON(register)
}

func GetID(c *fiber.Ctx) error {
	db := database.DBConn
	var register []m.Register

	db.Find(&register)
	return c.Status(200).JSON(register)
}

func DeleteID(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var register m.Register

	result := db.Delete(&register, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

// func GetDogsJson(c *fiber.Ctx) error {
// 	db := database.DBConn
// 	var dogs []m.Dogs

// 	db.Find(&dogs)

// 	type DogsRes struct {
// 		Name  string `json:"name"`
// 		DogID int    `json:"dog_id"`
// 	}
// 	type typeColor struct {
// 		Type string `json:"type"`
// 	}
// 	var dataResults []DogsRes
// 	var redResults []typeColor
// 	var greenResults []typeColor
// 	var pinkResults []typeColor
// 	var no_colorResults []typeColor
// 	for _, v := range dogs {
// 		typeStr := ""
// 		if v.DogID >= 10 && v.DogID <= 50 {
// 			typeStr = "Red"
// 			sred := typeColor{
// 				Type: typeStr,
// 			}
// 			redResults = append(redResults, sred)
// 		} else if v.DogID >= 100 && v.DogID <= 150 {
// 			typeStr = "green"
// 			sgreen := typeColor{
// 				Type: typeStr,
// 			}
// 			greenResults = append(greenResults, sgreen)
// 		} else if v.DogID >= 200 && v.DogID <= 250 {
// 			typeStr = "pink"
// 			sred := typeColor{
// 				Type: typeStr,
// 			}
// 			pinkResults = append(pinkResults, sred)
// 		} else {
// 			typeStr = "no color"
// 			sred := typeColor{
// 				Type: typeStr,
// 			}
// 			no_colorResults = append(no_colorResults, sred)
// 		}

// 		d := DogsRes{
// 			Name:  v.Name,
// 			DogID: v.DogID,
// 		}
// 		dataResults = append(dataResults, d)
// 	}
// 	return c.Status(200).JSON(map[string]interface{}{
// 		"data":        dataResults,
// 		"name":        "golang-test",
// 		"count":       len(dogs),
// 		"sum_red":     len(redResults),
// 		"sum_green":   len(greenResults),
// 		"sum_pink":    len(pinkResults),
// 		"sum_nocolor": len(no_colorResults),
// 	})
// }

// func GetDogsJson(c *fiber.Ctx) error {
// 	db := database.DBConn
// 	var dogs []m.Dogs

// 	db.Find(&dogs)

// 	type DogsRes struct {
// 		Name  string `json:"name"`
// 		DogID int    `json:"dog_id"`
// 	}
// 	type typeColor struct {
// 		Sum_red     int `json:"sum_red"`
// 		Sum_green   int `json:"sum_green"`
// 		Sum_pink    int `json:"sum_pink"`
// 		Sum_nocolor int `json:"sum_nocolor"`
// 	}
// 	var dataResults []DogsRes
// 	var sum_color []typeColor

// 	sum_red := 0
// 	sum_green := 0
// 	sum_pink := 0
// 	sum_nocolor := 0
// 	for _, v := range dogs {
// 		if v.DogID >= 10 && v.DogID <= 50 {
// 			sum_red++
// 		} else if v.DogID >= 100 && v.DogID <= 150 {
// 			sum_green++
// 		} else if v.DogID >= 200 && v.DogID <= 250 {
// 			sum_pink++
// 		} else {
// 			sum_nocolor++
// 		}
// 		d := DogsRes{
// 			Name:  v.Name,
// 			DogID: v.DogID,
// 		}
// 		dataResults = append(dataResults, d)
// 	}
// 	s := typeColor{
// 		Sum_red:     sum_red,
// 		Sum_green:   sum_green,
// 		Sum_pink:    sum_pink,
// 		Sum_nocolor: sum_nocolor,
// 	}
// 	sum_color = append(sum_color, s)

// 	return c.Status(200).JSON(map[string]interface{}{
// 		"data":      dataResults,
// 		"name":      "golang-test",
// 		"count":     len(dogs),
// 		"sum_color": sum_color,
// 	})
// }

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	type DogsRes struct {
		Name  string `json:"name"`
		DogID int    `json:"dog_id"`
	}
	type typeColor struct {
		Sum_red     int `json:"sum_red"`
		Sum_green   int `json:"sum_green"`
		Sum_pink    int `json:"sum_pink"`
		Sum_nocolor int `json:"sum_nocolor"`
	}
	var dataResults []DogsRes
	var sum_color []typeColor

	sum_red := 0
	sum_green := 0
	sum_pink := 0
	sum_nocolor := 0
	for _, v := range dogs {
		if v.DogID >= 10 && v.DogID <= 50 {
			sum_red++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			sum_green++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			sum_pink++
		} else {
			sum_nocolor++
		}
		d := DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
		}
		dataResults = append(dataResults, d)
	}
	s := typeColor{
		Sum_red:     sum_red,
		Sum_green:   sum_green,
		Sum_pink:    sum_pink,
		Sum_nocolor: sum_nocolor,
	}
	sum_color = append(sum_color, s)

	return c.Status(200).JSON(map[string]interface{}{
		"data":        dataResults,
		"name":        "golang-test",
		"count":       len(dogs),
		"sum_red":     sum_red,
		"sum_green":   sum_green,
		"sum_pink":    sum_pink,
		"sum_nocolor": sum_nocolor,
	})
}

func AddProfile(c *fiber.Ctx) error {
	db := database.DBConn
	validate := validator.New()
	profile_user := new(m.Profile_User)

	if err := c.BodyParser(&profile_user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := validate.Struct(profile_user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	db.Create(&profile_user)
	return c.Status(201).JSON(profile_user)
}

func GetProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile_user []m.Profile_User

	db.Find(&profile_user)
	return c.Status(200).JSON(profile_user)
}

func UpdateProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile_user m.Profile_User
	id := c.Params("id")

	if err := c.BodyParser(&profile_user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&profile_user)
	return c.Status(200).JSON(profile_user)
}

func RemoveProfile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var profile_user m.Profile_User

	result := db.Delete(&profile_user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
