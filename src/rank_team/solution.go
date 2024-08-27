package rank_team

import (
	"sort"
	"strings"
)

// Description: https://leetcode.com/problems/rank-teams-by-votes/description

func rankTeams(votes []string) string {
	if len(votes) == 1 {
		return votes[0]
	}
	type Team struct {
		team  rune
		ranks []int
	}
	mp := make(map[rune]*Team)
	length := len(votes[0])

	for _, vote := range votes {
		for i, team := range vote {
			v, ok := mp[team]
			if ok {
				v.ranks[i] += 1
			} else {
				mp[team] = &Team{
					team:  team,
					ranks: make([]int, length),
				}
				mp[team].ranks[i] += 1
			}
		}
	}

	teams := make([]*Team, 0)
	for _, team := range mp {
		teams = append(teams, team)
	}

	sort.Slice(teams, func(i, j int) bool {
		for k := 0; k < len(teams[i].ranks) && k < len(teams[j].ranks); k++ {
			if teams[i].ranks[k] != teams[j].ranks[k] {
				return teams[i].ranks[k] > teams[j].ranks[k]
			}
		}
		return teams[i].team < teams[j].team
	})

	var res strings.Builder
	for _, team := range teams {
		res.WriteRune(team.team)
	}
	return res.String()
}

func rankTeams2(votes []string) string {
	if len(votes) == 1 {
		return votes[0]
	}
	var count [26][26]int
	for _, v := range votes {
		for i, ch := range v {
			count[ch-'A'][i]++
		}
	}
	cand := []byte(votes[0])
	sort.Slice(cand, func(i, j int) bool {
		a := count[cand[i]-'A']
		b := count[cand[j]-'A']
		for rank := 0; rank < 26; rank++ {
			if a[rank] == b[rank] {
				continue
			}
			return a[rank] > b[rank]
		}
		return cand[i] < cand[j]
	})
	return string(cand)
}
