with employee_department as (
    select d.name as Department, e.name as Employee, e.salary as Salary, rank() over (partition by departmentId order by salary desc) as rnk
    from Employee e
    join Department d
    on e.departmentId = d.id
)
select Department, Employee, Salary
from employee_department
where rnk = 1
