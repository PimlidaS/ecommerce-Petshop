import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { Container, Row, Col, Card, Button } from 'react-bootstrap';
import TopNav from '../components/TopNav';
import TopMenu from '../components/TopMenu';
import Footer from '../components/Footer';

const placeholderImage = '../assets/images/placeholder.jpg';

const ProductDetail = () => {
  const { productId } = useParams(); // รับ productId จาก URL
  const [product, setProduct] = useState(null);

  // ฟังก์ชันดึงข้อมูลรายละเอียดสินค้าจาก API
  const fetchProductDetails = (id) => {
    axios
      .get(`/api/v1/products/${id}`)
      .then((response) => {
        setProduct(response.data);
      })
      .catch((error) => {
        console.error('Error fetching product details:', error);
      });
  };

  useEffect(() => {
    fetchProductDetails(productId);
  }, [productId]);

  if (!product) {
    return <div>Loading...</div>; // แสดง Loading ระหว่างดึงข้อมูล
  }

  return (
    <div>
      <TopNav />
      <TopMenu />

      <Container className="my-5">
        <Row>
          <Col md={6}>
            {/* แสดงรูปภาพสินค้า */}
            <Card.Img
              src={product.images.find((img) => img.is_primary)?.image_url || placeholderImage}
              onError={(e) => {
                e.target.onerror = null;
                e.target.src = placeholderImage;
              }}
              style={{ maxHeight: '400px', objectFit: 'contain' }}
            />
          </Col>
          <Col md={6}>
            <h2>{product.name}</h2>
            <p><strong>ราคา:</strong> {product.price ? `฿${product.price}` : 'ราคาไม่ระบุ'}</p>
            <p><strong>รายละเอียดสินค้า:</strong> {product.description}</p>
            <p><strong>แบรนด์:</strong> {product.brand}</p>
            <p><strong>สต็อก:</strong> {product.inventory.quantity}</p>
            <Button style={{backgroundColor: '#e0a85e', border: '2px solid #e0a85e'}} className="me-2">เพิ่มในรถเข็น</Button>
          </Col>
        </Row>
      </Container>

      <Footer />
    </div>
  );
};

export default ProductDetail;