import React from 'react';
import TopNav from '../components/TopNav'; // ส่วนแสดง Navigation Bar
import TopMenu from '../components/TopMenu'; // ส่วนแสดงเมนูด้านบน
import BannerCarousel from '../components/BannerCarousel'; // ส่วนแสดง Banner
import NewProducts from '../components/NewProducts'; // ส่วนแสดงสินค้ามาใหม่
import RecommendedProducts from '../components/RecommendedProducts';
import Footer from '../components/Footer'; // ส่วนแสดง Footer
// import Category from '../components/Category';
import Shop from '../components/Shop';


const Home = () => {
  return (
    <div>
      <TopNav /> {/* ส่วนค้นหาจะอยู่ใน TopNav */}
      <TopMenu />
      <BannerCarousel />
      <Shop />
      {/* <Category /> */}
      <RecommendedProducts />
      <NewProducts /> {/* แสดงสินค้ามาใหม่ */}
      <Footer />
    </div>
  );
};

export default Home;
