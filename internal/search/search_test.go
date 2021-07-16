package search

import (
	"fmt"
	"testing"

	"github.com/nicholas-boyson/wordsearch/internal/organizations"
	"github.com/nicholas-boyson/wordsearch/internal/tickets"
	"github.com/nicholas-boyson/wordsearch/internal/users"
	"github.com/stretchr/testify/assert"
)

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
