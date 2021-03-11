package main

import (
	"fmt"
)

func main() {
	m := GetMatrix(10)
	for _, val := range m {
		fmt.Println(val)
	}
}

//Получение матрицы
func getMatrixNil(n int) [][]int {
	mainMatrix := make([][]int, n)
	for i := 0; i < len(mainMatrix); i++ {
		mainMatrix[i] = make([]int, n)
	}
	return mainMatrix
}

//Заполнение матрицы значениями
func filMatrix(m [][]int, counter int, lineTop int, lineRight int, lineBottom int, lineLeft int, stopFirstLine int) {
	sideCount := len(m[0])
	if counter > 1 {
		stopFirstLine = stopFirstLine - 1
	}
	for i := lineTop; i < stopFirstLine; i++ {
		m[lineTop][i] = counter
		counter++
	}
	for i := lineTop + 1; i <= lineRight; i++ {
		m[i][lineRight] = counter
		counter++
	}
	for i := lineBottom - 1; i >= lineTop; i-- {
		m[lineBottom][i] = counter
		counter++
	}
	for i := lineBottom - 1; i >= lineTop+1; i-- {
		m[i][lineLeft] = counter
		counter++
	}
	if sideCount*sideCount-counter >= 3 {
		lineTop = lineTop + 1
		lineRight = lineRight - 1
		lineBottom = lineBottom - 1
		lineLeft = lineLeft + 1
		filMatrix(m, counter, lineTop, lineRight, lineBottom, lineLeft, stopFirstLine)
	}

}

//Функция интерфейс, которая возвращает итоговую матрицу
func GetMatrix(n int) [][]int {
	m := getMatrixNil(n)
	filMatrix(m, 1, 0, len(m)-1, len(m[0])-1, 0, len(m[0]))
	return m
}
