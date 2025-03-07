select distinct t.num ConsecutiveNums
from (
    select num,
    lead(num, 1, null) over (order by id) as next_num,
    lead(num, 2, null) over (order by id) as next_next_num
    from logs
) t
where t.num = t.next_num and t.next_num = t.next_next_num
