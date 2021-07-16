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

Search by user will display the users details plus the linked organizations details.
Search by ticket will display the ticket information plus the linked organizations details.
Search by organization will display the organization details plus list of linked users (subset of information) and a list of linked tickets (subset of information).

Approach Idea 1:
- Search interface options to search or display list of search field options
- on search the user is presented with Search on users, Search on tickets, or Search on Organizations
- Read the file the user selected e.g. Users and search the file based on search input field and store the found information and pull out the organisation id.
- Then read the file of the related information i.e. organisation

Pros: code is simple to follow and maintain. 
Cons: for each search will need to constantly read the files.

Approach Idea 2:
- Load each source file into slices of structs
- Search interface options to search or display list of search field options
- on search the user is presented with Search on users, Search on tickets, or Search on Organizations
- Search on slice