package phzData

import "testing"

//func TestTryHu2(t *testing.T) {
//	gameData := &UserGameData{}
//	gameData.RoundHuXi = 15
//	//indexs := []int32{69, 70, 71, 10, 44, 45, 46, 65, 66, 9, 1, 9, 17}
//	//indexs := []int32{69, 70, 71, 10, 44, 45, 46, 65, 66, 67, 9, 1, 2, 3, 9, 10, 11, 17, 18, 19}
//
//	//indexs := []int32{5, 77, 57, 1, 61, 9, 10, 49, 50, 17, 13, 3, 13, 8, 74, 75, 57}
//	//indexs := []int32{13, 29, 2, 44, 49, 69, 41, 53, 45, 21, 77, 61, 42, 9, 17, 49}
//	//indexs := []int32{49, 2, 18, 17, 77, 5, 3, 74, 11, 79}
//	indexs := []int32{42, 63, 22, 29, 62, 28, 8, 31, 35, 32, 14}
//	pais := []*PHZPoker{}
//	for _, i := range indexs {
//		pai := InitPaiByIndex(i)
//		if pai != nil {
//			pais = append(pais, pai)
//		}
//	}
//	checkPai := InitPaiByIndex(int32(61))
//	gameData.HandPokers = pais
//	TryHu2(gameData, checkPai, true)
//}

func TestTryHu2(t *testing.T) {
	gameData := &UserGameData{}
	gameData.RoundHuXi = 0

	indexs := []int32{76, 15, 34, 70, 35, 9, 46, 33, 24, 20, 77, 39, 13, 73, 78, 27, 1, 57, 26, 51, 28}
	pais := []*PHZPoker{}
	for _, i := range indexs {
		pai := InitPaiByIndex(i)
		if pai != nil {
			pais = append(pais, pai)
		}
	}
	//checkPai := InitPaiByIndex(int32(28))
	gameData.HandPokers = pais
	//TryHu2(gameData, nil, true, false)
}
