import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useSearchParams, Link } from 'react-router-dom';
import { Container, Row, Col, Card, Button } from 'react-bootstrap';
import TopNav from '../components/TopNav';
import TopMenu from '../components/TopMenu';
import Footer from '../components/Footer';

const placeholderImage = '../assets/images/placeholder.jpg';

const SearchResults = () => {
  const [products, setProducts] = useState([]);
  const [searchParams] = useSearchParams();
  const [error, setError] = useState(null); // ใช้สำหรับเก็บข้อผิดพลาด

  const searchQuery = searchParams.get('q');

  // ฟังก์ชันดึงข้อมูลสินค้าจาก API
  const fetchProducts = (query) => {
    if (query) {
      axios
        .get(`/api/v1/products?search=${query}&limit=20`)
        .then((response) => {
          if (response.data.items) {
            setProducts(response.data.items); // ตั้งค่า products ด้วยข้อมูลจาก API
            setError(null); // ล้าง error หากมีข้อมูล
          } else if (response.data.error) {
            setError(response.data.error); // เก็บข้อความข้อผิดพลาด
            setProducts([]); // ล้างรายการ products
          }
        })
        .catch((error) => {
          console.error('Error fetching the products:', error);
          setError('ไม่สามารถดึงข้อมูลสินค้ามาได้'); // เก็บข้อความข้อผิดพลาด
        });
    } else {
      setProducts([]);
    }
  };

  useEffect(() => {
    fetchProducts(searchQuery);
  }, [searchQuery]);

  return (
    <div>
      <TopNav />
      <TopMenu />

      <Container className="my-5">
        <h2 className="text-center mb-4">ผลการค้นหาสำหรับ "{searchQuery}"</h2>
        <Row>
          {error ? (
            <p className="text-center">{error}</p> // แสดงข้อความข้อผิดพลาด
          ) : products.length > 0 ? (
            products.map((product) => {
              const primaryImage = product.images.find((img) => img.is_primary)?.image_url || placeholderImage;

              return (
                <Col md={12} key={product.id}>
                  <Card className="mb-4 p-3" style={{ border: '1px solid #e5e5e5' }}>
                    <Row>
                      <Col md={3}>
                        <Card.Img
                          src={primaryImage}
                          onError={(e) => {
                            e.target.onerror = null;
                            e.target.src = placeholderImage;
                          }}
                          style={{ maxHeight: '150px', objectFit: 'contain' }}
                        />
                      </Col>
                      <Col md={6}>
                        <Card.Body>
                          <Card.Title style={{ fontSize: '1.5rem', color: '#007bff' }}>
                            {product.name}
                          </Card.Title>
                          <div style={{ fontSize: '1.2rem', color: '#ff6600' }}>
                            <strong>{product.price ? `฿${product.price}` : 'ราคาไม่ระบุ'}</strong>
                          </div>
                          <Card.Text>
                            {product.description ? product.description : 'ไม่มีคำอธิบายสินค้า'}
                          </Card.Text>
                          {/* <div className="d-flex align-items-center">
                            <span className="badge bg-success me-2">New</span>
                            {product.price && (
                              <span className="badge bg-danger me-2">Hot</span>
                            )}
                          </div> */}
                        </Card.Body>
                      </Col>
                      <Col md={3} className="d-flex align-items-center justify-content-center">
                        <div>
                          <Button style={{backgroundColor: '#e0a85e', border: '2px solid #e0a85e'}} className="me-2">
                            เพิ่มลงในตะกร้า
                          </Button>
                          <Link to={`/product/${product.id}`}>
                            <Button variant="outline-secondary">
                              ดูรายละเอียด
                            </Button>
                          </Link>
                        </div>
                      </Col>
                    </Row>
                  </Card>
                </Col>
              );
            })
          ) : (
            <p className="text-center">ไม่พบสินค้าที่ค้นหา</p>
          )}
        </Row>
      </Container>
      <Footer />
    </div>
  );
};

export default SearchResults;