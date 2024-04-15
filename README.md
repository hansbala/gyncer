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

### Environment Setup

`cp .env.template .env` and make the necessary changes

Run this before starting docker

### Local API and database

Docker is amazing.

`docker-compose up --build -d` to start all services.

To get back:
`docker-compose down`

To also delete the persistent MySQL volume:
`docker-compose down --volumes`