package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"ocg.com/lec2/calculator"
	"ocg.com/lec2/sort"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fileName string) []int {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	file := string(content)
	var arrString []string
	arrString = strings.Split(file, " ")

	var arrInt []int
	for _, value := range arrString {
		val, err := strconv.Atoi(value)
		if err == nil {
			arrInt = append(arrInt, val)
		}
	}
	return arrInt
}


func Contain(num int, arr []int) bool {
	for _, item := range arr {
		if item == num {
			return true
		}
	}
	return false
}

func main() {
	arr := ReadFile("text.txt")
	n := len(arr)
	x := 10
	fmt.Print("Array: ")
	fmt.Println(arr)
	fmt.Print("Min: ")
	fmt.Println(calculator.FindMin(arr))
	fmt.Print("Max: ")
	fmt.Println(calculator.FindMax(arr))
	fmt.Print("Average: ")
	fmt.Println(calculator.Average(arr))
	fmt.Println( sort.QuickSort(arr, 0, n-1))
	//fmt.Print("Bubble sort: ")
	//fmt.Println(sort.BubbleSort(arr, n))
	fmt.Println("Check prime number: ")
	calculator.CheckArrPrime(arr)
	if Contain(x, arr) {
		fmt.Printf("%d already in array\n", x)
	} else {
		fmt.Printf("Array not contain %d\n", x)
	}

	file, err := os.Create("result")
	if err != nil {
		fmt.Println(err)
		err := file.Close()
		if err != nil {
			return
		}
		return
	}
	d := []string{"Min: " + strconv.Itoa(calculator.FindMin(arr)),
		"Max: " + strconv.Itoa(calculator.FindMax(arr)),
		"Average: " + fmt.Sprintf("%.2f", calculator.Average(arr)),
	}

	for _, v := range d {
		_, err := fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
