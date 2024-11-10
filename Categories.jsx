import React, { useState } from 'react'; // ลบ useEffect
// import { useParams } from 'react-router-dom'; // ลบ useParams
// import { useNavigate } from 'react-router-dom'; // ลบ useNavigate
import TopNav from '../components/TopNav';
import TopMenu from '../components/TopMenu';
import Footer from '../components/Footer';
import { Container, Row, Col, Card } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { Button } from '@mui/material';

const Categories = () => {
  const categories = [
    { id: 5, name: 'อาหารแมว' },
    { id: 6, name: 'ของเล่นแมว' },
    { id: 7, name: 'ยาสำหรับแมว' },
    { id: 8, name: 'บ้านแมว' }
  ];

  const [selectedCategory, setSelectedCategory] = useState(null);
  const [categoryProducts, setCategoryProducts] = useState([]);

  const handleCategoryClick = async (categoryId) => {
    setSelectedCategory(categoryId);
    try {
      const response = await fetch(`/api/v1/categories`);
      const data = await response.json();
      const categoryData = data.find(cat => cat.category.category_id === categoryId);
      if (categoryData) {
        setCategoryProducts(categoryData.products || []);
      } else {
        setCategoryProducts([]);
      }
    } catch (error) {
      console.error('Error fetching category products:', error);
    }
  };

  return (
    <div>
      <TopNav />
      <TopMenu />
      <div style={{ padding: '20px', textAlign: 'center' }}>
        <h1>หมวดหมู่สินค้า</h1>
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center', gap: '20px' }}>
          {categories.map((category) => (
            <div key={category.id} style={{ textAlign: 'center' }}>
              <Button
                variant="contained"
                onClick={() => handleCategoryClick(category.id)}
                style={{
                  padding: '10px 20px',
                  margin: '10px',
                  fontSize: '16px',
                  fontWeight: 'bold',
                  textTransform: 'none',
                  minWidth: '150px',
                  backgroundColor: 'orange',
                  color: 'white',
                  borderRadius: '30px',
                  transition: 'background-color 0.3s ease',
                }}
              >
                {category.name}
              </Button>
            </div>
          ))}
        </div>

        {selectedCategory && ( 
        <div style={{ marginTop: '40px' }}>
          <h2>สินค้าในหมวดหมู่: {categories.find(cat => cat.id === selectedCategory)?.name}</h2>
          <Container>
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
                            <Card.Body style={{ textAlign: 'left' }}>
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
        </div>
        )}
      </div>
      <Footer />
    </div>
  );
};

export default Categories;
