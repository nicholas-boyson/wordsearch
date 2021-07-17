package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nicholas-boyson/wordsearch/internal/display"
	"github.com/nicholas-boyson/wordsearch/internal/organizations"
	"github.com/nicholas-boyson/wordsearch/internal/search"
	"github.com/nicholas-boyson/wordsearch/internal/tickets"
	"github.com/nicholas-boyson/wordsearch/internal/users"
)

var orgList []organizations.Organization
var ticketList []tickets.Ticket
var userList []users.User

func init() {
	var err error
	orgList, err = organizations.LoadOrganizations("")
	if err != nil {
		fmt.Printf("Failed to load organizations: %s", err.Error())
		os.Exit(1)
	}
	ticketList, err = tickets.LoadTickets("")
	if err != nil {
		fmt.Printf("Failed to load tickets: %s", err.Error())
		os.Exit(1)
	}
	userList, err = users.LoadUsers("")
	if err != nil {
		fmt.Printf("Failed to load users: %s", err.Error())
		os.Exit(1)
	}
}

func process(scanner *bufio.Scanner) {
	display.Welcome()
	quit := false
	for !quit {
		display.SelectSearchOptions()
		scanner.Scan()
		switch scanner.Text() {
		case "1":
			// New search
			var searchRequest = search.Search{}
			searchRequest.Organizations = orgList
			searchRequest.Tickets = ticketList
			searchRequest.Users = userList

			knownGroup := false
			for !knownGroup {
				display.SelectGroupOptions()
				scanner.Scan()
				knownGroup = true
				switch scanner.Text() {
				case "1":
					searchRequest.Group = "Users"
				case "2":
					searchRequest.Group = "Tickets"
				case "3":
					searchRequest.Group = "Organizations"
				case "quit":
					quit = true
					knownGroup = false
				default:
					knownGroup = false
				}
			}

			if !quit {
				display.EnterSearchTerm()
				scanner.Scan()
				if scanner.Text() != "quit" {
					if search.ValidSearchTerms(searchRequest.Group, scanner.Text()) {
						searchRequest.Ident = scanner.Text()
						display.EnterSearchValue()
						scanner.Scan()
						if scanner.Text() != "quit" {
							searchRequest.Value = scanner.Text()
							searchResult := search.SearchData(searchRequest)
							search.SearchResultDisplay(searchRequest.Group, searchResult)
						} else {
							quit = true
						}
					} else {
						display.InvalidSearchTerm()
					}
				} else {
					quit = true
				}
			}
		case "2":
			display.ListSearchableFields()
		case "quit":
			quit = true
		default:
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	process(scanner)
}
