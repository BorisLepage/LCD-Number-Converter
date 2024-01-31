package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//    _  _     _  _  _  _  _
//  | _| _||_||_ |_   ||_||_|
//  ||_  _|  | _||_|  ||_| _|

// (each digit is 3 rows high by default)
// modif to test

const rows = 3

var (
	ErrorInvalidNumber = errors.New("expected integer between 1 and 9999 ")
	ErrorInvalidWidth  = errors.New("width expected integer between 1 and 10 ")
	ErrorInvalidHeight = errors.New("height expected integer between 1 and 10 ")
)

type ConversionLCDParams struct {
	Num    int
	Width  int
	Height int
}

func main() {
	var conversionLCDParams ConversionLCDParams
	err := ReadArgsFromTerminal(&conversionLCDParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := ConvertToLCD(conversionLCDParams)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

func ConvertToLCD(params ConversionLCDParams) (string, error) {
	strNums := strconv.Itoa(params.Num)

	var strBuilder strings.Builder
	var strBuffer string
	var addLine int
	var addLineBuffer string

	// We can decompose each number in 3 parts, each part corresponds to a row.
	for row := 0; row < rows; row++ {
		for _, strNum := range strNums {
			representations, ok := representationsLCD[strNum]
			if !ok {
				return "", fmt.Errorf("unknown number: %c", strNum)
			}

			for i, line := range representations[row] {
				switch {
				case line == " ":
					if !contains(representations[row], "_") && i == 1 {
						for w := 0; w < params.Width; w++ {
							strBuffer += line
							addLineBuffer += " "
						}
					} else {
						strBuffer += line
						addLineBuffer += line
					}
				case line == "_":
					for w := 0; w < params.Width; w++ {
						strBuffer += line
						addLineBuffer += " "
					}
				case line == "|":
					strBuffer += line
					addLineBuffer += line
					if params.Height > 1 {
						addLine = params.Height - 1
					}
				}
			}
		}
		// strBuffer += "\n"
		// addLineBuffer += "\n"

		if strings.Contains(addLineBuffer, "|") {
			for i := 0; i < addLine; i++ {
				_, err := strBuilder.WriteString(addLineBuffer)
				if err != nil {
					return "", fmt.Errorf("writing string in builder: %s", addLineBuffer)
				}
			}
		}
		_, err := strBuilder.WriteString(strBuffer)
		if err != nil {
			return "", fmt.Errorf("writing string in builder: %s", strBuffer)
		}

		// Reset var Buffer
		addLineBuffer = ""
		strBuffer = ""
	}

	return strBuilder.String(), nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ReadArgsFromTerminal reads the value from the terminal and validate it.
func ReadArgsFromTerminal(params *ConversionLCDParams) error {
	fmt.Print("Enter a number (1-9999) : ")
	_, err := fmt.Scanln(&params.Num)
	if err != nil {
		return fmt.Errorf("scanning params num : err : %w", err)
	}
	if params.Num < 1 || 9999 < params.Num {
		return fmt.Errorf(
			"scanning params num : err : %w : num : %d", ErrorInvalidNumber, params.Num)
	}

	fmt.Print("Enter the width (1-10) : ")
	_, err = fmt.Scanln(&params.Width)
	if err != nil {
		return fmt.Errorf("scanning params width : err : %w", err)
	}
	if params.Width < 1 || 10 < params.Width {
		return fmt.Errorf(
			"scanning params width : err : %w : num : %d", ErrorInvalidWidth, params.Width)
	}

	fmt.Print("Enter the height (1-10) : ")
	_, err = fmt.Scanln(&params.Height)
	if err != nil {
		return fmt.Errorf("scanning params height : err : %w", err)
	}
	if params.Height < 1 || 10 < params.Height {
		return fmt.Errorf(
			"scanning params height : err : %w : num : %d", ErrorInvalidHeight, params.Height)
	}

	return nil
}

// The representationsLCD map stores the LCD-style representations for the numbers from 0 to 9.
// Each number is associated with a 2D matrix of characters that represents its LCD appearance.
// Each matrix represents the rows of the number, where each character represents a part
// of the number (space, horizontal or vertical line).
// For example, the number '0' is represented by the matrix:
//
//	[" ", "_", " "]
//	["|", " ", "|"]
//	["|", "_", "|"]
//
// This map is used to convert a number into its corresponding LCD-style representation.
var representationsLCD = map[rune][][]string{
	'0': {
		{" ", "_", " "},
		{"|", " ", "|"},
		{"|", "_", "|"},
	},
	'1': {
		{" ", " ", " "},
		{" ", " ", "|"},
		{" ", " ", "|"},
	},
	'2': {
		{" ", "_", " "},
		{" ", "_", "|"},
		{"|", "_", " "},
	},
	'3': {
		{" ", "_", " "},
		{" ", "_", "|"},
		{" ", "_", "|"},
	},
	'4': {
		{" ", " ", " "},
		{"|", "_", "|"},
		{" ", " ", "|"},
	},
	'5': {
		{" ", "_", " "},
		{"|", "_", " "},
		{" ", "_", "|"},
	},
	'6': {
		{" ", "_", " "},
		{"|", "_", " "},
		{"|", "_", "|"},
	},
	'7': {
		{" ", "_", " "},
		{" ", " ", "|"},
		{" ", " ", "|"},
	},
	'8': {
		{" ", "_", " "},
		{"|", "_", "|"},
		{"|", "_", "|"},
	},
	'9': {
		{" ", "_", " "},
		{"|", "_", "|"},
		{" ", "_", "|"},
	},
}
