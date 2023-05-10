# catalogServer

catalogServer is an application that simulates the possible catalog of Fluidos, allowing users to view the available flavours of different clusters. It provides a RESTful API to retrieve flavour information from a MongoDB database.

## Routes

The following routes are currently available:

- GET ```/api/flavours```: Retrieves all flavours.
- GET ```/api/flavours/architecture/{architecture}```: Retrieves flavours with a specific architecture.
- GET ```/api/flavours/os/{os}```: Retrieves flavours with a specific operating system.

## Prerequisites

Before running the application, ensure you have the following dependencies installed:

- Go 1.16 or higher
- MongoDB

## About

In order to see how it works, see also [Producer](https://github.com/ilBarlo/FlavourGeneratorProducer) and [Consumer](https://github.com/ilBarlo/FlavourGeneratorConsumer). 
