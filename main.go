package main

import (
	"fmt"
	"sort"
)

func main() {
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
	criticMe, Others := getCritics()
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
				r.TotalScore += f.Rating
				r.SimilaritySum += wl.CorrelationScore
			} else {
				ratedFilmsList[f.Name] =
					&FilmRatedWithWeight{
						Name:              f.Name,
						TotalParticipants: 1,
						SimilaritySum:     wl.CorrelationScore,
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

	for _, rfl := range ss {
		fmt.Println("->", rfl.Name, rfl.TotalScore/rfl.SimilaritySum)
	}

}
