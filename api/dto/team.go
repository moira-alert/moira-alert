package dto

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/moira-alert/moira"
)

const (
	teamNameLimit        = 100
	teamDescriptionLimit = 1000
)

// TeamModel is a structure that represents team entity in HTTP transfer
type TeamModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NewTeamModel is a constructor function that creates a new TeamModel using moira.Team
func NewTeamModel(team moira.Team) TeamModel {
	return TeamModel{
		ID:          team.ID,
		Name:        team.Name,
		Description: team.Description,
	}
}

// Bind is a method that implements Binder interface from chi and checks that validity of data in request
func (t TeamModel) Bind(request *http.Request) error {
	if t.Name == "" {
		return fmt.Errorf("team name cannot be empty")
	}
	if utf8.RuneCountInString(t.Name) > teamNameLimit {
		return fmt.Errorf("team name cannot be longer than %d characters", teamNameLimit)
	}
	if utf8.RuneCountInString(t.Description) > teamDescriptionLimit {
		return fmt.Errorf("team description cannot be longer than %d characters", teamNameLimit)
	}
	return nil
}

// Render is a function that implements chi Renderer interface for TeamModel
func (TeamModel) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ToMoiraTeam is a method that converts dto.Team to general moira.Team datatype
func (t TeamModel) ToMoiraTeam() moira.Team {
	return moira.Team{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}
}

// SaveTeamResponse is a structure to return team creation result in HTTP response
type SaveTeamResponse struct {
	ID string `json:"id"`
}

// Render is a function that implements chi Renderer interface for SaveTeamResponse
func (SaveTeamResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// UserTeams is a structure that represents a set of teams of user
type UserTeams struct {
	Teams []TeamModel `json:"teams"`
}

// Render is a function that implements chi Renderer interface for UserTeams
func (UserTeams) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// TeamMembers is a structure that represents a team members in HTTP transfer
type TeamMembers struct {
	Usernames []string `json:"usernames"`
}

// Bind is a method that implements Binder interface from chi and checks that validity of data in request
func (m TeamMembers) Bind(request *http.Request) error {
	if len(m.Usernames) == 0 {
		return fmt.Errorf("at least one user should be specified")
	}
	return nil
}

// Render is a function that implements chi Renderer interface for TeamMembers
func (TeamMembers) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
