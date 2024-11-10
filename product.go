// ecommerce.go
package ecommerce

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Brand          string    `json:"brand"`
	ModelNumber    string    `json:"model_number"`
	SKU            string    `json:"sku"`
	Price          float64   `json:"price"`
	Availability   string    `json:"availability"`   // สถานะการใช้งาน เช่น 'active', 'inactive'
	Recommendation string    `json:"recommendation"` // สถานะแนะนำสินค้า เช่น 'recommended', 'normal'
	SellerID       string    `json:"seller_id"`
	ProductType    string    `json:"product_type"`
	CategoryID     int       `json:"category_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type NewProduct struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Brand          string          `json:"brand"`
	ModelNumber    string          `json:"model_number"`
	SKU            string          `json:"sku"`
	Price          float64         `json:"price"`
	Availability   string          `json:"availability"`   // สถานะการใช้งาน เช่น 'active', 'inactive'
	Recommendation string          `json:"recommendation"` // สถานะแนะนำสินค้า เช่น 'recommended', 'normal'
	SellerID       string          `json:"seller_id"`
	ProductType    string          `json:"product_type"`
	CategoryID     int             `json:"category_id"`
	Quantity       int             `json:"quantity"`
	Values         json.RawMessage `json:"values"`
	OptName        string          `json:"optname"`
}

type UpdateProduct struct {
	Price          float64 `json:"price"`
	Availability   string  `json:"availability"`   // สถานะการใช้งาน เช่น 'active', 'inactive'
	Recommendation string  `json:"recommendation"` // สถานะแนะนำสินค้า เช่น 'recommended', 'normal'
}

type ProductImage struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	ImageURL  string    `json:"image_url"`
	IsPrimary bool      `json:"is_primary"`
	AltText   string    `json:"alt_text"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
}

type NewProductImage struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
	AltText   string `json:"alt_text"`
	SortOrder int    `json:"sort_order"`
}

type UpdateProductImage struct {
	IsPrimary bool `json:"is_primary"`
	SortOrder int  `json:"sort_order"`
}

type ProductQueryParams struct {
	Cursor         string `json:"cursor"`
	Limit          int    `json:"limit"`
	Search         string `json:"search"`
	CategoryID     int    `json:"category_id"`
	SellerID       string `json:"seller_id"`
	Availability   string `json:"availability"`   // สถานะการใช้งาน เช่น 'active', 'inactive'
	Recommendation string `json:"recommendation"` // สถานะแนะนำสินค้า เช่น 'recommended', 'normal'
	ProductType    string `json:"product_type"`
	Sort           string `json:"sort"`
	Order          string `json:"order"`
}

type ProductResponse struct {
	Items      []ProductItem `json:"items"`
	NextCursor string        `json:"next_cursor"`
	Limit      int           `json:"limit"`
}

type ProductItem struct {
	Product
	Categories []Category      `json:"categories"`
	Inventory  Inventory       `json:"inventory"`
	Images     []ProductImage  `json:"images"`
	Options    []ProductOption `json:"options"`
}

type Category struct {
	ID   int    `json:"category_id"`
	Name string `json:"name"`
}

type CategoryWithProducts struct {
	Category Category      `json:"category"`
	Products []ProductItem `json:"products"`
}

type Inventory struct {
	Quantity  int       `json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Seller struct {
	SellerID    string    `json:"seller_id"`   // uuid
	Name        string    `json:"name"`        // character varying(255)
	CreatedAt   time.Time `json:"created_at"`  // timestamp with time zone, default CURRENT_TIME
	UpdatedAt   time.Time `json:"updated_at"`  // timestamp with time zone, default CURRENT_TIME
	Logo        string    `json:"logo"`        // character varying(255), optional
	Description string    `json:"description"` // text, optional
	Address     string    `json:"address"`     // character varying(255), optional
	Phone       string    `json:"phone"`       // character varying(20), optional
	Email       string    `json:"email"`       // character varying(255), optional
}

type ProductOption struct {
	ID      string          `json:"id"`
	OptName string          `json:"optname"`
	Values  json.RawMessage `json:"values"`
}

type NewFood struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"` // ใช้ float64 เพื่อให้ตรงกับ NUMERIC(10, 2)
}

type NewMedicine struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"` // ใช้ float64 เพื่อให้ตรงกับ NUMERIC(10, 2)
}

type NewToy struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"` // ใช้ float64 เพื่อให้ตรงกับ NUMERIC(10, 2)
}

