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

	abel := Critic{Name: "Abel"}
	baker := Critic{Name: "Baker"}
	charlie := Critic{Name: "Charlie"}
	delta := Critic{Name: "Delta"}
	echo := Critic{Name: "Echo"}
	foxtrot := Critic{Name: "Foxtrot"}

	abel.RatedFilms =
		[]RatedFilm{
			{"World War Z", 8},
			{"Saving Private Ryan", 10},
			{"Rogue One", 8},
			{"Arrival", 10},
		}

	baker.RatedFilms =
		[]RatedFilm{
			{"World War Z", 6},
			{"Saving Private Ryan", 8},
			{"Rogue One", 7},
			{"Arrival", 8},
			{"La La Land", 7},
			{"Doctor Strange", 6},
			{"Casino Royale", 8},
		}

	charlie.RatedFilms =
		[]RatedFilm{
			{"World War Z", 6},
			{"Saving Private Ryan", 8},
			{"Rogue One", 6},
			{"Arrival", 6},
		}

	delta.RatedFilms =
		[]RatedFilm{
			{"World War Z", 9},
			{"Saving Private Ryan", 2},
			{"Rogue One", 1},
			{"Arrival", 3},
		}

	echo.RatedFilms =
		[]RatedFilm{
			{"World War Z", 8},
			{"Saving Private Ryan", 10},
			{"Rogue One", 8},
			{"Arrival", 9},
			{"La La Land", 6},
			{"Blade Runner 2049", 9},
			{"Doctor Strange", 7},
			{"Casino Royale", 7},
		}

	foxtrot.RatedFilms =
		[]RatedFilm{
			{"World War Z", 6},
			{"Saving Private Ryan", 8},
			{"Rogue One", 6},
			{"Arrival", 6},
			{"La La Land", 8},
			{"Blade Runner 2049", 8},
			{"Doctor Strange", 7},
		}

	list = []Critic{}

	list = append(list, baker, charlie, delta, echo, foxtrot)

	return abel, list
}
