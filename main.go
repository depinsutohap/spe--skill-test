package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func narcissistic(text string) {
	if _, err := strconv.Atoi(text); err != nil {
		fmt.Printf("%q Must Be Number.\n", text)
	}

	input, _ := strconv.Atoi(text)
	calc := 0
	for _, ch := range text {
		num, _ := strconv.ParseFloat(string(ch), 64)
		res := int(math.Pow(num, float64(len([]rune(text)))))
		calc += res
	}
	if calc != input {
		fmt.Print("false\n")
	} else {
		fmt.Print("true\n")
	}
}

func parityoutliner(text string) {
	arr := strings.Split(text, ",")
	odd := 0
	even := 0
	oddnum := 0
	evennum := 0
	for _, val := range arr {
		num, _ := strconv.Atoi(val)
		if num%2 == 0 {
			even += 1
			evennum = num
		} else {
			odd += 1
			oddnum = num
		}
	}
	if even > odd {
		if odd != 0 {
			fmt.Print(oddnum, " (the only odd number)\n")
		} else {
			fmt.Print("false (all even integer, no outliner)\n")
		}
	} else {
		if even != 0 {
			fmt.Print(evennum, " (the only even number)\n")
		} else {
			fmt.Print("false (all odd integer, no outliner)\n")
		}

	}
}

func findneedle(text, text1 string) {
	arr := strings.Split(text, ",")
	for idx, val := range arr {
		if val == text1 {
			fmt.Println(idx)
		}
	}
}
func blueocean(text, text1 string) {
	arr := strings.Split(text, ",")
	for _, val := range arr {
		if val != text1 {
			fmt.Print(val, ",")
		}
	}
}
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Narcissistic Number")
	fmt.Println("---------------------")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	narcissistic(text)

	fmt.Println("Parity Outliner")
	fmt.Println("---------------------")
	fmt.Println("Example input: 1,2,3,4")
	fmt.Println("No Space")

	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	parityoutliner(text)

	fmt.Println("Find Needle")
	fmt.Println("---------------------")
	fmt.Println("Example input 1: red,blue,yellow")
	fmt.Println("Example input 2: blue")
	fmt.Println("No Space")

	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')
	fmt.Print("-> ")
	text1, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	text1 = strings.Replace(text1, "\n", "", -1)
	findneedle(text, text1)

	fmt.Println("The Blue Ocean Reverse")
	fmt.Println("---------------------")
	fmt.Println("Example input 1: 1,2,3,4,5,6")
	fmt.Println("Example input 2: 1")
	fmt.Println("No Space")

	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')
	fmt.Print("-> ")
	text1, _ = reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	text1 = strings.Replace(text1, "\n", "", -1)
	blueocean(text, text1)

}
