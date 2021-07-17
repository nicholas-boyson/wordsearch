package display

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWelcome(t *testing.T) {
	welcome := welcome()
	assert.Equal(t, "Welcome to Zendesk Search\nType 'quit' to exit at any time, Press 'Enter' to continue", welcome)
}

func TestSelectSearchOptions(t *testing.T) {
	selectSearchOptions := selectSearchOptions()
	assert.Equal(t, "          Select search options:\n          * Press 1 to search Zendesk\n          * Press 2 to view a list of searchable fields\n          * Type 'quit' to exit", selectSearchOptions)
}

func TestListSearchableFields(t *testing.T) {
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

	searchFieldsDisplay := listSearchableFields()
	assert.Equal(t, searchFields, searchFieldsDisplay)
}

func TestSelectGroupOptions(t *testing.T) {
	selectGroupOptions := selectGroupOptions()
	assert.Equal(t, "Select 1) Users or 2) Tickets or 3) Organizations", selectGroupOptions)
}

func TestEnterSearchTerm(t *testing.T) {
	enterTerm := enterSearchTerm()
	assert.Equal(t, "Enter search term", enterTerm)
}

func TestEnterSearchValue(t *testing.T) {
	enterValue := enterSearchValue()
	assert.Equal(t, "Enter search value", enterValue)
}

func TestNoResultFound(t *testing.T) {
	noResultMsg := noResultFound()
	assert.Equal(t, "No results found", noResultMsg)
}

func TestInvalidSearchTerm(t *testing.T) {
	msg := invalidSearchTerm()
	assert.Equal(t, "Invalid search term", msg)
}
