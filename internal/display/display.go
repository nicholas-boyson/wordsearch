package display

import "fmt"

// Welcome display welcome message
func Welcome() {
	fmt.Println(welcome())
}
func welcome() string {
	return "Welcome to Zendesk Search\nType 'quit' to exit at any time, Press 'Enter' to continue"
}

// SelectSearchOptions display search options message
func SelectSearchOptions() {
	fmt.Println(selectSearchOptions())
}
func selectSearchOptions() string {
	return "          Select search options:\n          * Press 1 to search Zendesk\n          * Press 2 to view a list of searchable fields\n          * Type 'quit' to exit"
}

// ListSearchableFields function to display the searchable fields
func ListSearchableFields() {
	fmt.Println(listSearchableFields())
}
func listSearchableFields() string {
	searchFields := "Searchable fields\n"
	searchFields = searchFields + "| Users           | Tickets         | Organizations  |\n"
	searchFields = searchFields + "|-----------------|-----------------|----------------|\n"
	searchFields = searchFields + "| _id             | _id             | _id            |\n"
	searchFields = searchFields + "| url             | url             | url            |\n"
	searchFields = searchFields + "| external_id     | external_id     | external_id    |\n"
	searchFields = searchFields + "| name            | created_at      | name           |\n"
	searchFields = searchFields + "| alias           | type            | domain_names   |\n"
	searchFields = searchFields + "| created_at      | subject         | created_at     |\n"
	searchFields = searchFields + "| active          | description     | details        |\n"
	searchFields = searchFields + "| verified        | priority        | shared_tickets |\n"
	searchFields = searchFields + "| shared          | status          | tags           |\n"
	searchFields = searchFields + "| locale          | submitter_id    |                |\n"
	searchFields = searchFields + "| timezone        | assignee_id     |                |\n"
	searchFields = searchFields + "| last_login_at   | organization_id |                |\n"
	searchFields = searchFields + "| email           | tags            |                |\n"
	searchFields = searchFields + "| phone           | has_incidents   |                |\n"
	searchFields = searchFields + "| signature       | due_at          |                |\n"
	searchFields = searchFields + "| organization_id | via             |                |\n"
	searchFields = searchFields + "| tags            |                 |                |\n"
	searchFields = searchFields + "| suspended       |                 |                |\n"
	searchFields = searchFields + "| role            |                 |                |\n"
	return searchFields
}

// SelectGroupOptions display group search options to user
func SelectGroupOptions() {
	fmt.Println(selectGroupOptions())
}
func selectGroupOptions() string {
	return "Select 1) Users or 2) Tickets or 3) Organizations"
}

// EnterSearchTerm display enter search term to user
func EnterSearchTerm() {
	fmt.Println(enterSearchTerm())
}
func enterSearchTerm() string {
	return "Enter search term"
}

// EnterSearchValue display enter search value to user
func EnterSearchValue() {
	fmt.Println(enterSearchValue())
}
func enterSearchValue() string {
	return "Enter search value"
}

// NoResultFound display no results found to user
func NoResultFound() {
	fmt.Println(noResultFound())
}
func noResultFound() string {
	return "No results found"
}
