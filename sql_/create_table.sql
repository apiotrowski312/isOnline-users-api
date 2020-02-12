create table users(
   id BIGINT(20) NOT NULL AUTO_INCREMENT,
   first_name VARCHAR(100),
   last_name VARCHAR(100),
   email VARCHAR(100) NOT NULL,
   date_created DATETIME,
   status VARCHAR(100) NOT NULL,
   password VARCHAR(32) NOT NULL,
   PRIMARY KEY ( id ),
   UNIQUE (email)
);