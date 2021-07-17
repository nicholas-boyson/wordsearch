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
	// load the data on start up and hold in memory
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

func process(scanner *bufio.Scanner) error {
	// Display welcome message
	display.Welcome()
	quit := false
	for !quit {
		// While the user has not quit repeat the search
		display.SelectSearchOptions()
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("reading input: %s", err)
		}

		switch scanner.Text() {
		case "1":
			// fresh search
			var searchRequest = search.Search{}
			searchRequest.Organizations = orgList
			searchRequest.Tickets = ticketList
			searchRequest.Users = userList

			knownGroup := false
			for !knownGroup {
				// prompt user for group to search on
				display.SelectGroupOptions()
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					return fmt.Errorf("reading input: %s", err)
				}
				knownGroup = true
				// repeat if group is unknown or input is quit
				switch scanner.Text() {
				case "1":
					searchRequest.Group = "Users"
				case "2":
					searchRequest.Group = "Tickets"
				case "3":
					searchRequest.Group = "Organizations"
				case "quit":
					// quit the search
					quit = true
				default:
					knownGroup = false
				}
			}

			if !quit {
				// request user to provide search term
				display.EnterSearchTerm()
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					return fmt.Errorf("reading input: %s", err)
				}

				if scanner.Text() != "quit" {
					// if the input is not quit validate search term based on search group
					if search.ValidSearchTerms(searchRequest.Group, scanner.Text()) {
						searchRequest.Ident = scanner.Text()
						// prompt user to search value blank is allowed
						display.EnterSearchValue()
						scanner.Scan()
						if err := scanner.Err(); err != nil {
							return fmt.Errorf("reading input: %s", err)
						}

						if scanner.Text() != "quit" {
							// if the input is not quit then perform search
							searchRequest.Value = scanner.Text()
							searchResult := search.SearchData(searchRequest)
							search.SearchResultDisplay(searchRequest.Group, searchResult)
						} else {
							quit = true
						}
					} else {
						// inform user of the invalid term, take the user back to the start of the search
						display.InvalidSearchTerm()
					}
				} else {
					quit = true
				}
			}
		case "2":
			// display a list of searchable field to the user
			display.ListSearchableFields()
		case "quit":
			// exit search option
			quit = true
		default:
			// unknown entry repeat the search options
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading input: %s", err)
	}
	return nil
}

// main function start the scanner
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	err := process(scanner)
	if err != nil {
		fmt.Printf("Hit an input error: %s", err.Error())
		os.Exit(1)
	}
}
