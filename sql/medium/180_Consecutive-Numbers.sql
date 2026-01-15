select distinct t.num ConsecutiveNums
from (
    select num,
    lead(num, 1, null) over (order by id) as next_num,
    lead(num, 2, null) over (order by id) as next_next_num
    from logs
) t
where t.num = t.next_num and t.next_num = t.next_next_num

-- select distinct l1.num ConsecutiveNums
-- from Logs l1 join Logs l2 on l1.id = l2.id + 1 join Logs l3 on l2.id = l3.id + 1
-- where l1.num = l2.num and l2.num = l3.num
