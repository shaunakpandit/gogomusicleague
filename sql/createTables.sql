CREATE DATABASE IF NOT EXISTS musicleague
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_0900_ai_ci;

USE musicleague;

CREATE TABLE competitors (
    id CHAR(32) NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE rounds (
    id CHAR(32) NOT NULL,
    created DATETIME NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    playlist_url VARCHAR(1024) NULL,
    PRIMARY KEY (id),
    KEY idx_rounds_created (created)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE submissions (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    spotify_uri VARCHAR(255) NOT NULL,
    title VARCHAR(500) NOT NULL,
    album VARCHAR(500) NULL,
    artists VARCHAR(500) NOT NULL,
    submitter_id CHAR(32) NOT NULL,
    created DATETIME NOT NULL,
    comment TEXT NULL,
    round_id CHAR(32) NOT NULL,
    visible_to_voters BOOLEAN NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT uq_submissions_spotify_round UNIQUE (spotify_uri, round_id),
    CONSTRAINT fk_submissions_submitter
        FOREIGN KEY (submitter_id) REFERENCES competitors(id),
    CONSTRAINT fk_submissions_round
        FOREIGN KEY (round_id) REFERENCES rounds(id),
    KEY idx_submissions_spotify_uri (spotify_uri),
    KEY idx_submissions_round_id (round_id),
    KEY idx_submissions_submitter_id (submitter_id),
    KEY idx_submissions_created (created)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE votes (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    spotify_uri VARCHAR(255) NOT NULL,
    voter_id CHAR(32) NOT NULL,
    created DATETIME NOT NULL,
    points_assigned INT NOT NULL,
    comment TEXT NULL,
    round_id CHAR(32) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_votes_voter
        FOREIGN KEY (voter_id) REFERENCES competitors(id),
    CONSTRAINT fk_votes_round
        FOREIGN KEY (round_id) REFERENCES rounds(id),
    KEY idx_votes_spotify_uri (spotify_uri),
    KEY idx_votes_round_id (round_id),
    KEY idx_votes_voter_id (voter_id),
    KEY idx_votes_created (created)
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
