# clean-architecture-telegram-bot

A sample project for creating telegram bots in Golang with clean architecture and repository pattern.

For running the app without building, use the following command:
```
go run main.go serve
```

[!NOTE] 
Although this is a telegram bot, clean architecture principles are the same for any types of web applications.

## Directory Explanation:

## /cmd

Main functionality for the application.

## /configs

Application's Configurations.

## /handlers

Handling all incoming requests.

This directory is responsible for defining endpoints to receive requests, and validating those incoming requests.

It contains no business logic and just passes validated requests to the appropriate services and returns the response to the client.

## /logs

Where all logs are saved.

## /models

Entities that are essential to the application, like structs and types.

## /pkg

Libraries or functions that help to maintain and improve the application.

## /repositories

Databases and their functionalities.

No business logic should be implemented here, just simple `CRUD`.

## /services

Services that application provides.

This is where core business logic exists, it works as a bridge between `handlers` and `repositories` directories.