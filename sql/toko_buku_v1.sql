-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 10 Jun 2024 pada 11.36
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
  `harga` int(255) NOT NULL,
  `stok` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `aksesoris`
--

INSERT INTO `aksesoris` (`kode`, `nama_barang`, `harga`, `stok`) VALUES
('cr332', 'Correction Tape', 12000, 0),
('pilot1222', 'Bolpoint Pilot', 4500, 0),
('standar2000', 'Bolpoint Standard', 3500, 0);

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
(12434, 'abidmaulana007@gmail.com', 'b7cilroy111', 'Abid Fakhri Maulana', 'OWNER'),
(12437, 'abidmaulana001@gmail.com', 'AAZsCIXVKYBpXa0DPPo', 'Abid Fakhri Maulana CASHIER', 'CASHIER'),
(12438, 'abidmaulana001@gmail.com', 'AAZsCIXVKYBpXa0DPPo', 'Abid Fakhri Maulana STAFF', 'STAFF');

-- --------------------------------------------------------

--
-- Struktur dari tabel `data_buku`
--

CREATE TABLE `data_buku` (
  `kode_buku` varchar(255) NOT NULL,
  `nama_buku` varchar(255) NOT NULL,
  `penerbit` varchar(255) NOT NULL,
  `gambar` varchar(255) NOT NULL,
  `pengarang` varchar(255) NOT NULL,
  `genre` enum('Kuliner','Pendidikan','Komik','Biografi','Novel','Keagamaan') NOT NULL,
  `rilis` date NOT NULL,
  `harga` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `data_buku`
--

INSERT INTO `data_buku` (`kode_buku`, `nama_buku`, `penerbit`, `gambar`, `pengarang`, `genre`, `rilis`, `harga`) VALUES
('ayamsayur', 'BH2111', 'Sugenks', 'jojojojo', 'gakManuks', 'Kuliner', '2024-05-02', 200000),
('ayamssssayur', 'BH2111', 'Sugenks', 'jojojojo', 'gakManuks', 'Kuliner', '2024-05-10', 200000),
('bhjsdabjhdasd', 'dnskajndjas', 'dnsadnsad', 'dskajdhsad', 'dnbshakbdksad', 'Kuliner', '2024-08-07', 3217932),
('BUKU1234', 'Mencari kebenaran', 'Anonymous', 'upload-312371505.png', 'Anonymous', 'Kuliner', '2024-08-27', 1000000),
('BUKUZZZZ', 'Mencari kebenaran', 'Anonymous', 'upload-3191251449.png', 'Anonymous', 'Kuliner', '2024-08-27', 1000000),
('nicholas2121', 'How To Be A Nichole', 'GedaGediGedaO', 'adve', 'WongKenthir', 'Kuliner', '2003-04-22', 50000),
('segitiga', 'Segitiga Bermuda', 'Daharjo', 'jasd', 'Julianto', 'Kuliner', '2002-06-21', 50000);

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
(210431, 'Asep Saipul Hamdan Jamaludin', '', '', ''),
(222334, 'SUK MAI', '', '', '');

-- --------------------------------------------------------

--
-- Struktur dari tabel `keranjang`
--

CREATE TABLE `keranjang` (
  `kode_transaksi` int(255) DEFAULT NULL,
  `kode_barang` varchar(255) NOT NULL,
  `total_barang` int(255) DEFAULT NULL,
  `total_harga` int(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `lokasi_buku`
--

CREATE TABLE `lokasi_buku` (
  `kode_buku` varchar(255) NOT NULL,
  `genre` varchar(255) NOT NULL,
  `nomor_rak` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `stok_barang`
--

CREATE TABLE `stok_barang` (
  `kode_barang` varchar(255) NOT NULL,
  `stok` int(12) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksi`
--

CREATE TABLE `transaksi` (
  `kode_transaksi` int(255) NOT NULL,
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
  ADD KEY `kode_transaksi` (`kode_transaksi`);

--
-- Indeks untuk tabel `lokasi_buku`
--
ALTER TABLE `lokasi_buku`
  ADD KEY `kode_buku` (`kode_buku`);

--
-- Indeks untuk tabel `stok_barang`
--
ALTER TABLE `stok_barang`
  ADD PRIMARY KEY (`kode_barang`);

--
-- Indeks untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`kode_transaksi`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `data_admin`
--
ALTER TABLE `data_admin`
  MODIFY `id_admin` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12441;

--
-- AUTO_INCREMENT untuk tabel `data_member`
--
ALTER TABLE `data_member`
  MODIFY `id_member` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=222336;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `keranjang`
--
ALTER TABLE `keranjang`
  ADD CONSTRAINT `keranjang_ibfk_1` FOREIGN KEY (`kode_transaksi`) REFERENCES `transaksi` (`kode_transaksi`);

--
-- Ketidakleluasaan untuk tabel `lokasi_buku`
--
ALTER TABLE `lokasi_buku`
  ADD CONSTRAINT `lokasi_buku_ibfk_1` FOREIGN KEY (`kode_buku`) REFERENCES `data_buku` (`kode_buku`);

--
-- Ketidakleluasaan untuk tabel `stok_barang`
--
ALTER TABLE `stok_barang`
  ADD CONSTRAINT `stok_barang_ibfk_1` FOREIGN KEY (`kode_barang`) REFERENCES `data_buku` (`kode_buku`),
  ADD CONSTRAINT `stok_barang_ibfk_2` FOREIGN KEY (`kode_barang`) REFERENCES `aksesoris` (`kode`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
