CREATE DATABASE Project1;

use Project1;

CREATE table user(
Name varchar(25) not null,
Phone varchar(15) not null,
Password varchar(25),
Gender varchar(10),
Address text,
Balance int,
primary key (Phone)
);