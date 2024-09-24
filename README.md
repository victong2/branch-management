# Enterprise Branch Management

An enterprise branch management API designed to follow these [specifications](specifications.md).

## Getting Started

### Database

Start Postgres and Adminer containers. Connect to Adminer: http://localhost:8080/

```sh
docker-compose up -d
```

#### Applying migrations

Install the [Goose](https://github.com/pressly/goose) migration tool.

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Install godotenv to handle the environment variables in [`.env`](.env) file.

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

### Interact with the API

I like to use [HTTPie](https://httpie.io/) CLI to interact with the API.

```sh
export API_URL=http://localhost:5000/api/v1

http $API_URL/branches
http POST $API_URL/branches name=Canada
http POST $API_URL/branches name=Montreal parent_id:=1
http PATCH $API_URL/branches/2 name=Montréal
http GET $API_URL/branches/2
http POST $API_URL/branches name=Merinio parent_id:=2
http POST $API_URL/requirements name=english details=C1
http POST $API_URL/requirements name=français details=C2
http POST $API_URL/requirements name=español details=B2
http POST $API_URL/requirements name=Golang details=Go
http POST $API_URL/branches/1/requirements requirements:='[1,2]'
http PUT $API_URL/branches/1/requirements requirements:='[1]'
http POST $API_URL/branches/2/requirements requirements:='[2]'
http POST $API_URL/branches/2/requirements requirements:='[2]'
http POST $API_URL/branches/3/requirements requirements:='[3,4]'
http GET $API_URL/branches/1/requirements
http GET $API_URL/branches/2/requirements
http GET $API_URL/branches/3/requirements
```

Last command returns:

```json
{
  "data": {
    "id": 3,
    "name": "Merinio",
    "parent_id": 2,
    "requirements": [
      {
        "details": "B2",
        "id": 3,
        "name": "español"
      },
      {
        "details": "Go",
        "id": 4,
        "name": "Golang"
      },
      {
        "details": "C2",
        "id": 2,
        "name": "français"
      },
      {
        "details": "C1",
        "id": 1,
        "name": "english"
      }
    ]
  }
}
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

### Write migrations

Create a migration with sequential number.

```sh
godotenv goose -s create add_some_column sql
```

## Left TODO

- Make sure there is no non-cyclical hierarchy among branches. Enforce when creating branches.
- Add [middlewares](middlewares) to handle concerns like authentication, authorization, etc.
