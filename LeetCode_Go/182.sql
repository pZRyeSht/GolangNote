1.
# Write your MySQL query statement below
select Email
from Person
group by Email
having count(Email) > 1;
2.
# Write your MySQL query statement below
select Email from
(
  select Email, count(Email) as num
  from Person
  group by Email
) as temp
where num > 1;