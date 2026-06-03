select sell_date, count(distinct product) num_sold, string_agg(distinct product, ',' order by product) products
from Activities
group by sell_date
order by sell_date
