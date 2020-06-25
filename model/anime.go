package model

import "github.com/darenliang/jikan-go"

type Anime struct {
	Series
	AnimeId    string  `db:"anime.id" json:"anime_id"`
	MalId      int     `db:"mal_reference" json:"mal_id"`
	Score      float64 `db:"score" json:"score"`
	Rank       int     `db:"rank" json:"rank"`
	Popularity int     `db:"popularity" json:"popularity"`
}

func FromMalAnime(malAnime *jikan.Anime) Anime {
	return Anime{
		Series: Series{
			Entry: Entry{
				Name:        malAnime.Title,
				ImageUrl:    malAnime.ImageURL,
				Description: malAnime.Synopsis,
			},
		},
		MalId:      malAnime.MalID,
		Score:      malAnime.Score,
		Rank:       malAnime.Rank,
		Popularity: malAnime.Popularity,
	}
}
