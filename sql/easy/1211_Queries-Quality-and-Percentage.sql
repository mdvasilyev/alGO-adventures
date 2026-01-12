select query_name,
round(avg(quality)::numeric, 2) quality,
round(avg(is_poor)::numeric * 100, 2) poor_query_percentage
from (select query_name, (rating::numeric / position) quality, case when rating < 3 then 1 else 0 end is_poor from Queries)
group by query_name
