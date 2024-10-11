package model

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Song struct {
	gorm.Model
	Title       string    `json:"Title" example:"Song title"`
	Author      string    `json:"Author" example:"Song author"`
	SongGroup   string    `json:"SongGroup" example:"Song group"`
	Link        string    `json:"Link" example:"https://www.youtube.com/watch?v=b_h8kh-PEfI9999"`
	Description string    `json:"Description" example:"Song description text"`
	ReleaseDate time.Time `json:"ReleaseDate" example:"2006-02-01T15:04:05Z" gorm:"default:current_timestamp"`
	Verses      []Verses  `json:"Verses" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Verses struct {
	gorm.Model
	Text   string `json:"Text" example:"Verse text"`
	SongID uint   `json:"SongID" swaggerignore:"true"`
}

var DB *gorm.DB

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

// CreateSong Function to create data for testing
func CreateSong(title string, author string, group string, link string, description string, releaseDate int64, verses []Verses) {
	DB.Create(&Song{
		Model:       gorm.Model{},
		Title:       title,
		Author:      author,
		SongGroup:   group,
		Link:        link,
		Description: description,
		ReleaseDate: time.Unix(releaseDate, 0),
		Verses:      verses,
	})
}

func InitDatabase() {
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		log.Fatal("Wrong DB_HOST")
	}
	port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		log.Fatal("Wrong DB_PORT")
	}
	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Fatal("Wrong DB_USER")
	}
	pass, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Fatal("Wrong DB_PASSWORD")
	}
	database, exists := os.LookupEnv("DATABASE_NAME")
	if !exists {
		log.Fatal("Wrong DATABASE_NAME")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	} else {
		DB = db
		err := db.AutoMigrate(&Song{}, &Verses{})
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
