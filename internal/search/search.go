package search

import (
	"strconv"
	"sync"

	"github.com/nicholas-boyson/wordsearch/internal/display"
	"github.com/nicholas-boyson/wordsearch/internal/organizations"
	"github.com/nicholas-boyson/wordsearch/internal/tickets"
	"github.com/nicholas-boyson/wordsearch/internal/users"
)

const (
	//SearchGroupUsers "Users"
	SearchGroupUsers = "Users"
	//SearchGroupTickets "Tickets"
	SearchGroupTickets = "Tickets"
	//SearchGroupOrganizations "Organizations"
	SearchGroupOrganizations = "Organizations"
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

const workerGrpMax = 10

// SearchData search across all data sources linking on organization id when single result or search by organization id
func SearchData(s Search) (result SearchResult) {
	switch s.Group {
	case SearchGroupOrganizations:
		var workerGrp sync.WaitGroup
		searchOrgChan := make(chan []organizations.Organization, workerGrpMax)
		split := (len(s.Organizations) / workerGrpMax)
		for i := 0; i < workerGrpMax; i++ {
			workerGrp.Add(1)
			var orgs []organizations.Organization
			if i == workerGrpMax-1 {
				orgs = s.Organizations[split*i:]
			} else {
				orgs = s.Organizations[split*i : split*(i+1)]
			}
			go func() {
				defer workerGrp.Done()
				searchOrgChan <- organizations.SearchOrganizations(orgs, s.Ident, s.Value)
			}()
			result.Organizations = append(result.Organizations, <-searchOrgChan...)
		}
		workerGrp.Wait()

		if len(result.Organizations) == 1 {
			// Only search when there is a single organization returned
			var workerGrp sync.WaitGroup
			searchTicketChan := make(chan []tickets.Ticket, workerGrpMax)
			searchUserChan := make(chan []users.User, workerGrpMax)
			splitTicket := (len(s.Tickets) / workerGrpMax)
			splitUsers := (len(s.Users) / workerGrpMax)
			for i := 0; i < workerGrpMax; i++ {
				workerGrp.Add(1)
				var ticketsInput []tickets.Ticket
				var usersInput []users.User
				if i == workerGrpMax-1 {
					ticketsInput = s.Tickets[splitTicket*i:]
					usersInput = s.Users[splitUsers*i:]
				} else {
					ticketsInput = s.Tickets[splitTicket*i : splitTicket*(i+1)]
					usersInput = s.Users[splitUsers*i : splitUsers*(i+1)]
				}
				go func() {
					defer workerGrp.Done()
					searchTicketChan <- tickets.SearchTickets(ticketsInput, "organization_id", strconv.Itoa(result.Organizations[0].Id))
					searchUserChan <- users.SearchUsers(usersInput, "organization_id", strconv.Itoa(result.Organizations[0].Id))
				}()
				result.Tickets = append(result.Tickets, <-searchTicketChan...)
				result.Users = append(result.Users, <-searchUserChan...)
			}
			workerGrp.Wait()
		}
	case SearchGroupTickets:
		var workerGrp sync.WaitGroup
		searchTicketChan := make(chan []tickets.Ticket, workerGrpMax)
		split := (len(s.Tickets) / workerGrpMax)
		for i := 0; i < workerGrpMax; i++ {
			workerGrp.Add(1)
			var ticketList []tickets.Ticket
			if i == workerGrpMax-1 {
				ticketList = s.Tickets[split*i:]
			} else {
				ticketList = s.Tickets[split*i : split*(i+1)]
			}
			go func() {
				defer workerGrp.Done()
				searchTicketChan <- tickets.SearchTickets(ticketList, s.Ident, s.Value)
			}()
			result.Tickets = append(result.Tickets, <-searchTicketChan...)
		}
		workerGrp.Wait()
		if len(result.Tickets) == 1 || (len(result.Tickets) > 0 && s.Ident == "organization_id") {
			// Only link organization details when there is a single ticket returned or the search was on the org id
			var workerGrp sync.WaitGroup
			searchOrgChan := make(chan []organizations.Organization, workerGrpMax)
			split := (len(s.Organizations) / workerGrpMax)
			for i := 0; i < workerGrpMax; i++ {
				workerGrp.Add(1)
				var orgs []organizations.Organization
				if i == workerGrpMax-1 {
					orgs = s.Organizations[split*i:]
				} else {
					orgs = s.Organizations[split*i : split*(i+1)]
				}
				go func() {
					defer workerGrp.Done()
					searchOrgChan <- organizations.SearchOrganizations(orgs, "_id", strconv.Itoa(result.Tickets[0].OrganizationId))
				}()
				result.Organizations = append(result.Organizations, <-searchOrgChan...)
			}
			workerGrp.Wait()
		}
	case SearchGroupUsers:
		var workerGrp sync.WaitGroup
		searchUsersChan := make(chan []users.User, workerGrpMax)
		split := (len(s.Users) / workerGrpMax)
		for i := 0; i < workerGrpMax; i++ {
			workerGrp.Add(1)
			var usersList []users.User
			if i == workerGrpMax-1 {
				usersList = s.Users[split*i:]
			} else {
				usersList = s.Users[split*i : split*(i+1)]
			}
			go func() {
				defer workerGrp.Done()
				searchUsersChan <- users.SearchUsers(usersList, s.Ident, s.Value)
			}()
			result.Users = append(result.Users, <-searchUsersChan...)
		}
		workerGrp.Wait()
		if len(result.Users) == 1 || (len(result.Users) > 0 && s.Ident == "organization_id") {
			// Only link organization details when there is a single user returned or the search was on the org id
			var workerGrp sync.WaitGroup
			searchOrgChan := make(chan []organizations.Organization, workerGrpMax)
			split := (len(s.Organizations) / workerGrpMax)
			for i := 0; i < workerGrpMax; i++ {
				workerGrp.Add(1)
				var orgs []organizations.Organization
				if i == workerGrpMax-1 {
					orgs = s.Organizations[split*i:]
				} else {
					orgs = s.Organizations[split*i : split*(i+1)]
				}
				go func() {
					defer workerGrp.Done()
					searchOrgChan <- organizations.SearchOrganizations(orgs, "_id", strconv.Itoa(result.Users[0].OrganizationId))
				}()
				result.Organizations = append(result.Organizations, <-searchOrgChan...)
			}
			workerGrp.Wait()
		}
	default:
		return SearchResult{}
	}
	return
}

// ValidSearchTerms return if ident is valid for a group
func ValidSearchTerms(group string, ident string) bool {
	switch group {
	case SearchGroupOrganizations:
		return organizations.ValidSearchTerms(ident)
	case SearchGroupTickets:
		return tickets.ValidSearchTerms(ident)
	case SearchGroupUsers:
		return users.ValidSearchTerms(ident)
	default:
		return false
	}
}

// SearchResultDisplay determines the display based on group and search results
func SearchResultDisplay(group string, sr SearchResult) {
	switch group {
	case SearchGroupOrganizations:
		display.DisplayOrganizations(sr.Organizations, sr.Tickets, sr.Users)
	case SearchGroupTickets:
		if len(sr.Organizations) > 0 {
			display.DisplayTickets(sr.Tickets, sr.Organizations[0])
		} else {
			display.DisplayTickets(sr.Tickets, organizations.Organization{})
		}
	case SearchGroupUsers:
		if len(sr.Organizations) > 0 {
			display.DisplayUsers(sr.Users, sr.Organizations[0])
		} else {
			display.DisplayUsers(sr.Users, organizations.Organization{})
		}
	default:
		display.NoResultFound()
	}
}
