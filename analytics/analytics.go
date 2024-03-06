package main

import (
  "fmt"
  "log"
	_ "github.com/lib/pq"
  "os"
	"github.com/joho/godotenv"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func main() {
  err := godotenv.Load("../.env")
  if err != nil {
    log.Fatal("Error loading .env file.")
  }
  dsn := fmt.Sprintf("host=%s user =%s password=%s dbname=%s port=%s sslmode=disable TimeZone=CET",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASS"),
    os.Getenv("DB_DATABASE"),
    os.Getenv("DB_PORT"),
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal("Error connecting to database.", err)
  }

  fmt.Println("Successful connection to database.")
  var numberOfUsers int 
  var averageFollowers float64
  query1 := `select count(*) from users;`
  result := db.Raw(query1).Scan(&numberOfUsers)
  if result.Error != nil {
    log.Fatalf("Query failed: %v", result.Error)
  }
  query2 := `select (select count (*) from follows) / (select count(*)::FLOAT from users) as average_followers;`
  result2 := db.Raw(query2).Scan(&averageFollowers)
  if result2.Error != nil {
    log.Fatalf("Query failed: %v", result2.Error)
  }

  fmt.Printf("Total number of users: %v\n", numberOfUsers)
  fmt.Printf("Average followers per user: %v\n", averageFollowers)
}
