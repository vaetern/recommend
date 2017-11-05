package main

import (
	"math"
	"fmt"
)

type critic struct {
	Name       string
	RatedFilms []ratedFilm
}

type ratedFilm struct {
	Name   string
	Rating int8
}

func main() {

	lyels := critic{Name: "lyels"}
	myself := critic{Name: "me"}

	lyelsRates :=
		[]ratedFilm{
			{"8.794 House, M.D.", 10},
			{"8.492 Shutter Island", 9},
			{"7.621 Kingsman: The Secret Service", 8},
			{"7.322 A Monster Calls", 8},
			{"7.847 Sin City", 8},
			{"7.484 Gravity", 8},
			{"7.476 Zombieland", 8},
			{"7.684 Prisoners", 8},
			{"7.795 RocknRolla", 8},
			{"7.740 Guardians of the Galaxy", 8},
			{"8.502 The Dark Knight", 8},
			{"8.139 The Dark Knight Rises", 8},
			{"7.834 Training Day", 8},
			{"8.284 Braveheart", 8},
			{"7.416 Finding Dory", 7},
			{"6.972 World War Z", 7},
			{"7.279 In Time", 7},
			{"6.607 The Golden Compass", 7},
			{"7.041 Furious Seven", 7},
			{"3,074", 1},
			{"8.656 Fight Club", 9},
			{"9.064 The Green Mile", 9},
			{"8.684 Léon", 9},
			{"8.491 The Matrix", 9},
			{"8.665 Inception", 9},
			{"8.527 The Prestige", 9},
			{"8.921 Forrest Gump", 9},
			{"9.114 The Shawshank Redemption", 9},
			{"7.857 Inglourious Basterds", 8},
			{"7.942 Edge of Tomorrow", 8},
			{"8.174 Django Unchained", 8},
			{"7.622 Real Steel", 8},
			{"8.039 The Curious Case of Benjamin Button", 8},
			{"7.739 Casino Royale", 8},
			{"7.425 Logan", 8},
			{"7.887 Jagten", 8},
			{"7.667 Phone Booth", 8},
			{"7.360 Fury", 8},
			{"7.432 Doctor Strange", 8},
			{"7.908 What Dreams May Come", 8},
			{"7.606 Warcraft", 7},
			{"7.021 Deepwater Horizon", 7},
			{"7.062 Public Enemies", 7},
			{"6.905 Surrogates", 7},
			{"6.939 Hugo", 7},
			{"6.948 The Equalizer", 7},
			{"7.318 Thor: The Dark World", 7},
			{"6.408 Grimsby", 6},
			{"5.925 Jack Reacher: Never Go Back", 6},
			{"6.343 Inferno", 6},
		}

	lyels.RatedFilms = lyelsRates

	myRatings :=
		[]ratedFilm{
			{"8.794 House, M.D.", 10},
			{"8.492 Shutter Island", 9},
			{"7.621 Kingsman: The Secret Service", 8},
			{"7.322 A Monster Calls", 8},
			{"7.847 Sin City", 8},
			{"7.484 Gravity", 8},
			{"7.476 Zombieland", 8},
			{"7.684 Prisoners", 8},
			{"7.795 RocknRolla", 8},
			{"7.740 Guardians of the Galaxy", 8},
			{"8.502 The Dark Knight", 8},
			{"8.139 The Dark Knight Rises", 8},
			{"7.834 Training Day", 8},
			{"8.284 Braveheart", 8},
			{"7.416 Finding Dory", 7},
			{"6.972 World War Z", 7},
			{"7.279 In Time", 7},
			{"6.607 The Golden Compass", 7},
			{"7.041 Furious Seven", 7},
			{"3,074", 1},
			{"8.656 Fight Club", 10},
			{"9.064 The Green Mile", 10},
			{"8.684 Léon", 10},
			{"8.491 The Matrix", 10},
			{"8.665 Inception", 10},
			{"8.527 The Prestige", 10},
			{"8.921 Forrest Gump", 10},
			{"9.114 The Shawshank Redemption", 8},
			{"7.857 Inglourious Basterds", 9},
			{"7.942 Edge of Tomorrow", 9},
			{"8.174 Django Unchained", 9},
			{"7.622 Real Steel", 9},
			{"8.039 The Curious Case of Benjamin Button", 9},
			{"7.739 Casino Royale", 9},
			{"7.425 Logan", 9},
			{"7.887 Jagten", 9},
			{"7.667 Phone Booth", 9},
			{"7.360 Fury", 9},
			{"7.432 Doctor Strange", 7},
			{"7.908 What Dreams May Come", 7},
			{"7.606 Warcraft", 8},
			{"7.021 Deepwater Horizon", 8},
			{"7.062 Public Enemies", 8},
			{"6.905 Surrogates", 8},
			{"6.939 Hugo", 8},
			{"6.948 The Equalizer", 6},
			{"7.318 Thor: The Dark World", 6},
			{"6.408 Grimsby", 7},
			{"5.925 Jack Reacher: Never Go Back", 5},
			{"6.343 Inferno", 5},
		}

	myself.RatedFilms = myRatings

	fmt.Println("Euclidian distance is ", euclidianDistance(lyels.RatedFilms, myself.RatedFilms))
	fmt.Println("Pearson Correlation score is ", pearsonScore(lyels.RatedFilms, myself.RatedFilms))

}

func euclidianDistance(ratings1 []ratedFilm, ratings2 []ratedFilm) float64 {

	ratingDistanceRaw := .0

	for _, x := range ratings1 {
		for _, y := range ratings2 {
			if x.Name == y.Name {
				ratingDistanceRaw = ratingDistanceRaw + math.Pow(float64(x.Rating-y.Rating), 2)
			}
		}
	}

	euclidianDistance := 1 / (1 + math.Sqrt(ratingDistanceRaw))

	return euclidianDistance
}

func pearsonScore(ratings1 []ratedFilm, ratings2 []ratedFilm) float64 {

	commonRatingCount := 0

	sum1, sum2 := .0, .0
	sum1Sq, sum2Sq := .0, .0
	totalSum := .0

	for _, x := range ratings1 {
		for _, y := range ratings2 {
			if x.Name == y.Name {

				commonRatingCount++

				sum1 += float64(x.Rating)
				sum2 += float64(y.Rating)
				sum1Sq += math.Pow(float64(x.Rating), 2)
				sum2Sq += math.Pow(float64(y.Rating), 2)
				totalSum += float64(x.Rating * y.Rating)

			}
		}
	}

	if commonRatingCount == 0 {
		return .0
	}

	n := float64(commonRatingCount)

	num := totalSum - (sum1 * sum2 / n)
	den := math.Sqrt((sum1Sq - math.Pow(sum1, 2)/n) * (sum2Sq - math.Pow(sum2, 2)/n))

	if den == 0 {
		return .0
	}

	r := num / den

	return r
}
