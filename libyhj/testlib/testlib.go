package testlib

//Music struct
//inform of music
type Music struct {
	title  string
	artist string
	time   int
}

var pop map[string]string

var mu0, mu1, mu2 Music

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin'"

	mu0.artist = "Maroon5"
	mu0.title = "Animals"
	mu0.time = 231

	mu1 = Music{"bad guy", "Billie Eilish", 195}

	mu2 = Music{"Thunder", "Imagine Dragons", 187}
}

//GetSong func
//return song of singer
func GetSong(singer string) string {

	var ret string
	song := pop[singer]

	if len(song) == 0 {
		ret = "no data"
	} else {
		ret = pop[singer]
	}
	return ret
}

//GetMusic func
//return music property (title, artist, play time)
func GetMusic(index int) Music {

	var ret Music

	switch index {
	case 0:
		ret = mu0
	case 1:
		ret = mu1
	case 2:
		ret = mu2
	}
	return ret
}
