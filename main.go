package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make(map[string]bool)

	for _, num := range arrayA {
		merged[num] = true
	}

	for _, num := range arrayB {
		merged[num] = true
	}

	mergedArray := make([]string, 0, len(merged))
	for num := range merged {
		mergedArray = append(mergedArray, num)
	}

	return mergedArray

}

func Mapping(slice []string) map[string]int {
	countMap := make(map[string]int)

	for _, str := range slice {
        countMap[str]++
    }

    return countMap
}

func MunculSekali(angka string) []int {
	countMap := make(map[int]int)

	for _, char := range angka {
        if digit := int(char - '0'); digit >= 0 && digit <= 9 {
            countMap[digit]++
        }
    }

	uniqueNumbers := make([]int, 0)

	for _, char := range angka {
        if digit := int(char - '0'); digit >= 0 && digit <= 9 {
            if countMap[digit] == 1 {
                uniqueNumbers = append(uniqueNumbers, digit)
            }
        }
    }

    return uniqueNumbers

}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))

	fmt.Println(Mapping([]string{"asd","qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{}))



	fmt.Println(MunculSekali("1234123"))
	fmt.Println(MunculSekali("76523752"))
	fmt.Println(MunculSekali("12345"))
	fmt.Println(MunculSekali("1122334455"))
	fmt.Println(MunculSekali("0872504"))
}
