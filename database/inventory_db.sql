-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 11, 2025 at 04:08 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `inventory_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `barang_keluar`
--

CREATE TABLE `barang_keluar` (
  `id` int(11) NOT NULL,
  `item_id` int(11) DEFAULT NULL,
  `jumlah` int(11) DEFAULT NULL,
  `tanggal` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `barang_keluars`
--

CREATE TABLE `barang_keluars` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `item_id` bigint(20) UNSIGNED DEFAULT NULL,
  `jumlah` bigint(20) DEFAULT NULL,
  `tanggal` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `barang_keluars`
--

INSERT INTO `barang_keluars` (`id`, `item_id`, `jumlah`, `tanggal`) VALUES
(2, 2, 10, '2025-07-11 00:26:36.291'),
(3, 6, 5, '2025-07-11 00:58:12.042'),
(4, 1, 5, '2025-07-11 00:58:23.113'),
(5, 8, 10, '2025-07-11 01:29:46.085'),
(6, 4, 5, '2025-07-11 01:37:36.552'),
(7, 6, 5, '2025-07-11 01:41:15.158'),
(8, 7, 5, '2025-07-11 00:00:00.000'),
(9, 9, 5, '2025-07-11 00:00:00.000');

-- --------------------------------------------------------

--
-- Table structure for table `barang_masuk`
--

CREATE TABLE `barang_masuk` (
  `id` int(11) NOT NULL,
  `item_id` int(11) DEFAULT NULL,
  `jumlah` int(11) DEFAULT NULL,
  `tanggal` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `barang_masuks`
--

CREATE TABLE `barang_masuks` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `item_id` bigint(20) UNSIGNED DEFAULT NULL,
  `jumlah` bigint(20) DEFAULT NULL,
  `tanggal` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `barang_masuks`
--

INSERT INTO `barang_masuks` (`id`, `item_id`, `jumlah`, `tanggal`) VALUES
(4, 1, 10, '2025-07-11 00:15:56.169'),
(5, 13, 10, '2025-07-11 00:32:07.056'),
(6, 12, 20, '2025-07-11 00:33:27.972'),
(7, 12, 30, '2025-07-11 00:50:01.338'),
(8, 11, 10, '2025-07-11 01:56:03.415');

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

CREATE TABLE `items` (
  `id` int(11) NOT NULL,
  `name` longtext DEFAULT NULL,
  `stock` bigint(20) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `items`
--

INSERT INTO `items` (`id`, `name`, `stock`, `price`, `created_at`) VALUES
(1, 'Sabun', 150, 3500, '2025-07-10 06:24:48.726'),
(2, 'Shampoo', 270, 19000, '2025-07-10 06:25:03.568'),
(3, 'Past Gigi', 80, 8000, '2025-07-10 07:44:13.615'),
(4, 'Handbody', 15, 17000, '2025-07-10 07:44:31.443'),
(5, 'Sunscreen', 70, 60000, '2025-07-10 07:45:01.778'),
(6, 'Conditioner', 70, 30000, '2025-07-10 08:23:56.012'),
(7, 'Parfum', 145, 30000, '2025-07-10 08:51:40.023'),
(8, 'Serum', 40, 40000, '2025-07-10 08:52:22.573'),
(9, 'Moisturizer', 75, 55000, '2025-07-10 08:52:44.384'),
(10, 'Body Scrub', 70, 33000, '2025-07-10 08:53:08.036'),
(11, 'Body Serum', 80, 60000, '2025-07-10 08:53:51.633'),
(12, 'Sunblock', 130, 60000, '2025-07-10 08:54:09.117'),
(13, 'Deodorant', 210, 20000, '2025-07-10 08:54:29.179'),
(14, 'Face Wash', 100, 23000, '2025-07-10 09:37:50.072'),
(15, 'Micellar Water', 80, 24000, '2025-07-10 09:56:59.086'),
(17, 'Sheet Mask', 70, 17000, '2025-07-11 00:24:29.165');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `barang_keluar`
--
ALTER TABLE `barang_keluar`
  ADD PRIMARY KEY (`id`),
  ADD KEY `item_id` (`item_id`);

--
-- Indexes for table `barang_keluars`
--
ALTER TABLE `barang_keluars`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `barang_masuk`
--
ALTER TABLE `barang_masuk`
  ADD PRIMARY KEY (`id`),
  ADD KEY `item_id` (`item_id`);

--
-- Indexes for table `barang_masuks`
--
ALTER TABLE `barang_masuks`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `barang_keluar`
--
ALTER TABLE `barang_keluar`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `barang_keluars`
--
ALTER TABLE `barang_keluars`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `barang_masuk`
--
ALTER TABLE `barang_masuk`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `barang_masuks`
--
ALTER TABLE `barang_masuks`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `items`
--
ALTER TABLE `items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `barang_keluar`
--
ALTER TABLE `barang_keluar`
  ADD CONSTRAINT `barang_keluar_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`);

--
-- Constraints for table `barang_masuk`
--
ALTER TABLE `barang_masuk`
  ADD CONSTRAINT `barang_masuk_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
