select round(count(*)::numeric / (select count(distinct player_id) from Activity), 2) fraction
from Activity
where (player_id, event_date) in (
    select player_id, min(event_date) + 1
    from Activity
    group by player_id)
