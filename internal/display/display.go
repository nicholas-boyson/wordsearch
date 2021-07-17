package display

import (
	"fmt"
	"strings"

	"github.com/nicholas-boyson/wordsearch/internal/organizations"
	"github.com/nicholas-boyson/wordsearch/internal/tickets"
	"github.com/nicholas-boyson/wordsearch/internal/users"
)

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

// InvalidSearchTerm display invalid search term to user
func InvalidSearchTerm() {
	fmt.Println(invalidSearchTerm())
}
func invalidSearchTerm() string {
	return "Invalid search term"
}

// DisplayOrganizations generate organization search result display
func DisplayOrganizations(orgList []organizations.Organization, ticketList []tickets.Ticket, userList []users.User) {
	if len(orgList) > 0 {
		if len(orgList) == 1 {
			fmt.Println(displayOrganizationDetails(orgList[0], ticketList, userList))
		} else {
			fmt.Println(displayOrganizationList(orgList))
		}
	} else {
		NoResultFound()
	}
}
func displayOrganizationList(orgList []organizations.Organization) string {
	result := "Multipe organizations found\n"
	result = result + fmt.Sprintf("%-20s|%-20s|%-100s\n", "Organization Id", "Organization Name", "Organization URL")
	result = result + fmt.Sprintf("%-20s|%-20s|%-100s\n", strings.Repeat("-", 20), strings.Repeat("-", 20), strings.Repeat("-", 100))
	for _, org := range orgList {
		result = result + fmt.Sprintf("%-20d|%-20s|%-100s\n", org.Id, org.Name, org.URL)
	}
	return result
}
func displayOrganizationDetails(org organizations.Organization, ticketList []tickets.Ticket, userList []users.User) string {
	result := fmt.Sprintf("Organization %s (Id %d)\n", org.Name, org.Id)
	result = result + "Details:\n"
	result = result + fmt.Sprintf("%-16s%s\n", "URL:", org.URL)
	result = result + fmt.Sprintf("%-16s%s\n", "External Id:", org.ExternalId)
	result = result + fmt.Sprintf("%-16s%s\n", "Created At:", org.CreatedAt)
	result = result + fmt.Sprintf("%-16s%v\n", "Shared Tickets:", org.SharedTickets)
	result = result + fmt.Sprintf("%-16s%s\n", "Details:", org.Details)
	for i, dm := range org.DomainNames {
		result = result + fmt.Sprintf("Domain Name %d: %s\n", i+1, dm)
	}
	for i, tag := range org.Tags {
		result = result + fmt.Sprintf("Tag %d: %s\n", i+1, tag)
	}
	for i, user := range userList {
		result = result + fmt.Sprintf("User %d: User Id: %-4d | User Name: %-30s | User Phone: %-15s | User Email: %-20s\n", i+1, user.Id, user.Name, user.Phone, user.Email)
	}
	for i, ticket := range ticketList {
		result = result + fmt.Sprintf("Ticket %d: Ticket Id: %s | Ticket Subject %s\n", i+1, ticket.Id, ticket.Subject)
	}
	return result
}

// DisplayTickets generate tickets search result display
func DisplayTickets(ticketList []tickets.Ticket, org organizations.Organization) {
	if len(ticketList) > 0 {
		if len(ticketList) == 1 {
			fmt.Println(displayTicketDetails(ticketList[0], org))
		} else {
			fmt.Println(displayTicketsList(ticketList, org))
		}
	} else {
		NoResultFound()
	}
}
func displayTicketsList(ticketList []tickets.Ticket, org organizations.Organization) string {
	result := "Multipe tickets found\n"
	result = result + fmt.Sprintf("%-50s|%-25s|%-100s\n", "Ticket Id", "Ticket subject", "Ticket Description")
	result = result + fmt.Sprintf("%-50s|%-25s|%-100s\n", strings.Repeat("-", 50), strings.Repeat("-", 25), strings.Repeat("-", 100))
	for _, ticket := range ticketList {
		result = result + fmt.Sprintf("%-50s|%-25s|%-100s\n", ticket.Id, ticket.Subject, ticket.Description)
	}
	result = result + fmt.Sprintf("%-50s|%-25s|%-100s\n", strings.Repeat("-", 50), strings.Repeat("-", 25), strings.Repeat("-", 100))
	if org.Id != 0 {
		result = result + displayOrganizationDetails(org, nil, nil)
	}
	return result
}
func displayTicketDetails(ticket tickets.Ticket, org organizations.Organization) string {
	result := fmt.Sprintf("Ticket %s (Id %s)\n", ticket.Subject, ticket.Id)
	result = result + "Details:\n"
	result = result + fmt.Sprintf("%-16s%s\n", "Description:", ticket.Description)
	result = result + fmt.Sprintf("%-16s%s\n", "URL:", ticket.URL)
	result = result + fmt.Sprintf("%-16s%s\n", "External Id:", ticket.ExternalId)
	result = result + fmt.Sprintf("%-16s%s\n", "Created At:", ticket.CreatedAt)
	result = result + fmt.Sprintf("%-16s%v\n", "Organization Id:", ticket.OrganizationId)
	result = result + fmt.Sprintf("%-16s%s\n", "Via:", ticket.Via)
	result = result + fmt.Sprintf("%-16s%s\n", "Type:", ticket.Type)
	result = result + fmt.Sprintf("%-16s%s\n", "Priority:", ticket.Priority)
	result = result + fmt.Sprintf("%-16s%s\n", "Status:", ticket.Status)
	result = result + fmt.Sprintf("%-16s%d\n", "Submitter Id:", ticket.SubmitterId)
	result = result + fmt.Sprintf("%-16s%d\n", "Assignee Id:", ticket.AssigneeId)
	result = result + fmt.Sprintf("%-16s%v\n", "Has Incidents:", ticket.HasIncidents)
	result = result + fmt.Sprintf("%-16s%s\n", "Due At:", ticket.DueAt)
	for i, tag := range ticket.Tags {
		result = result + fmt.Sprintf("Tag %d: %s\n", i+1, tag)
	}
	if org.Id != 0 {
		result = result + displayOrganizationDetails(org, nil, nil)
	}
	return result
}

