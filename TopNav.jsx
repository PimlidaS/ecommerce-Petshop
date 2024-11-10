import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { FaSearch, FaShoppingCart, FaUser } from 'react-icons/fa';
import { Button, Modal, Dropdown } from 'react-bootstrap';
import GoogleAuth from './GoogleAuth';

const TopNav = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [showLogin, setShowLogin] = useState(false);
  const [user, setUser] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const storedUser = sessionStorage.getItem('user');
    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
  }, []);

  const handleSearchSubmit = (e) => {
    e.preventDefault();
    if (searchTerm) {
      navigate(`/search?q=${searchTerm}`);
    }
  };

  const handleShow = () => setShowLogin(true);
  const handleClose = () => setShowLogin(false);

  const handleLogout = () => {
    sessionStorage.removeItem('user');
    setUser(null);
    navigate('/');
  };

  return (
    <nav className="navbar" style={{ backgroundColor: '#ffedde' }}>
      <div className="container-fluid d-flex justify-content-between align-items-center">
        <Link className="navbar-brand d-flex align-items-center" to="/">
          <img src="/assets/images/FurPet_1.png" alt="Logo" width="150" height="100" />
          <h2 className="ms-3" style={{ color: '#FF7F00', fontWeight: 'bold', fontSize: '50px' }}>FurPet-Shop</h2>
        </Link>

        <form onSubmit={handleSearchSubmit} className="d-flex ms-auto me-4">
          <input
            className="form-control me-2"
            type="search"
            placeholder="ค้นหาสินค้า"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            aria-label="Search"
            style={{ width: '300px' }}
          />
          <button 
            type="submit" 
            style={{
              backgroundColor: '#f8f9fa', // พื้นหลังสีขาว
              border: 'none',
              borderRadius: '8px',
              padding: '8px',
              cursor: 'pointer',
              transition: 'background-color 0.3s ease' // Smooth transition เมื่อเปลี่ยนพื้นหลัง
            }}
            onMouseEnter={(e) => {
              e.target.style.backgroundColor = '#e0e0e0'; // พื้นหลังสีเทาเมื่อ hover
            }}
            onMouseLeave={(e) => {
              e.target.style.backgroundColor = '#f8f9fa'; // พื้นหลังกลับไปสีขาวเมื่อออกจาก hover
            }}
          >
            <FaSearch style={{ color: '#000000', fontSize: '20px' }} />
          </button>
        </form>

        <div className="d-flex align-items-center">
          {/* ปุ่มตะกร้าสินค้า */}
          <Link 
            to="/cart"
            style={{
              display: 'flex',
              alignItems: 'center',
              backgroundColor: '#f8f9fa',
              borderRadius: '12px',
              padding: '8px 12px',
              marginRight: '10px',
              textDecoration: 'none',
              color: '#000',
              transition: 'background-color 0.3s ease' // Smooth transition
            }}
            onMouseEnter={(e) => {
              e.currentTarget.style.backgroundColor = '#e0e0e0'; // เปลี่ยนเป็นสีเทาเมื่อ hover
            }}
            onMouseLeave={(e) => {
              e.currentTarget.style.backgroundColor = '#f8f9fa'; // กลับเป็นสีขาวเมื่อออกจาก hover
            }}
          >
            <FaShoppingCart style={{ fontSize: '20px', color: '#000' }} />
            <span 
              style={{
                backgroundColor: '#dc3545',
                color: '#fff',
                borderRadius: '50%',
                padding: '2px 6px',
                fontSize: '12px',
                marginLeft: '8px'
              }}
            >
              2
            </span>
          </Link>

          {/* ปุ่มเข้าสู่ระบบ */}
          {user ? (
            <Dropdown>
              <Dropdown.Toggle 
                variant="light"
                id="dropdown-basic"
                className="d-flex align-items-center"
                style={{
                  backgroundColor: '#f8f9fa',
                  borderRadius: '12px',
                  padding: '8px 12px',
                  transition: 'background-color 0.3s ease', // Smooth transition
                  border: 'none' // เอาเส้นขอบออก
                }}
                onMouseEnter={(e) => {
                  e.currentTarget.style.backgroundColor = '#e0e0e0'; // เปลี่ยนเป็นสีเทาเมื่อ hover
                }}
                onMouseLeave={(e) => {
                  e.currentTarget.style.backgroundColor = '#f8f9fa'; // กลับเป็นสีขาวเมื่อออกจาก hover
                }}
              >
                <img
                  src={user.picture}
                  alt="user-profile"
                  style={{ borderRadius: '50%', width: '24px', marginRight: '8px' }}
                  referrerPolicy="no-referrer"
                />
              </Dropdown.Toggle>

              <Dropdown.Menu>
                <Dropdown.Item as={Link} to="/profile">My Profile</Dropdown.Item>
                <Dropdown.Item onClick={handleLogout}>Logout</Dropdown.Item>
              </Dropdown.Menu>
            </Dropdown>
          ) : (
            <Button 
              variant="light"
              onClick={handleShow}
              style={{
                display: 'flex',
                alignItems: 'center',
                backgroundColor: '#f8f9fa',
                borderRadius: '12px',
                padding: '8px 12px',
                color: '#000',
                transition: 'background-color 0.3s ease', // Smooth transition
                border: 'none' // เอาเส้นขอบออก
              }}
              onMouseEnter={(e) => {
                e.currentTarget.style.backgroundColor = '#e0e0e0'; // เปลี่ยนเป็นสีเทาเมื่อ hover
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.backgroundColor = '#f8f9fa'; // กลับเป็นสีขาวเมื่อออกจาก hover
              }}
            >
              <FaUser style={{ fontSize: '20px', marginRight: '8px' }} /> เข้าสู่ระบบ
            </Button>
          )}
        </div>
      </div>

      <Modal show={showLogin} onHide={handleClose} centered>
        <Modal.Header closeButton>
          <Modal.Title>เข้าสู่ระบบ</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <GoogleAuth
            setUser={(user) => {
              setUser(user);
              sessionStorage.setItem('user', JSON.stringify(user));
            }}
            handleClose={handleClose}
          />
        </Modal.Body>
      </Modal>
    </nav>
  );
};

export default TopNav;
