select person_name
from (select person_name, turn, sum(weight) over (order by turn) cumulative_weight from Queue)
where cumulative_weight <= 1000
order by turn desc limit 1
