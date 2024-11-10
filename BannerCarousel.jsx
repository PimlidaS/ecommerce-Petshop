// src/components/BannerCarousel.jsx
import React from 'react';
import { Carousel } from 'react-bootstrap';

const BannerCarousel = () => {
  return (
    <Carousel>
      <Carousel.Item>
        <img
          className="d-block w-100"
          src="/assets/images/Banner_Condo_Catplay_0.png"
          alt="First slide"
        />
        <Carousel.Caption>
          <h3>"คอนโดแมวขนาดใหญ่ คอนโดแมวไม้ ลับเล็บแมว หลุม อวกาศ"</h3>
          <p>"คอนโดแมวทำจากไม้คุณภาพสูง ไร้เหลี่ยมคมและไม่มีกลิ่นฟอร์มาลดีไฮด์ ประกอบง่าย แข็งแรงทนทาน มีแคปซูลใสขนาด 30 ซม. เพื่อสังเกตท่าทางแมว พร้อมรังและหอสังเกตการณ์เพื่อให้แมวพักผ่อนสะดวกสบาย เหมาะสำหรับแมวหลายตัว"</p>
        </Carousel.Caption>
      </Carousel.Item>
      <Carousel.Item>
        <img
          className="d-block w-100"
          src="/assets/images/Banner_KittenFood_PurinaOne_0.png"
          alt="Second slide"
        />
        <Carousel.Caption>
          <h3>"Purina One อาหารแมวชนิดเม็ด สูตรลูกแมวอายุ 3 สัปดาห์ถึง 1 ปี"</h3>
          <p>"ผลิตจากวัตถุดิบคุณภาพสูง มี DHA ใช้เนื้อไก่แท้ๆ และข้าว ทำให้มีโปรตีนสูงถึง 40% ย่อยง่าย อุดมไปด้วยสารอาหารที่ครบถ้วน เหมาะสำหรับลูกแมวที่กำลังเจริญเติบโตทุกสายพันธุ์ รสชาติอร่อยและกลิ่นหอม"</p>
        </Carousel.Caption>
      </Carousel.Item>
      <Carousel.Item>
        <img
          className="d-block w-100"
          src="/assets/images/Banner_Toy_Catplay_0.png"
          alt="Third slide"
        />
        <Carousel.Caption>
          <h3>"ที่ลับเล็บแมว ลายไม้"</h3>
          <p>"ที่ลับเล็บสำหรับแมว ช่วยให้น้องแมวได้ขีดข่วนและตะไบเล็บ ช่วยลดปัญหาแมวขีดข่วนทำลายเฟอร์นิเจอร์ และสร้างความสนุกสนาน เพลิดเพลินให้กับแมว ช่วยคลายเหงาและลดความเบื่อหน่ายเวลาคุณไม่อยู่บ้าน"</p>
        </Carousel.Caption>
      </Carousel.Item>
    </Carousel>
  );
};

export default BannerCarousel;