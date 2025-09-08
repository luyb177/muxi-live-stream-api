## 📖 木犀直播 API 使用说明

这是木犀直播的后端 API，提供登录、房间信息、用户信息、座位信息等功能。
整体使用流程是：

1. **登录** → 获取 Cookie（相当于登录凭证）
2. **带上 Cookie 调用其他接口**

    * 获取房间号列表
    * 查询用户是否在馆
    * 查询座位信息

---

### 🔑 1. 登录接口

**地址**:
`POST /library/login`

**参数**:

```json
{
  "username": "学号",
  "password": "密码"
}
```

**返回**:

```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "cookie": [
      "ASP.NET_SessionId=2qss5545blobz2vbdjceqxev"
    ]
  }
}
```

👉 **说明**:

* 登录成功后会返回一个 `cookie`。
* **后续请求都必须带上这个 cookie**，否则会提示`cookie`不能为空。

---

### 🏠 2. 获取房间列表

**地址**:
`GET /library/roomids`

**返回**:

```json
{
  "code": 200,
  "message": "ok",
  "data": {
    "room_ids": ["101", "102", "103"]
  }
}
```

👉 **说明**:

* 每个房间都有唯一的 `roomId`，后续接口会用到。

---

### 👤 3. 查询用户是否在馆

**地址**:
`POST /library/inlibrary`

**参数**:

```json
{
  "name": "张三",
  "room_ids": ["101", "102"]
}
```

**返回**:

```json
{
  "code": 200,
  "message": "xxx 在图书馆的xxx，09:40 - 13:55\n",
  "data": {
    "is_in_library": true,
    "area": "xxx",
    "start": "09:40",
    "end": "13:55"
  }
}
```

👉 **说明**:

* `is_in_library` 表示是否在馆
* 如果在馆，会返回具体的区域、开始时间和结束时间

---

### 💺 4. 获取房间的座位信息

**地址**:
`POST /library/seatinfo`

**参数**:

```json
{
  "room_ids": ["101"]
}
```

**返回**:

```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "seat_infos": [
      {
        "is_occupied": false,
        "owner": "",
        "start": "",
        "end": ""
      },
      {
        "is_occupied": true,
        "owner": "xxx",
        "start": "2025-09-08 08:38",
        "end": "2025-09-08 12:37"
      }
    ]
  }
}
```

👉 **说明**:

* 每个 `seat_info` 表示一个座位
* `is_occupied = true` 时，说明座位有人，并返回占用者及时间段

---

### 🚀 整体流程示例

1. 调用 **登录接口**，得到 cookie。
2. 请求时在 header 加上

   ```
   Cookie: JSESSIONID=xxxxxx
   ```
3. 调用 **获取房间列表**，拿到可用房间号。
4. 调用 **查询用户是否在馆** 或 **获取座位信息**。

---

## 💡 Tips

* 接口没有做缓存，所以请求可能会比较慢，请耐心等待。
* 登录得到的 `cookie` 会有有效期，过期后需要重新登录获取新的 `cookie`。
* 不同接口对 `cookie` 的要求不同：

   * **获取房间列表**（`/library/roomids`）：不会检查 `cookie` 是否过期，所以即使 `cookie` 失效，这个接口也可能还能正常返回。
   * **查询用户是否在馆**、**获取座位信息**：会严格检查 `cookie`，如果过期会返回空数据或失败。

👉 举个例子：
你调用 `/library/roomids` 得到了房间号，但在 `/library/seatinfo` 里查不到座位信息，这很可能是因为 `cookie` 已经过期了。可以尝试重新登录获取新的 `cookie`。

