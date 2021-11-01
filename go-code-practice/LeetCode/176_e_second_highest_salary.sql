SQL Schema
/*
Create table If Not Exists Employee (Id int, Salary int)
Truncate table Employee
insert into Employee (Id, Salary) values ('1', '100')
insert into Employee (Id, Salary) values ('2', '200')
insert into Employee (Id, Salary) values ('3', '300')
*/

Write a SQL query to get the second highest salary from the Employee table.

+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
For example, given the above Employee table, the query should return 200 as the second highest salary.
If there is no second highest salary, then the query should return null.

+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+

--

-- 先寫 SQL 把情境列出來，再依照情境下的結果去篩選
/*
1. 先找出 Employee 表中的最高 Salary
SELECT max(Salary) FROM Employee

2. 假設最高的 Salary 值為 maxSalary，在小於最高Salary的情況下，尋找最大Salary
SELECT max(Salary) FROM Employee WHERE Salary < maxSalary
*/

-- solution 1.
SELECT max(Salary) AS SecondHighestSalary	
FROM Employee	
WHERE Salary < (SELECT max(Salary)
    FROM Employee);
	
	
-- solution 2. 可以用ORDER BY DESC取代max	
SELECT max(Salary) AS SecondHighestSalary	
FROM Employee	
WHERE Salary < ( SELECT Salary 
	FROM Employee
	ORDER BY Salary DESC
	LIMIT 1);


/*
solution:
- https://leetcode.com/problems/second-highest-salary/discuss/520757/Faster-than-select-MAX-solution-and-can-work-for-any-nth-highest

SELECT (CASE 
        WHEN (SELECT COUNT(DISTINCT Salary) FROM Employee) < 2 
        THEN NULL 
        ELSE (SELECT Salary FROM Employee 
              ORDER BY Salary DESC LIMIT 1,1) 
        END) AS SecondHighestSalary
*/