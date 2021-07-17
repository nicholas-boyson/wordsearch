package search

import (
	"strconv"

	"github.com/nicholas-boyson/wordsearch/internal/organizations"
	"github.com/nicholas-boyson/wordsearch/internal/tickets"
	"github.com/nicholas-boyson/wordsearch/internal/users"
)

type Search struct {
	Ident         string
	Group         string
	Value         string
	Organizations []organizations.Organization
	Tickets       []tickets.Ticket
	Users         []users.User
}

type SearchResult struct {
	Organizations organizations.Organization
	Tickets       []tickets.Ticket
	Users         []users.User
}

func SearchData(s Search) (result SearchResult) {
	switch s.Group {
	case "Organizations":
		result.Organizations = organizations.SearchOrganizations(s.Organizations, s.Ident, s.Value)
		result.Tickets = tickets.SearchTickets(s.Tickets, "organization_id", strconv.Itoa(result.Organizations.Id))
		result.Users = users.SearchUsers(s.Users, "organization_id", strconv.Itoa(result.Organizations.Id))
	case "Tickets":
		result.Tickets = tickets.SearchTickets(s.Tickets, s.Ident, s.Value)
		if len(result.Tickets) > 0 {
			result.Organizations = organizations.SearchOrganizations(s.Organizations, "_id", strconv.Itoa(result.Tickets[0].OrganizationId))
		}
	case "Users":
		result.Users = users.SearchUsers(s.Users, "organization_id", strconv.Itoa(result.Organizations.Id))
		if len(result.Tickets) > 0 {
			result.Organizations = organizations.SearchOrganizations(s.Organizations, "_id", strconv.Itoa(result.Users[0].OrganizationId))
		}
	default:
		return SearchResult{}
	}
	return
}

// ValidSearchTerms return if ident is valid for a group
func ValidSearchTerms(group string, ident string) bool {
	switch group {
	case "Organizations":
		return organizations.ValidSearchTerms(ident)
	case "Tickets":
		return tickets.ValidSearchTerms(ident)
	case "Users":
		return users.ValidSearchTerms(ident)
	default:
		return false
	}
}
