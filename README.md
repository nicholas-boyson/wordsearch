# wordsearch
Application to perform multiple searchs across a data source of users, tickets and orginazations.
The applications supported search fields:
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
Users and tickets are associated to an organization by the organization_id field. 
When performing a search by the organization group the user will be presented with the
details of the organization record plus a list of associated tickets and users.
When performing a search by the ticket group the user will be presented with the details of 
that ticket, and the details of the linked organization.
When performing a search by the user group the user will be presented with the details of the
user, and the details of the linked organization.
For any group if there is more than one result found the user will be presented with a list of
the found information plus additional details to help refine the search.
No results found will result in a message back to the user and return them to the start of the search.
You can exit the application anytime by entering 'quit'

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
// run all tests from main directory and subdirectories
go test ./...
// run all tests from the main directory and subdirectories with more information on the tests
go test -v ./...
// run all benchmark tests from the main directory and across the subdirectoris.
go test -bench . ./...
```