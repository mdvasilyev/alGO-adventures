select e1.employee_id, e1.name, count(e2.employee_id) reports_count, round(avg(e2.age)) average_age
from Employees e1 cross join Employees e2
where e2.reports_to is not null and e1.employee_id = e2.reports_to
group by e1.employee_id, e1.name
order by e1.employee_id
