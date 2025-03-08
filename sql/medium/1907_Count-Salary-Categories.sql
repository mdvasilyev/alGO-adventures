(select 'Low Salary' category, count(*) accounts_count from Accounts where income < 20000)
union all
(select 'Average Salary' category, count(*) accounts_count from Accounts where income between 20000 and 50000)
union all
(select 'High Salary' category, count(*) accounts_count from Accounts where income > 50000)
