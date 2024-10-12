# Description

Small go service for managing song library with REST API.

## Installation

Clone the repo using git clone:

```bash
git clone https://github.com/ladovod444/songs-library.git
```

Create postgress database 'songs' and set parameters for database in .env,
e.g.:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=db
DATABASE_NAME=songs
```
## Usage
Run the command in the root
```bash
go run main.go
```

## Swagger

All endpoints could be reviewed by visiting documentation page: http://localhost:8080/swagger/index.html

## License

[MIT](https://choosealicense.com/licenses/mit/)