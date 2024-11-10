// main.go

package main

import (
	"context"
	"log"
	"productproject/internal/config"
	"productproject/internal/handlers"

	product "productproject/internal/product"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := product.NewPostgresDatabase(cfg.GetConnectionString())
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	store := product.NewStore(db)
	h := handlers.NewProductHandlers(store)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			if err := db.Ping(); err != nil {
				log.Printf("Database connection lost: %v", err)
				// พยายามเชื่อมต่อใหม่
				if reconnErr := db.Reconnect(cfg.GetConnectionString()); reconnErr != nil {
					log.Printf("Failed to reconnect: %v", reconnErr)
				} else {
					log.Printf("Successfully reconnected to the database")
				}
			}
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// กำหนดค่า CORS
	configCors := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:4000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(configCors))

	r.Use(TimeoutMiddleware(5 * time.Second))

	r.GET("/health", h.HealthCheck)

	// API v1
	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("", h.GetProducts)
			products.POST("", h.AddProduct)
			products.GET("/:id", h.GetProduct)
			products.PUT("/:id", h.UpdateProduct)
			products.DELETE("/:id", h.DeleteProduct)

			// แก้ไขเส้นทางสำหรับแนะนำสินค้า
			recommendedProducts := products.Group("/Recommendproducts")
			{
				recommendedProducts.GET("", h.GetRecommendedProduct)
			}
			// เส้นทางสำหรับแสดงสินค้าล่าสุดจากแต่ละ Seller
			newProducts := products.Group("/Newproducts")
			{
				newProducts.GET("", h.GetNewProductSeller)
			}
			// เส้นทางสำหรับแสดงสินค้าของผู้ขาย
			sellerProducts := products.Group("/seller/:sellerID")
			{
				sellerProducts.GET("", h.GetDetailProductSeller)
			}

			// Nested resources - Images
			images := products.Group("/:id/images")
			{
				images.GET("", h.GetProductImages)
				images.POST("", h.AddProductImage)
				images.PUT("/:image_id", h.UpdateProductImage)
				images.DELETE("/:image_id", h.DeleteProductImage)
			}

		}
		v1.GET("/shops", h.GetAllShops)
		// เพิ่มเส้นทางสำหรับแสดงรายละเอียดของร้านค้า (Shop Detail)
		v1.GET("/shops/:id", h.GetShopDetail)

		v1.GET("/images", h.GetAllProductImages)
		// Categories
		categories := v1.Group("/categories")
		{
			categories.GET("", h.GetCategories)
		}
	}

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Printf("Failed to run server: %v", err)
	}
}