type NewShelter struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"` // ใช้ float64 เพื่อให้ตรงกับ NUMERIC(10, 2)
}

type EcommerceDatabase interface {
	GetProduct(ctx context.Context, id string) (ProductItem, error)
	AddProduct(ctx context.Context, product NewProduct) (Product, error)
	UpdateProduct(ctx context.Context, id string, update UpdateProduct) (Product, error)
	DeleteProduct(ctx context.Context, id string) error
	GetProducts(ctx context.Context, params ProductQueryParams) (*ProductResponse, error)
	GetCategories(ctx context.Context) ([]CategoryWithProducts, error)
	GetProductImages(ctx context.Context, productID string) ([]ProductImage, error)
	AddProductImage(ctx context.Context, productID string, image NewProductImage) (ProductImage, error)
	UpdateProductImage(ctx context.Context, productID, imageID string, update UpdateProductImage) (ProductImage, error)
	DeleteProductImage(ctx context.Context, productID, imageID string) error
	GetShopDetail(ctx context.Context, sellerID string) (Seller, error)
	GetRecommendedProduct(ctx context.Context) ([]ProductItem, error)                   // เพิ่ม method ใหม่
	GetNewProductSeller(ctx context.Context) ([]ProductItem, error)                     // method ใหม่สำหรับสินค้าล่าสุดจากแต่ละ Seller
	GetAllProductImages(ctx context.Context) ([]ProductImage, error)                    // เพิ่ม method ใหม่สำหรับดึงรูปสินค้าทั้งหมด
	GetDetailProductSeller(ctx context.Context, sellerID string) ([]ProductItem, error) // New method for seller product details
	GetAllShops(ctx context.Context) ([]Seller, error)                                  // เพิ่ม method สำหรับดึงข้อมูลร้านค้าทั้งหมด
	Close() error
	Ping() error
	Reconnect(connStr string) error
}

type PostgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase(connStr string) (*PostgresDatabase, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &PostgresDatabase{db: db}, nil
}

func (pdb *PostgresDatabase) GetCategories(ctx context.Context) ([]CategoryWithProducts, error) {
	// สร้าง slice สำหรับหมวดหมู่และสินค้า
	var categoryWithProducts []CategoryWithProducts

	// ดึงข้อมูลหมวดหมู่
	categories, err := pdb.getCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %v", err)
	}

	for _, category := range categories {
		// ดึงสินค้าสำหรับแต่ละหมวดหมู่
		products, err := pdb.getProductsByCategory(ctx, category.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get products for category %d: %v", category.ID, err)
		}

		// สร้าง CategoryWithProducts struct และเพิ่มลงใน slice
		categoryWithProducts = append(categoryWithProducts, CategoryWithProducts{
			Category: category,
			Products: products,
		})
	}

	return categoryWithProducts, nil
}

func (pdb *PostgresDatabase) getProductsByCategory(ctx context.Context, categoryID int) ([]ProductItem, error) {
	var products []ProductItem

	// ปรับ query ให้รวมข้อมูลจากตาราง inventory และ categories
	rows, err := pdb.db.QueryContext(ctx, `
		SELECT 
			p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, 
			p.price, p.availability, p.recommendation, p.seller_id, p.product_type, p.created_at, p.updated_at,
			i.quantity,
			c.category_id, c.name as category_name
		FROM 
			products p
		LEFT JOIN 
			inventory i ON p.product_id = i.product_id
		LEFT JOIN 
			categories c ON p.category_id = c.category_id
		WHERE 
			p.category_id = $1`, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductItem
		var category Category

		// เพิ่มการสแกนข้อมูลที่เกี่ยวข้องกับหมวดหมู่
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.ModelNumber, &product.SKU, &product.Price, &product.Availability,
			&product.Recommendation, &product.SellerID, &product.ProductType, &product.CreatedAt, &product.UpdatedAt,
			&product.Inventory.Quantity, &category.ID, &category.Name); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}

		product.Categories = []Category{category} // เพิ่มข้อมูลหมวดหมู่ไปยังผลิตภัณฑ์

		// ดึงข้อมูลรูปภาพ
		product.Images, err = pdb.GetProductImages(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product images: %v", err)
		}

		// ดึงข้อมูลตัวเลือก
		product.Options, err = pdb.getProductOptions(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product options: %v", err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return products, nil
}

