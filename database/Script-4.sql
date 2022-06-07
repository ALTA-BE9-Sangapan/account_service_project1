CREATE database db_proyek1;

use db_proyek1;

CREATE table users(
user_phone int(13) not null,
user_id int(11) not null,
saldo int,
password varchar(25),
date_birth date,
gender varchar(10),
alamat text,
primary key (user_phone, user_id)
);

CREATE table top_up(
top_up_id int(11) auto_increment primary key,
user_phone int(13) not null,
saldo int,
status text,
created_at timestamp default now(),
foreign key (user_phone) references users(user_phone)
);

CREATE table transfer(
transfer_id int(11) primary key auto_increment,
user_phone int(13) not null,
user_phone_receiver int(13) not null,
saldo int,
status text,
created_at timestamp default now(),
foreign key (user_phone) references users(user_phone),
foreign key (user_phone_receiver) references users(user_phone)
);
