import React, { useEffect, useState } from 'react';   
import axios from 'axios';
import { useParams, Link } from 'react-router-dom';
import { Container, Row, Col, Card, Button } from 'react-bootstrap';
import TopNav from '../components/TopNav';
import TopMenu from '../components/TopMenu';
import Footer from '../components/Footer';

const ShopDetail2 = () => {
  const { shopId } = useParams();
  const [shop, setShop] = useState(null);
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const placeholderImage = '/path/to/placeholder-image.jpg';

  useEffect(() => {
    axios.get(`/api/v1/shops/${shopId}`)
      .then((response) => {
        setShop(response.data);
        setLoading(false);
      })
      .catch((err) => {
        setError('Error fetching shop details');
        setLoading(false);
      });

    axios.get(`/api/v1/products/seller/${shopId}`)
      .then((response) => {
        setProducts(response.data);
      })
      .catch((err) => {
        setError('Error fetching products');
        setLoading(false);
      });
  }, [shopId]);

  if (loading) {
    return <div style={{ textAlign: 'center', marginTop: '50px' }}>Loading...</div>;
  }

  if (error) {
    return <div style={{ color: 'red', textAlign: 'center', marginTop: '50px' }}>{error}</div>;
  }

  return (
    <div>
      <TopNav />
      <TopMenu />

      {/* ส่วนรายละเอียดร้านค้า */}
      <Container className="my-5">
        <Row>
          <Col md={6} style={{ display: 'flex', justifyContent: 'center' }}>
            <Card style={{ boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.1)', borderRadius: '10px', width: '300px', height: '300px' }}>
              <Card.Img
                variant="top"
                src={shop.logo || placeholderImage}
                alt={shop.name}
                style={{
                  borderRadius: '10px',
                  objectFit: 'contain',
                  width: '100%',
                  height: '100%'
                }}
              />
            </Card>
          </Col>

          <Col md={6} style={{ display: 'flex', flexDirection: 'column', justifyContent: 'center' }}>
            <h2 style={{ fontWeight: 'bold' }}>{shop.name}</h2>
            <p><strong>รายละเอียด:</strong> {shop.description}</p>
            <p><strong>ที่อยู่:</strong> {shop.address}</p>
            <p><strong>เบอร์โทร:</strong> {shop.phone}</p>
            <p><strong>อีเมล:</strong> {shop.email}</p>
            <Button 
              variant="primary" 
              className="mt-3" 
              style={{ backgroundColor: '#e0a85e', border: '2px solid #e0a85e' }}
            >
              ติดต่อร้านค้า
            </Button>
          </Col>
        </Row>
      </Container>

      {/* ส่วนรายการสินค้าที่ขายในร้าน */}
      <Container className="my-5">
        <h3 style={{ fontWeight: 'bold', marginBottom: '20px' }}>สินค้าที่มีในร้าน</h3>
        <Row>
          {products.length > 0 ? (
            products.map((product) => (
              <Col md={4} key={product.id} className="mb-4 d-flex">
                <Card style={{ boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.1)', borderRadius: '10px', width: '100%', display: 'flex', flexDirection: 'column', justifyContent: 'space-between', minHeight: '100%' }}>
                  <Card.Img
                    variant="top"
                    src={product.images.find((img) => img.is_primary)?.image_url || placeholderImage}
                    alt={product.name}
                    style={{ borderRadius: '10px 10px 0 0', objectFit: 'cover', maxHeight: '300px' }}
                    onError={(e) => {
                      e.target.onerror = null;
                      e.target.src = placeholderImage;
                    }}
                  />
                  <Card.Body style={{ flex: '1' }}>
                    <Card.Title style={{ fontWeight: 'bold' }}>{product.name}</Card.Title>
                    <Card.Text><strong>ราคา:</strong> {product.price ? `฿${product.price}` : 'ราคาไม่ระบุ'}</Card.Text>
                    <Card.Text style={{ fontSize: '14px', color: '#555' }}><strong>รายละเอียด:</strong> {product.description}</Card.Text>
                    <Card.Text><strong>แบรนด์:</strong> {product.brand}</Card.Text>
                    <Card.Text><strong>สต็อก:</strong> {product.inventory?.quantity || 'ไม่ระบุ'}</Card.Text>
                  </Card.Body>
                  <Card.Footer style={{ backgroundColor: 'transparent', borderTop: 'none' }}>
                    <Link to={`/product/${product.id}`}>
                      <Button style={{backgroundColor: '#e0a85e', border: '2px solid #e0a85e'}} className="me-2">
                        ดูรายละเอียดสินค้า
                      </Button>
                    </Link>
                  </Card.Footer>
                </Card>
              </Col>
            ))
          ) : (
            <p style={{ textAlign: 'center', width: '100%', marginTop: '20px' }}>ไม่พบสินค้าที่ขายในร้าน</p>
          )}
        </Row>
      </Container>

      <Footer />
    </div>
  );
};

export default ShopDetail2;
