package main

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	//recommendItems()

	criticMe, Others := getFilms()

	for _, c  := range Others {
		printPearsonCorrelationFor(criticMe, c)
	}

	fmt.Println("\n----\n")

	recommendItems()
}

type FilmRatedWithWeight struct {
	Name              string
	TotalScore        float64
	SimilaritySum     float64
	TotalParticipants int
}

func printPearsonCorrelationFor(c1 Critic, c2 Critic) {
	fmt.Println("Pearson Correlation for", c1.Name, "/", c2.Name)
	fmt.Println("->", pearsonScore(c1.RatedFilms, c2.RatedFilms))
}

func printEuclideanDistanceFor(c1 Critic, c2 Critic) {
	fmt.Println("Euclidian distance for ", c1.Name, "/", c2.Name)
	fmt.Println("->", euclidianDistance(c1.RatedFilms, c2.RatedFilms))
}

func recommendItems() {
	criticMe, Others := getFilms()
	var weightedList []RatedFilmWeightedList
	for _, critic := range Others {
		weightedList = append(weightedList, buildRatedFilmWeightedList(pearsonScore(criticMe.RatedFilms, critic.RatedFilms), critic))
	}

	//for _, rfl := range weightedList {
	//	fmt.Println("->", rfl.CorrelationScore, rfl.RatedFilmsWeighted, "\n\n")
	//}

	ratedFilmsList := make(map[string]*FilmRatedWithWeight)

	for _, wl := range weightedList {
		for _, f := range wl.RatedFilmsWeighted {
				r, ok := ratedFilmsList[f.Name]
				if ok {
					r.TotalParticipants += 1
					r.SimilaritySum += math.Abs(wl.CorrelationScore)
					if wl.CorrelationScore > 0 {
						r.TotalScore += f.Rating
					} else {
						r.TotalScore -= f.Rating
					}

				} else {
					ratedFilmsList[f.Name] =
						&FilmRatedWithWeight{
							Name:              f.Name,
							TotalParticipants: 1,
							SimilaritySum:     math.Abs(wl.CorrelationScore),
							TotalScore:        f.Rating,
						}
				}
		}
	}

	var ss []FilmRatedWithWeight
	for _, v := range ratedFilmsList {
		ss = append(ss, *v)
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].TotalScore/ss[i].SimilaritySum > ss[j].TotalScore/ss[j].SimilaritySum
	})

	//log.Printf("%v", ratedFilmsList);

	var simScore float64
	for _, rfl := range ss {
		if rfl.SimilaritySum<1 {
			simScore = rfl.TotalScore * rfl.SimilaritySum
		}else{
			simScore = rfl.TotalScore/rfl.SimilaritySum
		}

		fmt.Println("->", rfl.Name, simScore)
	}

}
