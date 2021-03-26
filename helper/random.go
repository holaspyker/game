package helper

import (
	"fmt"
	"math/rand"
)

func GetRandomSlice(number int, total int) []int {

	q := rand.Perm((total))
	q = q[:number]
	fmt.Println(q)
	return q

}

func ArrayChunk(arr []int, limit int) [][]int {
	batchTotal := make([][]int, 0)
	for i := 0; i < len(arr); i += limit {
		batch := arr[i:min(i+limit, len(arr))]
		batchTotal = append(batchTotal, batch)
	}

	return batchTotal
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
