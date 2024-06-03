![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/KingMichaelPark/gcal-resource-lister)


# Office Room Availability Checker

This project is a Go-based REST API that interacts with the Google Calendar API
to poll all available rooms that are bookable today and tomorrow. The API is
designed to eventually integrate with Slack via a bot, allowing users to easily
check if there are desks available in the office.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Slack Integration](#slack-integration)
- [Contributing](#contributing)
- [License](#license)

## Features
- Polls Google Calendar for available rooms.
- Checks availability for today and tomorrow.
- RESTful API endpoints for querying room availability.
- Future integration with Slack for easy desk availability checks.

## Prerequisites
- Go 1.22 or later
- Google Cloud Project with Google Calendar API enabled
- Service Account with access to the Google Calendar API
- Slack workspace (for future integration)

## Installation
1. **Clone the repository:**
    ```sh
    git clone https://github.com/KingMichaelPark/gcal-resource-lister.git
    cd gcal-resource-lister
    ```
2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Set up Google Calendar API credentials:**

    - Create a service account in your Google Cloud Project.
    - Download the JSON key file and save it in the project directory.

## Configuration

1. **Environment Variables:**

    Set some environment variable for the API key. You can have it automatically
    with `mise`.

    ```env
    GOOGLE_CALENDAR_API_KEY=YOUR_API_KEY
    ```

2. **Google Calendar API:**
    Share your calendar with the service account email to grant access.

## Usage

1. **Run the server:**
    ```sh
    go run src/*
    ```

## API Endpoints

- **GET /rooms**
    Returns the list of available rooms for today.
    ```sh
    curl -X GET http://localhost:8080/rooms?day=today
    ```
- **GET /rooms**
    Returns the list of available rooms for tomorrow.
    ```sh
    curl -X GET http://localhost:8080/rooms?day=tomorrow
    ```

- **GET /rooms/:calendarId**
    Returns the specific rooms availability
    ```sh
    curl -X GET http://localhost:8080/rooms/:calendarId
    ```

## Slack Integration

*Note: This feature is under development.*

The API will eventually integrate with Slack using a bot. Users will be able to
check room availability directly from Slack by interacting with the bot.

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for
more details.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file
for details.

--- Feel free to open an issue or submit a pull request if you have any
questions or suggestions. Happy coding! I am very new to `go` so let me
know if I can be better!


