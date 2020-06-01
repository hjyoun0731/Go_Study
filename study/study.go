package main

import (
	"fmt"
	"libyhj/testlib"
)

func main() {

	//fmt.Println("hello go")

	getsong()
	getmusic()
}

func getsong() {

	song := testlib.GetSong("Adele")
	fmt.Println(song)
}

func getmusic() {

	var muInfo testlib.Music

	muInfo = testlib.GetMusic(0)
	fmt.Println(muInfo)
	muInfo = testlib.GetMusic(1)
	fmt.Println(muInfo)
	muInfo = testlib.GetMusic(2)
	fmt.Println(muInfo)
}
