package main

type Critic struct {
	Name       string
	RatedFilms []RatedFilm
}

type RatedFilmWeightedList struct{
	CriticSecondary Critic
	CorrelationScore float64
	RatedFilmsWeighted []RatedFilm
}

type RatedFilmScore struct{
	Critic Critic
	Rating int8
}

type RatedFilm struct {
	Name   string
	Rating float64
}

func getCritics() (me Critic, list []Critic) {

	criticMe := Critic{Name: "Myself"}
	criticLyels := Critic{Name: "Lyels"}
	criticHagalazia := Critic{Name: "Hagalazia"}
	criticChiz := Critic{Name: "Chiz50"}

	criticMe.RatedFilms =
		[]RatedFilm{
			{"World War Z", 7},
			{"Saving Private Ryan", 10},
			{"Rogue One", 7},
			{"Arrival", 10},
			{"Doctor Strange", 7},
			{"No Country for Old Men", 10},
			{"La La Land", 5},
			{"Fantastic Beasts and Where to Find Them", 7},
			{"Logan", 9},
			{"Sucker Punch", 5},
			{"Бегущий по лезвию 2049", 9},
			{"The LEGO Batman Movie", 8},
		}

	criticLyels.RatedFilms =
		[]RatedFilm{
			{"World War Z", 7},
			{"Saving Private Ryan", 9},
			{"Rogue One", 7},
			{"Arrival", 9},
			{"Doctor Strange", 8},
			{"No Country for Old Men", 9},
			{"La La Land", 8},
			{"Fantastic Beasts and Where to Find Them", 7},
			{"Logan", 8},
			{"Sucker Punch", 4},
			{"Blade Runner", 4},
			{"Rush", 3},
			{"Casino Royale", 3},
			{"Thor: Ragnarek", 4},
		}

	criticHagalazia.RatedFilms =
		[]RatedFilm{
			{"World War Z", 8},
			{"Saving Private Ryan", 9},
			{"Rogue One", 9},
			{"Arrival", 8},
			{"Doctor Strange", 5},
			{"No Country for Old Men", 4},
			{"La La Land", 4},
			{"Fantastic Beasts and Where to Find Them", 8},
			{"Logan", 7},
			{"Sucker Punch", 5},
			{"Бегущий по лезвию 2049", 6},
			{"The LEGO Batman Movie", 10},
			{"Blade Runner", 7},
			{"Rush", 4},
			{"Casino Royale", 7},
			{"Thor: Ragnarek", 9},
		}

	criticChiz.RatedFilms =
		[]RatedFilm{
			{"World War Z", 5},
			{"Saving Private Ryan", 7},
			{"Rogue One", 7},
			{"Arrival", 8},
			{"Doctor Strange", 9},
			{"No Country for Old Men", 9},
			{"La La Land", 9},
			{"Fantastic Beasts and Where to Find Them", 6},
			{"Logan", 7},
			{"Sucker Punch", 6},
			{"Бегущий по лезвию 2049", 9},
			{"The LEGO Batman Movie", 7},
			{"Blade Runner", 5},
			{"Rush", 8},
			{"Casino Royale", 7},
			{"Thor: Ragnarek", 8},
		}

	list = []Critic{}

	list = append(list, criticLyels, criticHagalazia, criticChiz)


	return criticMe, list
}
