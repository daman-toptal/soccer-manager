# Can you kick it

## Overview

This is an example project which demonstrates the use of microservices for a fictional soccer manager. The backend is
powered by 2 microservices, all of which happen to be written in Go, using MongoDB for managing the database and Docker
to isolate and deploy the ecosystem.

* External Service : Manages authentication, middlewares, routing etc. for REST APIs
* Internal Service : Manages all entities like player, team, user etc

## Index

* [Deployment](#deployment)
* [Available Endpoints](#available-endpoints)

## Deployment

The application can be deployed in **local machine**. You can find the appropriate documentation in the following link:

* [local machine (docker compose)](./docs/localhost.md)

## Available Endpoints

* [endpoints](./docs/endpoints.md)

