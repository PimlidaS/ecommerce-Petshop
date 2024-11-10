import React, { useEffect, useState } from 'react';  
import { useParams } from 'react-router-dom';
import TopNav from '../components/TopNav';
import TopMenu from '../components/TopMenu';
import Footer from '../components/Footer';
import { Container, Row, Col, Card } from 'react-bootstrap';
import { Link } from 'react-router-dom';

const Category = () => {
  const { category_id } = useParams();
  const [categoryProducts, setCategoryProducts] = useState([]);
  const [categoryName, setCategoryName] = useState('');

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const response = await fetch(`/api/v1/categories`);
        const data = await response.json();
        
        // หาหมวดหมู่ที่ตรงกับ category_id ที่ได้รับ
        const category = data.find(cat => cat.category.category_id === parseInt(category_id));
        if (category) {
          setCategoryName(category.category.name);
          setCategoryProducts(category.products || []); // เก็บสินค้าของหมวดหมู่นั้น
        }
      } catch (error) {
        console.error('Error fetching categories:', error);
      }
    };

    fetchCategories();
  }, [category_id]);

  return (
    <div>
      <TopNav />
      <TopMenu />
      <Container className="my-5">
        {/* แสดงชื่อหมวดหมู่พร้อมจำนวนสินค้าในวงเล็บ */}
        <h1>{categoryName} ({categoryProducts.length} รายการ)</h1>

        {categoryProducts.length > 0 ? (
          <Row>
            {categoryProducts.map((product) => (
              <Col md={12} key={product.id} className="mb-4">
                <Card className="d-flex flex-row">
                  
                  <Link to={`/product/${product.id}`} style={{ textDecoration: 'none', color: 'inherit' }} className="d-flex w-100">
                    
                    {/* รูปภาพอยู่ทางด้านซ้าย */}
                    <Col md={4}>
                      <Card.Img
                        variant="top"
                        src={product.images.length > 0 ? product.images[0].image_url : 'https://example.com/path/to/placeholder-image.jpg'}
                        onError={(e) => {
                          e.target.onerror = null; 
                          e.target.src = 'https://example.com/path/to/placeholder-image.jpg'; // เปลี่ยน URL เป็นที่ถูกต้อง
                        }}
                      />
                    </Col>

                    {/* ข้อมูลอยู่ทางด้านขวา */}
                    <Col md={8} className="d-flex flex-column justify-content-between">
                      <Card.Body>
                        <Card.Title>{product.name}</Card.Title>
                        <Card.Text>
                          <strong>ราคา:</strong> ฿{product.price}<br />
                          <strong>รายละเอียด:</strong> {product.description}<br />
                          <strong>แบรนด์:</strong> {product.brand}<br />
                          <strong>สต็อก:</strong> {product.inventory.quantity}
                        </Card.Text>
                      </Card.Body>
                    </Col>
                  </Link>
                </Card>
              </Col>
            ))}
          </Row>
        ) : (
          <p>ไม่มีสินค้าในหมวดหมู่นี้</p>
        )}
      </Container>

      <Footer />
    </div>
  );
};

export default Category;
