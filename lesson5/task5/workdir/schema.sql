/*
 Проект существующий сейчас который переезжает на рельсы postgres.
Итак.
БД hives это пул основаный по территориальному признаку.
БД nodes это нода на которой крутится прокси.
БД operators это информация об мобильном операторе.
Это пул мобильных проксей, разных операторов, который размещен территориально.
Есть тщвуы, назовем их bee101,bee102 итд (на деле это одноплатные компьютеры OrangePi с модемами 4G).
Есть так назывемые соты Hive, в каждой соте есть свой пул nodes.
heve100 относятся прокси bee101, bee102 и тд.
heve200 относятся прокси bee201, bee202 и тд.
 */

create table hives (
       id integer generated by default as identity,
       name varchar(50) not null,
       created_at timestamp default current_timestamp,
       updated_at timestamp,

       constraint hives_id_pkey primary key (id)
);

create table nodes (
       id integer generated by default as identity,
       hive_id integer not null,
       hostname varchar (50) not null,
       address varchar (50) not null,
       created_at timestamp default current_timestamp,
       updated_at timestamp,
	
       constraint nodes_id_pkey primary key (id),
       constraint nodes_hostname_key unique (hostname),
       constraint nodes_hive_id_fkey foreign key (hive_id) references hives(id) on delete restrict
);


create table operators (
       id integer generated by default as identity,
       node_id integer not null,
       operator varchar(50) not null,
       phone varchar(50) not null,
       created_at timestamp default current_timestamp,
       updated_at timestamp,

       constraint operators_id_pkey primary key (id),
       constraint operators_node_id_fkey foreign key (node_id) references nodes(id) on delete cascade
);



/* 
 продолжение следует.......
 */