// DisplayUsers generate users search result display
func DisplayUsers(userList []users.User, org organizations.Organization) {
	if len(userList) > 0 {
		if len(userList) == 1 {
			fmt.Println(displayUserDetails(userList[0], org))
		} else {
			fmt.Println(displayUsersList(userList, org))
		}
	} else {
		NoResultFound()
	}
}
func displayUsersList(userList []users.User, org organizations.Organization) string {
	result := "Multipe users found\n"
	result = result + fmt.Sprintf("%-20s|%-20s|%-20s\n", "User Id", "User Name", "User Active")
	result = result + fmt.Sprintf("%-20s|%-20s|%-20s\n", strings.Repeat("-", 20), strings.Repeat("-", 20), strings.Repeat("-", 20))
	for _, user := range userList {
		result = result + fmt.Sprintf("%-20d|%-20s|%-20v\n", user.Id, user.Name, user.Active)
	}
	result = result + fmt.Sprintf("%-20s|%-20s|%-20s\n", strings.Repeat("-", 20), strings.Repeat("-", 20), strings.Repeat("-", 20))
	if org.Id != 0 {
		result = result + displayOrganizationDetails(org, nil, nil)
	}
	return result
}
func displayUserDetails(user users.User, org organizations.Organization) string {
	result := fmt.Sprintf("User %s (Alias %s) (Id %d)\n", user.Name, user.Alias, user.Id)
	result = result + "Details:\n"
	result = result + fmt.Sprintf("%-16s%s\n", "URL:", user.URL)
	result = result + fmt.Sprintf("%-16s%s\n", "External Id:", user.ExternalId)
	result = result + fmt.Sprintf("%-16s%s\n", "Email:", user.Email)
	result = result + fmt.Sprintf("%-16s%s\n", "Phone:", user.Phone)
	result = result + fmt.Sprintf("%-16s%s\n", "Signature:", user.Signature)
	result = result + fmt.Sprintf("%-16s%s\n", "Created At:", user.CreatedAt)
	result = result + fmt.Sprintf("%-16s%d\n", "Organization Id:", user.OrganizationId)
	result = result + fmt.Sprintf("%-16s%v\n", "Active:", user.Active)
	result = result + fmt.Sprintf("%-16s%s\n", "Role:", user.Role)
	result = result + fmt.Sprintf("%-16s%v\n", "Verified:", user.Verified)
	result = result + fmt.Sprintf("%-16s%v\n", "Shared:", user.Shared)
	result = result + fmt.Sprintf("%-16s%s\n", "Local:", user.Locale)
	result = result + fmt.Sprintf("%-16s%s\n", "Timezone:", user.Timezone)
	result = result + fmt.Sprintf("%-16s%s\n", "Last Login At:", user.LastLoginAt)
	result = result + fmt.Sprintf("%-16s%v\n", "Suspended:", user.Suspended)

	for i, tag := range user.Tags {
		result = result + fmt.Sprintf("Tag %d: %s\n", i+1, tag)
	}
	if org.Id != 0 {
		result = result + displayOrganizationDetails(org, nil, nil)
	}
	return result
}
