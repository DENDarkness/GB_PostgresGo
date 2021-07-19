/*
Для этого запроса индекс уже создан так-как столбец hostname является уникальным;
*/
select hostname, address from nodes where hostname='bee101';
/*
Для этого запроса можно проиндексировать поле operator (при условии большого колличества записей)
*/
select node_id, phone from operators where operator='Megafon';
/*
Для этого запроса можно проиндексировать поле phone (при условии большого колличества записей)
*/
select operator, phone from operators where phone like '+7(965)%';
/*
Для этого запроса можно проиндексировать поле address (при условии большого колличества записей)
*/
select hive_id, hostname from nodes where address like '1.1.1.1%';