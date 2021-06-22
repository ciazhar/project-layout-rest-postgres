# Project Layout
> Project Layout merupakan standar projek yang saya gunakan untuk membuat program menggunakan bahasa golang.   

## Referensi
Struktur dan technology stack project ini sendiri diilhami dari :
- [Clean Architecture Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Golang Standard Projecr Layout](https://github.com/golang-standards/project-layout) 
- [Golang Developer Roadmap](https://github.com/Alikhll/golang-developer-roadmap)
- [Golang Web Framework Benchmark](https://github.com/smallnest/go-web-framework-benchmark)


## Requirement
- Mock Utility using [mockery](https://github.com/vektra/mockery)
- Database Migration using [sql-migrate](https://github.com/rubenv/sql-migrate)

## Instalasi
```sql
create schema if not exists project_layout_rest_postgres
```
 

## Stack
- [CLI](https://github.com/spf13/cobra)
- [Config File](https://github.com/spf13/viper)
- [Database (PostgreSQL)](https://github.com/jackc/pgx)
  https://github.com/georgysavva/scany
- Logger 
