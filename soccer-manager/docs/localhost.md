# Service Localhost Deployment

## Overview

This project can be deployed in a single machine (localhost) using docker compose in order to know the behavior of
microservices.

## Index

- [Service Localhost Deployment](#service-localhost-deployment)
  - [Overview](#overview)
  - [Index](#index)
  - [Requirements](#requirements)
  - [Starting services](#starting-services)
  - [Stoping services](#stoping-services)

## Requirements

* Docker Engine 20.10.11
* Docker Compose 1.29.2
* Create [keyfile](https://www.digitalocean.com/community/tutorials/how-to-configure-keyfile-authentication-for-mongodb-replica-sets-on-ubuntu-20-04#step-2-creating-and-distributing-an-authentication-keyfile) in monodb directory
  
## Starting services

Use the following command to deploy all services in your local environment.

```bash
$ docker-compose up -d

Creating datastore                                   ... done
Creating soccer-manager_soccer-manager-internal_1    ... done
Creating soccer-manager_soccer-manager-external_1    ... done
```

Once the services have started, you can access the web through the following link: <http://localhost:3000/v1/login>.

The following command is an example of how to get the user:

```bash
$ curl --location --request GET 'localhost:3000/v1/user/<userId>' \
--header 'Authorization: Bearer <token from login api>'

{
    "id": "usr-ab196a1d-f576-40f5-92d1-d37f04ed00f1",
    "name": "abc",
    "email": "abc@xyz.com",
    "teamId": "tea-583ea984-c3aa-43e9-beff-6acb9cd60f6c",
    "createdAt": "2022-04-21T22:26:22.532Z"
}
```

## Stoping services

```bash
$ docker-compose stop

Stopping datastore                                   ... done
Stopping soccer-manager_soccer-manager-internal_1    ... done
Stopping soccer-manager_soccer-manager-external_1    ... done
```

