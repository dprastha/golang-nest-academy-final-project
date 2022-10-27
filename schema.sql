CREATE TABLE `users` (
  `id` varchar(255) PRIMARY KEY,
  `fullname` varchar(255) NOT NULL,
  `gender` varchar(255),
  `contact` varchar(255),
  `street` varchar(255),
  `city_id` varchar(255),
  `province_id` varchar(255),
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(255) DEFAULT "user",
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `products` (
  `id` varchar(255) PRIMARY KEY,
  `name` varchar(255),
  `category` varchar(255),
  `desc` varchar(255),
  `price` int,
  `stock` int,
  `weight` int,
  `img_url` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `transactions` (
  `id` varchar(255) PRIMARY KEY,
  `user_id` varchar(255),
  `product_id` varchar(255),
  `origin` varchar(255),
  `destination` varchar(255),
  `quantity` int,
  `weight` int,
  `total_price` int,
  `courier` varchar(255),
  `courier_cost` double,
  `status` varchar(255),
  `estimation_arrived` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `transactions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
