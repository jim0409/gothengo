SQL Schema

/*
Create table Person (PersonId int, FirstName varchar(255), LastName varchar(255))
Create table Address (AddressId int, PersonId int, City varchar(255), State varchar(255))
Truncate table Person
insert into Person (PersonId, LastName, FirstName) values ('1', 'Wang', 'Allen')
Truncate table Address
insert into Address (AddressId, PersonId, City, State) values ('1', '2', 'New York City', 'New York')
*/


Table: Person
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| PersonId    | int     |
| FirstName   | varchar |
| LastName    | varchar |
+-------------+---------+
PersonId is the primary key column for this table.


Table: Address
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| AddressId   | int     |
| PersonId    | int     |
| City        | varchar |
| State       | varchar |
+-------------+---------+
AddressId is the primary key column for this table.
 

Write a SQL query for a report that provides the following information for each person in the Person table,
regardless if there is an address for each of those people:

FirstName, LastName, City, State


--
SELECT tb1.FirstName AS FirstName, tb1.LastName AS LastName, tb2.City AS City, tb2.State AS State
FROM Person AS tb1
LEFT JOIN Address AS tb2
ON tb1.PersonID = tb2.PersonID;


-- GOAL

PersonID	FirstName	LastName	AddressID	City	State
1	Juilie	Lin	11	Las Vegas	Nevada
3	Emilie	Liu	36	Seattle	Washington
4	Peter	Wang	69	Boston	Massachusetts
2	Vincent	Li	26	Washington	DC


/*
solution:
- https://leetcode.com/problems/combine-two-tables/discuss/520683/MSSQL-Left-Join

select firstname, lastname, city, state from person p left join address a on p.personid=a.personid
*/