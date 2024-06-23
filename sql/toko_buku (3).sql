-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 22 Jun 2024 pada 09.42
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
  `stok` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `aksesoris`
--

INSERT INTO `aksesoris` (`kode`, `nama_barang`, `deskripsi`, `gambar`, `harga`, `stok`) VALUES
('aksesoris1', 'Pensil V2', 'lorem ipsum dolor sit amet pensil v2', 'upload-3607547374.png', 4000, -5);

-- --------------------------------------------------------

--
-- Struktur dari tabel `data_admin`
--

CREATE TABLE `data_admin` (
  `id_admin` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nama_lengkap` varchar(255) NOT NULL,
  `role` enum('OWNER','CASHIER','STAFF','') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `data_admin`
--

INSERT INTO `data_admin` (`id_admin`, `username`, `password`, `nama_lengkap`, `role`) VALUES
(12442, 'abidmaulana001@gmail.com', 'AAZsCIXVKYBpXa0DPPo', 'Abid Fakhri Maulana OWNER', 'OWNER'),
(12444, 'abidmaulana002@gmail.com', 'AAZsCIXVKYBpXa0DPPo', 'Abid Fakhri Maulana STAFF', 'STAFF'),
(12445, 'abidmaulana003@gmail.com', 'AAZsCIXVKYBpXa0DPPo', 'Abid Fakhri Maulana CASHIER', 'CASHIER');

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
  `stok` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `data_buku`
--

INSERT INTO `data_buku` (`kode_buku`, `nama_buku`, `penerbit`, `gambar`, `deskripsi`, `pengarang`, `genre`, `rilis`, `harga`, `stok`) VALUES
('buku1', 'Test Buku 1', 'Anonymous', 'upload-27844852.png', 'lorem ipsum dolor sit amet buku 1', 'Pengarang V1', 'Kuliner', '2024-08-27', 100000, 1959),
('buku2', 'Test Buku 2', 'Anonymous', 'upload-2436998733.png', 'lorem ipsum dolor sit amet buku 2', 'Pengarang V1', 'Komik', '2024-08-27', 25000, 490),
('buku3', 'Test Buku 3', 'Anonymous', 'upload-3395764876.png', 'lorem ipsum dolor sit amet buku 3', 'Pengarang V9', 'Biografi', '2024-08-27', 65000, 8);

-- --------------------------------------------------------

--
-- Struktur dari tabel `data_member`
--

CREATE TABLE `data_member` (
  `id_member` int(11) NOT NULL,
  `nama_lengkap` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `telepon` varchar(255) NOT NULL,
  `alamat` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `data_member`
--

INSERT INTO `data_member` (`id_member`, `nama_lengkap`, `email`, `telepon`, `alamat`) VALUES
(222337, 'Member New Update', 'member004@gmail.com', '08123456789112233', 'lorem ipsum dolor sit amet alamat new'),
(222338, 'Member 2', 'member002@gmail.com', '0812345678922', 'lorem ipsum dolor sit amet alamat v2');

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
  `harga` int(11) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `keranjang`
--

INSERT INTO `keranjang` (`kode_keranjang`, `kode_transaksi`, `tipe_barang`, `kode_barang`, `total_barang`, `harga`) VALUES
(5, NULL, 'BUKU', 'buku1', 10, 100000),
(6, NULL, 'BUKU', 'buku2', 10, 25000),
(7, NULL, 'AKSESORIS', 'aksesoris1', 5, 4000),
(10, NULL, 'BUKU', 'buku3', 1, 65000),
(15, NULL, 'BUKU', 'buku1', 1, 100000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksi`
--

CREATE TABLE `transaksi` (
  `kode_transaksi` int(11) NOT NULL,
  `nomor_member` int(255) NOT NULL,
  `id_admin` int(255) NOT NULL,
  `tanggal` date NOT NULL,
  `total_harga` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
  MODIFY `id_admin` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12446;

--
-- AUTO_INCREMENT untuk tabel `data_member`
--
ALTER TABLE `data_member`
  MODIFY `id_member` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=222340;

--
-- AUTO_INCREMENT untuk tabel `keranjang`
--
ALTER TABLE `keranjang`
  MODIFY `kode_keranjang` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `kode_transaksi` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

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
