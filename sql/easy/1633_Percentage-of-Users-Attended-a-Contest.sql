select contest_id, round(count(contest_id) * 100.00 / (select count(*) from Users), 2) percentage
from Register
group by contest_id
order by percentage desc, contest_id asc
