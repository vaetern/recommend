package main

func getFilms() (me Critic, list []Critic) {

	filmWWZ := Critic{Name: "World War Z"}
	filmPrivateRyan := Critic{Name: "Saving Private Ryan"}
	filmRogueOne := Critic{Name: "Rogue One"}
	filmArrival := Critic{Name: "Arrival"}
	filmLaLa := Critic{Name: "La La Land"}
	filmDoctor := Critic{Name: "Doctor Strange"}
	filmBladeRunner := Critic{Name: "Blade Runner 2049"}
	filmRoyale := Critic{Name: "Casino Royale"}

	filmWWZ.RatedFilms =
		[]RatedFilm{
			{"Abel", 8},
			{"Baker", 6},
			{"Charlie", 6},
			{"Delta", 9},
			{"Echo", 8},
			{"Foxtrot", 6},
		}

	filmPrivateRyan.RatedFilms =
		[]RatedFilm{
			{"Abel", 10},
			{"Baker", 8},
			{"Charlie", 8},
			{"Delta", 2},
			{"Echo", 10},
			{"Foxtrot", 8},
		}

	filmRogueOne.RatedFilms =
		[]RatedFilm{
			{"Abel", 8},
			{"Baker", 6},
			{"Charlie", 6},
			{"Delta", 1},
			{"Echo", 8},
			{"Foxtrot", 6},
		}

	filmArrival.RatedFilms =
		[]RatedFilm{
			{"Abel", 10},
			{"Baker", 8},
			{"Charlie", 6},
			{"Delta", 3},
			{"Echo", 9},
			{"Foxtrot", 6},
		}

	filmLaLa.RatedFilms =
		[]RatedFilm{
			{"Baker", 7},
			{"Echo", 6},
			{"Foxtrot", 8},
		}

	filmDoctor.RatedFilms =
		[]RatedFilm{
			{"Baker", 6},
			{"Echo", 7},
			{"Foxtrot", 6},
		}

	filmBladeRunner.RatedFilms =
		[]RatedFilm{
			{"Echo", 9},
			{"Foxtrot", 8},
		}

	filmRoyale.RatedFilms =
		[]RatedFilm{
			{"Baker", 8},
			{"Echo", 7},
		}




	list = []Critic{}

	//list = append(list, filmPrivateRyan, filmArrival, filmWWZ, filmLaLa, filmRogueOne, filmBladeRunner, filmRoyale)

	//return filmDoctor, list

	list = append(list, filmPrivateRyan, filmArrival, filmWWZ, filmLaLa, filmRogueOne, filmDoctor, filmRoyale)

	return filmBladeRunner, list
}
