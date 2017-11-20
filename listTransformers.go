package main

func buildRatedFilmWeightedList(distance float64, criticSecondary Critic) RatedFilmWeightedList {

	ratingsSecondary := criticSecondary.RatedFilms
	weightedList := RatedFilmWeightedList{
		CriticSecondary:  criticSecondary,
		CorrelationScore: distance,
	}

	for _, film := range ratingsSecondary {
		weightedList.RatedFilmsWeighted = append(
			weightedList.RatedFilmsWeighted,
			RatedFilm{Name: film.Name, Rating: distance*film.Rating},
		)
	}

	return weightedList
}
