package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ConvertToLCD(t *testing.T) {

	t.Run("digit with default params equal 1", func(t *testing.T) {
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
			"nominal case 2": {
				param: ConversionLCDParams{
					Num:    2,
					Width:  1,
					Height: 1,
				},
				expectedResult: ` _ 
 _|
|_ 
`,
			},
			"nominal case 3": {
				param: ConversionLCDParams{
					Num:    3,
					Width:  1,
					Height: 1,
				},
				expectedResult: ` _ 
 _|
 _|
`,
			},
			"nominal case 4": {
				param: ConversionLCDParams{
					Num:    4,
					Width:  1,
					Height: 1,
				},
				expectedResult: `   
|_|
  |
`,
			},
			"nominal case 5": {
				param: ConversionLCDParams{
					Num:    5,
					Width:  1,
					Height: 1,
				},
				expectedResult: ` _ 
|_ 
 _|
`,
			},
			"nominal case 6": {
				param: ConversionLCDParams{
					Num:    6,
					Width:  1,
					Height: 1,
				},
				expectedResult: ` _ 
|_ 
|_|
`,
			},
			"nominal case 7": {
				param: ConversionLCDParams{
					Num:    7,
					Width:  1,
					Height: 1,
				},
				expectedResult: ` _ 
  |
  |
`,
			},
			"nominal case 8": {
				param: ConversionLCDParams{
					Num:    8,
					Width:  1,
					Height: 1,
				},
				expectedResult: ` _ 
|_|
|_|
`,
			},
			"nominal case 9": {
				param: ConversionLCDParams{
					Num:    9,
					Width:  1,
					Height: 1,
				},
				expectedResult: ` _ 
|_|
 _|
`,
			},
			"nominal case 0": {
				param: ConversionLCDParams{
					Num:    0,
					Width:  1,
					Height: 1,
				},
				expectedResult: ` _ 
| |
|_|
`,
			},
		} {
			t.Run(name, func(t *testing.T) {
				res, err := ConvertToLCD(c.param)
				require.NoError(t, err, "ConvertToLCD should not failed")
				assert.Equal(t, c.expectedResult, res)
			})
		}
	})

	t.Run("number with different params", func(t *testing.T) {
		for name, c := range map[string]struct {
			param          ConversionLCDParams
			expectedResult string
			expectedError  error
		}{
			"number with params 3-3": {
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
			"number with params 5-2": {
				param: ConversionLCDParams{
					Num:    9857,
					Width:  5,
					Height: 2,
				},
				expectedResult: ` _____  _____  _____  _____ 
|     ||     ||            |
|_____||_____||_____       |
      ||     |      |      |
 _____||_____| _____|      |
`,
			},
			"number with params 2-6": {
				param: ConversionLCDParams{
					Num:    5786,
					Width:  2,
					Height: 6,
				},
				expectedResult: ` __  __  __  __ 
|      ||  ||   
|      ||  ||   
|      ||  ||   
|      ||  ||   
|      ||  ||   
|__    ||__||__ 
   |   ||  ||  |
   |   ||  ||  |
   |   ||  ||  |
   |   ||  ||  |
   |   ||  ||  |
 __|   ||__||__|
`,
			},
		} {
			t.Run(name, func(t *testing.T) {
				res, err := ConvertToLCD(c.param)
				require.NoError(t, err, "ConvertToLCD should not failed")
				assert.Equal(t, c.expectedResult, res)
			})
		}
	})
}
