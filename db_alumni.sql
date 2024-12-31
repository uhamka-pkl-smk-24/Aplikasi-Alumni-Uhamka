-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 31 Des 2024 pada 02.29
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_alumni`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `lamaran`
--

CREATE TABLE `lamaran` (
  `id` int(11) NOT NULL,
  `nama_pekerjaan` varchar(100) NOT NULL,
  `perusahaan` varchar(100) NOT NULL,
  `surat_lamaran` varchar(100) NOT NULL,
  `approve` int(11) NOT NULL DEFAULT 0,
  `nim` int(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `lamaran`
--

INSERT INTO `lamaran` (`id`, `nama_pekerjaan`, `perusahaan`, `surat_lamaran`, `approve`, `nim`) VALUES
(3, 'ofice boy', 'PT Terserah', 'Cara Membangun Jaringan LAN.docx', 1, 8976),
(28, 'Instansi Jaringan', 'MegaCrop', 'Cryptarithm.docx', 1, 8976),
(29, 'cyber security', 'cyberia', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 1, NULL),
(30, 'Back End Dev', 'FWP', 'KLINIK 24.docx', 1, NULL),
(31, 'Data Analyst', 'Megacorp', 'NODE JS.docx', 1, NULL),
(32, '123', 'pthbs', 'LIMIT - Limit Fungsi Aljabar.pdf', 1, NULL),
(33, 'Programmer', 'Cyber Corporation', 'Calculus 1.docx', 1, NULL),
(34, 'Front End Dev', 'MGP', 'Apa Itu Kalkulus Dasar.docx', 0, NULL),
(35, 'Programmer', 'Cyber Corporation', 'Cara Membangun Jaringan LAN.docx', 0, NULL),
(36, 'Cyber', 'GGWP', 'Polymorphism.docx', 0, NULL),
(37, 'Instansi Jaringan', 'MGP', 'Cara Menghitung Matriks.docx', 0, 8976),
(38, 'Barista', 'Breaker.co', 'Cara Menggunakan Mendeley di Word.docx', 1, 9899),
(39, 'Customer Service ', 'MGP.co', 'Doc2.docx', 1, 9899),
(40, '3D Designer', 'Terserah', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 1, 5134),
(41, 'Customer Service', 'Matahari.co', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 0, 124235),
(42, 'Customer Service', 'Matahari.co', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 0, 124235),
(43, 'Customer Service', 'Matahari.co', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 0, 124235),
(44, 'Proggramer', 'ITku', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 0, 124235),
(45, 'Data Analyst', 'Megacorp', 'Materi Dimensi Tiga.docx', 0, 6803),
(46, 'teh pucuk', 'fh', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 1, 51634),
(47, 'Data Analyst', 'Megacorp', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 1, 124235),
(48, 'Data Analyst', 'Megacorp', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 0, 124235),
(49, 'Proggramer', 'ITku', 'Apa Itu Ransomware, cara kerja, Jenis-Jenis, dan Cara Mencegahnya.docx', 0, 51634),
(50, 'sd', 'popopo', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 1, 51634),
(51, 'teh pucuk', 'fh', 'NODE JS.docx', 1, 51634),
(52, 'Customer Service', 'Matahari.co', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 1, 124235),
(53, 'Proggramer', 'ITku', 'CARA MEMBUAT QRIS ALL PAYMENT.docx', 1, 51634),
(54, 'Office Boy', 'PT Anti Nganggur', 'E-Book Aplikasi Alumni.docx', 0, 51634),
(55, 'Customer Service', 'Matahari.co', 'FORMAT DOKUMEN TEKNIS PEMBUATAN APLIKASI.docx', 0, 51634);

-- --------------------------------------------------------

--
-- Struktur dari tabel `lowongan`
--

CREATE TABLE `lowongan` (
  `id` int(11) NOT NULL,
  `nama_pekerjaan` varchar(100) NOT NULL,
  `perusahaan` varchar(100) NOT NULL,
  `lokasi` varchar(100) NOT NULL,
  `gaji` int(100) NOT NULL,
  `deskripsi` text NOT NULL,
  `syarat` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `lowongan`
--

INSERT INTO `lowongan` (`id`, `nama_pekerjaan`, `perusahaan`, `lokasi`, `gaji`, `deskripsi`, `syarat`) VALUES
(10, 'Customer Service', 'Matahari.co', 'Tanggerang', 1000000, 'Melayani kebutuhan nasabah dengan sepenuh hati\r\nMelayani pembukaan rekening koran tabungan dan deposito, menerima dan membuat permintaan buku cek dan giro serta mengaktifkan resinya, menangani dan mengelola keluhan customer\r\nMelayani masalah yang dihadapi oleh nasabah mengenai transaksi atau lainnya\r\nMembantu mencarikan solusi untuk masalah yang dihadapi oleh nasabah', 'Usia 23-25 Tahun.\r\nPendidikan minimal D3 (IPK minimal 2,75), S1 (IPK minimal 3,00) semua jurusan.\r\nTerbiasa mengoperasikan komputer.\r\nTeliti, jujur, dan bertanggung jawab pada pekerjaan.\r\nTerbiasa bekerja di bawah tekanan.\r\nBerpenampilan menarik, ramah, dan komunikatif.\r\nLebih diutamakan yang menguasai Bahasa Inggris dan Mandarin.\r\nBersedia ditempatkan di daerah Jabodetabek dan Cabang BBA Luar Kota lainnya.'),
(11, 'Customer Service ', 'MGP.co', 'Bandung', 5000000, 'Harus mengerti apa pun soal customer service\r\nMempunyai sikap ramah\r\nBerani bekerja di bawah tekanan\r\nMampu berkomunikasi dengan baik\r\n', 'Minimal usia 20, maksimal 30\r\nTidak memiliki tato\r\nTidak merokok\r\nPendidikan minimal D3'),
(16, 'Office Boy', 'PT Anti Nganggur', 'Depok', 3000000, 'Membersihkan kantor', 'Manusia'),
(18, 'Proggramer', 'ewew', 'ewew', 232323, '23232', '3232');

-- --------------------------------------------------------

--
-- Struktur dari tabel `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `nim` int(100) NOT NULL,
  `no_ijazah` int(100) NOT NULL,
  `nama_lengkap` varchar(100) NOT NULL,
  `tempat_lahir` varchar(100) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `agama` varchar(100) NOT NULL,
  `alamat` varchar(100) NOT NULL,
  `no_telp` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `ipk` int(100) NOT NULL,
  `bidang_studi` varchar(100) NOT NULL,
  `photo` varchar(1000) NOT NULL,
  `angkatan` int(20) NOT NULL,
  `tahun_lulus` int(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `mahasiswa`
--

INSERT INTO `mahasiswa` (`nim`, `no_ijazah`, `nama_lengkap`, `tempat_lahir`, `tanggal_lahir`, `agama`, `alamat`, `no_telp`, `email`, `ipk`, `bidang_studi`, `photo`, `angkatan`, `tahun_lulus`) VALUES
(223, 21321123, 'rapip', 'Jakarta', '2024-09-03', 'Islam', 'Petamburan', '0812362', 'rapip@gmail.com', 4, 'Informatika', 'Screenshot (116).png', 2010, 2015),
(1247, 2357, 'joko', 'Jawa Tengah', '2024-07-30', 'Islam', 'Petamburan', '082', 'joko@gmail.com', 4, 'Informatika', '', 2018, 2023),
(2342, 4652, 'arif', 'Bali', '2024-08-26', 'Islam', 'Bekasi', '08524', 'arifgans@gmail.com', 4, 'Kedokteran', '', 2020, 2024),
(2351, 3152, 'darno', 'Kalimantan Barat', '2024-02-25', 'Islam', 'Sawangan', '084213', 'darno@gmail.com', 4, 'Informatika', '', 2014, 2019),
(4124, 1234, 'sabil', 'jakarta', '2006-03-25', 'islam', 'jakarta', '012342', 'sabil@gmail.com', 4, 'informatika', '', 20020, 2024),
(5134, 2353, 'sabil', 'Jakarta', '2006-11-11', 'Islam', 'Cilangkap', '081585993109', 'sabil@gmail.com', 4, 'Informatika', 'images3.png', 2017, 2021),
(6452, 51232, 'ucup', 'Jawa Barat', '2024-08-26', 'Islam', 'Cipayung', '082132', 'ucup@gmail.com', 4, 'Informatika', '', 2014, 2019),
(6803, 876807668, 'Marsha Lenathea', 'Jakarta', '2006-01-09', 'Kristen', 'Jakarta Selatan', '089878468845', 'Marsha@gmail.com', 4, 'Sastra Inggris', '', 2019, 2024),
(7542, 5633, 'jarwo', 'Jawa Timur', '2024-05-23', 'Islam', 'Depok', '08124', 'jarwo@gmail.com', 4, 'Pertanian', '', 2013, 2017),
(8976, 7654, 'George Adam', 'Medan', '2006-07-11', 'Buddha', 'Pasar Minggu', '08876577860', 'george@gmail.com', 4, 'Tehnik Kelautan', 'asdadssa', 2020, 2024),
(9899, 8737658, 'Azizi Shafaa Asadel', 'Jakarta', '2000-06-08', 'Islam', 'Tangerang Selatan', '089876567843', 'zeeasadel@gmail.com', 3, 'Ilmu Komunikasi', '', 2018, 2023),
(21365, 1234, 'asep', 'Jawa Barat', '2024-02-08', 'Islam', 'Bandung', '089532', 'asep@gmail.com', 4, 'Keguruan', '', 2020, 2024),
(51634, 51234, 'Pria Sigma Lv100', 'Ohio', '2024-01-08', 'Skibidi', 'Ohio, jalan rizz Rt.mewing Rw.Skibidi', '8479837', 'sigmamale@gmail.com', 4, 'PerMewingan', 'WhatsApp Image 2024-09-04 at 10.24.31.jpeg', 2020, 2024),
(76315, 13512, 'wadhi', 'Bali', '2024-09-03', 'Islam', 'Petamburan', '08151', 'wadhi@gmail.com', 4, 'Informatika', 'yrry', 2015, 2020),
(124235, 2352, 'mahasiswa', 'depok', '2024-09-04', 'islam', 'depok', '08667333', 'mhs@gmail.com', 4, 'informatika', 'download.png', 2019, 2023),
(141212, 141212, 'SAMSUNG', 'Bekasi', '2024-09-03', 'islam', 'bekasi', '0895565', 'test@gmail.com', 4, 'tekni', '', 2009, 2014),
(505050, 505050, 'XLXLXL', 'Bekasi', '2024-09-03', 'islam', 'bekasi', '895565', 'XLXL@gmail.com', 4, 'tekni', '', 2009, 2014),
(908070, 123455, 'ucup lawrence', 'Bekasi', '2024-09-03', 'islam', 'bekasi', '08821123', 'nfsjbjsf@gmail.com', 4, 'tekni', '', 2009, 2014),
(1231312, 12331, 'test', 'Bekasi', '2024-09-03', 'islam', 'bekasi', '01238183', 'test@gmail.com', 4, 'tekni', 'Screenshot (116).png', 2009, 2014),
(2344134, 1231, 'test2', 'Bekasi', '2024-09-03', 'islam', 'bekasi', '089123', 'test2@gmai.com', 4, 'teknik', 'images.png', 2018, 2022),
(5423543, 13512123, 'fadli', 'Jakarta', '2024-09-02', 'Islam', 'Bekasi', '082132', 'fadli.hardiyanto.p11@gmail.com', 4, 'Informatika', 'Screenshot (119).png', 2019, 2024),
(6666666, 12331, 'test', 'Bekasi', '2024-09-03', 'islam', 'bekasi', '1238183', '666@gmail.com', 4, 'tekni', '', 2009, 2014),
(7090100, 123455, 'kotak', 'Bekasi', '2024-09-03', 'islam', 'bekasi', '08821123', 'nfsjbjsf@gmail.com', 4, 'tekni', '', 2009, 2014),
(73875879, 425646, 'al hitam', 'Bekasi', '2024-09-03', 'islam', 'bekasi', '0895565', 'nfsjbjsf@gmail.com', 4, 'tekni', '', 2009, 2014);

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `nim` int(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `role` enum('admin','admin_lowongan','mahasiswa') NOT NULL,
  `no_telp` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`nim`, `username`, `email`, `password`, `role`, `no_telp`) VALUES
(1, 'ucup', 'ucup@gmail.com', '$2a$10$B8J8aocDoEPtLUll0kdd/uoGlaJl1era7nVrKVnn/kkVnxkhLLxBm', 'mahasiswa', '0812362'),
(2, 'jarwo', 'jarwo@gmail.com', '$2a$10$CmgWkbz63JqH.dsbW6sPJ.E4dIG2TPixjJKnlozV9CvelBiWVK0Rm', 'mahasiswa', '082461'),
(111, 'megumii', 'megumii@gmail.com', '$2a$10$u6yDPIS5DkYSVMdSbr0v3.PlM5CscGQeOsfvNpzFLz5AbjneGCFOW', 'mahasiswa', ''),
(223, 'rapip', 'rapip@gmail.com', '$2a$10$4IIaO6IigXC/T1NPQcv1ROa3Vq6jvnm4034S2.s6KP4TDHJ3PZbNS', 'mahasiswa', '0812362'),
(321, 'ucup', 'ucupgans01@gmail.com', '$2a$12$IYIlazoZsXAyE9N8UpwKx.o4lIYKvQvscoOmEmjWlxpzhovT1U8XK', 'mahasiswa', ''),
(987, 'megumii', 'megumi@gmail.com', '$2a$10$U4o0lAkpZ5pwEZy75UJCde2ZrPS01H/n7Fmpf303aTr4ADbj8Jqeu', 'mahasiswa', '0812'),
(1001, 'Grego', 'gregoria@gmail.com', '$2a$10$1bS4gZsiZohuENc41XG.eu6QmYqubbNq6uSQAIaDeY0AvBqtiFE.u', 'mahasiswa', '0987657786'),
(1247, 'joko', 'joko@gmail.com', '$2a$10$CaGIU0K7MIPKecbUebbk4u0WHM60slnrx3Lq9MwT5VXVBsvALgmA2', 'mahasiswa', ''),
(2342, 'arif', 'arifgans@gmail.com', '$2a$10$CmvwV7BsOJAY2VfJqLsOguZu0aDV24yP3QD/fozdCuA/gRSWa0ONa', 'mahasiswa', '08524'),
(2351, 'darno', 'darno@gmail.com', '$2a$10$g7csah8.SJwAzodUo0DheOnJxRj1AhCAuUkgvpZtfoqo9LCT5k596', 'mahasiswa', '084213'),
(4124, 'sabil', 'sabil@gmail.com', '123', 'mahasiswa', '012342'),
(5134, 'sabil', 'sabil@gmail.com', '$2a$12$3TBoQpmDVH4xoWnnnDgEEugmO8xHWd32hjEcsRMAjwxJje5m.x/VO', 'mahasiswa', '081585993109'),
(5341, 'sumanto', 'sumanto@gmail.com', '$2a$10$ld2ZbPbGK4Dm.M.NFQnxfOT0.AhjuxTtXIEJsK5ojogPhIQTGubWG', 'mahasiswa', '081585'),
(6452, 'ucup', 'ucup@gmail.com', '$2a$10$NjqhUMyy/SlaGeWNV02x/.gaTqFBBWGpADCVQIPSIhcPoahTVBDye', 'mahasiswa', '082132'),
(6532, 'endang', 'endang@gmail.com', '$2a$10$Em5Ojrae6ioHdiEdzn3lZeXk01/xFogEokEFa/fO7OaqHZ8oeTE9K', 'mahasiswa', '083806'),
(6803, 'marsha', 'Marsha@gmail.com', '$2a$10$acuRGT60oArO2Rny6fqGIeRs7/98Nsh3mISPpav5WlBaWufESfXZK', 'mahasiswa', ''),
(7542, 'jarwo', 'jarwo@gmail.com', '$2a$10$EP1wyDf2zyVGfcrFfdGmOew4u81Giw2Ktf7KWb2JuJNrSHORPHf8W', 'mahasiswa', '08124'),
(8976, 'jorji', 'george@gmail.com', '$2a$10$u7fmcVCeCWHLYRphyhGwwelHUbMW/yz72aDQIMswI2VvHXKyBn.AS', 'mahasiswa', '08876577860'),
(9899, 'zee', 'zeeasadel@gmail.com', '$2a$10$JlVYA3183I8Q0suo3RYWBeYGTuzpYXMXcoeDyGdGhh55YQJQNcV.e', 'mahasiswa', ''),
(12345, 'admin', 'admin123@gmail.com', '$2y$10$NLArdUfJUWLEkyA/umP/Le.t3IgDh08QqUSsv2iwJhPwayEEacDb.', 'admin', '12121'),
(21365, 'asep', 'asep@gmail.com', '$2a$10$VzIbEMHT/b2fi/n7cXmzyuzVA8gkp7eKPRycngdYEmC1i8PRpoie2', 'mahasiswa', '089532'),
(24124, 'jarwo', 'jarwo@gmail.com', '$2a$10$8uPvY.J1tx23UdBbEQ.F0.Kn0uh7YJS488GmoH3pI20C/B7uTy7yu', 'mahasiswa', '081241'),
(32331, 'mhs', 'mhs123@gmail.com', '$2a$12$kodMZK1SB7NePQxwIQ326.ip0SBHz1jToeh1VUM2jUnV3YR/58IDS', 'mahasiswa', '86343'),
(42141, 'jarwo', 'jarwo@gmail.com', 'jarwo123', 'mahasiswa', '82132'),
(45364, 'admin2', 'admin2@gmail.com', '$2a$12$ahPZ/AOZW3K/M7Z4moapo.XwtoAdp01Mz8jPYW22OMIma2Vx1x6wG', 'admin_lowongan', '8904566464'),
(51634, 'sigma', 'sigmamale@gmail.com', '$2a$10$4Sb2Kk.cCiFl.bQ98ff2b.NcCFOc2OijvdMYzUC.BBkfeb3AKeHUS', 'mahasiswa', '8479837'),
(76315, 'wadhi', 'wadhi@gmail.com', '$2a$10$YKV/p6nQJ0a60wHCT5asJ.BmYwPZ5xoj5FUmc.DAMs7uIPxBKVgfy', 'mahasiswa', '08151'),
(86241, 'joko', 'joko@gmail.com', '$2a$10$8H/aLiqXz.4RpiVYzD1yse9/MxOgJcJCk.LSliaPoJNlbLqINvTne', 'mahasiswa', '084562'),
(124235, 'mahasiswa', 'mhs@gmail.com', '$2a$10$zjBvjglPtF7xl6z06KwPgePiO9yCBp7GNBngCdqNhMLHi9uFmpree', 'mahasiswa', '08667333'),
(141212, 'SAMSUNG', 'test@gmail.com', '123', 'mahasiswa', '0895565'),
(177013, 'Joko', 'jokonotoboto@gmail.com', '$2a$12$wcgsw9taSmczPe7YKjUZauCRKLhZFxVOljckB4UYBZh28bc66E5x.', 'mahasiswa', '8080834'),
(505050, 'XLXLXL', 'XLXL@gmail.com', '123', 'mahasiswa', '895565'),
(908070, 'ucup lawrence', 'nfsjbjsf@gmail.com', '123', 'mahasiswa', '08821123'),
(1231312, 'test', 'test@gmail.com', '$2a$10$yyEXzpLOutfZ3k.yP3.9S.Xcw9R7k2AumwvwZZoFx3f3X8UJ.oCZG', 'mahasiswa', '01238183'),
(2344134, 'test2', 'test2@gmai.com', '$2a$10$prTOCYfWzhszTJ5m4bZU/.kZ7CEZ7R84F/i3bpPfj6j0YWL9yNuju', 'mahasiswa', '089123'),
(5423543, 'fadli', 'fadli.hardiyanto.p11@gmail.com', '$2a$10$vC1oo7DdIIT9TThHS8XOeu.xxX//0AeSj5SefWaoWaoxVuw9176By', 'mahasiswa', '082132'),
(6666666, 'test', '666@gmail.com', '123', 'mahasiswa', '1238183'),
(7090100, 'kotak', 'nfsjbjsf@gmail.com', '$2a$12$8.PbJh49qw8FBMHvX7mvbe1vEMDn3hUq8BqubzPT/w49XSV3EFd9m', 'mahasiswa', '08821123'),
(9999999, 'fadli', 'fadli@gmail.com', '$2y$10$NLArdUfJUWLEkyA/umP/Le.t3IgDh08QqUSsv2iwJhPwayEEacDb.', 'admin', '08888888888'),
(73875879, 'al hitam', 'nfsjbjsf@gmail.com', '123', 'mahasiswa', '0895565');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `lamaran`
--
ALTER TABLE `lamaran`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_mahasiswa_lamaran` (`nim`);

--
-- Indeks untuk tabel `lowongan`
--
ALTER TABLE `lowongan`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `mahasiswa`
--
ALTER TABLE `mahasiswa`
  ADD PRIMARY KEY (`nim`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`nim`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `lamaran`
--
ALTER TABLE `lamaran`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=56;

--
-- AUTO_INCREMENT untuk tabel `lowongan`
--
ALTER TABLE `lowongan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT untuk tabel `mahasiswa`
--
ALTER TABLE `mahasiswa`
  MODIFY `nim` int(100) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=577866679;

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `nim` int(100) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=577866679;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `lamaran`
--
ALTER TABLE `lamaran`
  ADD CONSTRAINT `fk_mahasiswa_lamaran` FOREIGN KEY (`nim`) REFERENCES `user` (`nim`);

--
-- Ketidakleluasaan untuk tabel `mahasiswa`
--
ALTER TABLE `mahasiswa`
  ADD CONSTRAINT `mahasiswa_ibfk_1` FOREIGN KEY (`nim`) REFERENCES `user` (`nim`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
