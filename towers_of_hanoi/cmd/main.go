package main

import (
	"towersofhanoi"
)

const NUM_DISKS = 6

func main() {
	// Make three posts.
	towers := towersofhanoi.NewTowersOfHanoi(NUM_DISKS, 0)
	towers.DrawPosts()

	towers.MoveDisks(NUM_DISKS, 0, 1, 2)

}
