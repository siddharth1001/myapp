CREATE DATABASE IF NOT EXISTS bird_encyclopedia;
USE bird_encyclopedia;
CREATE TABLE IF NOT EXISTS birds (
  id INT NOT NULL AUTO_INCREMENT,
  species varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  PRIMARY KEY (id)
);

-- GRANT ALL PRIVILEGES ON *.* TO 'myuser'@'%';