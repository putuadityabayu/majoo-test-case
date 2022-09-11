# Majoo Test Cases 1

[Run in Postman](https://www.postman.com/portalnesia/workspace/majoo-test-cases)

## Installation

1. Clone Repository

    git clone https://github.com/putuadityabayu/majoo-test-case.git

2. Install Mysql >= 8.0

    [Sample Data](sample.sql)

3. Go to project folder

    ```bash
    cd "Majoo Test Case 1"
    ```
    
4. Run mod tidy
  
    ```bash
    go mod tidy
    ```

5. Change Configuration Variables

    Database: [database.go](config/database.go)   
    JWT: [jwt.go](config/jwt.go)

6. Run Server

    ```bash
    go run main.go
    ```

7. Server running on ***localhost:3000***

---

## Fungsi-Fungsi

### A. Fungsi Login

- Endpoint [/login](http://localhost:3000/login)
- Menggunakan username dan md5(password)

### B. Authorization
- Middleware

### C. Endpoint laporan **name merchant**, **omzet** per hari pada bulan november mulai tanggal 1 sampai dengan tanggal 30 dengan pagination, Apabila tidak ada transaksi pada tanggal itu, maka **omzet** bernilai 0

- Endpoint [/opsi-c](http://localhost:3000/opsi-c)
- Pagination menggunakan `page` & `page_size` query

### D. Endpoint laporan **name merchant**, **nama outlet**, **omzet** per hari pada bulan november mulai tanggal 1 sampai dengan tanggal 30 dengan pagination (`page` & `page_size` query), Apabila tidak ada transaksi pada tanggal itu, maka **omzet** bernilai 0

- Endpoint [/opsi-d](http://localhost:3000/opsi-d)
- Pagination menggunakan `page` & `page_size` query

### E. Poin C, user tidak bisa melakukan akses pada merchant_id yang bukan miliknya

- users.id = ?
- users.id = merchants.user_id
- transactions.merchant_id = merchants.id

### F. Poin D, user tidak bisa melakukan akses pada outlet_id yang bukan miliknya

- users.id = ?
- users.id = merchants.user_id
- transactions.merchant_id = merchants.id
- outlets.id = transactions.outlet_id
- merchants.id = outlets.merchant_id


### G. Dari test case pada point C dan D, apakah ERD yang dibentuk sudah optimal? Berikan penjelasannya

### H. Dokumen teknis Data Manipulation Language (DML)

---

## JAWABAN

* Jawaban dapat dilihat pada file [MAJOO TEST CASE 1.pdf](MAJOO%20TEST%20CASE%201.pdf)