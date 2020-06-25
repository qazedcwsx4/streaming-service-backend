package service

import "github.com/darenliang/jikan-go"

func CreateAnime() {
	anime, _ := jikan.GetAnime(1)

	println(anime.Title)
}
