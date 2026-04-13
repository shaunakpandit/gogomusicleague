-- all submissions with total points_awarded
SELECT
    s.title,
    SUM(v.points_assigned) AS total_votes
FROM submissions AS s
JOIN competitors AS c
    ON c.id = s.submitter_id
JOIN votes AS v
    ON v.spotify_uri = s.spotify_uri
   AND v.round_id = s.round_id
WHERE c.name = 'shaunakpandit'
GROUP BY s.id, s.title
ORDER BY total_votes DESC;

-- total points awarded by competitor
select
    v.voter_id,
    vc.name as voter_name,
    sum(v.points_assigned) as points
from competitors as c
join submissions as s
    on s.submitter_id = c.id
join votes as v
    on s.spotify_uri = v.spotify_uri
   and s.round_id = v.round_id
join competitors as vc
    on vc.id = v.voter_id
where c.name = 'shaunakpandit'
group by v.voter_id, vc.name
order by points desc;


