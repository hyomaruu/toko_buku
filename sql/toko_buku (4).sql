-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 26 Jun 2024 pada 15.40
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `toko_buku`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `aksesoris`
--

CREATE TABLE `aksesoris` (
  `kode` varchar(255) NOT NULL,
  `nama_barang` varchar(255) NOT NULL,
  `deskripsi` text DEFAULT NULL,
  `gambar` varchar(255) NOT NULL,
  `harga` int(255) NOT NULL,
  `stok` int(11) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `aksesoris`
--

INSERT INTO `aksesoris` (`kode`, `nama_barang`, `deskripsi`, `gambar`, `harga`, `stok`, `created_at`, `updated_at`, `deleted_at`) VALUES
('dsbakhjdbasjdsad', 'Barang B', 'dnsajkndjskandknsajdnsad', 'upload-3977523033.png', 1000000, 0, '2024-06-26 06:23:11', '2024-06-26 06:23:11', NULL),
('njknkasd', 'AKSESORIS A', 'dnsajkndkjsandsa', 'upload-1304492031.png', 100000, 100, '2024-06-26 06:14:04', '2024-06-26 06:14:04', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `data_admin`
--

CREATE TABLE `data_admin` (
  `id_admin` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nama_lengkap` varchar(255) NOT NULL,
  `role` enum('OWNER','CASHIER','STAFF','') NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `data_admin`
--

INSERT INTO `data_admin` (`id_admin`, `username`, `password`, `nama_lengkap`, `role`, `created_at`, `updated_at`, `deleted_at`) VALUES
(12442, 'abidmaulana001@gmail.com', 'dnsakjdnsnadsad', 'Abid Fakhri Maulana OWNER', 'OWNER', '2024-06-25 16:48:23', '2024-06-25 16:48:23', NULL),
(12444, 'abidmaulana002@gmail.com', 'AAZsCIXVKYBpXa0DPPo', 'Abid Fakhri Maulana STAFF', 'STAFF', '2024-06-25 16:48:23', '2024-06-25 16:48:23', NULL),
(12445, 'abidmaulana003@gmail.com', 'asdfghjkl', 'Abid Fakhri Maulana CASHIER', 'CASHIER', '2024-06-25 16:48:23', '2024-06-25 16:48:23', NULL),
(12463, 'abidmaulana0072n@gmail.com', 'asdfghjkl', 'Abid Maulana FMA', 'OWNER', '2024-06-25 16:48:23', '2024-06-26 17:27:28', NULL),
(12465, 'dnsjaknads@gmail.com', 'dnsajkndkasnd', 'dnsjaknadsdnsajkndjkasnjd', 'OWNER', '2024-06-25 16:51:18', '2024-06-25 16:52:00', '2024-06-25 16:54:37'),
(12466, 'dsadsad@gmail.com', 'dsadadsa', 'dsadasd', 'OWNER', '2024-06-26 17:07:06', '2024-06-26 17:07:06', '2024-06-26 17:12:44'),
(12467, 'dsadsanjkmkldsamlkdas@gmail.com', 'dsandsaddsa', 'dnsajkdnsakld as', 'OWNER', '2024-06-26 17:08:23', '2024-06-26 17:08:23', '2024-06-26 17:12:39'),
(12468, 'dsabdhas@gmail.com', 'dsad', 'dsadasd', 'STAFF', '2024-06-26 17:11:36', '2024-06-26 17:11:36', '2024-06-26 17:12:41'),
(12469, 'dsad@gmail.com', 'ndkjnkjnkjn', 'jsjsjsjs', 'CASHIER', '2024-06-26 17:12:56', '2024-06-26 17:12:56', '2024-06-26 17:21:38'),
(12470, 'dnskjandasdnk@gmmdksla.com', 'fdnajkfas', 'dsadsad', 'STAFF', '2024-06-26 17:13:32', '2024-06-26 17:13:32', '2024-06-26 17:21:35'),
(12471, 'c@gmail.com', 'bhjdsadsa', 'ddfdsf', 'OWNER', '2024-06-26 17:14:18', '2024-06-26 17:17:46', '2024-06-26 17:21:30'),
(12472, 'dsds@gmail.com', 'dsd', 'dsds', 'CASHIER', '2024-06-26 17:15:10', '2024-06-26 17:15:10', '2024-06-26 17:21:24');

-- --------------------------------------------------------

--
-- Struktur dari tabel `data_buku`
--

CREATE TABLE `data_buku` (
  `kode_buku` varchar(255) NOT NULL,
  `nama_buku` varchar(255) NOT NULL,
  `penerbit` varchar(255) NOT NULL,
  `gambar` varchar(255) NOT NULL,
  `deskripsi` text DEFAULT NULL,
  `pengarang` varchar(255) NOT NULL,
  `genre` enum('Kuliner','Pendidikan','Komik','Biografi','Novel','Keagamaan') NOT NULL,
  `rilis` date NOT NULL,
  `harga` int(255) NOT NULL,
  `stok` int(11) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `data_buku`
--

INSERT INTO `data_buku` (`kode_buku`, `nama_buku`, `penerbit`, `gambar`, `deskripsi`, `pengarang`, `genre`, `rilis`, `harga`, `stok`, `created_at`, `updated_at`, `deleted_at`) VALUES
('budibsauhidsa', 'ABCGD', 'Penerbit A', 'upload-2183997362.png', 'dmksamkldmlaskmdlmaslkmd\r\ndsandknsakdjnsa\r\ndsa\r\ndas\r\nd\r\nsadsad', 'Test Pengarang', 'Komik', '2024-06-12', 10000, 0, '2024-06-25 19:19:18', '2024-06-25 19:19:18', '2024-06-25 22:39:37'),
('buku1', 'Test Buku 1 v2', 'Anonymous', 'upload-608559348.png', 'lorem ipsum dolor sit amet buku 1', 'Pengarang V1', 'Kuliner', '2024-08-27', 100000, 1949, '2024-06-25 19:19:18', '2024-06-26 20:38:31', NULL),
('buku12345d', 'Mencari Alam', 'Test Penerbit', 'upload-3555934773.png', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'Test Pengarang', 'Kuliner', '2024-06-29', 120000, 1009, '2024-06-25 19:19:18', '2024-06-26 06:04:35', '2024-06-26 06:04:55'),
('buku2', 'Test Buku 2', 'Anonymous', 'upload-2436998733.png', 'lorem ipsum dolor sit amet buku 2', 'Pengarang V1', 'Komik', '2024-08-27', 25000, 486, '2024-06-25 19:19:18', '2024-06-26 20:06:39', NULL),
('BUKUABC', 'ABC ', 'Erlangga', 'upload-3954447388.png', 'lorem ipsum dolor sit amet', 'Test Pengarang 2', 'Kuliner', '2024-06-18', 2000000, 0, '2024-06-25 19:19:18', '2024-06-25 19:19:18', '2024-06-26 18:00:50'),
('dsad', 'dsa', 'dsadsa', 'upload-247638025.png', 'dsad', 'dsada', 'Kuliner', '2024-07-04', 123213213, 100, '2024-06-25 22:22:06', '2024-06-26 18:00:01', '2024-06-26 18:00:12'),
('hbshjbasd67', 'asdfghjkl', 'Penerbit HHHH', 'upload-832758716.png', 'sdfghjkl', 'Pengarang jjj', 'Kuliner', '2024-06-29', 10000000, 0, '2024-06-26 17:40:09', '2024-06-26 17:40:09', '2024-06-26 17:40:19');

-- --------------------------------------------------------

--
-- Struktur dari tabel `data_member`
--

CREATE TABLE `data_member` (
  `id_member` int(11) NOT NULL,
  `nama_lengkap` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `telepon` varchar(255) NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `data_member`
--

INSERT INTO `data_member` (`id_member`, `nama_lengkap`, `email`, `telepon`, `alamat`, `created_at`, `updated_at`, `deleted_at`) VALUES
(222337, 'Member New Update', 'member004@gmail.com', '08123456789112233', 'lorem ipsum dolor sit amet alamat new', '2024-06-25 18:41:34', '2024-06-25 18:41:34', NULL),
(222338, 'Member 2', 'member002@gmail.com', '0812345678922', 'lorem ipsum dolor sit amet alamat v2', '2024-06-25 18:41:34', '2024-06-25 18:41:34', NULL),
(222340, 'dnsajkndkjsad', 'dnksajd@gmail.com', '01293218093', 'lorem ipsum dolor sit amaetlorem ipsum dolor sit amaetlorem ipsum dolor sit amaetlorem ipsum dolor sit amaetlorem ipsum dolor sit amaetlorem xbxhbsahxbashbxhasb', '2024-06-25 18:41:34', '2024-06-25 18:41:34', '2024-06-25 19:05:08'),
(222342, 'dsadsad', 'dsadsa@gmail.com', '123456789034567', 'dsahuidhuihuadsiuhasduhidsauhihuiasdhuadsiuhiasduhisadiuhdsa', '2024-06-25 19:05:28', '2024-06-26 17:28:56', NULL),
(222343, 'dsaddsa', 'dsadasd@gmial.com', '2134', 'dsadsad', '2024-06-26 17:24:42', '2024-06-26 17:33:35', '2024-06-26 17:33:45');

-- --------------------------------------------------------

--
-- Struktur dari tabel `keranjang`
--

CREATE TABLE `keranjang` (
  `kode_keranjang` int(11) NOT NULL,
  `kode_transaksi` int(11) DEFAULT NULL,
  `tipe_barang` enum('AKSESORIS','BUKU') NOT NULL,
  `kode_barang` varchar(255) NOT NULL,
  `total_barang` int(255) DEFAULT 0,
  `harga` int(11) DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `keranjang`
--

INSERT INTO `keranjang` (`kode_keranjang`, `kode_transaksi`, `tipe_barang`, `kode_barang`, `total_barang`, `harga`, `created_at`, `updated_at`, `deleted_at`) VALUES
(16, NULL, 'BUKU', 'buku2', 1, 25000, '2024-06-26 12:33:20', '2024-06-26 12:33:20', '2024-06-26 18:19:14'),
(17, 27, 'BUKU', 'buku1', 7, 100000, '2024-06-26 15:47:39', '2024-06-26 15:47:39', NULL),
(18, NULL, 'BUKU', 'buku1', 1, 100000, '2024-06-26 18:03:42', '2024-06-26 18:03:42', '2024-06-26 18:21:17'),
(19, NULL, 'BUKU', 'buku1', 1, 100000, '2024-06-26 18:05:38', '2024-06-26 18:05:38', '2024-06-26 18:21:16'),
(20, NULL, 'BUKU', 'buku2', 1, 25000, '2024-06-26 18:06:08', '2024-06-26 18:06:08', '2024-06-26 18:21:15'),
(21, NULL, 'BUKU', 'buku2', 1, 25000, '2024-06-26 18:06:30', '2024-06-26 18:06:30', '2024-06-26 18:21:14'),
(22, NULL, 'BUKU', 'buku2', 1, 25000, '2024-06-26 18:06:33', '2024-06-26 18:06:33', '2024-06-26 18:21:14'),
(23, NULL, 'BUKU', 'buku1', 1, 100000, '2024-06-26 18:07:13', '2024-06-26 18:07:13', '2024-06-26 18:21:13'),
(24, NULL, 'BUKU', 'buku2', 1, 25000, '2024-06-26 18:08:43', '2024-06-26 18:08:43', '2024-06-26 18:20:21'),
(25, 27, 'BUKU', 'buku2', 4, 25000, '2024-06-26 19:15:02', '2024-06-26 19:15:02', NULL),
(26, NULL, 'BUKU', 'buku2', 1, 25000, '2024-06-26 20:25:38', '2024-06-26 20:25:38', '2024-06-26 20:38:14'),
(27, NULL, 'BUKU', 'buku2', 1, 25000, '2024-06-26 20:34:20', '2024-06-26 20:34:20', '2024-06-26 20:38:15'),
(28, 28, 'BUKU', 'buku1', 3, 100000, '2024-06-26 20:38:20', '2024-06-26 20:38:20', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksi`
--

CREATE TABLE `transaksi` (
  `kode_transaksi` int(11) NOT NULL,
  `nomor_member` int(255) NOT NULL,
  `id_admin` int(255) NOT NULL,
  `tanggal` date NOT NULL,
  `total_harga` int(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp(),
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `transaksi`
--

INSERT INTO `transaksi` (`kode_transaksi`, `nomor_member`, `id_admin`, `tanggal`, `total_harga`, `created_at`, `updated_at`, `deleted_at`) VALUES
(27, 222338, 12445, '0000-00-00', 800000, '2024-06-26 20:06:39', '2024-06-26 20:06:39', NULL),
(28, 222338, 12445, '0000-00-00', 300000, '2024-06-26 20:38:31', '2024-06-26 20:38:31', NULL);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `aksesoris`
--
ALTER TABLE `aksesoris`
  ADD PRIMARY KEY (`kode`) USING BTREE;

--
-- Indeks untuk tabel `data_admin`
--
ALTER TABLE `data_admin`
  ADD PRIMARY KEY (`id_admin`);

--
-- Indeks untuk tabel `data_buku`
--
ALTER TABLE `data_buku`
  ADD PRIMARY KEY (`kode_buku`);

--
-- Indeks untuk tabel `data_member`
--
ALTER TABLE `data_member`
  ADD PRIMARY KEY (`id_member`);

--
-- Indeks untuk tabel `keranjang`
--
ALTER TABLE `keranjang`
  ADD PRIMARY KEY (`kode_keranjang`);

--
-- Indeks untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`kode_transaksi`),
  ADD KEY `nomor_member` (`nomor_member`),
  ADD KEY `id_admin` (`id_admin`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `data_admin`
--
ALTER TABLE `data_admin`
  MODIFY `id_admin` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12473;

--
-- AUTO_INCREMENT untuk tabel `data_member`
--
ALTER TABLE `data_member`
  MODIFY `id_member` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=222344;

--
-- AUTO_INCREMENT untuk tabel `keranjang`
--
ALTER TABLE `keranjang`
  MODIFY `kode_keranjang` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `kode_transaksi` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD CONSTRAINT `admin_key` FOREIGN KEY (`id_admin`) REFERENCES `data_admin` (`id_admin`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `member_key` FOREIGN KEY (`nomor_member`) REFERENCES `data_member` (`id_member`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
