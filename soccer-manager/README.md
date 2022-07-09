# Can you kick it

## Overview

This is an example project which demonstrates the use of microservices for a fictional soccer manager. The backend is
powered by 2 microservices, all of which happen to be written in Go, using MongoDB for managing the database and Docker
to isolate and deploy the ecosystem.

* External Service : Manages authentication, middlewares, routing etc. for REST APIs
* Internal Service : Manages all entities like player, team, user etc

## Functionalities

RESTful APIs for a simple application where football/soccer fans will create fantasy teams and will be able to sell or buy players.  
User is able to create an account and log in.
Each user can have only one team (user is identified by an email).  
When the user is signed up, they get a team of 20 players.
Each player has an initial value of $1.000.000.
Each team has an additional $5.000.000 to buy other players.  
When logged in, a user can see their team and player information.  
A team owner can set the player on a transfer list
When a user places a player on a transfer list, they must set the asking price/value for this player.  
When another user/team buys this player, they must be bought for this price.
With each transfer, team budgets are updated in atomic fashion.  
When a player is transferred to another team, their value is increased randomly between 10 and 100 percent.

## Deployment

The application can be deployed in **local machine**. You can find the appropriate documentation in the following link:

* [local machine (docker compose)](./docs/localhost.md)

## Available Endpoints

* [endpoints](./docs/endpoints.md)

