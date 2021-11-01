Write a SQL query to delete all duplicate email entries in a table named `Person`, 
keeping only unique emails based on its smallest Id.

+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
Id is the primary key column for this table.
For example, after running your query, the above Person table should have the following rows:

+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+
Note:

Your output is the whole Person table after executing your sql. Use delete statement.
--

DELETE tb1 
FROM Person AS tb1
-- 要用 INNER JOIN 來
INNER JOIN Person AS tb2
ON tb1.Email = tb2.Email AND tb1.Id > tb2.Id


-- solution:
-- DELETE P1
-- FROM person P1 INNER JOIN person P2
-- ON P1.email = P2.email AND P1.id > P2.id;

