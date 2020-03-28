# myapp

First Go app - Bird Encyclopedia

## 1. Docker:
```
docker-compose -up

http://localhost:8090/
```

## 2. Non Docker:

### create database
```
CREATE DATABASE bird_encyclopedia;
```
### create table
```
CREATE TABLE birds (
  id INT NOT NULL AUTO_INCREMENT,
  species varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  PRIMARY KEY (id)
);
```
### Build and Run
```
go build

./myapp

http://localhost:8090/
```

