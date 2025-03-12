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

ALTER SEQUENCE packages_id_seq RESTART WITH 1;
ALTER SEQUENCE photographers_id_seq RESTART WITH 1;
ALTER SEQUENCE users_id_seq RESTART WITH 1;
ALTER SEQUENCE citizen_cards_id_seq RESTART WITH 1;
ALTER SEQUENCE tags_id_seq RESTART WITH 1;
ALTER SEQUENCE media_id_seq RESTART WITH 1;
ALTER SEQUENCE reviews_id_seq RESTART WITH 1;
ALTER SEQUENCE categories_id_seq RESTART WITH 1;
ALTER SEQUENCE quotations_id_seq RESTART WITH 1;

-- Insert Users
INSERT INTO Users (name, email, phone_number, profile_picture_url, role, facebook, instagram, bank, account_no, bank_branch, created_at, updated_at)
VALUES
('User 1', 'user1@example.com', '1234567890', 'https://cdn-icons-png.flaticon.com/512/10337/10337609.png', 'CUSTOMER', 'Fookbace', 'ig', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 2', 'user2@example.com', '2345678901', 'https://img.freepik.com/free-vector/blue-circle-with-white-user_78370-4707.jpg', 'CUSTOMER', 'bookface', 'graminsta', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 3', 'user3@example.com', '3456789012', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRuGFjsxZCvbMuKnsJHFywAKXzJh6SsPWVsifY_z36wVT9p38WQ3IQPDPDjhFPDyxv6YQY&usqp=CAU', 'PHOTOGRAPHER', 'vlllqw sq', 'IG', 'BAY', '', 'branch', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 4', 'user4@example.com', '4567890123', 'https://img.freepik.com/premium-vector/user-profile-icon-flat-style-member-avatar-vector-illustration-isolated-background-human-permission-sign-business-concept_157943-15752.jpg', 'PHOTOGRAPHER', 'face book', 'GI', 'KKP', '', 'bchnaf', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 5', 'user5@example.com', '5678901234', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRuGFjsxZCvbMuKnsJHFywAKXzJh6SsPWVsifY_z36wVT9p38WQ3IQPDPDjhFPDyxv6YQY&usqp=CAU', 'ADMIN', '', '', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Citizen Cards
INSERT INTO Citizen_Cards (citizen_id, laser_id, picture, expire_date, created_at, updated_at)
VALUES
('1519999567819', 'LASER123', 'https://www.visa.com.vn/dam/VCOM/regional/ap/vietnam/global-elements/images/vn-visa-gold-card-498x280.png', '2026-12-31', LOCALTIMESTAMP, LOCALTIMESTAMP),
('4819999567819', 'LASER234', 'https://www.visa.com.vn/dam/VCOM/regional/ap/vietnam/global-elements/images/vn-visa-classic-card-498x280.png', '2027-11-30', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Photographers
INSERT INTO Photographers (user_id, is_verified, active_status, citizen_card_id, created_at, updated_at)
VALUES
(3, true, true, 1, LOCALTIMESTAMP, LOCALTIMESTAMP),
(4, true, true, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Categories
INSERT INTO Categories (name, description, created_at, updated_at)
VALUES
('Wedding', 'Packages related to wedding photography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Portrait', 'Packages related to portrait photography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Event', 'Packages related to event photography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Landscape', 'Packages related to landscape photography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Sports', 'Packages related to sports photography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Street', 'Packages related to street photography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Astro', 'Packages related to astrophotography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Family', 'Packages related to family photography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Underwater', 'Packages related to underwater photography', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Nature', 'Packages related to nature and landscapes', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Packages
INSERT INTO Packages (photographer_id, name, description, price, category_id, created_at, updated_at)
VALUES
(1, 'Sunset Package', 'A collection of sunset images', 150.00, 10, LOCALTIMESTAMP, LOCALTIMESTAMP),
(1, 'Wedding Package', 'A collection of wedding images', 300.00, 1, LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'Nature Package', 'A collection of nature images', 200.00, 10, LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'Portrait Package', 'A collection of portrait images', 250.00, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Tags
INSERT INTO Tags (package_id, name, created_at, updated_at)
VALUES
(1, 'Sunset', LOCALTIMESTAMP, LOCALTIMESTAMP),
(1, 'Nature', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'Wedding', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'Portrait', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Media
INSERT INTO Media (package_id, picture_url, description, created_at, updated_at)
VALUES
(1, 'https://t4.ftcdn.net/jpg/01/04/78/75/360_F_104787586_63vz1PkylLEfSfZ08dqTnqJqlqdq0eXx.jpg', 'Sunset image', LOCALTIMESTAMP, LOCALTIMESTAMP),
(1, 'https://media.istockphoto.com/id/1172427455/photo/beautiful-sunset-over-the-tropical-sea.jpg?s=612x612&w=0&k=20&c=i3R3cbE94hdu6PRWT7cQBStY_wknVzl2pFCjQppzTBg=', 'Sunset image 2', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'https://media.istockphoto.com/id/587197548/photo/beautiful-setting-for-outdoors-wedding-ceremony.jpg?s=612x612&w=0&k=20&c=E46nXAiNpnREvNNPUvc-4tQZhzdjJb6PSPasNFvNsOs=', 'Wedding image', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'https://media.istockphoto.com/id/1043755348/photo/romantic-wedding-ceremony.jpg?s=612x612&w=0&k=20&c=pXjKa-aTfh3oxYzc06HkYw19f-Ez9q-bPpElZmwlFKw=', 'Wedding image 2', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'https://media.istockphoto.com/id/681119612/photo/wedding-birthday-reception-decoration-chairs-tables-and-flowers.jpg?s=612x612&w=0&k=20&c=8K-WOBrUC9KrrQbuD8LwDgAH7g3KyEvbe1jOsfdsE6w=', 'Wedding image 3', LOCALTIMESTAMP, LOCALTIMESTAMP),
(3, 'https://media.istockphoto.com/id/517188688/photo/mountain-landscape.jpg?s=1024x1024&w=0&k=20&c=z8_rWaI8x4zApNEEG9DnWlGXyDIXe-OmsAyQ5fGPVV8=', 'Nature image', LOCALTIMESTAMP, LOCALTIMESTAMP),
(4, 'https://t4.ftcdn.net/jpg/05/23/62/91/360_F_523629123_RpAModBJXgCTPfilfYaCIbPaalFIjbvv.jpg', 'Portrait image', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Reviews
INSERT INTO Reviews (package_id, customer_id, rating, comment, created_at, updated_at)
VALUES
(1, 1, 4.5, 'Amazing!', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 2, 5.0, 'Beautiful shots', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Link Packages to Categories
INSERT INTO Packages_Categories (package_id, category_id)
VALUES
(1, 1),
(2, 1);

-- Insert Quotations
INSERT INTO Quotations (package_id, customer_id, photographer_id, status, price, description, from_date, to_date, created_at, updated_at)
VALUES
(1, 1, 1, 'PENDING', 1500.00, 'I would like to book this package', '2025-03-03T17:33:00+07:00', '2025-03-03T20:33:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP),
(3, 2, 1, 'PAID', 540.00, 'I would like to book this package', '2025-03-03T17:23:00+07:00', '2025-03-03T17:40:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 2, 1, 'PENDING', 1530.00, 'I would like to book this package', '2025-03-03T17:23:00+07:00', '2025-03-03T17:40:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 2, 2, 'ACCEPTED', 6300.00, 'I would like to book this package', '2025-03-03', '2025-03-04', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 1, 2, 'CANCELLED', 300.00, 'I would like to book this package', '2025-03-03T17:33:00+07:00', '2025-03-03T19:50:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP);

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