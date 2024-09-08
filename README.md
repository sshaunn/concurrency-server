## Installation

To set up the project, follow these steps:

1. Clone the repository:
   git clone https://github.com/yourusername/your-repo-name.git
   cd your-repo-name

2. Install dependencies:
   go mod download
   go mod tidy

## Configuration

Before running the application, make sure to set up the environment:
export ENVIRONMENT=local

You can create different environment configurations by creating `.env.{environment}` files.

## Usage

To run the application:
go run ./cmd/api/main.go

## Gatling performance test

Go to the gatling directory:
cd performance

Build the app:
./gradlew build

To run the gatling app:
./gradlew gatlingRun

## Mock server
To set up the mock server repo 

git clone https://github.com/sshaunn/eventing-tools

cd path/to/your/project

go mod download
go mode tidy

go run main.go
