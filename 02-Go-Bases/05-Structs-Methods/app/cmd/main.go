package main

import (
	"go-bases/structs-methods/app/internal"
	"go-bases/structs-methods/app/internal/repository"
	"go-bases/structs-methods/app/internal/service"
)

func main() {
	// dependencies
	var rp internal.MovieRepository
	db := []internal.Movie{
		{ID: 1, MovieAttributes: internal.MovieAttributes{
			Title:  "The Shawshank Redemption",
			Rating: 9.0,
		},
		},
		{ID: 2, MovieAttributes: internal.MovieAttributes{
			Title:  "The Godfather",
			Rating: 7.0,
		},
		},
		{ID: 3, MovieAttributes: internal.MovieAttributes{
			Title:  "The Dark Knight",
			Rating: 5.0,
		},
		},
	}
	rp = repository.NewMoviesSlice(db)
	sv := service.NewMovieDefault(rp)

	// use the service
	avg, err := sv.AverageRating()
	if err != "" {
		println(err)
		return
	}

	println(avg)
}
