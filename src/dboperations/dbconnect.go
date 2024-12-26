package dboperations

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Link struct {
	ID       uint   `gorm:"primaryKey"`
	FullURL  string `gorm:"type:varchar(255)"`
	ShortURL string `gorm:"type:varchar(255)"`
	// Created string `gorm:"type:varchar(100)"`
}

func ConnectToDB() *gorm.DB {
	dsn := "root:secret@tcp(localhost:3306)/shortlink?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func FetchLinks(db *gorm.DB) ([]map[string]interface{}, error) {
	var links []map[string]interface{}
	result := db.Table("links").Find(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}

func PrintDB() {
	db := ConnectToDB()
	links, err := FetchLinks(db)
	if err != nil {
		log.Fatalf("fail to load db %v", err)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}

func ValidateShortInBD(shortURL string) (bool, error) {
	result := true
	db := ConnectToDB()
	var result_link Link
	db.Where("short_url = ?", shortURL[:6]).First(&result_link)
	if result_link.ShortURL == shortURL {
		result = false
		return result, nil
	}
	fmt.Println(result_link.ShortURL, result_link.FullURL, "you are here")
	return result, nil
}

func GetShortURl(fullURL string) (string, error) {
	db := ConnectToDB()
	var result Link
	db.Where("full_url = ?", fullURL).Find(&result)
	if result.ShortURL == "" {
		return "", fmt.Errorf("error: not find shortlink in db")
	}
	shortURL := result.ShortURL
	fmt.Println(shortURL)
	return shortURL, nil
}
