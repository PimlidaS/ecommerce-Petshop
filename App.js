import React from "react"; 
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { GoogleOAuthProvider } from "@react-oauth/google";
import 'bootstrap/dist/js/bootstrap.bundle.min';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from "./pages/Home";
import SearchResults from "./pages/SearchResults";
import ProductDetail from "./pages/ProductDetail";
// import PrivacyPolicy from './pages/PrivacyPolicy'; // นำเข้า Privacy Policy
// import TermsOfService from './pages/TermsOfService'; // นำเข้า Terms of Service
import Profile from "./pages/Profile"; // นำเข้า Profile page
import Category from "./components/Category";
import ShopDetail2 from './pages/ShopDetail2';
import Shop from "./components/Shop";
import AllShop from './pages/AllShop';
import Categories from "./pages/Categories"; // หน้าที่แสดงหมวดหมู่ทั้งหมด

function App() {
  return (
    <GoogleOAuthProvider clientId="618419592763-j24l6q083madahfa8rfg2orkv705fr0c.apps.googleusercontent.com">
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/search" element={<SearchResults />} />
          <Route path="/product/:productId" element={<ProductDetail />} />
          {/*<Route path="/privacy-policy" element={<PrivacyPolicy />} />  เส้นทาง Privacy Policy */}
          {/*<Route path="/terms-of-service" element={<TermsOfService />} />  เส้นทาง Terms of Service */}
          <Route path="/category/:category_id" element={<Category />} />
           {/* เส้นทาง Profile */}
          <Route path="/profile" element={<Profile />} />{" "}
          <Route path="/shops" element={<Shop />} />
          <Route path="/shop/:shopId" element={<ShopDetail2 />} />
          <Route path="/store" element={<AllShop />} /> {/* หน้าแสดงร้านค้า */}
          <Route path="/categories" element={<Categories />} /> {/* หน้าหมวดหมู่ทั้งหมด */}
        </Routes>
      </Router>
    </GoogleOAuthProvider>
  );
}

export default App;
