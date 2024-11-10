package handlers

import (
	"encoding/base64"
	"log"
	"net/http"
	product "productproject/internal/product"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	store *product.Store
}

func NewProductHandlers(store *product.Store) *ProductHandlers {
	return &ProductHandlers{store: store}
}

func convertTimesToUserTimezone(product *product.ProductItem, loc *time.Location) {
	product.CreatedAt = product.CreatedAt.In(loc)
	product.UpdatedAt = product.UpdatedAt.In(loc)
	product.Inventory.UpdatedAt = product.Inventory.UpdatedAt.In(loc)

	for j := range product.Images {
		product.Images[j].CreatedAt = product.Images[j].CreatedAt.In(loc)
	}
}

func (h *ProductHandlers) GetProducts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	cursor := c.Query("cursor")
	var decodedCursor string
	if cursor != "" {
		decodedCursor, err = decodeCursor(cursor)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor"})
			return
		}
	}

	categoryIDStr := c.Query("category")
	var categoryID int
	if categoryIDStr != "" {
		categoryID, err = strconv.Atoi(categoryIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}
	}

	params := product.ProductQueryParams{
		Cursor:         decodedCursor,
		Limit:          limit,
		Search:         c.Query("search"),
		CategoryID:     categoryID,
		SellerID:       c.Query("seller_id"),
		Availability:   c.Query("availability"),   // replaces Status to match table
		Recommendation: c.Query("recommendation"), // new filter for product recommendation status
		ProductType:    c.Query("product_type"),
		Sort:           c.Query("sort"),
		Order:          c.Query("order"),
	}

	response, err := h.store.GetProducts(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode the NextCursor before sending the response
	if response.NextCursor != "" {
		response.NextCursor = encodeCursor(response.NextCursor)
	}

	userTimezone := "Asia/Bangkok"
	loc, err := time.LoadLocation(userTimezone)
	if err != nil {
		log.Fatal("ไม่สามารถโหลด timezone ได้:", err)
	}

	for i := range response.Items {
		convertTimesToUserTimezone(&response.Items[i], loc)
	}

	c.JSON(http.StatusOK, response)
}

func encodeCursor(cursor string) string {
	return base64.StdEncoding.EncodeToString([]byte(cursor))
}

func decodeCursor(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (h *ProductHandlers) GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.store.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userTimezone := "Asia/Bangkok"
	loc, err := time.LoadLocation(userTimezone)
	if err != nil {
		log.Fatal("ไม่สามารถโหลด timezone ได้:", err)
	}

	convertTimesToUserTimezone(&product, loc)

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandlers) AddProduct(c *gin.Context) {
	var product product.NewProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.store.AddProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandlers) GetAllProductImages(c *gin.Context) {
	images, err := h.store.GetAllProductImages(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, images)
}

func (h *ProductHandlers) GetProductImages(c *gin.Context) {
	id := c.Param("id") // รับค่า id เป็น string

	images, err := h.store.GetProductImages(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, images)
}

func (h *ProductHandlers) AddProductImage(c *gin.Context) {
	id := c.Param("id") // รับค่า id เป็น string

	var image product.NewProductImage
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdImage, err := h.store.AddProductImage(c.Request.Context(), id, image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdImage)
}

