# Gyncer

## Description

Gyncer (short for Go Music Syncer) is a collection of REST API endpoints implemented in Go to faciliate music syncing across various platform.

The though is to support these platforms out of the box:
* Spotify
* Youtube Music
* Tidal (mainly because I'm a Hi-Fi music fan and need to push my speakers)

## API structure

Unprotected routes:
- `/users` POST: Creates a new user
- `/login` POST: Logs a user in and returns a JWT token to access protected routes

Protected Routes:
- `/sync` POST: Creates a new sync


## Local Development

### `.env`

`cp .env.template .env` and make the necessary changes

Run this before any of the below steps

### Database

You need a local database setup to test the API. This command sets up a local SQL database

```
docker-compose up -d
```

Some more helpful database utilities are included in `database/local_db_tools.sh`

### Local API

`go run .` sets up the API running locally on `localhost:8080`.
TODO: Migrate this to docker too
