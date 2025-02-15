DELETE FROM Photographers;
DELETE FROM Users;
DELETE FROM Citizen_Cards;

ALTER SEQUENCE photographers_id_seq RESTART WITH 1;
ALTER SEQUENCE users_id_seq RESTART WITH 1;
ALTER SEQUENCE citizen_cards_id_seq RESTART WITH 1;

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
(5, false, false),
(6, true, true),
(7, false, true),
(8, true, false),
(9, true, true),
(10, false, false),
(11, true, true),
(12, false, true);

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
UPDATE Photographers SET Citizen_Card_ID = 1 WHERE User_ID = 3;
UPDATE Photographers SET Citizen_Card_ID = 2 WHERE User_ID = 4;
UPDATE Photographers SET Citizen_Card_ID = 3 WHERE User_ID = 5;
UPDATE Photographers SET Citizen_Card_ID = 4 WHERE User_ID = 6;
UPDATE Photographers SET Citizen_Card_ID = 5 WHERE User_ID = 7;
UPDATE Photographers SET Citizen_Card_ID = 6 WHERE User_ID = 8;
UPDATE Photographers SET Citizen_Card_ID = 7 WHERE User_ID = 9;
UPDATE Photographers SET Citizen_Card_ID = 8 WHERE User_ID = 10;
UPDATE Photographers SET Citizen_Card_ID = 9 WHERE User_ID = 11;
UPDATE Photographers SET Citizen_Card_ID = 10 WHERE User_ID = 12;

SELECT * FROM Users;
SELECT * FROM Photographers;
SELECT * FROM Citizen_Cards;