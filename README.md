# go-i18n-dictionary

`go-i18n-dictionary` is a Go package for internationalization (i18n) with support for managing dictionaries and formatting messages in multiple languages.

## Installation

To use this package in your Go project, you can install it using `go get`:

```bash
go get -u github.com/adehikmatfr/go-i18n-dictionary
```

## Usage

Here is an example of how to use the i18n-dictionary package:

```go
    dic := dictionary.New()
	fmt.Println(dic.Get(dictionary.Hallo))
	dic.SetLang(dictionary.JP)
	fmt.Println(dic.Get(dictionary.Hallo))
	fmt.Println(dic.Format(dictionary.HalloNamef, "ade"))
```

## Contributing

Feel free to contribute by opening issues or submitting pull requests on the GitHub repository: https://github.com/adehikmatfr/go-i18n-dictionary
