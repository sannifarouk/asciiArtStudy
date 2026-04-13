package main

import (
	"fmt"
	"os"
	"strings"
)	

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error! Enter at least a letter.")
		return 
	}

	bannerType := ""
	if len(args) < 3 {
		bannerType = "standard.txt"
		// fmt.Println("The Art was printed with 'standard' banner. Kindly input the banner of choice if you wish to change this" )
	} else {
		bannerType = args[2]
	}


	bannerFile, err := os.ReadFile(bannerType)
	if err != nil {
		fmt.Println("Error!", err)
		return
	}

	if args[1] == "" {
		fmt.Print("")
		return 
	}

	splitBanner := strings.Split(string(bannerFile), "\n")
	splitInput := strings.Split(args[1], "\\n")
	
	for _, word := range splitInput {

		for i := 1; i <= 8; i++ {

			for _, char := range word {
				fmt.Print(splitBanner[((int(char) - 32) * 9) + i])
			}
			fmt.Print("\n")
		}
	}

}