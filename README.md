# OXO Game API 文档

本文档描述了 OXO 游戏后端的 RESTful API 接口，用于玩家和等级管理。

##  1. 玩家管理系统 (Player Management System)

### 列出所有玩家

**请求**

- Method: GET
- Endpoint: `/players`

**响应示例**

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
注册一个新玩家
请求

- Method: POST
- Endpoint: /players
- Body:

```json
{
    "name": "Carol",
    "level": "Beginner"
}
```
- 响应示例

```json

{
    "id": 3
}
```
### 获取特定 ID 的玩家详细信息
- 请求

- Method: GET
- Endpoint: /players/{id}
- 响应示例

```json
{
    "id": 3,
    "name": "Carol",
    "level": "Beginner",
    "balance": 0.0
}
```
### 更新特定 ID 的玩家信息
请求

- Method: PUT
- Endpoint: /players/{id}
Body:
```json
{
    "name": "Carol Updated",
    "level": "Intermediate"
}
```
响应示例

```json
{
    "message": "player updated successfully"
}
```
### 删除特定 ID 的玩家
请求

- Method: DELETE
- Endpoint: /players/{id}
响应示例

```json
{
    "message": "player deleted successfully"
}
```
### 等级 (Levels)
列出所有等级
请求

- Method: GET
- Endpoint: /levels
响应示例

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
### 新增一个等级
请求

- Method: POST
- Endpoint: /levels
Body:
```json
{
    "name": "Expert"
}
```
响应示例

```json
{
    "id": 4
}
```

##  2.  游戏房间管理系统 (Game Room Management System)
### 1. 列出所有游戏房间
   请求

- Method: GET
- Endpoint: /rooms
回应
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
### 2. 获取特定 ID 的游戏房间详细资讯
   请求

- Method: GET
- Endpoint: /rooms/{id}
回应
```
Status: 200 OK
```
ID 为 1 的房间：
```json
{
"id": 1,
"name": "Room A",
"description": "A cozy room for beginners",
"status": "available"
}
```
### 3. 新增一个游戏房间
   请求

- Method: POST
- Endpoint: /rooms
Request Body:
```json
{
"name": "Room C",
"description": "A challenging room for experts"
}
```
回应
```
Status: 201 Created
```
```json
{
"id": 3
}
```
### 4. 更新特定 ID 的游戏房间资讯
   请求

- Method: PUT
- Endpoint: /rooms/{id}
- Request Body:
```json
{
"name": "Room A (Updated)",
"description": "An updated description"
}
```
回应
``
Status: 204 No Content
``

### 5. 删除特定 ID 的游戏房间
   请求

- Method: DELETE
- Endpoint: /rooms/{id}
回应
```
Status: 204 No Content
```
### 6. 查询游戏房间预约
   请求

- Method: GET
- Endpoint: /reservations
- Query Parameters:
- room_id (optional): 房间 ID，指定要查询的房间 ID。
- date (optional): 查询日期，格式为 yyyy-mm-dd。
- limit (optional): 最大返回的预约条数。
回应
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
### 7. 新增一个游戏房间预约
   请求

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
## 3. 无尽挑战系统 (Endless Challenge System)
接口列表
### 参与挑战

- 方法： POST
- URL： /challenges
- 描述： 玩家支付固定金额参与挑战。
请求体：
```json
{
"player_id": 123
}
```
响应：
HTTP 状态码：200 OK
返回体：
```json
{
"won_jackpot": true
}
```
描述：指示玩家是否在此次挑战中赢得了奖池。
### 最近挑战结果列表

- 方法： GET
- URL： /challenges/results
- 描述： 获取最近挑战的结果列表，包括每次挑战是否赢得了奖池。
- 查询参数：
- n（可选）：要检索的最新挑战数量（如果未指定，默认为 10）。

- 响应：
```
HTTP 状态码：200 OK
```

返回体：
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
描述：返回最近挑战的详细信息列表。
## 4. 游戏日志收集器 (Game Log Collector)
接口列表
### 查询游戏日志

- 方法： GET
- URL： /logs
- 描述： 获取游戏操作日志。
- 查询参数：
player_id（可选）：指定要查询的玩家 ID。
action（可选）：指定要查询的操作类型。
注册
登入
登出
进入房间
退出房间
参加挑战
挑战结果
- start_time、end_time（可选）：指定查询的时间范围。
- limit（可选）：指定最多返回的日志条数。
- 响应：
HTTP 状态码：200 OK
返回体：
```json
[
{
"id": 1,
"player_id": 123,
"action": "登入",
"timestamp": 1656739200,
"details": "玩家123登入了游戏"
},
{
"id": 2,
"player_id": 456,
"action": "参加挑战",
"timestamp": 1656739500,
"details": "玩家456参加了挑战"
}
]
```
描述：返回符合查询条件的游戏操作日志列表。
### 新增游戏操作日志

- 方法： POST
- URL： /logs
- 描述： 记录新的游戏操作日志。
请求体：
```json
{
"player_id": 123,
"action": "登出",
"details": "玩家123登出游戏"
}
```
响应：
HTTP 状态码：201 Created
返回体：
```json 
{
"id": 100
}
```
描述：返回新创建的游戏操作日志的 ID。
