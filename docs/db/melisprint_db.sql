-- DDL
DROP DATABASE IF EXISTS `melisprint`;

CREATE DATABASE `melisprint`;

USE `melisprint`;

-- table `locality`
CREATE TABLE `locality` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `locality_name` varchar(50) NOT NULL,
    `province_name` varchar(50) NOT NULL,
    `country_name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
);

-- table `sellers`
CREATE TABLE `sellers` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `cid` int NOT NULL,
    `company_name` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `locality_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sellers_cid` (`cid`),
    KEY `idx_sellers_locality_id` (`locality_id`),
    CONSTRAINT `fk_sellers_locality_id` FOREIGN KEY (`locality_id`) REFERENCES `locality` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `carriers`
CREATE TABLE `carriers` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `cid` int NOT NULL,
    `company_name` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `locality_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_carriers_cid` (`cid`),
    KEY `idx_carriers_locality_id` (`locality_id`),
    CONSTRAINT `fk_carriers_locality_id` FOREIGN KEY (`locality_id`) REFERENCES `locality` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `warehouses`
CREATE TABLE `warehouses` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `warehouse_code` varchar(25) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `minimum_capacity` int NOT NULL,
    `minimum_temperature` float NOT NULL,
    `locality_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_warehouses_warehouse_code` (`warehouse_code`),
    KEY `idx_warehouses_locality_id` (`locality_id`),
    CONSTRAINT `fk_warehouses_locality_id` FOREIGN KEY (`locality_id`) REFERENCES `locality` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `sections`
CREATE TABLE `sections` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `section_number` int NOT NULL,
    `current_temperature` float NOT NULL,
    `minimum_temperature` float NOT NULL,
    `current_capacity` int NOT NULL,
    `minimum_capacity` int NOT NULL,
    `maximum_capacity` int NOT NULL,
    `warehouse_id` int unsigned NOT NULL,
    `product_type_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sections_section_number` (`section_number`),
    KEY `idx_sections_warehouse_id` (`warehouse_id`),
    CONSTRAINT `fk_sections_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `products`
CREATE TABLE `products` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `product_code` varchar(25) NOT NULL,
    `description` text NOT NULL,
    `height` float NOT NULL,
    `length` float NOT NULL,
    `width` float NOT NULL,
    `weight` float NOT NULL,
    `expiration_rate` float NOT NULL,
    `freezing_rate` float NOT NULL,
    `recom_freez_temp` float NOT NULL,
    `seller_id` int unsigned NOT NULL,
    `product_type_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_products_product_code` (`product_code`),
    KEY `idx_products_seller_id` (`seller_id`),
    CONSTRAINT `fk_products_seller_id` FOREIGN KEY (`seller_id`) REFERENCES `sellers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `product_records`
CREATE TABLE `product_records` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `last_update_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `purchase_price` float NOT NULL,
    `sale_price` float NOT NULL,
    `product_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_product_records_product_id` (`product_id`),
    CONSTRAINT `fk_product_records_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `product_batches`
CREATE TABLE `product_batches` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `batch_number` int NOT NULL,
    `due_date` date NOT NULL,
    `minimum_temperature` float NOT NULL,
    `current_temperature` float NOT NULL,
    `initial_quantity` int NOT NULL,
    `current_quantity` int NOT NULL,
    `manufacturing_date` date NOT NULL,
    `manufacturing_hour` int NOT NULL,
    `section_id` int unsigned NOT NULL,
    `product_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_product_batches_batch_number` (`batch_number`),
    KEY `idx_product_batches_section_id` (`section_id`),
    KEY `idx_product_batches_product_id` (`product_id`),
    CONSTRAINT `fk_product_batches_section_id` FOREIGN KEY (`section_id`) REFERENCES `sections` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_product_batches_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `employees`
CREATE TABLE `employees` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `card_number_id` int unsigned NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    `warehouse_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_employees_card_number_id` (`card_number_id`),
    KEY `idx_employees_warehouse_id` (`warehouse_id`),
    CONSTRAINT `fk_employees_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `inbound_orders`
CREATE TABLE `inbound_orders` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `order_number` int NOT NULL,
    `order_date` date NOT NULL,
    `warehouse_id` int unsigned NOT NULL,
    `employee_id` int unsigned NOT NULL,
    `product_batch_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_inbound_orders_order_number` (`order_number`),
    KEY `idx_inbound_orders_warehouse_id` (`warehouse_id`),
    KEY `idx_inbound_orders_employee_id` (`employee_id`),
    KEY `idx_inbound_orders_product_batch_id` (`product_batch_id`),
    CONSTRAINT `fk_inbound_orders_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_inbound_orders_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_inbound_orders_product_batch_id` FOREIGN KEY (`product_batch_id`) REFERENCES `product_batches` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- table `buyers`
CREATE TABLE `buyers` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `card_number_id` int unsigned NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_buyers_card_number_id` (`card_number_id`)
);

-- table `purchase_orders`
CREATE TABLE `purchase_orders` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `order_number` int NOT NULL,
    `order_date` date NOT NULL,
    `tracking_code` varchar(25) NOT NULL,
    `buyer_id` int unsigned NOT NULL,
    `product_record_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_purchase_orders_order_number` (`order_number`),
    KEY `idx_purchase_orders_buyer_id` (`buyer_id`),
    KEY `idx_purchase_orders_product_record_id` (`product_record_id`),
    CONSTRAINT `fk_purchase_orders_buyer_id` FOREIGN KEY (`buyer_id`) REFERENCES `buyers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_purchase_orders_product_record_id` FOREIGN KEY (`product_record_id`) REFERENCES `product_records` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);