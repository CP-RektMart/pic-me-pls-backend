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
('User 5', 'user5@example.com', '5678901234', 'url5.jpg', 'Photographer');

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
('3319999567819', 'LASER345', 'citizen5_pic.jpg', '2028-10-31');

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
INSERT INTO Quotations (Package_ID, Customer_ID, Photographer_ID, Status, Price, Description)
VALUES
(1, 1, 1, 'pending', 200.00, 'Quotation for sunset photography'),
(2, 2, 2, 'confirmed', 250.00, 'Nature photo session'),
(3, 1, 3, 'paid', 300.00, 'City life photo package');

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
