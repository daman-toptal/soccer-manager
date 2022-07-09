# Service Endpoints

## Login

This endpoint is used to sign up or login using email and password. It returns `Bearer` token which allows access to
other endpoints.

| Service | Method | Endpoint       |
|---------|--------|----------------|
| Login| `POST` | `/v1/login` |

## User

These endpoints return information about the users of our service.

| Service | Method | Endpoint       |
|---------|--------|----------------|
| Get user by Id | `GET` | `/v1/user/{id}` |
| Update user by Id | `PATCH` | `/v1/user/{id}` |

## Player

These endpoints are used to get and update information about a player, check listed players and buy a player.

| Service | Method | Endpoint       |
|---------|--------|----------------|
| Get player by Id | `GET` | `/v1/player/{id}` |
| Update player by Id | `PATCH` | `/v1/player/{id}` |
| Get listed players | `GET` | `/v1/player/listed` |
| Buy player | `POST` | `/v1/player/buy` |

## Team

These endpoints are used to get and update information about a team, get team's players and transactions.

| Service | Method | Endpoint       |
|---------|--------|----------------|
| Get team by Id | `GET` | `/v1/team/{id}` |
| Get players for team | `GET` | `/v1/team/{id}/players` |
| Get transactions for team | `GET` | `/v1/team/{id}/transactions` |

## Transaction

This endpoint is used to get information about a transaction.

| Service | Method | Endpoint       |
|---------|--------|----------------|
| Get transaction by Id | `GET` | `/v1/transaction/{id}` |
