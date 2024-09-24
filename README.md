# Enterprise Branch Management

An enterprise branch management API designed to follow these [specifications](specifications.md)

## Getting Started

### DB

Start Postgres and Adminer containers. Connect to Adminer: http://localhost:8080/

```sh
docker-compose up -d
```

#### Applying migrations

Install the Goose migration tool.

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Install godotenv to handle the environment variables in .env file.

```sh
 go install github.com/joho/godotenv/cmd/godotenv@latest
```

Apply the migrations.

```sh
godotenv goose up
```

### Start the API

```
go mod download
go run .
```

I like to use [HTTPie](https://httpie.io/) CLI to interact with the API.

```sh
export API_URL=http://localhost:5000/api/v1

http $API_URL/branches
http POST $API_URL/branches name=Canada
http POST $API_URL/branches name=Montreal parent_id:=1
http PATCH $API_URL/branches/2 name=Montréal
http GET $API_URL/branches/2
http POST $API_URL/requirements name=anglais details=C1
http POST $API_URL/requirements name=french details=C2
http POST $API_URL/requirements name=spanish details=B2
http POST $API_URL/branches/1/requirements requirements:='[1,2]'
http PUT $API_URL/branches/1/requirements requirements:='[1]'
http POST $API_URL/branches/2/requirements requirements:='[3]'
```

## Development

Use [air](https://github.com/air-verse/air) for live reload.

```sh
# simply...
air
```

Or vscode **Run and Debug** to use the debugger fonctionalities.

### Tests

#### Integration

At the moment, we use the local database for the integration tests. We might want to separate have separate test and dev databases.

Start from a clean state:

```sh
# psql -h localhost -U postgres -d postgres -f scripts/drop_all_tables.sql
godotenv goose down
godotenv goose up
```

Run all tests in verbose mode.

```sh
go test -v ./...
```

### Write migration

Create a migration with sequential number.

```sh
godotenv goose -s create add_some_column sql
```

## TODO

- Make children branches inherit requirements from parents.
- Add middleware to handle cross-cutting concerns like authentication.
