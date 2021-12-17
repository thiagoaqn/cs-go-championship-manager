# cs-go-championship-manager

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Prisma](https://img.shields.io/badge/Prisma-3982CE?style=for-the-badge&logo=Prisma&logoColor=white)

## 🐘 Running Database
```bash
docker-compose up -d
```

## ✨ Database migration
```bash
go run github.com/prisma/prisma-client-go migrate dev --name init
```