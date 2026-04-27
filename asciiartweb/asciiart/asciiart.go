package asciiart

import (
	"os"
	"strings"
)	

func AsciiArt(text string, banner string) (string, error) {

	bannerFile, err := os.ReadFile("asciiart/" + banner + ".txt")
	if err != nil {
		return "", err
	}

	splitBanner := strings.Split(string(bannerFile), "\n")
	splitInput := strings.Split(text, "\n")
	
	output := ""

	for _, word := range splitInput {

		for i := 1; i <= 8; i++ {

			for _, char := range word {
				output += (splitBanner[((int(char) - 32) * 9) + i])
			}
			output += "\n"
		}
	}

	// fmt.Println("splitInput:", splitInput)

	return output, err

}


	// args := os.Args
	// if len(args) < 2 {
	// 	fmt.Println("Error! Enter at least a letter.")
	// 	return ""
	// }

	// bannerType := banner
	// // if len(args) < 3 {
	// 	bannerType = "standard.txt"
	// 	// fmt.Println("The Art was printed with 'standard' banner. Kindly input the banner of choice if you wish to change this" )
	// } else {
	// 	bannerType = args[2]
	// }

	// if args[1] == "" {
	// 	fmt.Print("")
	// 	return ""
	// }
