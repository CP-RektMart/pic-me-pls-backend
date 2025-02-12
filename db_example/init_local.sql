Delete From Photographers;
Delete From Users;
Delete From Citizen_Cards;

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

select * from users;
select * from photographers;
select * from citizen_cards;


