package search

import (
	"strconv"

	"github.com/nicholas-boyson/wordsearch/internal/display"
	"github.com/nicholas-boyson/wordsearch/internal/organizations"
	"github.com/nicholas-boyson/wordsearch/internal/tickets"
	"github.com/nicholas-boyson/wordsearch/internal/users"
)

//Search search definition
type Search struct {
	Ident         string
	Group         string
	Value         string
	Organizations []organizations.Organization
	Tickets       []tickets.Ticket
	Users         []users.User
}

// SearchResult search result definition
type SearchResult struct {
	Organizations []organizations.Organization
	Tickets       []tickets.Ticket
	Users         []users.User
}

// SearchData search across all data sources linking on organization id when single result or search by organization id
func SearchData(s Search) (result SearchResult) {
	switch s.Group {
	case "Organizations":
		result.Organizations = organizations.SearchOrganizations(s.Organizations, s.Ident, s.Value)
		if len(result.Organizations) == 1 {
			// Only search when there is a single organization returned
			result.Tickets = tickets.SearchTickets(s.Tickets, "organization_id", strconv.Itoa(result.Organizations[0].Id))
			result.Users = users.SearchUsers(s.Users, "organization_id", strconv.Itoa(result.Organizations[0].Id))
		}
	case "Tickets":
		result.Tickets = tickets.SearchTickets(s.Tickets, s.Ident, s.Value)
		if len(result.Tickets) == 1 || (len(result.Tickets) > 0 && s.Ident == "organization_id") {
			// Only link organization details when there is a single ticket returned or the search was on the org id
			result.Organizations = organizations.SearchOrganizations(s.Organizations, "_id", strconv.Itoa(result.Tickets[0].OrganizationId))
		}
	case "Users":
		result.Users = users.SearchUsers(s.Users, s.Ident, s.Value)
		if len(result.Users) == 1 || (len(result.Users) > 0 && s.Ident == "organization_id") {
			// Only link organization details when there is a single user returned or the search was on the org id
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

// SearchResultDisplay determines the display based on group and search results
func SearchResultDisplay(group string, sr SearchResult) {
	switch group {
	case "Organizations":
		display.DisplayOrganizations(sr.Organizations, sr.Tickets, sr.Users)
	case "Tickets":
		if len(sr.Organizations) > 0 {
			display.DisplayTickets(sr.Tickets, sr.Organizations[0])
		} else {
			display.DisplayTickets(sr.Tickets, organizations.Organization{})
		}
	case "Users":
		if len(sr.Organizations) > 0 {
			display.DisplayUsers(sr.Users, sr.Organizations[0])
		} else {
			display.DisplayUsers(sr.Users, organizations.Organization{})
		}
	default:
		display.NoResultFound()
	}
}
