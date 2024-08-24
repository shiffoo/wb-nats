-- SELECT 'CREATE DATABASE dbWB' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'dbWB')\gexec
-- CREATE USER IF NOT EXISTS postgres IDENTIFIED BY '1234';
-- GRANT ALL PRIVILEGES ON DATABASE dbWB TO postgres;

-- \c dbWB;

-- CREATE TABLE IF NOT EXISTS orders (
-- 	order_uid						VARCHAR PRIMARY KEY,
-- 	track_number				VARCHAR NOT NULL,
-- 	entry								VARCHAR NOT NULL,
-- 	locale							VARCHAR NOT NULL,
-- 	internal_signature	VARCHAR,
-- 	customer_id					VARCHAR NOT NULL,
-- 	delivery_service		VARCHAR NOT NULL,
-- 	shardkey						VARCHAR NOT NULL,
-- 	sm_id								INT NOT NULL,
-- 	date_created				TIMESTAMP NOT NULL,
-- 	oof_shard						VARCHAR NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS delivery (
-- 	name					VARCHAR NOT NULL,
-- 	phone					VARCHAR NOT NULL,
-- 	zip						VARCHAR NOT NULL,
-- 	city					VARCHAR NOT NULL,
-- 	address				VARCHAR NOT NULL,
-- 	region				VARCHAR NOT NULL,
-- 	email					VARCHAR NOT NULL,
-- 	order_uid VARCHAR REFERENCES orders(order_uid) ON DELETE CASCADE NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS payment (
-- 	payment_id		SERIAL PRIMARY KEY,
-- 	transaction		VARCHAR NOT NULL,
-- 	request_id		VARCHAR NOT NULL,
-- 	currency			VARCHAR NOT NULL,
-- 	provider			VARCHAR NOT NULL,
-- 	amount				INT NOT NULL,
-- 	payment_dt		INT NOT NULL,
-- 	bank					VARCHAR NOT NULL,
-- 	delivery_cost	INT NOT NULL,
-- 	goods_total		INT NOT NULL,
-- 	custom_fee		INT NOT NULL,
-- 	order_uid VARCHAR REFERENCES orders(order_uid) ON DELETE CASCADE NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS item (
-- 	chrt_id				INT PRIMARY KEY,
-- 	track_number	VARCHAR NOT NULL,
-- 	price					INT NOT NULL,
-- 	rid						VARCHAR NOT NULL,
-- 	name					VARCHAR NOT NULL,
-- 	sale					INT NOT NULL,
-- 	size					VARCHAR NOT NULL,
-- 	total_price		INT NOT NULL,
-- 	nm_id					INT NOT NULL,
-- 	brand					VARCHAR NOT NULL,
-- 	status				INT NOT NULL,
-- 	order_uid VARCHAR REFERENCES orders(order_uid) ON DELETE CASCADE NOT NULL
-- );

-- -- INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
-- -- SELECT * FROM (SELECT 'b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 'en', '', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1') AS tmp
-- -- WHERE NOT EXISTS (
-- -- 	SELECT name FROM users WHERE order_uid = 'b563feb7b2b84b6test' AND track_number = 'WBILMTESTTRACK' AND entry = 'WBIL' AND locale= 'en' AND 
-- -- 															internal_signature = '' AND customer_id = 'test' AND delivery_service = 'meest' AND shardkey = '9' AND sm_id = 99 AND 
-- -- 															date_created = '2021-11-26T06:22:19Z' AND oof_shard = '1'
-- -- ) LIMIT 1;

-- -- INSERT INTO delivery (name, phone, zip, city, address, region, email, order_uid)
-- -- SELECT * FROM (SELECT 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com', 'b563feb7b2b84b6test') AS tmp
-- -- WHERE NOT EXISTS (
-- -- 	SELECT name FROM users WHERE name = 'Test Testov' AND phone = '+9720000000' AND zip = '2639809' AND city = 'Kiryat Mozkin' AND 
-- -- 																address = 'Ploshad Mira 15' AND region = 'Kraiot' AND email = 'test@gmail.com' AND order_uid = 'b563feb7b2b84b6test'
-- -- ) LIMIT 1;

-- -- INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid)
-- -- SELECT * FROM (SELECT 'b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500, 317, 0, 'b563feb7b2b84b6test') AS tmp
-- -- WHERE NOT EXISTS (
-- -- 	SELECT name FROM users WHERE transaction = 'b563feb7b2b84b6test' AND request_id = '' AND currency = 'USD' AND provider = 'wbpay' AND amount = 1817 AND payment_dt = 1637907727 AND 
-- -- 															bank = 'alpha' AND delivery_cost = 1500 AND goods_total = 317 AND custom_fee = 0 AND order_uid = 'b563feb7b2b84b6test'
-- -- ) LIMIT 1;

-- -- INSERT INTO item (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid)
-- -- SELECT * FROM (SELECT 9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202, 'b563feb7b2b84b6test') AS tmp
-- -- WHERE NOT EXISTS (
-- -- 	SELECT name FROM users WHERE chrt_id = 9934930, AND track_number = 'WBILMTESTTRACK' AND price = 453 AND rid = 'ab4219087a764ae0btest' AND
-- -- 																name = 'Mascaras' AND 'sale' = 30 AND size = "0" AND total_price = 317 AND nm_id = 2389212 AND 
-- -- 																brand = 'Vivienne Sabo' AND status = 202 AND order_uid = 'b563feb7b2b84b6test'
-- -- ) LIMIT 1;


