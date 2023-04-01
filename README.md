# notice-latest-program-version

## Overview
This application sends notifications when programming language versions are updated. It uses Google Cloud Functions to automatically detect new releases and send notifications to the targeted users.

## Used Stack

|                | Detail                                                                                        |
|----------------|-----------------------------------------------------------------------------------------------|
| Language       | Go v1.20.0                                                                                    |
| Framework      | Functions Framework for Go v1.6.1                                                             |
| Infrastructure | Google Cloud Functions, Google Cloud Scheduler, Google Pub/Sub, Google Memory Store for Redis |


## Setup

### Clone the repository

```sh
$ git clone https://github.com/your-repo/notice-latest-program-version.git
$ cd notice-latest-program-version
```

### Use Docker to launch the application in your local environment.
```sh
$ docker compose up -d --build
```

## Usage
```sh
$ curl --location --request POST 'http://localhost:8080' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": {
        "data": "eyJmb28iOiAiYmFyIn0="
    }
}'
```

## Infra
https://github.com/S-Ryouta/notice-latest-program-version-tf

