package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var lowerText string = strings.ToLower(text)
	lowerText = strings.Join(strings.Fields(lowerText), "")

	var symbolNumber int
	var symbolPovtor = make([] int, 0, len(lowerText))

	for _, s := range lowerText {
		if unicode.IsLetter(s) {
			symbolNumber++
		}
	}

	keySlice := strings.Split(lowerText, "")

	duplicateFrequency := make(map[string]int)
	for _, item := range keySlice {
        _, exists := duplicateFrequency[item]
        if exists {
            duplicateFrequency[item] += 1
        } else {
            duplicateFrequency[item] = 1
        }
	symbolPovtor = append(symbolPovtor, duplicateFrequency[item])
    }

	var procent = make([]float64, 0, len(lowerText))
	for i :=range (symbolPovtor) {
		procent = append(procent, float64(symbolPovtor[i]) / float64(symbolNumber))
	}

	myMap1 := make(map[string]string)
	for i, key := range keySlice {
		myMap1[key] = "кол-во повторов - " + strconv.Itoa(symbolPovtor[i])+ " частота в % - " + strconv.FormatFloat(procent[i], 'f', 2, 64)
	}
	
	for k, v := range myMap1 {
		fmt.Println(k, v)
	}
}