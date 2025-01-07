# Golang Weather

A simple weather application written in Go.

## Features

- Fetch current weather information for a current or specified city
- Supports multiple output formats

## Requirements

- Go (version specified in `go.mod`)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/altynbek07/golang-weather.git
    ```

2. Navigate to the project directory:
   ```sh
   cd golang-weather
   ```

3. Install dependencies:
    ```sh
    go mod tidy
    ```

4. Build the project:
   ```sh
   go build -o weather-app
   ```

5. Run the application:
   ```sh
   ./weather-app -city="London" -format=3
   ```

## Usage

After running the application, you can fetch the weather information by specifying the city and format through optional command-line flags:

- **City**:
  Specify the city for which you want to fetch the weather information.
- **Format**:
  Specify the format of the weather output (1 to 4).

## Example

```bash
./weather-app -format=3
Atyrau: ☁️   -2°C

./weather-app -city="London" -format=4
London: ⛅️  🌡️+3°C 🌬️↓4km/h
```

## Author
Altynbek Kazezov

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
