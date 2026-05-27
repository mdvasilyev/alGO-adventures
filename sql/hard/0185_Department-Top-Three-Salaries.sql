with employee_department as (
    select d.name as Department, e.name as Employee, e.salary as Salary, dense_rank() over (partition by e.departmentId order by e.salary desc) rnk
    from Employee e
    join Department d
    on e.departmentId = d.id
)
select Department, Employee, Salary
from employee_department
where rnk <= 3
