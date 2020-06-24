# learningGo

**db ddl**

create schema kartProject collate utf8mb4_0900_ai_ci;

create table kart
(
	id mediumint auto_increment,
	userId mediumint not null,
	status varchar(100) default 'OPEN' not null,
	createDateTime datetime default CURRENT_TIMESTAMP not null,
	constraint kart_id_uindex
		unique (id)
);

alter table kart
	add primary key (id);

create table product
(
	name varchar(200) not null,
	price float not null,
	constraint product_name_uindex
		unique (name)
);

alter table product
	add primary key (name);

create table kartitem
(
	id mediumint auto_increment,
	kartId mediumint not null,
	quantity int not null,
	productName varchar(200) not null,
	constraint kartitem_id_uindex
		unique (id),
	constraint kartitem_kart_id_fk
		foreign key (kartId) references kart (id),
	constraint kartitem_product_name_fk
		foreign key (productName) references product (name)
);

alter table kartitem
	add primary key (id);
	
**db dml**

INSERT INTO kartProject.product (name, price) VALUES ('banana', 1);


