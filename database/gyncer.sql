-- TODO: split into multiple databases

CREATE DATABASE IF NOT EXISTS Gyncer;
USE Gyncer;

CREATE TABLE IF NOT EXISTS Users (
    -- id is the SHA256 hash of the email
    id VARCHAR(64),
    email VARCHAR(255) NOT NULL UNIQUE,
    -- Length appropriate for bcrypt
    hashed_password CHAR(255) NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Syncs (
    -- unique id of the sync
    id INT NOT NULL AUTO_INCREMENT,
    -- user id (hash of the user email)
    user_id VARCHAR(64) NOT NULL,
    source_datasource VARCHAR(20) NOT NULL,
    -- TODO: 100 might be too small
    source_playlist_id VARCHAR(100) NOT NULL,
    destination_datasource VARCHAR(20) NOT NULL,
    -- TODO: 100 might be too small
    destination_playlist_id VARCHAR(100) NOT NULL,
    -- frequency is always defined in hours
    -- default is once every 24 hours
    -- so if sync frequency is 2 it syncs once every 2 hours
    sync_frequency INT NOT NULL,
    -- timestamp the last sync happened
    last_sync_time TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE (source_playlist_id, destination_playlist_id),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS SpotifyKeys (
    -- unique id for this key
    id INT NOT NULL AUTO_INCREMENT,
    -- user id associated with this key (hash of the user email)
    user_id VARCHAR(64) UNIQUE NOT NULL,
    -- access token (this is used to make the API call)
    access_token VARCHAR(500) NOT NULL,
    -- `refresh_token` along with `access_token` is used to retrieve new `access_token`
    refresh_token VARCHAR(500) NOT NULL,
    -- `expiry` is the time when this token expires
    expiry TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS GoogleKeys (
    -- unique id for this key
    id INT NOT NULL AUTO_INCREMENT,
    -- user id associated with this key (hash of the user email)
    user_id VARCHAR(64) UNIQUE NOT NULL,
    -- auth code spotify gave us
    auth_code VARCHAR(1000) NOT NULL,
    -- access token (this is used to make the API call)
    access_token VARCHAR(500) NOT NULL,
    -- `refresh_token` along with `auth_code` is used to retrieve new `access_token`
    refresh_token VARCHAR(500) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);