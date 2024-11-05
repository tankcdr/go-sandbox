package towersofhanoi_test

import (
	"testing"
	"towersofhanoi"

	"gotest.tools/assert"
)

func TestTower_StackOperations(t *testing.T) {
	tower := &towersofhanoi.Tower{}
	tower.Push(3)
	tower.Push(2)
	tower.Push(1)
	assert.Equal(t, 1, tower.Pop())
	assert.Equal(t, 2, tower.Pop())
	assert.Equal(t, 3, tower.Pop())
	assert.Equal(t, -1, tower.Pop())
}

func TestTowersOfHanoi_MoveDisks(t *testing.T) {
	towers := towersofhanoi.NewTowersOfHanoi(3, 0)

	towers.MoveDisk(0, 2)
	assert.DeepEqual(t, []int{3, 2}, towers.Towers[0].Disks)
	assert.DeepEqual(t, []int{1}, towers.Towers[2].Disks)
}
