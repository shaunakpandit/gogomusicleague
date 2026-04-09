-- run these commands via cli. workbench has lacking permissions for local infile operations

LOAD DATA LOCAL INFILE '/Users/shaunak/Documents/Projects/musicLeague/season1/competitors.csv'
INTO TABLE competitors
CHARACTER SET utf8mb4
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(@id, @name)
SET
  id = @id,
  name = @name;

LOAD DATA LOCAL INFILE '/Users/shaunak/Documents/Projects/musicLeague/season1/rounds.csv'
INTO TABLE rounds
CHARACTER SET utf8mb4
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(@id, @created, @name, @description, @playlist_url)
SET
  id = @id,
  created = STR_TO_DATE(@created, '%Y-%m-%dT%H:%i:%sZ'),
  name = @name,
  description = @description,
  playlist_url = @playlist_url;

LOAD DATA LOCAL INFILE '/Users/shaunak/Documents/Projects/musicLeague/season1/submissions.csv'
INTO TABLE submissions
CHARACTER SET utf8mb4
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(@spotify_uri, @title, @album, @artists, @submitter_id, @created, @comment, @round_id, @visible_to_voters)
SET
  spotify_uri = @spotify_uri,
  title = @title,
  album = NULLIF(@album, ''),
  artists = @artists,
  submitter_id = @submitter_id,
  created = STR_TO_DATE(@created, '%Y-%m-%dT%H:%i:%sZ'),
  comment = NULLIF(@comment, ''),
  round_id = @round_id,
  visible_to_voters = CASE
    WHEN @visible_to_voters = 'Yes' THEN 1
    WHEN @visible_to_voters = 'No' THEN 0
    ELSE NULL
  END;

LOAD DATA LOCAL INFILE '/Users/shaunak/Documents/Projects/musicLeague/season1/votes.csv'
INTO TABLE votes
CHARACTER SET utf8mb4
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(@spotify_uri, @voter_id, @created, @points_assigned, @comment, @round_id)
SET
  spotify_uri = @spotify_uri,
  voter_id = @voter_id,
  created = STR_TO_DATE(@created, '%Y-%m-%dT%H:%i:%sZ'),
  points_assigned = @points_assigned,
  comment = NULLIF(@comment, ''),
  round_id = @round_id;
