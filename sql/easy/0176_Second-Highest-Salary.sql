with ratedSalaries as (
    select salary, dense_rank() over (order by salary desc) rnk
    from Employee
)
select max(salary) SecondHighestSalary
from ratedSalaries
where rnk = 2
