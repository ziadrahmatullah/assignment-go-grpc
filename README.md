# assignment-go-grpc

Asumsi:
- reset expire 1 menit
- Saat Pembuatan user, auto generate wallet dan juga attempt row
- pada GetTransactions jika user hanya menginput salah satu antara end dan start pada endpoint, maka error invalid filter format
- pada GetTransactions jika user hanya meninput sort pada endpoint tanpa sortBy, maka error invalid sort format
- pada GetTransactions jika user menginput page yang melebihi total page, maka error page not found
- sort pada GetTransaction, date: created_at, amount: amount, to : receiver (wallet number)
- search menggunakan case insensitive

Seeding: /sql

Unit Test Coverage:
- Handler : 93.8%
- Usecase : 100%
- Util : 31.9%

How To Use Code:
1. Run the third party server: ./cmd/calculator/api_linux
2. Run the GRPC server: go run cmd/grpc/main.go
3. Open /request/auth.http
4. Can do register and login
5. After login, you will get access token
6. Copy that token to another http file in default header for using that file

Documentation : https://documenter.getpostman.com/view/31472691/2s9YeN1U24