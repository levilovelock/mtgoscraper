package parser

import (
	"errors"

	"github.com/levilovelock/mtgoscraper/magic"
)

// Parser A parser of webpages to pull decklists
type Parser struct {
	responseBody []byte
}

// ParseEvent is a function that will parse an entire magic event
// from the MTGO website, given a url
func (p *Parser) ParseEvent(url string) (*magic.Event, error) {
	// Grab html from url - stores response body of event in parser
	// func (p *Parser) fetchEventFromURL() (error)

	// Parse event title from stored responseBody
	// func (p *Parser) parseEventTitle() (string, error)

	// Isolate/extract decklists html from responseBody
	// func (p *Parser) extractDeckLists() ([]magic.Deck, error)

	// Locally compose a magic.Event and return

	return nil, errors.New("ParseEvent has not been implemented yet!")
}
