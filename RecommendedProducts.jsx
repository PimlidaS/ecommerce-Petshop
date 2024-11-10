import React, { useState, useEffect } from 'react';  
import axios from 'axios';
import { Card, Container, Row, Col } from 'react-bootstrap';
import { Link } from 'react-router-dom';

const RecommendedProducts = () => {
  const [products, setProducts] = useState([]);

  // Placeholder image URL for missing product images
  const placeholderImage = '/assets/images/placeholder.jpg';

  useEffect(() => {
    axios
      .get('/api/v1/products/Recommendproducts')
      .then((response) => {
        console.log('Fetched products:', response.data); // ตรวจสอบข้อมูลที่ได้รับ
        setProducts(response.data); // อัปเดต state
      })
      .catch((error) => {
        console.error('Error fetching the products:', error);
      });
  }, []);

  return (
    <Container className="my-5">
      <h2 className="text-center mb-4">รายการสินค้าแนะนำ</h2>
      <Row>
        {products.length > 0 ? (
          products.map((product) => {
            const primaryImage = product.images && product.images.length > 0
              ? product.images.find((img) => img.is_primary)?.image_url
              : placeholderImage;

            return (
              <Col md={4} key={product.id}>
                <Card className="mb-4">
                  <Link to={`/product/${product.id}`} style={{ textDecoration: 'none', color: 'inherit' }}>
                    <Card.Img 
                      variant="top" 
                      src={primaryImage} 
                      alt={product.name} 
                      onError={(e) => {
                        e.target.onerror = null;
                        e.target.src = placeholderImage;
                      }} 
                    />
                    <Card.Body>
                      <Card.Title>{product.name}</Card.Title>
                      <Card.Text>{product.description}</Card.Text>
                      <Card.Text>
                        <strong>Price: ฿{product.price}</strong>
                      </Card.Text>
                    </Card.Body>
                  </Link>
                </Card>
              </Col>
            );
          })
        ) : (
          <p className="text-center">ไม่พบสินค้าที่ค้นหา</p>
        )}
      </Row>
    </Container>
  );
};

export default RecommendedProducts;