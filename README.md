# Instagram Bot (Go)

A Go-based Instagram bot that monitors comments and sends predefined DMs using Gin and PostgreSQL.

## Setup
1. Clone the repository.
2. Run `docker-compose up --build` to start the services.
3. Apply migrations: `docker exec instagram-bot_web_1 migrate -path db/migrations -database $DATABASE_URL up`
4. Access the API at `http://localhost:8080`.

## Requirements
- Go 1.20
- Docker

## License
MIT