func (h *ProductHandlers) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var update product.UpdateProduct
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.store.UpdateProduct(c.Request.Context(), id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandlers) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := h.store.DeleteProduct(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

func (h *ProductHandlers) UpdateProductImage(c *gin.Context) {
	id := c.Param("id")            // รับค่า id เป็น string
	imageID := c.Param("image_id") // รับค่า imageID เป็น string

	var update product.UpdateProductImage
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedImage, err := h.store.UpdateProductImage(c.Request.Context(), id, imageID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedImage)
}

func (h *ProductHandlers) DeleteProductImage(c *gin.Context) {
	id := c.Param("id")            // รับค่า id เป็น string
	imageID := c.Param("image_id") // รับค่า imageID เป็น string

	if err := h.store.DeleteProductImage(c.Request.Context(), id, imageID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ProductImage deleted successfully"})
}

func (h *ProductHandlers) GetDetailProductSeller(c *gin.Context) {
	// รับ sellerID จากพารามิเตอร์ URL
	sellerID := c.Param("sellerID")

	// เรียกใช้ GetDetailProductSeller จาก store
	products, err := h.store.GetDetailProductSeller(c.Request.Context(), sellerID)
	if err != nil {
		// ถ้ามีข้อผิดพลาดเกิดขึ้นให้ส่ง error กลับ
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่าไม่มีสินค้าสำหรับผู้ขายนี้
	if len(products) == 0 {
		// ถ้าไม่มีสินค้า ส่งข้อความว่าไม่มีสินค้า
		c.JSON(http.StatusOK, gin.H{"message": "No products available for this seller"})
		return
	}

	// ส่งข้อมูลสินค้าของผู้ขายในรูปแบบ JSON
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandlers) GetRecommendedProduct(c *gin.Context) {
	// เรียกใช้ GetRecommendedProduct จาก store
	products, err := h.store.GetRecommendedProduct(c.Request.Context())
	if err != nil {
		// ถ้ามีข้อผิดพลาดเกิดขึ้นให้ส่ง error กลับ
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่าไม่มีสินค้าแนะนำ
	if len(products) == 0 {
		// ถ้าไม่มีสินค้าแนะนำ ส่งข้อความว่าไม่มีสินค้าแนะนำ
		c.JSON(http.StatusOK, gin.H{"message": "No recommended products available"})
		return
	}

	// ส่งข้อมูลสินค้าที่แนะนำ 3 รายการกลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandlers) GetShopDetail(c *gin.Context) {
	// รับค่า sellerID จากพารามิเตอร์ใน URL
	sellerID := c.Param("id")

	// เรียกใช้ GetShopDetail จาก store
	shopDetail, err := h.store.GetShopDetail(c.Request.Context(), sellerID)
	if err != nil {
		// ถ้ามีข้อผิดพลาดเกิดขึ้นให้ส่ง error กลับ
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่าร้านค้านี้ไม่มีข้อมูล
	if shopDetail.SellerID == "" {
		// ถ้าไม่มีข้อมูลร้านค้า ส่งข้อความว่าร้านค้าไม่พบ
		c.JSON(http.StatusNotFound, gin.H{"message": "Shop not found"})
		return
	}

	// ส่งข้อมูลร้านค้ากลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, shopDetail)
}

func (h *ProductHandlers) GetAllShops(c *gin.Context) {
	// เรียกใช้ GetAllShops จาก store
	shops, err := h.store.GetAllShops(c.Request.Context())
	if err != nil {
		// ถ้ามีข้อผิดพลาดเกิดขึ้นให้ส่ง error กลับ
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่าไม่มีร้านค้า
	if len(shops) == 0 {
		// ถ้าไม่มีร้านค้า ส่งข้อความว่าไม่มีข้อมูลร้านค้า
		c.JSON(http.StatusNotFound, gin.H{"message": "No shops available"})
		return
	}

	// ส่งข้อมูลร้านค้าทั้งหมดกลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, shops)
}

func (h *ProductHandlers) GetNewProductSeller(c *gin.Context) {
	// เรียกใช้ GetNewProductSeller จาก store
	products, err := h.store.GetNewProductSeller(c.Request.Context())
	if err != nil {
		// ถ้ามีข้อผิดพลาดเกิดขึ้นให้ส่ง error กลับ
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่าไม่มีสินค้าใหม่จากแต่ละ Seller
	if len(products) == 0 {
		// ถ้าไม่มีสินค้าใหม่จากแต่ละ Seller ส่งข้อความว่าไม่มีสินค้า
		c.JSON(http.StatusOK, gin.H{"message": "No new products available from sellers"})
		return
	}

	// ส่งข้อมูลสินค้าล่าสุดจากแต่ละ Seller กลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandlers) GetCategories(c *gin.Context) {
	// ดึงรายการหมวดหมู่พร้อมสินค้าจากฐานข้อมูล
	categoryWithProducts, err := h.store.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// สร้าง slice เพื่อเก็บหมวดหมู่และสินค้าที่เกี่ยวข้อง
	var categories []product.CategoryWithProducts
	for _, categoryWithProduct := range categoryWithProducts {
		categories = append(categories, categoryWithProduct)
	}

	// ส่งข้อมูลหมวดหมู่พร้อมสินค้ากลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, categories)
}

func (h *ProductHandlers) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
