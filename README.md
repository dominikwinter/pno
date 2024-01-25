# Swedish Personnummer Generator

A simple command-line interface (CLI) application written in Go that generates valid [Swedish "Personnummer" (personal identification numbers)](https://en.wikipedia.org/wiki/Personal_identity_number_(Sweden)).

## Usage

a) Download the binary from the releases page or

b) Build the binary from source
```bash
go install github.com/dominikwinter/genpno@latest

genpno
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

## Features

- **Easy to Use:** Simply execute the binary in the CLI to generate a valid Swedish "Personnummer."
- **Planned Flags:** Future updates will include optional flags to specify gender, birthdate, and output format.

## How to Contribute

If you'd like to contribute to the project, feel free to fork the repository and submit a pull request. Contributions are welcome!


## License

This project is licensed under the [MIT License](LICENSE).
