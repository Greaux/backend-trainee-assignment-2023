package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Username string    `gorm:"unique"`
	Segments []Segment `gorm:"many2many:user_segments;"`
}

type Segment struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Users []User `gorm:"many2many:user_segments;"`
}

type App struct {
	DB *gorm.DB
}

func (app *App) Initialize() {
	dsn := "host=db user=postgres password=BESOPASNIYPAROL dbname=AvitoTech port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	app.DB = db

	// Автоматическая миграция в БД
	app.DB.AutoMigrate(&User{}, &Segment{})
}

// Создаем юзера
func (app *App) CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	result := app.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(user)
}

// Создаем сегмент
func (app *App) CreateSegment(c *fiber.Ctx) error {
	segment := new(Segment)
	if err := c.BodyParser(segment); err != nil {
		return err
	}

	result := app.DB.Create(&segment)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(segment)
}

func (app *App) AddUserToSegment(c *fiber.Ctx) error {
	userID := c.FormValue("userid")
	var user User

	if err := app.DB.First(&user, userID).Error; err != nil {
		return err
	}

	segmentNames := strings.Split(c.FormValue("segments"), ",")
	var segments []Segment

	// Получение сегментов по их именам
	if err := app.DB.Where("name IN ?", segmentNames).Find(&segments).Error; err != nil {
		return err
	}

	// Добавление пользователя к сегментам
	for _, segment := range segments {
		// Проверка, не добавлен ли уже пользователь к данному сегменту
		alreadyAdded := false
		for _, seg := range user.Segments {
			if seg.ID == segment.ID {
				alreadyAdded = true
				break
			}
		}

		if !alreadyAdded {
			app.DB.Model(&user).Association("Segments").Append(&segment)
		}
	}

	return c.JSON(user)
}

// Удаление сегментов у пользователя
func (app *App) RemUserFromSegment(c *fiber.Ctx) error {
	userID := c.FormValue("userid")
	var user User

	if err := app.DB.First(&user, userID).Error; err != nil {
		return err
	}

	segmentNames := c.FormValue("segments")
	segmentNamesSlice := strings.Split(segmentNames, ",")

	var segments []Segment

	// Получение сегментов по их именам
	if err := app.DB.Where("name IN ?", segmentNamesSlice).Find(&segments).Error; err != nil {
		return err
	}

	// Удаление сегментов пользователя
	for _, segment := range segments {
		app.DB.Model(&user).Association("Segments").Delete(&segment)
	}

	return c.JSON(user)
}

// Удаление сегмента
func (app *App) DeleteSegment(c *fiber.Ctx) error {
	segmentName := c.FormValue("name")
	var segment Segment

	if err := app.DB.Where("name = ?", segmentName).First(&segment).Error; err != nil {
		return err
	}

	// Удаление связанных записей из таблицы "user_segments"
	app.DB.Table("user_segments").Where("segment_id = ?", segment.ID).Delete(nil)

	// Удаление сегмента
	app.DB.Delete(&segment)

	return c.JSON(segment)
}

// Получение сегментов пользователя по его ID
func (app *App) GetUserSegments(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user User

	if err := app.DB.Preload("Segments").First(&user, userID).Error; err != nil {
		return err
	}

	return c.JSON(user.Segments)
}

// Получение сегментов пользователя по его имени
func (app *App) GetUserSegmentsByName(c *fiber.Ctx) error {
	username := c.Params("UserName")
	var user User

	if err := app.DB.Preload("Segments").Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}

	return c.JSON(user.Segments)
}

func (app *App) MainPage(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.SendString("Тут ничего нет.")
}

func main() {
	app := &App{}

	time.Sleep(2 * time.Second)
	app.Initialize()

	appFiber := fiber.New()

	appFiber.Post("/users", app.CreateUser)
	//Для создания юзера нужно в Body запроса добавить ключ Username и значение(уникальное)

	appFiber.Post("/segments", app.CreateSegment)
	// Для создания сегмента нужно в Body запроса добавить ключ Name и значение(уникальное)

	appFiber.Post("/editUserSegments", app.AddUserToSegment)
	// Для добавления пользователя в сегмент нужно в Body запроса добавить ключ userid и значение(UserID), а так же список slug (названий) сегментов которые нужно добавить пользователю

	appFiber.Delete("/editUserSegments", app.RemUserFromSegment)
	// Для удаления у пользователя сегментов нужно в Body запроса добавить ключ userid и значение(UserID), а так же список slug (названий) сегментов которые нужно добавить пользователю

	appFiber.Delete("/segments", app.DeleteSegment)
	// Для удаления сегмента нужно в Body запроса добавить ключ Name и значение(уникальное)

	appFiber.Get("/user/:id", app.GetUserSegments)
	// Для получения сегментов нужно вписать localhost:3000/user/ID

	appFiber.Get("/UserName/:Username", app.GetUserSegmentsByName)
	// Для получения сегментов нужно вписать localhost:3000/UserName/Name

	appFiber.Get("/", app.MainPage)
	// Start the server
	log.Fatal(appFiber.Listen(":3000"))
}
