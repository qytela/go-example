# Belajar Go + Gin + Gorm + Auth

Ini adalah projek belajar Bahasa Go dengan CRUD simple dan Auth JWT (Go JWT).

## **Installation**

### **Setting up .env**

```bash
cp .env.example .env
```

### **Download Modules**

```bash
make download
```

### **Migration**

Lakukan migration dengan dengan menggunakan [go-migrate](https://github.com/golang-migrate/migrate) dan Makefile.

Up

```bash
make migration_up
```

Down

```bash
make migration_down
```

### **Run**

Running Go (windows):

```bash
.\run.bat
```

Selain windows? google please. :)

### **Postman Collection**

Download postman collection [disini](https://api.postman.com/collections/6184613-f0f5af2a-fe99-4764-a5bf-107af0be79b9?access_key=PMAT-01HCMDBTWR9K7V8KJ1FAF4VXSD).
