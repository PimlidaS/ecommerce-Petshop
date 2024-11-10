// src/components/Footer.jsx
import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';

const Footer = () => {
  return (
    <footer className="text-black py-4" style={{ backgroundColor: '#ffedde' }}>
      <Container>
        <Row>
          <Col md={4}>
            <h5>FurPet</h5>
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Auctor libero id, in gravida.</p>
          </Col>
          <Col md={2}>
            <h5>About Us</h5>
            <ul className="list-unstyled">
              <li>Careers</li>
              <li>Our Stores</li>
              <li>Terms & Conditions</li>
              <li>Privacy Policy</li>
            </ul>
          </Col>
          <Col md={2}>
            <h5>Customer Care</h5>
            <ul className="list-unstyled">
              <li>Help Center</li>
              <li>How to Buy</li>
              <li>Track Your Order</li>
              <li>Returns & Refunds</li>
            </ul>
          </Col>
          <Col md={4}>
            <h5>Contact Us</h5>
            <p>
              70 Washington Square South, New York, NY 10012, United States <br/>
              Email: <a href="mailto:me-mart@memartmail.com" className="text-brown">FurPet@gmail.com</a> <br/>
              Phone: +1 123-456-3583
            </p>
          </Col>
        </Row>
      </Container>
    </footer>
  );
};

export default Footer;
