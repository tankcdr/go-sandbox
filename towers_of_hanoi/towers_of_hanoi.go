package towersofhanoi

import "fmt"

/**********************************************
 * Tower
 **********************************************/
type Tower struct {
	Disks []int
}

func (t *Tower) Push(disk int) {
	t.Disks = append(t.Disks, disk)
}

func (t *Tower) Pop() int {
	if len(t.Disks) == 0 {
		return -1
	}
	disk := t.Disks[len(t.Disks)-1]
	t.Disks = t.Disks[:len(t.Disks)-1]
	return disk
}

/**********************************************
 * TowersOfHanoi
 **********************************************/
type TowersOfHanoi struct {
	NumDisks int
	Towers   [3]Tower
}

func NewTowersOfHanoi(numDisks, initPost int) *TowersOfHanoi {
	towers := &TowersOfHanoi{}
	towers.Towers[0] = Tower{Disks: []int{}}
	towers.Towers[1] = Tower{Disks: []int{}}
	towers.Towers[2] = Tower{Disks: []int{}}
	towers.NumDisks = numDisks
	for i := numDisks; i > 0; i-- {
		towers.Towers[initPost].Push(i)
	}

	return towers
}

func (t *TowersOfHanoi) MoveDisk(fromPost, toPost int) {
	disk := t.Towers[fromPost].Pop()
	t.Towers[toPost].Push(disk)
}

func (t *TowersOfHanoi) DrawPosts() {

	for i := t.NumDisks - 1; i >= 0; i-- {
		for j := 0; j < 3; j++ {
			if i < len(t.Towers[j].Disks) {
				fmt.Printf("%v ", t.Towers[j].Disks[i])
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Print("\n")
	}

	fmt.Print("______\n")
}

func (t *TowersOfHanoi) MoveDisks(numToMove, fromPost, toPost, tempPost int) {

	if numToMove > 1 {
		t.MoveDisks(numToMove-1, fromPost, tempPost, toPost)
	}

	t.MoveDisk(fromPost, toPost)
	t.DrawPosts()

	if numToMove > 1 {
		t.MoveDisks(numToMove-1, tempPost, toPost, fromPost)
	}
}
