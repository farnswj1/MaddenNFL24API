# Madden NFL 24 API
This is an API that provides data on players from Madden NFL 24.

## Setup
The project uses the following:
- Go
- Gin
- PostgreSQL
- Redis
- Docker
- Docker Compose

For additional information on project specifications, see the ```go.mod``` file in ```app```.

### Setting up PostgreSQL
In the ```postgres/``` directory, create a ```.env``` file
that contains the following environment variables:
```
POSTGRES_DB=madden_24
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
```

### Setting up the API
In the ```app/``` directory, create a ```.env``` file
that contains the following environment variables:
```
GIN_MODE=release
CORS_ALLOWED_ORIGINS=http://localhost http://127.0.0.1
DATABASE_URL=postgresql://postgres:password@postgres:5432/madden_24
REDIS_URL=redis://redis:6379/0
SECRET_KEY=secret
```

## Building
The project uses Docker. Ensure Docker and Docker Compose are installed before continuing.

To build, run ```docker compose build```

## Running
To run the web API, run ```docker compose up -d```, then go to http://localhost using your web browser.
