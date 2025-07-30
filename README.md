# WeatherCLI

A simple command-line interface (CLI) for fetching weather data using the Weather API. It allows users to get current weather information for a specified city. Made with Go and Cobra.

## Features
- Fetch current weather data for a specified city.

## Requirements
- Go 1.18 or later
- [Weather API key](https://www.weatherapi.com/)
- Cobra library for CLI commands

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/anmol420/WeatherCLI.git
    cd WeatherCLI
    ```
2. Create a `.env` file in the root directory and add your [Weather API key](https://www.weatherapi.com/):
    ```bash
    WEATHER_API_KEY=your_api_key_here
    ```
3. Build the project:
    ```bash
    go build -o weatherCLI.exe
    ```
4. Run the CLI:
    ```bash
    ./weatherCLI.exe current New York
    ```

