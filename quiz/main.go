package main


import "fmt"

func main() {

	// scores := [8]int{100, 80, 75, 90, 70, 93, 88,67}
    // //var result int
	// var med float64
	// panjangArray := len(scores)
	// result := 0
	// for _, score := range scores {
	// 	result = score + result
	// }
	// med = float64(result / panjangArray)
	// fmt.Println(med)

	scores := [8]int{100, 80, 75, 90, 70, 93, 88,67}
    // pakai slice
	var goodScores []int

	for _,  score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores)





    

}