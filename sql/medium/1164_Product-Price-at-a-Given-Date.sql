select distinctProducts.product_id, coalesce(price, 10) as price
from
(
    select distinct product_id
    from Products
) distinctProducts
left join
(
    select distinct
        product_id,
        first_value(new_price) over (partition by product_id order by change_date desc) as price
    from Products
    where change_date <= '2019-08-16'
) as lastPrice
on distinctProducts.product_id = lastPrice.product_id
