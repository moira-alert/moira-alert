package msteams

import (
	"fmt"
)

/*
Fact models a fact in a MessageCard, contains a timestamp and trigger data
 {
		"name": "10:45",
    "value": "someServer = 0.11 (NODATA to WARN)"
 }
*/
type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
Section models a section in a MessageCard, contains Facts and the Trigger description
 {
		"activityTitle": "Description",
		"activityText": "A trigger description",
    "facts": [
			 {
					"name": "10:45",
					"value": "someServer = 0.11 (NODATA to WARN)"
			 }
		]
 }
*/
type Section struct {
	ActivityTitle string `json:"activityTitle"`
	ActivityText  string `json:"activityText"`
	Facts         []Fact `json:"facts"`
}

/*
OpenURITarget creates a clickable target back to the trigger URI in a MessageCard
 {
		"os": "default",
    "value": "http://moira.tld/trigger/ABCDEF-GH"
 }
*/
type OpenURITarget struct {
	Os  string `json:"os"`
	URI string `json:"uri"`
}

/*
Actions models possible actions in a MessageCard, currently limited to OpenURI actions
 {
		"@type": "OpenUri",
    "name": "Open in Moira"
		"targets": [
			{
				"os": "default",
				"value": "http://moira.tld/trigger/ABCDEF-GH"
 			}
		]
 }
*/
type Actions struct {
	Type    string          `json:"@type"`
	Name    string          `json:"name"`
	Targets []OpenURITarget `json:"targets"`
}

/*
MessageCard models an MSTeams compatible MessageCard
 {
		"@context": "https://schema.org/extensions",
    "@type": "MessageCard",
		"summary": "Moira Alert"
		"title" : "WARN Trigger Name [tag1]"
		"themeColor": "ffa500"
		"sections": [
			 {
					"activityTitle": "Description",
					"activityText": "A trigger description",
					"facts": [
						 {
								"name": "10:45",
								"value": "someServer = 0.11 (NODATA to WARN)"
						 }
					]
			 }
		]
		"potentialAction": [
			{
				"@type": "OpenUri",
				"name": "Open in Moira"
				"targets": [
					{
						"os": "default",
						"value": "http://moira.tld/trigger/ABCDEF-GH"
					}
				]
			}
		]
 }
*/
type MessageCard struct {
	Context         string    `json:"@context"`
	MessageType     string    `json:"@type"`
	Summary         string    `json:"summary"`
	ThemeColor      string    `json:"themeColor"`
	Title           string    `json:"title"`
	Sections        []Section `json:"sections"`
	PotentialAction []Actions `json:"potentialAction,omitempty"`
}

type ErrTeamsError struct {
	error       string
	description string
	err         error
}

func (e *ErrTeamsError) Error() string {
	return fmt.Sprintf("%s : %s", e.description, e.error)
}

func (e *ErrTeamsError) Unwrap() error { return e.err }
