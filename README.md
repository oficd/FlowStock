# FlowStock
# Warehouse & Inventory Management System

Sistem informasi inventori berbasis web yang dikembangkan untuk membantu pengelolaan data barang dan mendukung proses pemantauan persediaan dalam gudang.

Aplikasi ini dikembangkan menggunakan **Golang** sebagai backend, **PostgreSQL** sebagai basis data, serta **Vue.js** dan **Tailwind CSS** sebagai teknologi frontend.

Proyek ini dikembangkan secara bertahap, dimulai dari autentikasi pengguna, pengelolaan data master, dashboard inventori, hingga pengembangan fitur transaksi persediaan.

## 1. Pengenalan Aplikasi

Warehouse & Inventory Management System merupakan aplikasi yang dirancang untuk mendukung pengelolaan inventori secara terstruktur melalui antarmuka web.

Tujuan pengembangan aplikasi ini adalah menyediakan sistem yang dapat dikembangkan lebih lanjut untuk mengelola data barang, mengorganisasi kategori produk, dan memantau informasi inventori. Pengembangan selanjutnya akan berfokus pada pencatatan transaksi serta pelacakan perubahan stok.

### Teknologi yang Digunakan

| Komponen | Teknologi |
|---|---|
| Backend | Golang |
| Framework Backend | Gin |
| Frontend | Vue.js |
| Styling | Tailwind CSS |
| Database | PostgreSQL |
| ORM | GORM |

## 2. Fitur Inti

### Selesai Dikembangkan

- **Autentikasi dan Login** — menyediakan mekanisme login untuk mengakses aplikasi.
- **Manajemen Produk** — mendukung operasi CRUD (*Create, Read, Update, Delete*) untuk mengelola data produk.
- **Manajemen Kategori** — mendukung operasi CRUD untuk mengelola kategori produk.
- **Dashboard Inventori** — menyediakan halaman dashboard sebagai bagian dari antarmuka utama aplikasi.
- **Backend API** — struktur backend Golang telah disiapkan dan koneksi ke PostgreSQL telah dikonfigurasi.
- **Frontend** — struktur aplikasi frontend telah disiapkan untuk mendukung interaksi pengguna dengan sistem.

### Sedang Dikembangkan

- **Fitur Transaksi Persediaan** — pengembangan mekanisme transaksi untuk mendukung pencatatan dan pengelolaan perubahan stok barang.

Fitur transaksi masih dalam tahap pengembangan dan belum dianggap selesai.

## 3. Milestone Fitur Tambahan

Fitur berikut merupakan rencana pengembangan selanjutnya dan belum dinyatakan tersedia pada versi saat ini.

### Milestone 1 — Penyelesaian Transaksi Persediaan

- **Stock In** — pencatatan barang masuk.
- **Stock Out** — pencatatan barang keluar.
- **Stock Adjustment** — penyesuaian stok berdasarkan hasil pemeriksaan.
- **Stock Transfer** — pemindahan barang antar gudang.

### Milestone 2 — Pengelolaan Data Pendukung

- **Manajemen Supplier** — pengelolaan data pemasok barang.
- **Manajemen Gudang** — pengelolaan data gudang penyimpanan.

### Milestone 3 — Pemantauan dan Riwayat Inventori

- **Stock Card** — melihat riwayat perubahan stok barang.
- **Transaction History** — melihat riwayat transaksi persediaan.
- **Low Stock Alert** — memberikan informasi ketika stok barang berada di bawah batas minimum.

Urutan milestone dapat disesuaikan dengan kebutuhan pengembangan dan prioritas implementasi.

## 4. Cara Penggunaan

### Persyaratan

Sebelum menjalankan aplikasi, siapkan lingkungan pengembangan berikut:

- Go
- Node.js dan npm
- PostgreSQL
- Git

Pastikan seluruh dependensi dan konfigurasi lingkungan telah tersedia.

### 4.1. Clone Repository

Clone repository proyek ke komputer lokal:

```bash
git clone <URL_REPOSITORY>
cd <NAMA_FOLDER_PROYEK>
```

Ganti placeholder dengan URL repository dan nama folder proyek yang sesuai.

### 4.2. Konfigurasi Database

1. Jalankan PostgreSQL.
2. Buat database untuk aplikasi.
3. Konfigurasikan parameter koneksi database sesuai pengaturan backend.
4. Pastikan backend dapat terhubung ke database sebelum menjalankan aplikasi.

Nama database, parameter koneksi, dan mekanisme migrasi mengikuti konfigurasi yang digunakan dalam proyek.

### 4.3. Menjalankan Backend

Masuk ke direktori backend:

```bash
cd backend
```

Unduh dependensi Go:

```bash
go mod download
```

Jalankan aplikasi sesuai entry point backend yang tersedia. Jika entry point berada di `main.go`, perintahnya adalah:

```bash
go run main.go
```

Pastikan konfigurasi database telah diatur sebelum menjalankan backend.

### 4.4. Menjalankan Frontend

Buka terminal baru, kemudian masuk ke direktori frontend:

```bash
cd frontend
```

Instal dependensi:

```bash
npm install
```

Jalankan server pengembangan:

```bash
npm run dev
```

Buka alamat lokal yang ditampilkan oleh Vite di terminal.

### 4.5. Mengakses Aplikasi

Setelah backend dan frontend berjalan:

1. Buka aplikasi melalui alamat lokal frontend.
2. Login menggunakan akun yang telah tersedia atau dikonfigurasi sesuai mekanisme autentikasi proyek.
3. Akses dashboard inventori.
4. Gunakan fitur pengelolaan produk dan kategori yang telah tersedia.

**Catatan:** Perintah menjalankan aplikasi, konfigurasi lingkungan, pembuatan akun, dan migrasi database perlu disesuaikan dengan implementasi aktual repository.

---

## Status Pengembangan

Proyek saat ini telah memiliki fondasi backend, frontend, koneksi PostgreSQL, autentikasi, CRUD produk, CRUD kategori, dan dashboard inventori. Fokus pengembangan berikutnya adalah menyelesaikan fitur transaksi persediaan sebelum melanjutkan ke fitur pendukung dan pemantauan inventori.
