# myapp

First Go app

# create database
CREATE DATABASE bird_encyclopedia;
# create table
CREATE TABLE birds (
  id INT NOT NULL AUTO_INCREMENT,
  species varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  PRIMARY KEY (id)
);
