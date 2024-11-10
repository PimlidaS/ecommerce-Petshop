import React from 'react'; 
import { Link } from 'react-router-dom';

const TopMenu = () => {
  const categories = [
    { id: 5, name: 'อาหารแมว' },
    { id: 6, name: 'ของเล่นแมว' },
    { id: 7, name: 'ยาสำหรับแมว' },
    { id: 8, name: 'บ้านแมว' }
  ];

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-white shadow-sm">
      <div className="container-fluid">
        {/* ลิงก์หน้าแรก */}
        <Link className="navbar-brand fw-bold" to="/">
          หน้าแรก
        </Link>

        {/* ลิงก์ร้านค้า */}
        <Link className="navbar-brand fw-bold" to="/store">
          ร้านค้า
        </Link>

        {/* ลิงก์หมวดหมู่สินค้า */}
        <Link className="navbar-brand fw-bold" to="/categories">
          หมวดหมู่สินค้า
        </Link>

        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon" />
        </button>

        <div className="collapse navbar-collapse" id="navbarSupportedContent">
          <ul className="navbar-nav ms-auto">
            {categories.map(category => (
              <li className="nav-item" key={category.id}>
                <Link className="nav-link text-dark fw-semibold" to={`/category/${category.id}`}>
                  {category.name}
                </Link>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default TopMenu;
