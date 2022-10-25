CREATE TABLE `users` (
  `id` varchar(255) PRIMARY KEY,
  `fullname` varchar(255),
  `gender` varchar(255),
  `contact` varchar(255),
  `street` varchar(255),
  `city_id` varchar(255),
  `province_id` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `role` varchar(255),
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
  `img_url` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `transactions` (
  `id` varchar(255) PRIMARY KEY,
  `user_id` varchar(255),
  `origin` varchar(255),
  `destination` varchar(255),
  `weight` int,
  `courier` varchar(255),
  `status` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `items` (
  `id` varchar(255) PRIMARY KEY,
  `transaction_id` varchar(255),
  `product_id` varchar(255),
  `quantity` int
);

ALTER TABLE `transactions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `items` ADD FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);

ALTER TABLE `items` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
