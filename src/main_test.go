package main

import (
	"testing"
)

func Test_ConvertToLCD(t *testing.T) {
	for name, c := range map[string]struct {
		param          ConversionLCDParams
		expectedResult string
		expectedError  error
	}{
		"nominal case 1": {
			param: ConversionLCDParams{
				Num:    1,
				Width:  1,
				Height: 1,
			},
			expectedResult: `   
  |
  |
`,
		},
		"multiple number with variable params": {
			param: ConversionLCDParams{
				Num:    1478,
				Width:  3,
				Height: 3,
			},
			expectedResult: `           ___  ___ 
    ||   |    ||   |
    ||   |    ||   |
    ||___|    ||___|
    |    |    ||   |
    |    |    ||   |
    |    |    ||___|
`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			res, err := ConvertToLCD(c.param)
			if err != nil {
				t.Error("ConvertToLCD should not failed", err)
			}

			if c.expectedResult != res {
				t.Error("ConvertToLCD unexpected result\n", res, c.expectedResult)
			}
		})
	}
}
