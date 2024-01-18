-- DDL
DROP DATABASE IF EXISTS `melisprint`;

CREATE DATABASE `melisprint`;

USE `melisprint`;

-- table `sellers`
CREATE TABLE `sellers` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `cid` int(11) NOT NULL,
    `company_name` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sellers_cid` (`cid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `warehouses`
CREATE TABLE `warehouses` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `warehouse_code` varchar(25) NOT NULL,
    `address` varchar(255) NOT NULL,
    `telephone` varchar(15) NOT NULL,
    `minimum_capacity` int NOT NULL,
    `minimum_temperature` float NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_warehouses_warehouse_code` (`warehouse_code`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `sections`
CREATE TABLE `sections` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `section_number` int(11) NOT NULL,
    `current_temperature` float NOT NULL,
    `minimum_temperature` float NOT NULL,
    `current_capacity` int NOT NULL,
    `minimum_capacity` int NOT NULL,
    `maximum_capacity` int NOT NULL,
    `warehouse_id` int(11) NOT NULL,
    `product_type_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sections_section_number` (`section_number`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `products`
CREATE TABLE `products` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `product_code` varchar(25) NOT NULL,
    `description` text NOT NULL,
    `height` float NOT NULL,
    `length` float NOT NULL,
    `width` float NOT NULL,
    `weight` float NOT NULL,
    `expiration_rate` float NOT NULL,
    `freezing_rate` float NOT NULL,
    `recom_freez_temp` float NOT NULL,
    `seller_id` int(11) NOT NULL,
    `product_type_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_products_product_code` (`product_code`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `employees`
CREATE TABLE `employees` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `card_number_id` int(11) NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    `warehouse_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_employees_card_number_id` (`card_number_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- table `buyers`
CREATE TABLE `buyers` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `card_number_id` int(11) NOT NULL,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_buyers_card_number_id` (`card_number_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;