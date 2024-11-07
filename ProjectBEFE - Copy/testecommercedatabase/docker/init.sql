-- ตรวจสอบว่า database มีอยู่แล้วหรือไม่
SELECT 'CREATE DATABASE ecommerce'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'ecommerce');

-- เชื่อมต่อกับ Database ที่สร้าง
\c ecommerce

-- สร้าง Extension สำหรับ UUID (เฉพาะ PostgreSQL)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- สร้าง ENUM สำหรับ status และ product_type
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_status') THEN
        CREATE TYPE product_status AS ENUM ('active', 'inactive');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_type') THEN
        CREATE TYPE product_type AS ENUM ('food', 'toy', 'medicine', 'shelter');
    END IF;

END$$;

-- เริ่ม Transaction
BEGIN;

-- สร้างตาราง categories
CREATE TABLE IF NOT EXISTS categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    parent_category_id INTEGER,
    FOREIGN KEY (parent_category_id) REFERENCES categories(category_id) ON DELETE SET NULL
);

-- สร้างตาราง sellers
CREATE TABLE IF NOT EXISTS sellers (
    seller_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    contact_info VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


-- สร้างตาราง products
CREATE TABLE IF NOT EXISTS products (
    product_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    brand VARCHAR(255),
    model_number VARCHAR(100), --เก็บข้อมูลหมายเลขรุ่น เช่น ยี่ห้ออาหารแมว หมา 
    sku VARCHAR(100) UNIQUE NOT NULL, --รหัสเฉพาะการระบุสินค้า
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    status product_status NOT NULL,
    seller_id UUID NOT NULL,
    category_id INTEGER NOT NULL,
    -- pet_type pet_type NOT NULL, -- เช่น 'cat', 'dog', 'bird', 'fish', 'rodent', 'rabbit'
    product_type product_type NOT NULL, -- เช่น 'food', 'toy', 'medicine', 'shelter'
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (seller_id) REFERENCES sellers(seller_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE
);

-- สร้างตาราง foods
CREATE TABLE IF NOT EXISTS foods (
    food_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL, -- ชื่อของอาหาร
    description TEXT, -- รายละเอียดอาหาร
    brand VARCHAR(255), -- แบรนด์ของอาหาร
    price NUMERIC(10, 2) NOT NULL, -- ราคาของอาหาร
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);
-- สร้างตาราง toys
CREATE TABLE IF NOT EXISTS toys (
    toy_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL, -- ชื่อของอาหาร
    description TEXT, -- รายละเอียดอาหาร
    brand VARCHAR(255), -- แบรนด์ของอาหาร
    price NUMERIC(10, 2) NOT NULL, -- ราคาของอาหาร
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);
-- สร้างตาราง medicines
CREATE TABLE IF NOT EXISTS medicines (
    medicine_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL, -- ชื่อของอาหาร
    description TEXT, -- รายละเอียดอาหาร
    brand VARCHAR(255), -- แบรนด์ของอาหาร
    price NUMERIC(10, 2) NOT NULL, -- ราคาของอาหาร
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);
-- สร้างตาราง shelter(บ้าน)
CREATE TABLE IF NOT EXISTS shelters (
    shelters_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL, -- ชื่อของอาหาร
    description TEXT, -- รายละเอียดอาหาร
    brand VARCHAR(255), -- แบรนด์ของอาหาร
    price NUMERIC(10, 2) NOT NULL, -- ราคาของอาหาร
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง product_images
CREATE TABLE IF NOT EXISTS product_images (
    image_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    alt_text VARCHAR(255),
    is_primary BOOLEAN DEFAULT FALSE,
    sort_order INTEGER DEFAULT 0,
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง product_options
CREATE TABLE IF NOT EXISTS product_options (
    option_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    values JSONB NOT NULL, -- ใช้ JSONB สำหรับเก็บค่าตัวเลือก
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง inventory
CREATE TABLE IF NOT EXISTS inventory (
    product_id UUID PRIMARY KEY,
    quantity INTEGER NOT NULL CHECK (quantity >= 0),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างฟังก์ชันสำหรับอัปเดตฟิลด์ updated_at อัตโนมัติ
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC';
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- สร้าง Trigger สำหรับตารางที่มีฟิลด์ updated_at
CREATE TRIGGER update_products_updated_at BEFORE UPDATE ON products
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_sellers_updated_at BEFORE UPDATE ON sellers
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_product_images_updated_at BEFORE UPDATE ON product_images
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_product_options_updated_at BEFORE UPDATE ON product_options
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_inventory_updated_at BEFORE UPDATE ON inventory
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();


-- เพิ่มข้อมูลในตาราง categories (หมวดหมู่หลัก)
INSERT INTO categories (name, description, parent_category_id)
VALUES 
    ('อาหารสัตว์', 'หมวดหมู่สำหรับอาหารสัตว์เลี้ยง', NULL),
    ('ของเล่นสัตว์เลี้ยง', 'หมวดหมู่สำหรับของเล่นที่เหมาะกับสัตว์เลี้ยง', NULL),
    ('ยาสัตว์', 'หมวดหมู่สำหรับยาและอาหารเสริมสำหรับสัตว์เลี้ยง', NULL),
    ('ที่อยู่สัตว์', 'หมวดหมู่สำหรับบ้านและที่พักสำหรับสัตว์เลี้ยง', NULL);

-- เพิ่มข้อมูลในตาราง categories (หมวดหมู่ย่อยสำหรับสัตว์แต่ละประเภทในหมวดหมู่หลัก)
-- หมวดหมู่ย่อยของ "อาหารสัตว์"
INSERT INTO categories (name, description, parent_category_id)
VALUES 
    ('อาหารแมว', 'หมวดหมู่สำหรับอาหารที่เหมาะสมกับแมว', 1);

-- หมวดหมู่ย่อยของ "ของเล่นสัตว์เลี้ยง"
INSERT INTO categories (name, description, parent_category_id)
VALUES 
    ('ของเล่นแมว', 'หมวดหมู่สำหรับของเล่นที่เหมาะกับแมว', 2);

-- หมวดหมู่ย่อยของ "ยาสัตว์"
INSERT INTO categories (name, description, parent_category_id)
VALUES 
    ('ยาสำหรับแมว', 'หมวดหมู่สำหรับยาและอาหารเสริมสำหรับแมว', 3);

-- หมวดหมู่ย่อยของ "ที่อยู่สัตว์"
INSERT INTO categories (name, description, parent_category_id)
VALUES 
    ('บ้านแมว', 'หมวดหมู่สำหรับบ้านและที่พักสำหรับแมว', 4);

-- เพิ่มข้อมูลในตาราง sellers จำนวน 10 คน โดยระบุ seller_id เอง
INSERT INTO sellers (seller_id, name, contact_info) 
VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'บริษัท KitCat จำกัด', 'contact@KitCat.com'),            --อาหาร 2 ของเล่น 1 ยา 3 บ้าน 2
    ('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'ศูนย์สัตว์เลี้ยง Pet Center', 'support@petcenter.com'),     --อาหาร 1 ของเล่น 2 ยา 3 บ้าน 1
    ('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 'ร้านค้าออนไลน์ Pet Shop', 'info@petshop.com'),           -- อาหาร 4 ของเล่น 1 ยา 1 บ้าน 1
    ('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a44', 'ร้านค้า Happy Paws', 'contact@happypawsclinic.com'),    -- อาหาร 3 ของเล่น 2 ยา 1 บ้าน 2
    ('e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a55', 'บริษัท สัตว์เลี้ยงสุขภาพดี', 'info@healthyPets.com')      -- อาหาร 2 ของเล่น 3 บ้าน 2 บ้าน 1
    ;
-- foods


-- ปิด Transaction
COMMIT;

-- สร้าง ENUM สำหรับ user_status และ user_role
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
        CREATE TYPE user_status AS ENUM ('active', 'inactive', 'suspended');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        CREATE TYPE user_role AS ENUM ('customer', 'seller', 'admin');
    END IF;
END$$;

-- สร้างตาราง users
CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    google_id VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    profile_picture_url VARCHAR(255),
    email_verified BOOLEAN DEFAULT FALSE,
    status user_status NOT NULL DEFAULT 'active',
    role user_role NOT NULL DEFAULT 'customer',
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง user_sessions
CREATE TABLE IF NOT EXISTS user_sessions (
    session_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    id_token TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- สร้างตาราง user_login_history
CREATE TABLE IF NOT EXISTS user_login_history (
    login_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    login_timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    ip_address INET,
    user_agent TEXT,
    success BOOLEAN NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- สร้างตาราง api_keys
CREATE TABLE IF NOT EXISTS api_keys (
    key_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    api_key VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_sessions_updated_at
BEFORE UPDATE ON user_sessions
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_api_keys_updated_at
BEFORE UPDATE ON api_keys
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- สร้าง Indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_google_id ON users(google_id);
CREATE INDEX idx_user_sessions_user_id ON user_sessions(user_id);
CREATE INDEX idx_user_login_history_user_id ON user_login_history(user_id);
CREATE INDEX idx_api_keys_api_key ON api_keys(api_key);