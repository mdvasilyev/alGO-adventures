select user_id, round(avg(case when c.action = 'confirmed' then 1 else 0 end)::numeric, 2) confirmation_rate
from Signups s left join Confirmations c
using(user_id)
group by user_id