// ฟังก์ชันช่วยในการดึงข้อมูลหมวดหมู่
func (pdb *PostgresDatabase) getCategories(ctx context.Context) ([]Category, error) {
	categories := []Category{}

	rows, err := pdb.db.QueryContext(ctx, `SELECT category_id, name FROM categories`)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return categories, nil
}

func (pdb *PostgresDatabase) GetDetailProductSeller(ctx context.Context, sellerID string) ([]ProductItem, error) {
	var products []ProductItem

	// Query ดึงสินค้าทั้งหมดของผู้ขายโดยเรียงตามวันที่ล่าสุดลงไป
	rows, err := pdb.db.QueryContext(ctx, `
		SELECT p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price,
		       p.availability, p.recommendation, p.seller_id, p.product_type, p.created_at, p.updated_at,
		       c.category_id, c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.category_id
		WHERE p.seller_id = $1
		ORDER BY p.created_at DESC
	`, sellerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products for seller: %v", err)
	}
	defer rows.Close()

	// อ่านข้อมูลจาก rows
	for rows.Next() {
		var product ProductItem
		var category Category

		err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.ModelNumber, &product.SKU, &product.Price, &product.Availability,
			&product.Recommendation, &product.SellerID, &product.ProductType,
			&product.CreatedAt, &product.UpdatedAt,
			&category.ID, &category.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}

		// ใส่ข้อมูลหมวดหมู่
		product.Categories = []Category{category}

		// ดึงข้อมูล inventory
		err = pdb.db.QueryRowContext(ctx, `
			SELECT quantity, updated_at
			FROM inventory
			WHERE product_id = $1
		`, product.ID).Scan(&product.Inventory.Quantity, &product.Inventory.UpdatedAt)

		if err != nil && err != sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get inventory: %v", err)
		}

		// ดึงข้อมูลรูปภาพ
		product.Images, err = pdb.GetProductImages(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product images: %v", err)
		}

		// ดึงข้อมูลตัวเลือก
		product.Options, err = pdb.getProductOptions(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product options: %v", err)
		}

		// เพิ่มสินค้าที่อ่านได้ลงในรายการ
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return products, nil
}

func (pdb *PostgresDatabase) GetNewProductSeller(ctx context.Context) ([]ProductItem, error) {
	var products []ProductItem

	// Query ดึงสินค้าที่ใหม่ล่าสุดจากแต่ละ Seller
	rows, err := pdb.db.QueryContext(ctx, `
		SELECT DISTINCT ON (p.seller_id) p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price,
		       p.availability, p.recommendation, p.seller_id, p.product_type, p.created_at, p.updated_at,
		       c.category_id, c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.category_id
		WHERE p.availability = 'active'
		ORDER BY p.seller_id, p.created_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get new products from each seller: %v", err)
	}
	defer rows.Close()

	// อ่านข้อมูลจาก rows
	for rows.Next() {
		var product ProductItem
		var category Category

		err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.ModelNumber, &product.SKU, &product.Price, &product.Availability,
			&product.Recommendation, &product.SellerID, &product.ProductType,
			&product.CreatedAt, &product.UpdatedAt,
			&category.ID, &category.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}

		// ใส่ข้อมูลหมวดหมู่
		product.Categories = []Category{category}

		// ดึงข้อมูล inventory
		err = pdb.db.QueryRowContext(ctx, `
			SELECT quantity, updated_at
			FROM inventory
			WHERE product_id = $1
		`, product.ID).Scan(&product.Inventory.Quantity, &product.Inventory.UpdatedAt)

		if err != nil && err != sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get inventory: %v", err)
		}

		// ดึงข้อมูลรูปภาพ
		product.Images, err = pdb.GetProductImages(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product images: %v", err)
		}

		// ดึงข้อมูลตัวเลือก
		product.Options, err = pdb.getProductOptions(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product options: %v", err)
		}

		// เพิ่มสินค้าที่อ่านได้ลงในรายการ
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return products, nil
}

func (pdb *PostgresDatabase) GetShopDetail(ctx context.Context, sellerID string) (Seller, error) {
	var seller Seller

	// Query ดึงข้อมูลของร้านค้า (Seller)
	err := pdb.db.QueryRowContext(ctx, `
		SELECT seller_id, name, created_at, updated_at, logo, description, address, phone, email
		FROM sellers
		WHERE seller_id = $1
	`, sellerID).Scan(
		&seller.SellerID, &seller.Name, &seller.CreatedAt, &seller.UpdatedAt,
		&seller.Logo, &seller.Description, &seller.Address, &seller.Phone, &seller.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Seller{}, fmt.Errorf("shop not found")
		}
		return Seller{}, fmt.Errorf("failed to get shop details: %v", err)
	}

	return seller, nil
}

func (pdb *PostgresDatabase) GetAllShops(ctx context.Context) ([]Seller, error) {
	var sellers []Seller

	// Query ดึงข้อมูลของร้านค้าทั้งหมด
	rows, err := pdb.db.QueryContext(ctx, `
		SELECT seller_id, name, created_at, updated_at, logo, description, address, phone, email
		FROM sellers
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all shop details: %v", err)
	}
	defer rows.Close()

	// Loop ผ่านแต่ละ row และ map ข้อมูลเข้ากับ struct Seller
	for rows.Next() {
		var seller Seller
		err := rows.Scan(
			&seller.SellerID, &seller.Name, &seller.CreatedAt, &seller.UpdatedAt,
			&seller.Logo, &seller.Description, &seller.Address, &seller.Phone, &seller.Email,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan shop row: %v", err)
		}
		sellers = append(sellers, seller)
	}

	// ตรวจสอบ error ที่เกิดขึ้นระหว่างการ loop
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return sellers, nil
}

