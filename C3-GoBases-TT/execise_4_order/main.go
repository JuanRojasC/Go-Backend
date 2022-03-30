package main

import (
	. "fmt"
	"math/rand"
	"time"
)

// Ejercicio 4 - Ordenamiento
/*
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores
Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo
*/

func insertionSort(arr []int, msg string, c chan int) {
	startTime := printTimeStamp(Sprint("\nStart ", msg))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	finishTime := printTimeStamp(Sprint("Finish ", msg))
	Printf("Duracion: %v\n", finishTime.Sub(startTime))
	c <- 1
}

func bubbleSort(arr []int, msg string, c chan int) {
	startTime := printTimeStamp(Sprint("\nStart ", msg))
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	finishTime := printTimeStamp(Sprint("Finish ", msg))
	Printf("Duracion: %v\n", finishTime.Sub(startTime))
	c <- 1
}

func selectionSort(arr []int, msg string, c chan int) {
	startTime := printTimeStamp(Sprint("\nStart ", msg))
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	finishTime := printTimeStamp(Sprint("Finish ", msg))
	Printf("Duracion: %v\n", finishTime.Sub(startTime))
	c <- 1
}

func printTimeStamp(msg string) time.Time {
	time := time.Now()
	Printf("%s: %s\n", msg, time.Format("15:04:05.000000"))
	return time
}

func sortArray(arr []int, msg string, c chan int, sort func(arr []int, msg string, c chan int)) {
	go sort(arr, msg, c)
	<-c
}

func copySlice(slice []int, originalSlice []int) []int {
	copy(slice, originalSlice)
	return slice
}

func main() {
	arr100 := rand.Perm(100)
	arr1000 := rand.Perm(1000)
	arr10000 := rand.Perm(10000)

	channel := make(chan int)

	Println("\nINSERTION SORT METHOD")
	sortArray(copySlice(make([]int, 100), arr100), "insertion sort 100 elements", channel, insertionSort)
	sortArray(copySlice(make([]int, 1000), arr1000), "insertion sort 1000 elements", channel, insertionSort)
	sortArray(copySlice(make([]int, 10000), arr10000), "insertion sort 10000 elements", channel, insertionSort)

	Println("\nBUBBLE SORT METHOD")
	sortArray(copySlice(make([]int, 100), arr100), "bubble sort 100 elements", channel, bubbleSort)
	sortArray(copySlice(make([]int, 1000), arr1000), "bubble sort 1000 elements", channel, bubbleSort)
	sortArray(copySlice(make([]int, 10000), arr10000), "bubble sort 10000 elements", channel, bubbleSort)

	Println("\nSELECTION SORT METHOD")
	sortArray(copySlice(make([]int, 100), arr100), "selection sort 100 elements", channel, selectionSort)
	sortArray(copySlice(make([]int, 1000), arr1000), "selection sort 1000 elements", channel, selectionSort)
	sortArray(copySlice(make([]int, 10000), arr10000), "selection sort 10000 elements", channel, selectionSort)

}
