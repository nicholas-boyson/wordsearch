package search

import (
	"fmt"
	"testing"

	"github.com/nicholas-boyson/wordsearch/internal/organizations"
	"github.com/nicholas-boyson/wordsearch/internal/tickets"
	"github.com/nicholas-boyson/wordsearch/internal/users"
	"github.com/stretchr/testify/assert"
)

func TestValidSearchTerms(t *testing.T) {
	tests := []struct {
		test   string
		group  string
		ident  string
		result bool
	}{
		{
			test:   "OrganizationValidIdent",
			group:  "Organizations",
			ident:  "_id",
			result: true,
		},
		{
			test:   "OrganizationInvalidIdent",
			group:  "Organizations",
			ident:  "invalid",
			result: false,
		},
		{
			test:   "TicketsValidIdent",
			group:  "Tickets",
			ident:  "_id",
			result: true,
		},
		{
			test:   "TicketsInvalidIdent",
			group:  "Tickets",
			ident:  "invalid",
			result: false,
		},
		{
			test:   "UsersValidIdent",
			group:  "Users",
			ident:  "_id",
			result: true,
		},
		{
			test:   "UsersInvalidIdent",
			group:  "Users",
			ident:  "invalid",
			result: false,
		},
		{
			test:   "UnknownGroup",
			group:  "UnknownGroup",
			ident:  "UnknownGroup",
			result: false,
		},
		{
			test:   "EmptyGroup",
			group:  "",
			ident:  "_id",
			result: false,
		},
		{
			test:   "ValidGroupEmptyIdent",
			group:  "Users",
			ident:  "",
			result: false,
		},
	}

	for _, tt := range tests {
		result := ValidSearchTerms(tt.group, tt.ident)
		assert.Equal(t, tt.result, result)
	}
}

func BenchmarkOrganizationsSearch(b *testing.B) {
	orgs, err := organizations.LoadOrganizations("")
	assert.Nil(b, err)
	tickets, err := tickets.LoadTickets("")
	assert.Nil(b, err)
	users, err := users.LoadUsers("")
	assert.Nil(b, err)
	searchRequest := Search{
		Group:         "Organizations",
		Value:         "125",
		Ident:         "_id",
		Organizations: orgs,
		Tickets:       tickets,
		Users:         users,
	}

	var searchResult = SearchResult{}
	for i := 0; i < b.N; i++ {
		searchResult = SearchData(searchRequest)
	}
	fmt.Println(searchResult)
}

func BenchmarkTicketsSearch(b *testing.B) {
	orgs, err := organizations.LoadOrganizations("")
	assert.Nil(b, err)
	tickets, err := tickets.LoadTickets("")
	assert.Nil(b, err)
	users, err := users.LoadUsers("")
	assert.Nil(b, err)
	searchRequest := Search{
		Group:         "Tickets",
		Value:         "125",
		Ident:         "organization_id",
		Organizations: orgs,
		Tickets:       tickets,
		Users:         users,
	}

	var searchResult = SearchResult{}
	for i := 0; i < b.N; i++ {
		searchResult = SearchData(searchRequest)
	}
	fmt.Println(searchResult)
}

func BenchmarkUsersSearch(b *testing.B) {
	orgs, err := organizations.LoadOrganizations("")
	assert.Nil(b, err)
	tickets, err := tickets.LoadTickets("")
	assert.Nil(b, err)
	users, err := users.LoadUsers("")
	assert.Nil(b, err)
	searchRequest := Search{
		Group:         "Users",
		Value:         "125",
		Ident:         "organization_id",
		Organizations: orgs,
		Tickets:       tickets,
		Users:         users,
	}

	var searchResult = SearchResult{}
	for i := 0; i < b.N; i++ {
		searchResult = SearchData(searchRequest)
	}
	fmt.Println(searchResult)
}
