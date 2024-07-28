
# ToFileTextArt

This project takes input text, an optional output flag from the command-line arguments, and writes ASCII-Art graphical representation of the input text in a file specified by the output flag. If the option is not provided, it prints the graphical art in the terminal.

## Documentation

This part contains instructions on how to make use of the program.

## Features

- This program writes asciiart to a file.
- This program allows the user to specify the color of the dispayed text.
- This program allows the user to specify the format/banner to be used in displaying the output.
- We have included an extra banner file that we created by ourselves to add more options to the project

### Usage

To use this program, download and install the latest version of Go from [here](https://go.dev/doc/install).

After cloning the projects repository, navigate to the ascii-art-output directory and execute the following command in your terminal.
```bash
go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```
Provide the output flag and enter the name of the text file where you want to output the graphical ascii art.

Only text files are written to. No other formats will be accepted, an appropriate error message will be displayed and the ascii art will be printed to the terminal instead.

For example:


```bash
go run main.go --output=newfile.txt "Hello world" thinkertoy #Writes the asciiart in Thinkertoy format to the specified text file.
```

## Contributions
Users of this program are allowed to contribute to this project in terms of adding features, or fixing bugs. Just make a pull request.

## Authors

- [@JosephOkumu](https://github.com/JosephOkumu)


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Credits

[Zone01Kisumu](https://zone01kisumu.ke)
