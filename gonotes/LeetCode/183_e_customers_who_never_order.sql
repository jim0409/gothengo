SQL Schema
/*
Create table If Not Exists Customers (Id int, Name varchar(255))
Create table If Not Exists Orders (Id int, CustomerId int)
Truncate table Customers
insert into Customers (Id, Name) values ('1', 'Joe')
insert into Customers (Id, Name) values ('2', 'Henry')
insert into Customers (Id, Name) values ('3', 'Sam')
insert into Customers (Id, Name) values ('4', 'Max')
Truncate table Orders
insert into Orders (Id, CustomerId) values ('1', '3')
insert into Orders (Id, CustomerId) values ('2', '1')
*/
Suppose that a website contains two tables, the Customers table and the Orders table.
Write a SQL query to find all customers who never order anything.

Table: Customers.

+----+-------+
| Id | Name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+
Table: Orders.

+----+------------+
| Id | CustomerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+
Using the above tables as example, return the following:

+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+

--
SELECT tb1.Name AS Customers
FROM Customers AS tb1
LEFT JOIN Orders AS tb2
ON tb1.Id = tb2.CustomerID
WHERE tb2.Id is null;

/*
solution:
- https://leetcode.com/problems/customers-who-never-order/discuss/509129/simple-left-join

select
name as customers
from customers c
left join orders d
on c.id = d.customerid
where d.id is null;
*/