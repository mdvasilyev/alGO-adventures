select
    to_char(trans_date, 'YYYY-MM') as month, country,
    count(*) trans_count,
    count(state) filter (where state = 'approved') approved_count,
    sum(amount) trans_total_amount,
    coalesce(sum(amount) filter (where state = 'approved'), 0) approved_total_amount
from Transactions t
group by country, month
order by month asc, trans_count desc
