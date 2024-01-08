# github-langs-go

`github-langs-go` is a Go package that provides information about programming languages sourced from [Github Linguist](https://github.com/github-linguist/linguist/blob/master/lib/linguist/languages.yml). It allows users to retrieve details such as color, extensions, and aliases for a given programming language.

## Usage

1. **Installation:**
    ```bash
    go get github.com/NDoolan360/github-langs-go
    ```

2. **Import in Your Code:**
    ```go
    import "github.com/NDoolan360/github-langs-go"
    ```

3. **Example Usage:**
    ```go
    func main() {
        language, err := githublangsgo.GetLanguage("go")
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        fmt.Println("Language Color:", language.Color)  // Language Color: #00ADD8
    }
    ```

## Features

- Retrieve information about programming languages.
- Search by case-insensitive language name or by alias.

## Contributing

If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

Special thanks to the [go-toml](https://github.com/pelletier/go-toml) library for TOML parsing.

## Author

Nathan Doolan
