-- grab all submissions and total votes by user
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
