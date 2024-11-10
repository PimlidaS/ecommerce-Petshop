import React, { useEffect, useState } from 'react'; 
import { useNavigate } from 'react-router-dom';
import { Figure } from 'react-bootstrap';

const CategoryFigure = () => {
  const [categories, setCategories] = useState([]);
  const navigate = useNavigate();

  // ดึงข้อมูลหมวดหมู่จาก API
  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const response = await fetch('/api/v1/categories');
        const data = await response.json();
        setCategories(data); // เก็บข้อมูลหมวดหมู่ใน state
      } catch (error) {
        console.error('Error fetching categories:', error);
      }
    };

    fetchCategories();
  }, []);

  const handleCategoryClick = (categoryId) => {
    navigate(`/category/${categoryId}`); // นำทางไปยังหน้าหมวดหมู่ที่เลือก
  };

  return (
    <section className="categories-section">
      <h4 className="section-title">หมวดหมู่สินค้า</h4>
      <div className="category-list d-flex justify-content-around flex-wrap">
        {categories.length > 0 ? (
          categories.map((category) => (
            <Figure
              key={category.category_id}
              className="category-item text-center"
              style={{ cursor: 'pointer' }}
              onClick={() => handleCategoryClick(category.category_id)}
            >
              <Figure.Image
                width={150}
                height={150}
                alt={`รูปภาพสำหรับ ${category.name}`}
                src={category.image || '/assets/images/placeholder.jpg'} // ใช้ placeholder หากไม่พบรูปภาพ
                onError={(e) => {
                  e.target.onerror = null;
                  e.target.src = '/assets/images/placeholder.jpg'; // ใช้ placeholder หากรูปโหลดไม่สำเร็จ
                }}
              />
              <Figure.Caption>
                <div className="category-name">{category.name}</div>
              </Figure.Caption>
            </Figure>
          ))
        ) : (
          <p>ไม่มีหมวดหมู่สินค้า</p>
        )}
      </div>
    </section>
  );
};

export default CategoryFigure;
