# Swedish Personnummer Generator

A simple command-line interface (CLI) application written in Go that generates valid [Swedish "Personnummer" (personal identification numbers)](https://en.wikipedia.org/wiki/Personal_identity_number_(Sweden)).

## Usage

a) Download the binary from the releases page or

b) Build the binary from source:
```bash
go install github.com/dominikwinter/genpno/cmd/genpno@latest
```

You can now run it with some options:
```bash
genpno -h
```
Which will list possible options:
```
Usage of genpno:
  -c string
    	Country code (default random)
  -d string
    	Date (format: yyyymmdd) (default random)
  -f string
    	Output forma (default 3):
    		1. yymmddccgn
    		2. yymmdd-ccgn
    		3. yyyymmddccgn
    		4. yyyymmdd-ccgn
  -g string
    	Gender (m/f) (default random)
  -h	Help
  -v	Verbose
```

## Development

1. Clone the repository:

```bash
git clone https://github.com/dominikwinter/genpno.git
```

2. Navigate to the project directory:

```bash
cd genpno
```

3. Build and install the application:

```bash
make
make install
```

4. Run the application:

```bash
genpno
```

## How to Contribute

I'm thrilled that you're considering contributing to this project! No matter how large or small, each contribution makes a difference and I certainly appreciate it!

Here are some ways you can contribute:

- **Bug Reports:** Have you encountered an issue? Please let me know! Be sure to describe the issue, steps to reproduce it, and your environment. This information is all very helpful in coming to a resolution.

- **New Features:** Got an idea for improving the tool? I'd love to hear about it! Let's discuss it in a thread or simply fork the repository, add your feature, and submit a pull request. Do ensure that the feature is described in detail in your pull request, and accompanied by tests that illustrate its correctness.

- **Code Cleanup:** Our code is a constant work in progress and there's always room for improvement in making it more efficient, readable, and maintainable. Feel free to suggest any modification!

- **Improving Tests:** Quality code needs quality tests. If you can provide better test coverage, improve test efficiency, or have any other improvements in mind, I'd be happy to hear about it.

- **Documentation:** If you've noticed a typo or an area that needs better explanation, updates to our documentation are very welcome.

## License

This project is licensed under the [MIT License](LICENSE).
