package flyweight

import "time"

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}
type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}
type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}
type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TEAM_A = "A"
	TEAM_B = "B"
)

type teamFlyweightFactory struct {
	createdTeams map[string]*Team
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[string]*Team),
	}
}

func (t *teamFlyweightFactory) GetTeam(teamID string) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(team string) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: TEAM_B,
		}
	default:
		return Team{
			ID:   1,
			Name: TEAM_A,
		}
	}
}
