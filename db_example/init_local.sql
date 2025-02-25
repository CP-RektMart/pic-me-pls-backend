-- Reset the database for testing
DELETE FROM Quotations;
DELETE FROM Reviews;
DELETE FROM Media;
DELETE FROM Tags;
DELETE FROM Packages_Categories;
DELETE FROM Packages;
DELETE FROM Photographers;
DELETE FROM Users;
DELETE FROM Citizen_Cards;
DELETE FROM Categories;

ALTER SEQUENCE Packages_id_seq RESTART WITH 1;
ALTER SEQUENCE photographers_id_seq RESTART WITH 1;
ALTER SEQUENCE users_id_seq RESTART WITH 1;
ALTER SEQUENCE citizen_cards_id_seq RESTART WITH 1;
ALTER SEQUENCE tags_id_seq RESTART WITH 1;
ALTER SEQUENCE media_id_seq RESTART WITH 1;
ALTER SEQUENCE reviews_id_seq RESTART WITH 1;
ALTER SEQUENCE categories_id_seq RESTART WITH 1;
ALTER SEQUENCE quotations_id_seq RESTART WITH 1;

-- 1. Insert Users
INSERT INTO Users (Name, Email, Phone_Number, Profile_Picture_URL, Role)
VALUES
('User 1', 'user1@example.com', '1234567890', 'url1.jpg', 'Customer'),
('User 2', 'user2@example.com', '2345678901', 'url2.jpg', 'Customer'),
('User 3', 'user3@example.com', '3456789012', 'url3.jpg', 'Photographer'),
('User 4', 'user4@example.com', '4567890123', 'url4.jpg', 'Photographer'),
('User 5', 'user5@example.com', '5678901234', 'url5.jpg', 'Photographer'),
('User 6', 'user6@example.com', '6789012345', 'url6.jpg', 'Photographer'),
('User 7', 'user7@example.com', '7890123456', 'url7.jpg', 'Photographer'),
('User 8', 'user8@example.com', '8901234567', 'url8.jpg', 'Photographer'),
('User 9', 'user9@example.com', '9012345678', 'url9.jpg', 'Photographer'),
('User 10', 'user10@example.com', '1123456789', 'url10.jpg', 'Photographer'),
('User 11', 'user11@example.com', '2234567890', 'url11.jpg', 'Photographer'),
('User 12', 'user12@example.com', '3345678901', 'url12.jpg', 'Photographer');

-- 2. Insert Photographers (only for Users with Role 'Photographer')
INSERT INTO Photographers (User_ID, Is_Verified, Active_Status)
VALUES
(3,  true, true),
(4,  true, true),
(5, false, false);

-- 3. Insert Citizen Cards for Photographers
INSERT INTO Citizen_Cards (Citizen_ID, Laser_ID, Picture, Expire_Date)
VALUES
('1519999567819', 'LASER123', 'citizen3_pic.jpg', '2026-12-31'),
('4819999567819', 'LASER234', 'citizen4_pic.jpg', '2027-11-30'),
('3319999567819', 'LASER345', 'citizen5_pic.jpg', '2028-10-31'),
('2219999567819', 'LASER456', 'citizen6_pic.jpg', '2029-09-30'),
('7719999567819', 'LASER567', 'citizen7_pic.jpg', '2030-08-31'),
('8819999567819', 'LASER678', 'citizen8_pic.jpg', '2031-07-31'),
('9919999567819', 'LASER789', 'citizen9_pic.jpg', '2032-06-30'),
('6619999567819', 'LASER890', 'citizen10_pic.jpg', '2033-05-31'),
('5519999567819', 'LASER901', 'citizen11_pic.jpg', '2034-04-30'),
('4419999567819', 'LASER012', 'citizen12_pic.jpg', '2035-03-31');

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

-- 5. Insert Packages
INSERT INTO Packages (Photographer_ID, Name, Description, Price)
VALUES
(1, 'Sunset Photography', 'A collection of sunset images', 150.00),
(2, 'Nature Wonders', 'Beautiful natural scenery', 200.00),
(3, 'Urban Life', 'Capturing the essence of the city', 180.00);

-- 6. Insert Tags
INSERT INTO Tags (Package_ID, Name)
VALUES
(1, 'Sunset'),
(1, 'Landscape'),
(2, 'Nature'),
(3, 'City');

-- 7. Insert Media
INSERT INTO Media (Package_ID, Picture_URL)
VALUES
(1, 'https://example.com/sunset1.jpg'),
(1, 'https://example.com/sunset2.jpg'),
(2, 'https://example.com/nature1.jpg'),
(3, 'https://example.com/city1.jpg');

-- 8. Insert Reviews
INSERT INTO Reviews (Package_ID, Customer_ID, Rating, Comment)
VALUES
(1, 1, 4.5, 'Amazing!'),
(2, 2, 5.0, 'Beautiful shots'),
(3, 1, 4.0, 'Great urban photography');

-- 9. Insert Categories
INSERT INTO Categories (Name, Description)
VALUES
('Nature', 'Packages related to nature and landscapes'),
('City Life', 'Packages capturing city life');

-- 10. Link Packages to Categories
INSERT INTO Packages_Categories (Package_ID, Category_ID)
VALUES
(1, 1),
(2, 1),
(3, 2);

-- 11. Insert Quotations
INSERT INTO Quotations (Package_ID, Customer_ID, Photographer_ID, Status, Price, Description, from_date, to_date)
INSERT INTO Quotations (Package_ID, Customer_ID, Photographer_ID, Status, Price, Description, from_date, to_date)
VALUES
(1, 1, 1, 'pending', 200.00, 'Quotation for sunset photography', '2026-12-31', '2026-12-31'),
(2, 2, 2, 'confirmed', 250.00, 'Nature photo session', '2026-12-31', '2026-12-31'),
(3, 1, 3, 'paid', 300.00, 'City life photo package', '2026-12-31', '2026-12-31');

-- Verify the data
SELECT * FROM Users;
SELECT * FROM Photographers;
SELECT * FROM Citizen_Cards;
SELECT * FROM Packages;
SELECT * FROM Tags;
SELECT * FROM Media;
SELECT * FROM Reviews;
SELECT * FROM Categories;
SELECT * FROM Packages_Categories;
SELECT * FROM Quotations;

-- Query packages including package details, photographerID, userID, username 
select p.name as Package_Name, p.description as Package_Description, p.price as Package_Price, ph.id as Photographer_ID,
u.id as User_ID, u.name as userName from Packages as p
join photographers as ph on p.photographer_id = ph.id
join users as u on u.id = ph.user_id;

-- -- Query quotations including quotations, photographer_username, customer_username
SELECT 
    q.id as quotation_id, 
    p.name AS package_name, 
    q.description, 
    q.price, 
    q.status, 
    q.photographer_id, 
    u2.name AS photographer_name, 
    q.customer_id, 
    u1.name AS customer_name
FROM quotations AS q
JOIN packages AS p ON p.id = q.package_id
JOIN users AS u1 ON u1.id = q.customer_id
JOIN photographers AS ph ON ph.id = q.photographer_id
JOIN users AS u2 ON u2.id = ph.user_id;
