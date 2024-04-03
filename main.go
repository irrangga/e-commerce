package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	product_handler "e-commerce/internal/handler/rest/product"
	user_handler "e-commerce/internal/handler/rest/user"
	product_repo "e-commerce/internal/repo/product"
	stock_repo "e-commerce/internal/repo/stock"
	user_repo "e-commerce/internal/repo/user"
	product_uc "e-commerce/internal/usecase/product"
	user_uc "e-commerce/internal/usecase/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// SETUP DATABASE
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	dbName := os.Getenv("DB_NAME")

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to open database")
	}

	_, err = sqlDB.Exec("CREATE DATABASE " + dbName)
	if err != nil && err.Error() != fmt.Sprintf(`pq: database "%s" already exists`, dbName) {
		panic("failed to initiate database")
	}

	dsn += " dbname=" + dbName
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// REPO INITIALIZATION
	userRepo := user_repo.New(db)
	productRepo := product_repo.New(db)
	stockRepo := stock_repo.New(db)

	// USECASE INITIALIZATION
	userUsecase := user_uc.New(userRepo)
	productUsecase := product_uc.New(productRepo, stockRepo)

	// HANDLER INITIALIZATION
	userHandler := user_handler.New(userUsecase)
	productHandler := product_handler.New(productUsecase)

	// SETUP ROUTER
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to e-commerce project!",
		})
	})

	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	router.GET("/products", productHandler.GetListProducts)
	router.GET("/products/:id", productHandler.GetProduct)
	router.POST("/products", productHandler.CreateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	log.Printf("server listening at %s", serverPort)

	router.Run(serverPort)
}
