package service

import "github.com/darenliang/jikan-go"

func GetById(id int) (*jikan.Anime, error){
	malAnime, err := jikan.GetAnime(id)
	if err != nil {
		return nil, err
	}

	return malAnime, nil
}