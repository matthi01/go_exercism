// Package tournament tallies the results of a small football competition.
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type teamScore struct {
	team          string
	matchesPlayed int
	matchesWon    int
	matchesDrawn  int
	matchesLost   int
	points        points
}

type teamGameOutcome struct {
	team    string
	outcome gameOutcome
}

type gameResult []teamGameOutcome

type gameOutcome string

type points int

const (
	winPoints  points = 3
	lossPoints points = 0
	drawPoints points = 1
)

const (
	win  gameOutcome = "win"
	loss gameOutcome = "loss"
	draw gameOutcome = "draw"
)

const fmtString = "%-31s| %2v | %2v | %2v | %2v | %2v\n"

// Tally tallies up the scores of the teams given a string of games
func Tally(reader io.Reader, writer io.Writer) error {
	teamScoreMap := make(map[string]*teamScore)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ";")
		if len(game) == 1 && game[0] == "" {
			continue
		} else if len(game) > 0 && strings.HasPrefix(game[0], "#") {
			continue
		} else if len(game) != 3 {
			return errors.New("invalid game format")
		} else if !validGameOutcome(game[2]) {
			return errors.New("not a valid game outcome")
		} else if game[0] == game[1] {
			return errors.New("team 1 cannot be the same as Team 2")
		} else {
			analyseGame(game, teamScoreMap)
		}
	}

	sortedTeamKeys := getSortedTeamKeys(teamScoreMap)
	scoreBoard := formatScoreBoard(teamScoreMap, sortedTeamKeys)
	io.WriteString(writer, scoreBoard)

	return nil
}

func formatScoreBoard(teamScores map[string]*teamScore, sortedTeamKeys []string) string {
	scoreBoard := scoreboardHeader()
	for _, team := range sortedTeamKeys {
		if record, ok := teamScores[team]; ok {
			scoreBoard += fmt.Sprintf(fmtString, record.team, record.matchesPlayed, record.matchesWon, record.matchesDrawn, record.matchesLost, record.points)
		}
	}
	return scoreBoard
}

func scoreboardHeader() string {
	return fmt.Sprintf(fmtString, "Team", "MP", "W", "D", "L", "P")
}

func analyseGame(game []string, teamScoreMap map[string]*teamScore) {
	var team1, team2 string
	var outcome gameOutcome

	for index, x := range game {
		switch index {
		case 0:
			team1 = x
		case 1:
			team2 = x
		case 2:
			if x == "win" {
				outcome = win
			} else if x == "loss" {
				outcome = loss
			} else if x == "draw" {
				outcome = draw
			} else {
				// panic("Input in wrong format")
			}
		}
	}

	var results gameResult
	if outcome == win {
		results = append(
			results,
			teamGameOutcome{team1, win},
			teamGameOutcome{team2, loss},
		)
	} else if outcome == loss {
		results = append(
			results,
			teamGameOutcome{team2, win},
			teamGameOutcome{team1, loss},
		)
	} else {
		results = append(
			results,
			teamGameOutcome{team1, draw},
			teamGameOutcome{team2, draw},
		)
	}

	recordTeamResults(results, teamScoreMap)
}

func getSortedTeamKeys(teams map[string]*teamScore) []string {
	orderedTeams := make([]string, 0, len(teams))
	for team := range teams {
		orderedTeams = append(orderedTeams, team)
	}
	sort.Slice(orderedTeams, func(i, j int) bool {
		if teams[orderedTeams[i]].points == teams[orderedTeams[j]].points {
			return teams[orderedTeams[i]].team < teams[orderedTeams[j]].team
		}
		return teams[orderedTeams[i]].points > teams[orderedTeams[j]].points
	})

	return orderedTeams
}

func recordTeamResults(results gameResult, teamScoreMap map[string]*teamScore) {
	for _, result := range results {
		team := result.team
		outcome := result.outcome
		if _, ok := teamScoreMap[team]; !ok {
			teamScoreMap[team] = &teamScore{
				team: team,
			}
		}

		teamRecord := teamScoreMap[team]
		switch outcome {
		case win:
			teamRecord.matchesWon++
			teamRecord.points += winPoints
		case loss:
			teamRecord.matchesLost++
			teamRecord.points += lossPoints
		case draw:
			teamRecord.matchesDrawn++
			teamRecord.points += drawPoints
		}
		teamRecord.matchesPlayed++
	}
}

func validGameOutcome(x string) bool {
	for _, s := range []gameOutcome{win, loss, draw} {
		if string(s) == x {
			return true
		}
	}
	return false
}
