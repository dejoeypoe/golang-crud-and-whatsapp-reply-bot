-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 10 Jul 2022 pada 04.30
-- Versi server: 10.4.14-MariaDB
-- Versi PHP: 7.2.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `fikomdb`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `chats`
--

CREATE TABLE `chats` (
  `kode` varchar(20) NOT NULL,
  `balasan` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `chats`
--

INSERT INTO `chats` (`kode`, `balasan`) VALUES
('nadia', 'hallo , ini pesan otomatis');

-- --------------------------------------------------------

--
-- Struktur dari tabel `cutis`
--

CREATE TABLE `cutis` (
  `id` int(10) UNSIGNED NOT NULL,
  `nik` varchar(50) DEFAULT NULL,
  `tanggal_awal` date DEFAULT NULL,
  `tanggal_akhir` date DEFAULT NULL,
  `potongan_gaji` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `cutis`
--

INSERT INTO `cutis` (`id`, `nik`, `tanggal_awal`, `tanggal_akhir`, `potongan_gaji`) VALUES
(6789, '90456', '2022-05-06', '2022-06-07', 15000),
(1234567, '5678', '2022-01-02', '2022-02-03', 120);

-- --------------------------------------------------------

--
-- Struktur dari tabel `dokumens`
--

CREATE TABLE `dokumens` (
  `id` int(10) UNSIGNED NOT NULL,
  `nama_dokumen` varchar(150) DEFAULT NULL,
  `file_id` varchar(255) DEFAULT NULL,
  `file_url` varchar(255) DEFAULT NULL,
  `waktu_upload` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `dokumens`
--

INSERT INTO `dokumens` (`id`, `nama_dokumen`, `file_id`, `file_url`, `waktu_upload`) VALUES
(1, 'coba upload', '1-7EJAK1pdz2YGqULd3QgAnnHmQctseHN', 'https://docs.google.com/document/d/1-7EJAK1pdz2YGqULd3QgAnnHmQctseHN/edit?usp=drivesdk&ouid=101694989736610027652&rtpof=true&sd=true', '2022-06-03 09:23:16'),
(2, 'coba upload', '1Ran_dZCScynNH8ynJOMP5A4XHtIvHTCc', 'https://drive.google.com/file/d/1Ran_dZCScynNH8ynJOMP5A4XHtIvHTCc/view?usp=drivesdk', '2022-06-03 09:26:50');

-- --------------------------------------------------------

--
-- Struktur dari tabel `jabatans`
--

CREATE TABLE `jabatans` (
  `id` int(10) UNSIGNED NOT NULL,
  `nama` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `jabatans`
--

INSERT INTO `jabatans` (`id`, `nama`) VALUES
(123, 'Nadia Amalia ');

-- --------------------------------------------------------

--
-- Struktur dari tabel `pegawais`
--

CREATE TABLE `pegawais` (
  `nik` varchar(255) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `jk` varchar(100) DEFAULT NULL,
  `no_telp` varchar(100) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `pegawais`
--

INSERT INTO `pegawais` (`nik`, `nama`, `jk`, `no_telp`, `alamat`) VALUES
('3311', 'Nadia Nur Amalia ', 'Perempuan', '085728379434', 'Karanganyar, Palur');

-- --------------------------------------------------------

--
-- Struktur dari tabel `riwayat_jabatans`
--

CREATE TABLE `riwayat_jabatans` (
  `id` int(10) UNSIGNED NOT NULL,
  `nik` varchar(255) DEFAULT NULL,
  `jabatan_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `nama` varchar(150) DEFAULT NULL,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `nama`, `username`, `password`) VALUES
(1, 'Sopingi', 'sopingi', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `chats`
--
ALTER TABLE `chats`
  ADD PRIMARY KEY (`kode`);

--
-- Indeks untuk tabel `cutis`
--
ALTER TABLE `cutis`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `dokumens`
--
ALTER TABLE `dokumens`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `jabatans`
--
ALTER TABLE `jabatans`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `riwayat_jabatans`
--
ALTER TABLE `riwayat_jabatans`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `cutis`
--
ALTER TABLE `cutis`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1234568;

--
-- AUTO_INCREMENT untuk tabel `dokumens`
--
ALTER TABLE `dokumens`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `jabatans`
--
ALTER TABLE `jabatans`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=124;

--
-- AUTO_INCREMENT untuk tabel `riwayat_jabatans`
--
ALTER TABLE `riwayat_jabatans`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=124;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
