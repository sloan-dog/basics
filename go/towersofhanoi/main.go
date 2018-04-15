package main

import "fmt"
import "./bruteforce"
import "./util"
import "./gentle"

func main() {
	// driver code
	towers := hanoiutil.BuildTowers(5)
	hanoigentle.SolveHanoi(len(towers[0]), 0, 1, 2, &towers)
	fmt.Println(towers)

	ap := hanoibrute.NewAudioPlayer("/Users/sloancoffin/Downloads/just_like_honey_jamc.mp3")
	// go ap.Play()

	towers = hanoiutil.BuildTowers(3)
	r := hanoibrute.SolveHanoiBrutal(len(towers), towers)
	fmt.Println(r.GetTowers())
	defer ap.Stop()
}
