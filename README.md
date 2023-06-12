# LCD Number Converter

This is a command-line program that converts a given number into an LCD-style representation.
The program takes input for the number, width, and height from the user and generates the corresponding LCD-style representation.

This Project is from [Coding Dojo](https://codingdojo.org/kata/NumberToLCD/).

## Getting Started

To run the LCD Number Converter, follow these steps:

1. Clone the repository:
```shell
git clone <repository_url> lcd-number-converter
```
2. Navigate to the project directory:
```shell
cd lcd-number-converter
```
3. Build the program:
```shell
go build
```
4. Run the program:
```shell
./lcd-number-converter
```

## Usage

The program will prompt you to enter the number, width, and height for the LCD-style representation.
Follow the instructions on the command-line interface and provide the required input.

- Number: Enter an integer between 1 and 9999.
- Width: Enter an integer between 1 and 10.
- Height: Enter an integer between 1 and 10.

After providing the input, the program will generate the LCD-style representation and display it on the console.

## Example

Here's an example usage of the LCD Number Converter:

```shell
Enter a number (1-9999) : 1234
Enter the width (1-10) : 3
Enter the height (1-10) : 3
```
```shell
      ___  ___      
    |    |    ||   |
    |    |    ||   |
    | ___| ___||___|
    ||        |    |
    ||        |    |
    ||___  ___|    |

```

In this example, the input number is 1234, the width is set to 3, and the height is set to 3.
The program generates the corresponding LCD-style representation and displays it on the console.

## Notes

- The program supports numbers between 1 and 9999, with a customizable width and height.
- The LCD-style representations for each digit from 0 to 9 are predefined in the code.
- The program validates the user input to ensure it falls within the specified range.

## Contributing

Contributions to the LCD Number Converter project are welcome!
If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