func (pdb *PostgresDatabase) GetRecommendedProduct(ctx context.Context) ([]ProductItem, error) {
	var products []ProductItem

	rows, err := pdb.db.QueryContext(ctx, `
		SELECT p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price,
		       p.availability, p.recommendation, p.seller_id, p.created_at, p.updated_at,
		       c.category_id, c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.category_id
		WHERE p.recommendation = 'recommended'
		LIMIT 3
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get recommended products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductItem
		var category Category

		err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.ModelNumber, &product.SKU, &product.Price, &product.Availability,
			&product.Recommendation, &product.SellerID,
			&product.CreatedAt, &product.UpdatedAt,
			&category.ID, &category.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}

		// เพิ่มดีบักเพื่อตรวจสอบค่า recommendation
		fmt.Printf("Product ID: %s, Name: %s, Recommendation: %s\n", product.ID, product.Name, product.Recommendation)

		// ใส่ข้อมูลหมวดหมู่ของสินค้า
		product.Categories = []Category{category}
		// ดึงข้อมูลรูปภาพ
		product.Images, err = pdb.GetProductImages(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product images: %v", err)
		}
		// ดึงข้อมูลตัวเลือก
		product.Options, err = pdb.getProductOptions(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product options: %v", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return products, nil
}

func (pdb *PostgresDatabase) GetProduct(ctx context.Context, id string) (ProductItem, error) {
	var product ProductItem
	var category Category

	// ดึงข้อมูลหลักของสินค้าและหมวดหมู่
	err := pdb.db.QueryRowContext(ctx, `
		SELECT p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price, 
		       p.availability, p.recommendation, p.seller_id, p.product_type, p.created_at, p.updated_at,
		       c.category_id, c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.category_id
		WHERE p.product_id = $1
	`, id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Brand,
		&product.ModelNumber, &product.SKU, &product.Price, &product.Availability,
		&product.Recommendation, &product.SellerID, &product.ProductType, &product.CreatedAt, &product.UpdatedAt,
		&category.ID, &category.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return ProductItem{}, fmt.Errorf("product not found")
		}
		return ProductItem{}, fmt.Errorf("failed to get product: %v", err)
	}

	product.Categories = []Category{category}

	// ดึงข้อมูล inventory
	err = pdb.db.QueryRowContext(ctx, `
		SELECT quantity, updated_at
		FROM inventory
		WHERE product_id = $1
	`, id).Scan(&product.Inventory.Quantity, &product.Inventory.UpdatedAt)

	if err != nil && err != sql.ErrNoRows {
		return ProductItem{}, fmt.Errorf("failed to get inventory: %v", err)
	}

	// ดึงข้อมูลรูปภาพ
	product.Images, err = pdb.GetProductImages(ctx, id)
	if err != nil {
		return ProductItem{}, fmt.Errorf("failed to get product images: %v", err)
	}

	// ดึงข้อมูลตัวเลือก
	product.Options, err = pdb.getProductOptions(ctx, id)
	if err != nil {
		return ProductItem{}, fmt.Errorf("failed to get product options: %v", err)
	}

	return product, nil
}

func (pdb *PostgresDatabase) AddProduct(ctx context.Context, product NewProduct) (Product, error) {
	var createdProduct Product

	// เพิ่มสินค้าในตาราง products
	err := pdb.db.QueryRowContext(ctx, `
		INSERT INTO products (name, description, brand, model_number, sku, price, availability, recommendation, seller_id, product_type, category_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
		RETURNING product_id, name, description, brand, model_number, sku, price, availability, recommendation, seller_id, product_type, category_id, created_at, updated_at
	`,
		product.Name, product.Description, product.Brand, product.ModelNumber, product.SKU, product.Price, product.Availability, product.Recommendation, product.SellerID, product.ProductType, product.CategoryID,
	).Scan(
		&createdProduct.ID, &createdProduct.Name, &createdProduct.Description, &createdProduct.Brand,
		&createdProduct.ModelNumber, &createdProduct.SKU, &createdProduct.Price, &createdProduct.Availability,
		&createdProduct.Recommendation, &createdProduct.SellerID, &createdProduct.ProductType, &createdProduct.CategoryID,
		&createdProduct.CreatedAt, &createdProduct.UpdatedAt)

	if err != nil {
		return Product{}, fmt.Errorf("failed to add product: %v", err)
	}

	// เพิ่มสินค้านั้นในตาราง inventory
	_, err = pdb.db.ExecContext(ctx, `
		INSERT INTO inventory (product_id, quantity) 
		VALUES ($1, $2)
	`, createdProduct.ID, product.Quantity) // ใช้ quantity ที่ได้รับจาก NewProduct

	if err != nil {
		return Product{}, fmt.Errorf("failed to add product to inventory: %v", err)
	}

	// เช็คประเภทของสินค้าที่เพิ่ม และเพิ่มเข้าไปในตารางที่เกี่ยวข้อง
	switch product.ProductType {
	case "food":
		_, err = pdb.db.ExecContext(ctx, `
			INSERT INTO foods (product_id, name, description, brand, price) 
			VALUES ($1, $2, $3, $4, $5)
		`, createdProduct.ID, createdProduct.Name, createdProduct.Description, createdProduct.Brand, createdProduct.Price)
	case "medicine":
		_, err = pdb.db.ExecContext(ctx, `
			INSERT INTO medicines (product_id, name, description, brand, price)
			VALUES ($1, $2, $3, $4, $5)
		`, createdProduct.ID, createdProduct.Name, createdProduct.Description, createdProduct.Brand, createdProduct.Price)
	case "toy":
		_, err = pdb.db.ExecContext(ctx, `
			INSERT INTO toys (product_id, name, description, brand, price)
			VALUES ($1, $2, $3, $4, $5)
		`, createdProduct.ID, createdProduct.Name, createdProduct.Description, createdProduct.Brand, createdProduct.Price)
	case "shelter":
		_, err = pdb.db.ExecContext(ctx, `
			INSERT INTO shelters (product_id, name, description, brand, price)
			VALUES ($1, $2, $3, $4, $5)
		`, createdProduct.ID, createdProduct.Name, createdProduct.Description, createdProduct.Brand, createdProduct.Price)
	}

	if err != nil {
		return Product{}, fmt.Errorf("failed to add product to related table: %v", err)
	}

	// เพิ่มข้อมูลในตาราง product_options ถ้ามี options ให้เพิ่ม
	_, err = pdb.db.ExecContext(ctx, `
		INSERT INTO product_options (product_id, optname, values)
		VALUES ($1, $2, $3)
	`, createdProduct.ID, product.OptName, product.Values)
	if err != nil {
		return Product{}, fmt.Errorf("failed to add product option: %v", err)
	}

	return createdProduct, nil
}

func (pdb *PostgresDatabase) UpdateProduct(ctx context.Context, id string, update UpdateProduct) (Product, error) {
	var updatedProduct Product
	err := pdb.db.QueryRowContext(ctx, `
		UPDATE products 
		SET price = $1, availability = $2, recommendation = $3, updated_at = NOW() 
		WHERE product_id = $4
		RETURNING product_id, name, description, brand, model_number, sku, price, availability, recommendation, seller_id, product_type, category_id, created_at, updated_at
	`,
		update.Price, update.Availability, update.Recommendation, id,
	).Scan(
		&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Description, &updatedProduct.Brand,
		&updatedProduct.ModelNumber, &updatedProduct.SKU, &updatedProduct.Price, &updatedProduct.Availability,
		&updatedProduct.Recommendation, &updatedProduct.SellerID, &updatedProduct.ProductType, &updatedProduct.CategoryID,
		&updatedProduct.CreatedAt, &updatedProduct.UpdatedAt)
	if err != nil {
		return Product{}, fmt.Errorf("failed to update product: %v", err)
	}
	return updatedProduct, nil
}

func (pdb *PostgresDatabase) DeleteProduct(ctx context.Context, id string) error {
	result, err := pdb.db.ExecContext(ctx, "DELETE FROM products WHERE product_id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

func (pdb *PostgresDatabase) GetProducts(ctx context.Context, params ProductQueryParams) (*ProductResponse, error) {
	query := `
        SELECT p.product_id, p.name, p.description, p.brand, p.model_number, p.sku, p.price, 
               p.availability, p.recommendation, p.seller_id, p.product_type, p.created_at, p.updated_at,
               c.category_id, c.name as category_name,
               i.quantity, i.updated_at as inventory_updated_at
        FROM products p
        LEFT JOIN categories c ON p.category_id = c.category_id
        LEFT JOIN inventory i ON p.product_id = i.product_id
        WHERE 1=1`

	args := []interface{}{}
	placeholderCount := 1

	// Handle cursor parameter
	if params.Cursor != "" {
		cursor, err := decodeCursor(params.Cursor)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor: %v", err)
		}
		query += fmt.Sprintf(" AND (p.created_at, p.product_id) > ($%d, $%d)", placeholderCount, placeholderCount+1)
		args = append(args, cursor.CreatedAt, cursor.ProductID)
		placeholderCount += 2
	}

	// Handle search parameter
	if params.Search != "" {
		// Remove whitespace from the search term
		searchTerm := strings.ReplaceAll(params.Search, " ", "")

		// Modify the query to ignore whitespace in the 'name' field
		query += fmt.Sprintf(" AND REPLACE(p.name, ' ', '') ILIKE $%d", placeholderCount)
		args = append(args, "%"+searchTerm+"%")
		placeholderCount++
	}

	// Handle category_id parameter
	if params.CategoryID != 0 {
		query += fmt.Sprintf(" AND p.category_id = $%d", placeholderCount)
		args = append(args, params.CategoryID)
		placeholderCount++
	}

	// Handle seller_id parameter
	if params.SellerID != "" {
		query += fmt.Sprintf(" AND p.seller_id = $%d", placeholderCount)
		args = append(args, params.SellerID)
		placeholderCount++
	}

	// Handle status parameter (availability)
	if params.Availability != "" {
		query += fmt.Sprintf(" AND p.availability = $%d", placeholderCount)
		args = append(args, params.Availability)
		placeholderCount++
	}

	if params.Recommendation != "" {
		query += fmt.Sprintf(" AND p.recommendation = $%d", placeholderCount)
		args = append(args, params.Recommendation)
		placeholderCount++
	}

	// Handle product_type parameter
	if params.ProductType != "" {
		query += fmt.Sprintf(" AND p.product_type = $%d", placeholderCount)
		args = append(args, params.ProductType)
		placeholderCount++
	}

	// Handle ORDER BY with sort and order
	sortFields := map[string]string{
		"name":       "p.name",
		"price":      "p.price",
		"created_at": "p.created_at",
	}

	orderDirections := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	sortColumn, ok := sortFields[params.Sort]
	if !ok {
		sortColumn = "p.price"
	}

	orderDirection, ok := orderDirections[strings.ToLower(params.Order)]
	if !ok {
		orderDirection = "ASC"
	}

	query += fmt.Sprintf(" ORDER BY %s %s, p.product_id ASC", sortColumn, orderDirection)

	// Set limit
	limit := 20
	if params.Limit > 0 && params.Limit <= 100 {
		limit = params.Limit
	}
	query += fmt.Sprintf(" LIMIT $%d", placeholderCount)
	args = append(args, limit+1)
	placeholderCount++

	// Execute query and process results
	rows, err := pdb.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	defer rows.Close()

	var products []ProductItem
	for rows.Next() {
		var product ProductItem
		var category Category
		var inventory Inventory
		if err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Brand,
			&product.ModelNumber, &product.SKU, &product.Price, &product.Availability,
			&product.Recommendation, &product.SellerID, &product.ProductType,
			&product.CreatedAt, &product.UpdatedAt,
			&category.ID, &category.Name,
			&inventory.Quantity, &inventory.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		product.Categories = []Category{category}
		product.Inventory = inventory

		// Fetch images and options for the product
		product.Images, err = pdb.GetProductImages(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product images: %v", err)
		}
		product.Options, err = pdb.getProductOptions(ctx, product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product options: %v", err)
		}

		products = append(products, product)
		if len(products) == limit+1 {
			break
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over products: %v", err)
	}

	response := &ProductResponse{
		Items: products[:min(len(products), limit)],
		Limit: limit,
	}

	if len(products) > limit {
		lastProduct := products[limit-1]
		response.NextCursor = encodeCursor(Cursor{CreatedAt: lastProduct.CreatedAt, ProductID: lastProduct.ID})
	} else {
		response.NextCursor = ""
	}

	return response, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Cursor struct {
	CreatedAt time.Time
	ProductID string
}

func (pdb *PostgresDatabase) GetAllProductImages(ctx context.Context) ([]ProductImage, error) {
	rows, err := pdb.db.QueryContext(ctx, `
		SELECT image_id, product_id, image_url, is_primary, sort_order, created_at 
		FROM product_images
		ORDER BY product_id ASC, sort_order ASC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all product images: %v", err)
	}
	defer rows.Close()

	var images []ProductImage
	for rows.Next() {
		var image ProductImage
		if err := rows.Scan(
			&image.ID, &image.ProductID, &image.ImageURL, &image.IsPrimary,
			&image.SortOrder, &image.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan product image: %v", err)
		}
		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over product images: %v", err)
	}

	return images, nil
}

func (pdb *PostgresDatabase) GetProductImages(ctx context.Context, productID string) ([]ProductImage, error) {
	rows, err := pdb.db.QueryContext(ctx, `
		SELECT image_id, product_id, image_url, is_primary, sort_order, created_at 
		FROM product_images 
		WHERE product_id = $1 
		ORDER BY sort_order ASC
	`, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product images: %v", err)
	}
	defer rows.Close()

	var images []ProductImage
	for rows.Next() {
		var image ProductImage
		if err := rows.Scan(
			&image.ID, &image.ProductID, &image.ImageURL, &image.IsPrimary,
			&image.SortOrder, &image.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan product image: %v", err)
		}
		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over product images: %v", err)
	}

	return images, nil
}

func (pdb *PostgresDatabase) AddProductImage(ctx context.Context, productID string, image NewProductImage) (ProductImage, error) {
	var createdImage ProductImage
	err := pdb.db.QueryRowContext(ctx, `
        INSERT INTO product_images (product_id, image_url, is_primary, sort_order, alt_text) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING image_id, product_id, image_url, is_primary, sort_order, alt_text, created_at
    `,
		productID, image.ImageURL, image.IsPrimary, image.SortOrder, image.AltText,
	).Scan(
		&createdImage.ID, &createdImage.ProductID, &createdImage.ImageURL,
		&createdImage.IsPrimary, &createdImage.SortOrder, &createdImage.AltText, &createdImage.CreatedAt)
	if err != nil {
		return ProductImage{}, fmt.Errorf("failed to add product image: %v", err)
	}
	return createdImage, nil
}

func (pdb *PostgresDatabase) UpdateProductImage(ctx context.Context, productID string, imageID string, update UpdateProductImage) (ProductImage, error) {
	var updatedImage ProductImage
	err := pdb.db.QueryRowContext(ctx, `
		UPDATE product_images 
		SET is_primary = $1, sort_order = $2, updated_at = NOW() 
		WHERE product_id = $3 AND image_id = $4 
		RETURNING image_id, product_id, image_url, is_primary, sort_order, created_at
	`,
		update.IsPrimary, update.SortOrder, productID, imageID,
	).Scan(
		&updatedImage.ID, &updatedImage.ProductID, &updatedImage.ImageURL,
		&updatedImage.IsPrimary, &updatedImage.SortOrder, &updatedImage.CreatedAt)
	if err != nil {
		return ProductImage{}, fmt.Errorf("failed to update product image: %v", err)
	}
	return updatedImage, nil
}

func (pdb *PostgresDatabase) DeleteProductImage(ctx context.Context, productID string, imageID string) error {
	result, err := pdb.db.ExecContext(ctx, `
		DELETE FROM product_images 
		WHERE product_id = $1 AND image_id = $2
	`, productID, imageID)
	if err != nil {
		return fmt.Errorf("failed to delete product image: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product image not found")
	}

	return nil
}

func (pdb *PostgresDatabase) Close() error {
	return pdb.db.Close()
}

func (pdb *PostgresDatabase) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return pdb.db.PingContext(ctx)
}

