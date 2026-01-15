with offset_6_days as (
    select distinct visited_on
    from Customer
    order by visited_on asc
    offset 6
)

select c2.visited_on, sum(c1.amount) amount, round(sum(c1.amount) / 7.00, 2) average_amount
from Customer c1 join offset_6_days c2
on c1.visited_on between c2.visited_on - 6 and c2.visited_on
group by c2.visited_on
