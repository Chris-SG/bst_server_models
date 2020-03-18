package bst_models

import "time"

type DdrSongIds struct {
	Ids []string `json:"ids"`
}

type DdrSongIdWithJacket struct {
	Id string `json:"id"`
	Jacket string `json:"jacket"`
}

type DdrScoresDetailedTopScore struct {
	BestScore  int       `json:"score_record"`
	Lamp       string    `json:"lamp"`
	Rank       string    `json:"rank"`
	PlayCount  int       `json:"playcount"`
	ClearCount int       `json:"clearcount"`
	MaxCombo   int       `json:"maxcombo"`
	LastPlayed time.Time `json:"lastplayed"`

	Mode       string `json:"mode"`
	Difficulty string `json:"difficulty"`
}

type DdrScoresDetailedScore struct {
	Score int `json:"score"`
	ClearStatus bool `json:"cleared"`
	TimePlayed time.Time `json:"timeplayed"`
}

type DdrScoresDetailedDifficulty struct {
	Difficulty string `json:"difficulty"`
	Scores []DdrScoresDetailedScore `json:"scores"`
}

type DdrScoresDetailedMode struct {
	Mode string `json:"mode"`
	Difficulties []DdrScoresDetailedDifficulty `json:"difficulties"`
}

type DdrScoresDetailed struct {
	Id string `json:"id"`
	TopScores []DdrScoresDetailedTopScore `json:"top_scores"`
	Modes []DdrScoresDetailedMode `json:"modes"`
}

type DdrStatisticsTable struct {
	Level int `json:"level"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Mode string `json:"mode"`
	Difficulty string `json:"difficulty"`
	Lamp string `json:"lamp"`
	Rank string `json:"rank"`
	Score int `json:"score"`
	PlayCount int `json:"playcount"`
	ClearCount int `json:"clearcount"`
	MaxCombo int `json:"maxcombo"`
}