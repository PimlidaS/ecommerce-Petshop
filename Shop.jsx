// Shop.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const Shop = () => {
  const [shops, setShops] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios.get('/api/v1/shops')
      .then((response) => {
        setShops(response.data);
        setLoading(false);
      })
      .catch((err) => {
        setError('Error fetching shops');
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div style={{ padding: '20px', textAlign: 'center' }}>
      <h1>ร้านค้าทั้งหมด</h1>
      <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center', gap: '20px' }}>
        {shops.map((shop) => (
          <div key={shop.seller_id} style={{
            width: '200px',
            textAlign: 'center',
            padding: '10px',
            border: '1px solid #ddd',
            borderRadius: '8px',
            boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)'
          }}>
            <Link to={`/shop/${shop.seller_id}`}>
              <img src={shop.logo} alt={shop.name} style={{ width: '100%', height: 'auto', borderRadius: '8px' }} />
            </Link>
            <h2 style={{ marginTop: '10px', fontSize: '16px', fontWeight: 'bold', color: 'black' }}>{shop.name}</h2> {/* ปรับสีตัวหนังสือ */}

          </div>
        ))}
      </div>
    </div>
  );
};

export default Shop;
