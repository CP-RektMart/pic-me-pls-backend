Delete From Quotations;
Delete From Galleries;
Delete From Photographers;
Delete From Users;
Delete From Citizen_Cards;


ALTER SEQUENCE photographers_id_seq RESTART WITH 1;
ALTER SEQUENCE users_id_seq RESTART WITH 1;
ALTER SEQUENCE citizen_cards_id_seq RESTART WITH 1;
ALTER SEQUENCE galleries_id_seq RESTART WITH 1;
ALTER SEQUENCE quotations_id_seq RESTART WITH 1;

-- 1. Insert Users
INSERT INTO Users (Name, Email, Phone_Number, Profile_Picture_URL, Role)
VALUES
('User 1', 'user1@example.com', '1234567890', 'url1.jpg', 'CUSTOMER'),
('User 2', 'user2@example.com', '2345678901', 'url2.jpg', 'CUSTOMER'),
('User 3', 'user3@example.com', '3456789012', 'url3.jpg', 'PHOTOGRAPHER'),
('User 4', 'user4@example.com', '4567890123', 'url4.jpg', 'PHOTOGRAPHER'),
('User 5', 'user5@example.com', '5678901234', 'url5.jpg', 'PHOTOGRAPHER');

-- 2. Insert Photographers (only for Users with Role 'Photographer')
INSERT INTO Photographers (User_ID, Is_Verified, Active_Status)
VALUES
(3,  true, true),
(4,  true, true),
(5, false, false);

-- 3. Insert Citizen Cards for Photographers
INSERT INTO Citizen_Cards (Citizen_ID, Laser_ID, Picture, Expire_Date)
VALUES
('0123456789123', 'ME123456789', 'citizen3_pic.jpg', '2026-12-31'),
('0987654321234', 'ME123456111', 'citizen4_pic.jpg', '2027-11-30'),
('1029384756574', 'ME123456222', 'citizen5_pic.jpg', '2028-10-31');

-- 4. Link Citizen Cards to Photographers
UPDATE Photographers
SET Citizen_Card_ID = 1
WHERE User_ID = 3;

UPDATE Photographers
SET Citizen_Card_ID = 2
WHERE User_ID = 4;

UPDATE Photographers
SET Citizen_Card_ID = 3
WHERE User_ID = 5;

-- 5. Add Galleries
INSERT INTO galleries (photographer_id, name, description, price)
VALUES 
(1, 'Modern Art Gallery', 'A contemporary collection of modern abstract art pieces.', 1500.00),
(1, 'Classic Sculpture Gallery', 'A curated collection of classical sculptures from renowned artists.', 1800.00),
(1, 'Photography Wonders', 'A gallery featuring modern photography and street art.', 1200.00),
(1, 'Nature’s Beauty', 'A gallery focused on nature and wildlife photography.', 1600.00),
(2, 'Tech Art Revolution', 'A fusion of technology and art, showcasing digital masterpieces.', 2000.00),
(2, 'Architectural Wonders', 'Architectural photography displaying breathtaking urban landscapes.', 1700.00),
(3, 'Pop Culture Art', 'Pop art and contemporary media-inspired artwork.', 1300.00);

-- 5. Add Quotations
INSERT INTO quotations (gallery_id, photographer_id, customer_id, price, description, from_date, to_date, status)
VALUES (3, 1, 1,1800.50, 'Exclusive photoshoot session for a wedding.', '2025-07-05 10:00:00', '2025-07-07 18:00:00', 'PENDING'),
(1, 2, 2, 3200.00, 'Corporate event photography package for a company.', '2025-08-15 09:00:00', '2025-08-15 18:00:00', 'CONFIRMED');

-- Query Tables
select * from users;
select * from photographers;
select * from citizen_cards;
select * from galleries;
select * from quotations;

-- Query galleries with photographerID, userID, username informations
select g.name as Gallery_Name, g.description as Gallery_Description, g.price as Gallery_Price, p.id as Photographer_ID,
u.id as User_ID, u.name as userName from galleries as g
join photographers as p on g.photographer_id = p.id
join users as u on u.id = p.user_id;

-- Query quotations with gallery_name, photographer_username, customer_username
SELECT 
    q.id, 
    g.name AS gallery_name, 
    q.description, 
    q.price, 
    q.status, 
    q.photographer_id, 
    u2.name AS photographer_name, 
    q.customer_id, 
    u1.name AS customer_name
FROM quotations AS q
JOIN galleries AS g ON g.id = q.gallery_id
JOIN users AS u1 ON u1.id = q.customer_id
JOIN photographers AS p ON p.id = q.photographer_id
JOIN users AS u2 ON u2.id = p.user_id;




