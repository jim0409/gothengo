SQL Schema
/*
Create table If Not Exists Person (Id int, Email varchar(255))
Truncate table Person
insert into Person (Id, Email) values ('1', 'a@b.com')
insert into Person (Id, Email) values ('2', 'c@d.com')
insert into Person (Id, Email) values ('3', 'a@b.com')
*/
Write a SQL query to find all duplicate emails in a table named Person.

+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
For example, your query should return the following for the above table:

+---------+
| Email   |
+---------+
| a@b.com |
+---------+
Note: All emails are in lowercase.

--
SELECT DISTINCT tb1.Email
FROM Person AS tb1
INNER JOIN Person AS tb2
ON tb1.Email = tb2.Email AND tb1.Id != tb2.Id

/*
solution:
- https://leetcode.com/problems/duplicate-emails/discuss/520470/simple-and-work-pretty-good-(greater90)

select distinct p1.Email
from Person p1
inner join
(select Email,count(Email) as c
from person
group by Email)
as p2
on p2.Email = p1.Email
where p2.c >1
*/