func (pdb *PostgresDatabase) getProductOptions(ctx context.Context, productID string) ([]ProductOption, error) {
	rows, err := pdb.db.QueryContext(ctx, `
        SELECT optname, values 
        FROM product_options 
        WHERE product_id = $1
    `, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product options: %v", err)
	}
	defer rows.Close()

	var options []ProductOption
	for rows.Next() {
		var option ProductOption
		if err := rows.Scan(&option.OptName, &option.Values); err != nil {
			return nil, fmt.Errorf("failed to scan product option: %v", err)
		}
		options = append(options, option)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over product options: %v", err)
	}

	return options, nil
}

func encodeCursor(c Cursor) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s,%s", c.CreatedAt.Format(time.RFC3339Nano), c.ProductID)))
}

func decodeCursor(s string) (Cursor, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return Cursor{}, err
	}
	parts := strings.Split(string(b), ",")
	if len(parts) != 2 {
		return Cursor{}, fmt.Errorf("invalid cursor format")
	}
	createdAt, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil {
		return Cursor{}, err
	}
	return Cursor{CreatedAt: createdAt, ProductID: parts[1]}, nil
}

func (pdb *PostgresDatabase) Reconnect(connStr string) error {
	if pdb.db != nil {
		pdb.db.Close()
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// ตั้งค่า connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	pdb.db = db
	return nil
}

type Store struct {
	db EcommerceDatabase
}

func NewStore(db EcommerceDatabase) *Store {
	return &Store{db: db}
}

func (s *Store) GetCategories(ctx context.Context) ([]CategoryWithProducts, error) {
	return s.db.GetCategories(ctx) // ส่งคืนประเภทที่ตรงกัน
}

func (s *Store) GetProduct(ctx context.Context, id string) (ProductItem, error) {
	return s.db.GetProduct(ctx, id)
}

func (s *Store) GetDetailProductSeller(ctx context.Context, sellerID string) ([]ProductItem, error) {
	return s.db.GetDetailProductSeller(ctx, sellerID)
}

func (s *Store) GetRecommendedProduct(ctx context.Context) ([]ProductItem, error) {
	return s.db.GetRecommendedProduct(ctx)
}

func (s *Store) GetShopDetail(ctx context.Context, sellerID string) (Seller, error) {
	return s.db.GetShopDetail(ctx, sellerID)
}

func (s *Store) GetAllShops(ctx context.Context) ([]Seller, error) {
	return s.db.GetAllShops(ctx)
}

func (s *Store) GetNewProductSeller(ctx context.Context) ([]ProductItem, error) {
	return s.db.GetNewProductSeller(ctx)
}

func (s *Store) AddProduct(ctx context.Context, product NewProduct) (Product, error) {
	return s.db.AddProduct(ctx, product)
}

func (s *Store) UpdateProduct(ctx context.Context, id string, update UpdateProduct) (Product, error) {
	return s.db.UpdateProduct(ctx, id, update)
}

func (s *Store) DeleteProduct(ctx context.Context, id string) error {
	return s.db.DeleteProduct(ctx, id)
}

func (s *Store) GetAllProductImages(ctx context.Context) ([]ProductImage, error) {
	return s.db.GetAllProductImages(ctx)
}

func (s *Store) GetProducts(ctx context.Context, params ProductQueryParams) (*ProductResponse, error) {
	return s.db.GetProducts(ctx, params)
}

func (s *Store) GetProductImages(ctx context.Context, productID string) ([]ProductImage, error) {
	return s.db.GetProductImages(ctx, productID)
}

func (s *Store) AddProductImage(ctx context.Context, productID string, image NewProductImage) (ProductImage, error) {
	return s.db.AddProductImage(ctx, productID, image)
}

func (s *Store) UpdateProductImage(ctx context.Context, productID string, imageID string, update UpdateProductImage) (ProductImage, error) {
	return s.db.UpdateProductImage(ctx, productID, imageID, update)
}

func (s *Store) DeleteProductImage(ctx context.Context, productID string, imageID string) error {
	return s.db.DeleteProductImage(ctx, productID, imageID)
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Ping() error {
	if s.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	return s.db.Ping()
}

func (s *Store) Reconnect(connStr string) error {
	return s.db.Reconnect(connStr)
}
