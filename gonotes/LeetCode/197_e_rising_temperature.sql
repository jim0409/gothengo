SQL Schema

Create table If Not Exists Weather (Id int, RecordDate date, Temperature int)
Truncate table Weather
insert into Weather (Id, RecordDate, Temperature) values ('1', '2015-01-01', '10')
insert into Weather (Id, RecordDate, Temperature) values ('2', '2015-01-02', '25')
insert into Weather (Id, RecordDate, Temperature) values ('3', '2015-01-03', '20')
insert into Weather (Id, RecordDate, Temperature) values ('4', '2015-01-04', '30')


Given a Weather table, write a SQL query to find all dates' Ids with higher temperature compared to its previous (yesterday's) dates.

+---------+------------------+------------------+
| Id(INT) | RecordDate(DATE) | Temperature(INT) |
+---------+------------------+------------------+
|       1 |       2015-01-01 |               10 |
|       2 |       2015-01-02 |               25 |
|       3 |       2015-01-03 |               20 |
|       4 |       2015-01-04 |               30 |
+---------+------------------+------------------+
For example, return the following Ids for the above Weather table:

+----+
| Id |
+----+
|  2 |
|  4 |
+----+
-- 

-- only 48.28% faster
SELECT tb1.Id AS Id
FROM Weather AS tb1
LEFT JOIN Weather AS tb2
ON tb1.Temperature > tb2.Temperature AND DATEDIFF(tb1.RecordDate, tb2.RecordDate) =1
WHERE tb1.RecordDate > tb2.RecordDate;

-- faster than 95.84%
SELECT W1.Id
    FROM Weather W1
    INNER JOIN Weather W2
WHERE DATEDIFF(W1.RecordDate, W2.RecordDate) = 1 AND        
      W1.Temperature > W2.Temperature



-- solution1:
-- https://leetcode.com/problems/rising-temperature/discuss/534674/MS-SQL-Faster-than-72.44

-- select w2.Id
-- from Weather w1, Weather w2
-- where DATEDIFF(day, w1.RecordDate, w2.RecordDate) = 1
-- and w1.Temperature < w2.Temperature


-- solution2:
-- https://leetcode.com/problems/rising-temperature/discuss/549259/Faster-than-99.90-of-MySQL
-- FROM    
--     (SELECT Id,
--         @greaterthan := IF(DATEDIFF(RecordDate, @prevDate) = 1, 
--                            IF(@prevTemp < Temperature, 1, 0), 0) as Ind,
--         @prevTemp := Temperature as Temp,
--         @prevDate := RecordDate as RecordDate
--     FROM (SELECT @prevTemp := Null, @greaterthan := Null, @prevDate := Null) Temp, Weather W
--     ORDER BY RecordDate) Temp
-- WHERE Ind = 1


