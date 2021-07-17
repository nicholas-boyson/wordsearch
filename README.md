# wordsearch
Application to perform a search by user, tickets or orginazation.
Supported search fields
| Users           | Tickets         | Organizations  |
|-----------------|-----------------|----------------|
| _id             | _id             | _id            |
| url             | url             | url            |
| external_id     | external_id     | external_id    |
| name            | created_at      | name           |
| alias           | type            | domain_names   |
| created_at      | subject         | created_at     |
| active          | description     | details        |
| verified        | priority        | shared_tickets |
| shared          | status          | tags           |
| locale          | submitter_id    |                |
| timezone        | assignee_id     |                |
| last_login_at   | organization_id |                |
| email           | tags            |                |
| phone           | has_incidents   |                |
| signature       | due_at          |                |
| organization_id | via             |                |
| tags            |                 |                |
| suspended       |                 |                |
| role            |                 |                |

## Assumptions
Every field can be used in the search.
When searching by user the application will display the users details plus the linked organizations details.
When searching by ticket the application will display the ticket information plus the linked organizations details.
When searching by organization the applications will display the organization details plus list of linked users (subset of information) and a list of linked tickets (subset of information).
When a search returns multiple results and the search was not by organization id, the result will display a grid of those returned results.
When a search returns multiple results and the search was by organization id, the result will display a grid of those returned results and the linked organization id.

## Usage
```
// running using go, requires installation of golang
go run .

// running using exe if you have created the build for windows
.\wordsearch.exe
```

## Build for Windows (creates an exe)
```
go build .
```

## Running unit tests
```
go test ./...
go test -v ./...
```