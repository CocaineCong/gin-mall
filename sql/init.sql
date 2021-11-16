
CREATE TABLE `user` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`user_name` varchar(255) UNIQUE,`email` varchar(255),`password_digest` varchar(255),`nickname` varchar(255) NOT NULL,`status` varchar(255),`avatar` varchar(1000),`money` int , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_user_deleted_at ON `user`(deleted_at)


CREATE TABLE `product` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`name` varchar(255),`category_id` int unsigned,`title` varchar(255),`info` varchar(1000),`img_path` varchar(255),`price` varchar(255),`discount_price` varchar(255),`on_sale` boolean DEFAULT false,`num` int,`boss_id` int,`boss_name` varchar(255),`boss_avatar` varchar(255) , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_product_deleted_at ON `product`(deleted_at)

CREATE INDEX idx_product_name ON `product`(`name`)

CREATE TABLE `carousel` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`img_path` varchar(255),`product_id` int unsigned , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_carousel_deleted_at ON `carousel`(deleted_at)
CREATE TABLE `category` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`category_name` varchar(255) , PRIMARY KEY (`id`)) charset=utf8mb4


CREATE INDEX idx_category_deleted_at ON `category`(deleted_at)

CREATE TABLE `favorite` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`user_id` int unsigned,`product_id` int unsigned,`boss_id` int unsigned , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_favorite_deleted_at ON `favorite`(deleted_at)

CREATE TABLE `product_img` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`product_id` int unsigned,`img_path` varchar(255) , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_product_img_deleted_at ON `product_img`(deleted_at)

CREATE TABLE `product_info_img` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`product_id` int unsigned,`img_path` varchar(255) , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_product_info_img_deleted_at ON `product_info_img`(deleted_at)

CREATE TABLE `product_param_img` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`product_id` int unsigned,`img_path` varchar(255) , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_product_param_img_deleted_at ON `product_param_img`(deleted_at)


CREATE TABLE `order` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`user_id` int unsigned,`product_id` int unsigned,`boss_id` int unsigned,`address_id` int unsigned,`num` int unsigned,`order_num` bigint unsigned,`type` int unsigned,`money` int , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_order_deleted_at ON `order`(deleted_at)

CREATE TABLE `cart` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`user_id` int unsigned,`product_id` int unsigned,`boss_id` int unsigned,`num` int unsigned,`max_num` int unsigned,`check` boolean , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_cart_deleted_at ON `cart`(deleted_at)

CREATE TABLE `admin` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`user_name` varchar(255),`password_digest` varchar(255),`avatar` varchar(1000) , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_admin_deleted_at ON `admin`(deleted_at)

CREATE TABLE `address` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`user_id` int unsigned,`name` varchar(20),`phone` varchar(11),`address` varchar(50) , PRIMARY KEY (`id`)) charset=utf8mb4
CREATE INDEX idx_address_deleted_at ON `address`(deleted_at)

CREATE TABLE `notice` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`text` text , PRIMARY KEY (`id`)) charset=utf8mb4

CREATE INDEX idx_notice_deleted_at ON `notice`(deleted_at)


ALTER TABLE `favorite` ADD CONSTRAINT favorite_user_id_User_id_foreign FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `favorite` ADD CONSTRAINT favorite_product_id_Product_id_foreign FOREIGN KEY (product_id) REFERENCES Product(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `cart` ADD CONSTRAINT cart_user_id_User_id_foreign FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `cart` ADD CONSTRAINT cart_product_id_Product_id_foreign FOREIGN KEY (product_id) REFERENCES Product(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `order` ADD CONSTRAINT order_user_id_User_id_foreign FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `order` ADD CONSTRAINT order_address_id_Address_id_foreign FOREIGN KEY (address_id) REFERENCES Address(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `product` ADD CONSTRAINT product_category_id_Category_id_foreign FOREIGN KEY (category_id) REFERENCES Category(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `product_img` ADD CONSTRAINT product_img_product_id_Product_id_foreign FOREIGN KEY (product_id) REFERENCES Product(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `product_info_img` ADD CONSTRAINT product_info_img_product_id_Product_id_foreign FOREIGN KEY (product_id) REFERENCES Product(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `product_param_img` ADD CONSTRAINT product_param_img_product_id_Product_id_foreign FOREIGN KEY (product_id) REFERENCES Product(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `address` ADD CONSTRAINT address_user_id_User_id_foreign FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE ON UPDATE CASCADE;
