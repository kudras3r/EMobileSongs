
## Installation

Before starting you need to have the pg database server running.

```bash
  git clone https://github.com/kudras3r/EMobileSongs.git
  cd EMobileSongs/
```

In root create .env file:

```env
  DB_HOST=localhost
  DB_USER=user
  DB_PASS=pass
  DB_NAME=user_db
  DB_PORT=5432

  LOG_LEVEL=DEBUG
  HELP_API_HOST=localhost/api

  SERVER_ADDR="localhost:8080"
  SERVER_RW_TIMEOUT=5s
  SERVER_IDLE_TIMEOUT=60s
```

Run:
```bash
  cd cmd/emobile/
  go run main.go
```
    