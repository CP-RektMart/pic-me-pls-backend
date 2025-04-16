-- Reset the database for testing
DELETE FROM Messages;
DELETE FROM Previews;
DELETE FROM Tags;
DELETE FROM Reviews;
DELETE FROM Citizen_Cards;
DELETE FROM Quotations;
DELETE FROM Media;
DELETE FROM Packages_Categories;
DELETE FROM Packages;
DELETE FROM Categories;
DELETE FROM Photographers;
DELETE FROM Users;

ALTER SEQUENCE packages_id_seq RESTART WITH 1;
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
('User 1', 'user1@example.com', '0034567890', 'https://cdn-icons-png.flaticon.com/512/10337/10337609.png', 'CUSTOMER', 'Fookbace', 'ig', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 2', 'user2@example.com', '0045678901', 'https://img.freepik.com/free-vector/blue-circle-with-white-user_78370-4707.jpg', 'CUSTOMER', 'bookface', 'graminsta', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 3', 'user3@example.com', '0056789012', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRuGFjsxZCvbMuKnsJHFywAKXzJh6SsPWVsifY_z36wVT9p38WQ3IQPDPDjhFPDyxv6YQY&usqp=CAU', 'PHOTOGRAPHER', 'vlllqw sq', 'IG', 'BAY', '', 'branch', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 4', 'user4@example.com', '0067890123', 'https://img.freepik.com/premium-vector/user-profile-icon-flat-style-member-avatar-vector-illustration-isolated-background-human-permission-sign-business-concept_157943-15752.jpg', 'PHOTOGRAPHER', 'face book', 'GI', 'KKP', '', 'bchnaf', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 5', 'user5@example.com', '0078901234', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRuGFjsxZCvbMuKnsJHFywAKXzJh6SsPWVsifY_z36wVT9p38WQ3IQPDPDjhFPDyxv6YQY&usqp=CAU', 'ADMIN', '', '', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 6', 'user6@example.com', '0038901234', 'https://thumbs.dreamstime.com/b/customer-support-service-agent-headset-flat-vector-icon-design-designs-153069456.jpg', 'ADMIN', '', '', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 7', 'user7@example.com', '0068901234', 'https://static.vecteezy.com/system/resources/previews/043/900/708/non_2x/user-profile-icon-illustration-vector.jpg', 'ADMIN', '', '', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP),
('User 8', 'user8@example.com', '0063901234', 'https://t3.ftcdn.net/jpg/00/65/75/68/360_F_65756860_GUZwzOKNMUU3HldFoIA44qss7ZIrCG8I.jpg', 'ADMIN', '', '', '', '', '', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Photographers
INSERT INTO Photographers (user_id, is_verified, active_status, created_at, updated_at)
VALUES
(3, true, true, LOCALTIMESTAMP, LOCALTIMESTAMP),
(4, true, true, LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Citizen Cards
INSERT INTO Citizen_Cards (citizen_id, laser_id, photographer_id, picture, expire_date, created_at, updated_at)
VALUES
('1519999567819', 'LASER123', 3, 'https://www.visa.com.vn/dam/VCOM/regional/ap/vietnam/global-elements/images/vn-visa-gold-card-498x280.png', '2026-12-31', LOCALTIMESTAMP, LOCALTIMESTAMP),
('4819999567819', 'LASER234', 4, 'https://www.visa.com.vn/dam/VCOM/regional/ap/vietnam/global-elements/images/vn-visa-classic-card-498x280.png', '2027-11-30', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Categories
INSERT INTO Categories (name, description, created_at, updated_at)
VALUES
('Wedding', 'Capture the love, laughter, and lifetime vows in stunning detail.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Portrait', 'A timeless collection of personal and professional portraits.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Event', 'From birthdays to corporate galas—relive every special moment.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Landscape', 'Breathtaking vistas, majestic mountains, and serene countryside views.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Sports', 'High-speed action shots that freeze the thrill of the game.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Street', 'Raw and authentic snapshots of everyday life in motion.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Astro', 'Journey beyond the stars with stunning astrophotography.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Family', 'Cherish heartfelt family moments with beautifully composed images.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Underwater', 'Dive into the deep and capture marine wonders beneath the surface.', LOCALTIMESTAMP, LOCALTIMESTAMP),
('Nature', 'Reconnect with the beauty of the earth through vibrant nature photography.', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Packages
INSERT INTO Packages (photographer_id, name, description, price, category_id, created_at, updated_at)
VALUES
(3, 'Golden Hour Magic', 'Let the sun paint your memories with breathtaking sunset photography.', 150.00, 10, LOCALTIMESTAMP, LOCALTIMESTAMP),
(3, 'Ever After Wedding', 'Immortalize the happiest day of your life with a dreamy wedding shoot.', 300.00, 1, LOCALTIMESTAMP, LOCALTIMESTAMP),
(4, 'Wilderness Wonders', 'Lose yourself in the raw beauty of nature through this immersive package.', 200.00, 10, LOCALTIMESTAMP, LOCALTIMESTAMP),
(4, 'Timeless Portraits', 'Classic and elegant portraits designed to make you look your best.', 250.00, 2, LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Tags
INSERT INTO Tags (package_id, name, created_at, updated_at)
VALUES
(1, 'Golden Hour', LOCALTIMESTAMP, LOCALTIMESTAMP),
(1, 'Nature', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'Wedding Bliss', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 'Elegant Portraits', LOCALTIMESTAMP, LOCALTIMESTAMP);

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

-- Link Packages to Categories
INSERT INTO Packages_Categories (package_id, category_id)
VALUES
(1, 1),
(2, 1);

-- Insert Quotations
INSERT INTO Quotations (package_id, customer_id, photographer_id, status, price, description, from_date, to_date, created_at, updated_at)
VALUES
(1, 1, 3, 'SUBMIT', 1500.00, 'Eager to capture the golden hour magic—let’s make this happen!', '2025-03-03T17:33:00+07:00', '2025-03-03T20:33:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP),
(3, 2, 4, 'PAID', 540.00, 'Excited to freeze this special moment in time—booking confirmed!', '2025-03-03T17:23:00+07:00', '2025-03-03T17:40:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 2, 3, 'PENDING', 1530.00, 'Looking forward to a timeless portrait session—can’t wait!', '2025-03-03T17:23:00+07:00', '2025-03-03T17:40:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 2, 3, 'CONFIRMED', 6300.00, 'All set for a stunning shoot—excited to bring our vision to life!', '2025-03-03', '2025-03-04', LOCALTIMESTAMP, LOCALTIMESTAMP),
(2, 1, 3, 'CANCELLED', 300.00, 'Unfortunately, plans changed—hope to rebook soon!', '2025-03-03T17:33:00+07:00', '2025-03-03T19:50:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP),
(3, 2, 4, 'CANCELLED', 540.00, 'Excited to freeze this special moment in time—booking confirmed!', '2025-03-03T17:23:00+07:00', '2025-03-03T17:40:00+07:00', LOCALTIMESTAMP, LOCALTIMESTAMP);

-- Insert Reviews
INSERT INTO Reviews (package_id, customer_id, rating, comment, created_at, updated_at, quotation_id, is_edited)
VALUES
(3, 2, 4.5, 'Breathtaking colors and stunning shots! Worth every penny.', LOCALTIMESTAMP, LOCALTIMESTAMP, 2, false),
(2, 1, 5.0, 'Absolutely magical! Every moment was perfectly captured.', LOCALTIMESTAMP, LOCALTIMESTAMP, 3, true),
(3, 2, 5.0, 'Absolutely magical! Every moment was perfectly captured.', LOCALTIMESTAMP, LOCALTIMESTAMP, 6, true);

-- Insert Previews
INSERT INTO Previews (quotation_id, link) 
VALUES
(1, 'img1.link'),
(1, 'img2.link');

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
select p.name as Package_Name, p.description as Package_Description, p.price as Package_Price, ph.user_id as Photographer_ID,
u.id as User_ID, u.name as userName from Packages as p
join photographers as ph on p.photographer_id = ph.user_id
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
JOIN photographers AS ph ON ph.user_id = q.photographer_id
JOIN users AS u2 ON u2.id = ph.user_id;