package main

import (
	"math"
	"fmt"
)

func main() {

	criticMe, Others := getCritics()

	for _, critic := range Others{
		printEuclideanDistanceFor(criticMe, critic)
		printPearsonCorrelationFor(criticMe, critic)
	}
}

func printPearsonCorrelationFor(c1 Critic, c2 Critic) {
	fmt.Println("Pearson Correlation for", c1.Name,"/",c2.Name)
	fmt.Println("->", pearsonScore(c1.RatedFilms, c2.RatedFilms))
}

func printEuclideanDistanceFor(c1 Critic, c2 Critic) {
	fmt.Println("Euclidian distance for ", c1.Name,"/",c2.Name)
	fmt.Println("->", euclidianDistance(c1.RatedFilms, c2.RatedFilms))
}

func euclidianDistance(ratings1 []RatedFilm, ratings2 []RatedFilm) float64 {

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

func pearsonScore(ratings1 []RatedFilm, ratings2 []RatedFilm) float64 {

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
