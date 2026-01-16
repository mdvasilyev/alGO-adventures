select *
from Cinema c
where description != 'boring' and id % 2 = 1
order by rating desc
