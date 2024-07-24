# OXO Game API Documentation

This document describes the RESTful API interfaces for the OXO game backend, used for player and level management.

## 1. Player Management System

### List All Players

**Request**

- Method: GET
- Endpoint: `/players`

**Response Example**

```json
[
  {
    "id": 1,
    "name": "Alice",
    "level": "Beginner",
    "balance": 100.5
  },
  {
    "id": 2,
    "name": "Bob",
    "level": "Intermediate",
    "balance": 50.75
  }
]
```

Register a New Player
- Request

- Method: POST
- Endpoint: /players
- Body:

```json
{
    "name": "Carol",
    "level": "Beginner"
}
```
- Response Example

```json

{
    "id": 3
}
```
### Get Details of a Specific Player by ID
- Request

- Method: GET
- Endpoint: /players/{id}
- Response Example



```json
{
    "id": 3,
    "name": "Carol",
    "level": "Beginner",
    "balance": 0.0
}
```
### Update a Specific Player's Information
- Request

- Method: PUT
- Endpoint: /players/{id}
Body:
```json
{
    "name": "Carol Updated",
    "level": "Intermediate"
}
```
Response Example

```json
{
    "message": "player updated successfully"
}
```
### Delete a Specific Player
- Request

- Method: DELETE
- Endpoint: /players/{id}
  Response Example


```json
{
    "message": "player deleted successfully"
}
```
### 2. Levels
- List All Levels
- Request

- Method: GET
- Endpoint: /levels
Response Example

```json
[
    {
        "id": 1,
        "name": "Beginner"
    },
    {
        "id": 2,
        "name": "Intermediate"
    },
    {
        "id": 3,
        "name": "Advanced"
    }
]
```
### Add a New Level
- Request

- Method: POST
- Endpoint: /levels
Body:
```json
{
    "name": "Expert"
}
```
Response Example

```json
{
    "id": 4
}
```

##  Game Room Management System
### 1. List All Game Rooms
   - Request

- Method: GET
- Endpoint: /rooms
- Response Example
```
Status: 200 OK
```
```json
[
{
"id": 1,
"name": "Room A",
"description": "A cozy room for beginners",
"status": "available"
},
{
"id": 2,
"name": "Room B",
"description": "An advanced room for professionals",
"status": "available"
}
]
```
### 2. Get Details of a Specific Game Room by ID
   - Request

- Method: GET
- Endpoint: /rooms/{id}
- Response Example
```
Status: 200 OK
```
```json
{
"id": 1,
"name": "Room A",
"description": "A cozy room for beginners",
"status": "available"
}
```
### 3. Add a New Game Room
   - Request

- Method: POST
- Endpoint: /rooms
Request Body:
```json
{
"name": "Room C",
"description": "A challenging room for experts"
}
```
- Response Example
```
Status: 201 Created
```
```json
{
"id": 3
}
```
### 4. Update a Specific Game Room's Information
   - Request

- Method: PUT
- Endpoint: /rooms/{id}
- Request Body:
```json
{
"name": "Room A (Updated)",
"description": "An updated description"
}
```
- Response Example
``
Status: 204 No Content
``

### 5. Delete a Specific Game Room
   - Request

- Method: DELETE
- Endpoint: /rooms/{id}
- Response Example
```
Status: 204 No Content
```
### 6. List Game Room Reservations
   - Request

- Method: GET
- Endpoint: /reservations
- Query Parameters:
- room_id (optional): Room ID to query.
- date (optional):  Query date in the format yyyy-mm-dd.
- limit (optional): Maximum number of reservations to return.
- Response Example
```
Status: 200 OK
```

```json
[
{
"id": 1,
"room_id": 1,
"date": "2024-07-15",
"time": "14:00",
"player_id": 1
},
{
"id": 2,
"room_id": 1,
"date": "2024-07-15",
"time": "16:00",
"player_id": 2
}
]
```
### 7.Add a New Game Room Reservation

   - Request

- Method: POST
- Endpoint: /reservations
Request Body:
```json
{
"room_id": 1,
"date": "2024-07-20",
"time": "15:00",
"player_id": 3
}
```
```
Status: 201 Created
```

```json
{
"id": 3
}
```
## 3. Endless Challenge System

### Participate in a Challenge


- Method: POST
- Endpoint: /challenges
- Body:
```json
{
"player_id": 123
}
```
- Response Example：


```json
{
"won_jackpot": true
}
```
### List Recent Challenge Results

- Method: GET
- Endpoint: /challenges/results
- Query Parameters:
- n (optional): Number of recent challenges to retrieve (defaults to 10 if not specified).

- - Response Example：
```
Status: 200 OK

```

Response
```json
[
{
"id": 1,
"player_id": 123,
"won_jackpot": true
},
{
"id": 2,
"player_id": 456,
"won_jackpot": false
}
]
```


## 4. Game Log Collector

### Query Game Logs

**Request**

- Method: GET
- Endpoint: `/logs`
- Query Parameters:
    - `player_id` (optional): Player ID to query.
    - `action` (optional): Action type to query. Possible values:
        - Register
        - Login
        - Logout
        - Enter Room
        - Exit Room
        - Participate in Challenge
        - Challenge Result
    - `start_time` and `end_time` (optional): Time range to query.
    - `limit` (optional): Maximum number of logs to return.

**Response Example**


```json
[   
  {       
    "id": 1,        
    "player_id": 123,        
    "action": "Login",        
    "timestamp": 1656739200,        
    "details": "Player 123 logged in"
  },    
  {        
    "id": 2,        
    "player_id": 456,        
    "action": "Participate in Challenge",        
    "timestamp": 1656739500,        
    "details": "Player 456 participated in a challenge"
  }
]
```
###  Add a New Game Log
**Request**

- Method: POST
- Endpoint: /logs
- Body:
```json
{
"player_id": 123,
"action": "Logout",
"details": "Player 123 logged out"
}
```
Response Example

Status: 201 Created
```json
{
"id": 100
}
```
