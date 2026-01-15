select min(name) results
from Users
where user_id in (
    select user_id from MovieRating
    group by user_id
    having count(*) = (select max(cnt) from (select count(*) cnt from MovieRating group by user_id))
)
union all
select min(title) results
from Movies
where movie_id in (
    select movie_id
    from MovieRating
    where created_at >= '2020-02-01' and created_at < '2020-03-01'
    group by movie_id
    having avg(rating) = (
        select max(rating) from (
            select avg(rating) rating
            from MovieRating
            where created_at >= '2020-02-01' and created_at < '2020-03-01'
            group by movie_id
        )
    )
)
