## 爬虫实战

### 基本知识

爬虫大体上分为两种类型，一种是爬取网页源码，另一种是模拟请求API获取数据。今天我们主要是学习后者，直接向网站后台的API发送请求。



#### 什么是API

API（应用程序接口）是一组用于构建和集成应用软件的定义和协议。它充当两个程序之间的“中间人”，想象你去一家餐厅吃饭，你想要吃到好吃的菜品，不可能直接走进后厨去自己做饭，而是打开菜单去点单，等后厨做好了送到你手上。



### 如何进行爬虫

前面直播中，我们学习了如何编写代码发送请求

请求的基本步骤简单总结一下

```text
1.明确请求地址和请求方法

2.携带必要的请求参数

3.处理请求头

4.发送请求并接收响应

5.解析响应内容并根据需求进行处理
```



#### 上手前还需要了解什么



##### JSON

JSON(JavaScript Object Notation),是轻量级的数据交换格式，它可读性强、结构清晰，目前几乎所有编程语言都能解析JSON

API返回的数据通常都是JSON，所以学习爬虫前，我们需要了解JSON格式，学会如何解析

###### JSON基本组成

* 对象（Object）

  用大括号包裹,数据以键值对的形式存在

* 数组（Array）

  用中括号包裹，里面可以有多个元素

* 值（Value）

  值可以是各种数据类型

一个简单的JSON格式返回体

```json
{
  "city": "Wuhan",
  "temperature": 26,
  "forecast": [
    {"date": "2025-09-20", "high": 25, "low": 19, "condition": "Cloudy"},
    {"date": "2025-09-21", "high": 23, "low": 22, "condition": "Rainy"}
  ]
}
```



###### go语言操作JSON

Go内置了encoding/json包，可以很方便地处理JSON



**结构体映射**

结构体(struct)是Go用来定义复杂数据类型的方式，JSON数据可以和Go结构体一一对应

一个简单的示例

```go
type UserMessage struct{
  Username string `json:"name"`
  Password string `json:"password"`
}
```

定义这个结构体（注意字段首字母大写），并在标签里给出与之对应的json字段名。



**Marshal**

通过```json.Marshal```我们可以将Go的数据结构(结构体，map等)序列化成JSON格式的字节切片

函数签名

```go
func Marshal(v interface{})([]byte,error)
```



**Unmarshal**

通过```json.Unmarshal```我们可以将JSON字节切片解析为Go数据结构

函数签名

```go
func Unmarshal(data []byte,v interface{})error
```



了解这两种基本的方法，我们就可以简单的处理请求体和响应体了



##### Cookie

Cookie是网站存储在用户浏览器里的小块数据，它记录用户的状态信息(是否已经登录等)。Cookie一般被放在请求头里，可以用来帮助网站判断你的登录状态。

而爬虫中，获取并在后续请求上带上正确的Cookie，就可以像真实用户一样访问数据

```json
req.Header.Set("Cookie", "sessionid=abc123")
```



学习完这些理论知识，接下来我们开始实战，上手完成第一个爬虫项目吧

---

### **Muxi Library Crawler**

这是木犀直播的后端 API，提供登录、房间信息、用户信息、座位信息等功能。
整体使用流程是：

1. **登录** → 获取 Cookie（相当于登录凭证）
2. **带上 Cookie 调用其他接口**
    * 获取房间号列表
    * 查询用户是否在馆
    * 查询座位信息

---

### 1. 登录接口

**地址**:
`POST https://demo.muxixyz.com/library/login`

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

###  2. 获取房间列表

**地址**:
`GET https://demo.muxixyz.com/library/roomids`

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

### 3. 查询用户是否在馆

**地址**:
`POST https://demo.muxixyz.com/library/inlibrary`

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
    "seat":"xxx",
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

### 4. 获取房间的座位信息

**地址**:
`POST https://demo.muxixyz.com/library/seatinfo`

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
        "seat":"xxx",
        "is_occupied": false,
        "owner": "",
        "start": "",
        "end": ""
      },
      {
        "seat":"xxx",
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

###  整体流程示例

1. 调用 **登录接口**，得到 cookie。
2. 请求时在 header 加上

   ```
   Cookie: JSESSIONID=xxxxxx
   ```
3. 调用 **获取房间列表**，拿到可用房间号。
4. 调用 **查询用户是否在馆** 或 **获取座位信息**。

---

## Tips

* 接口没有做缓存，所以请求可能会比较慢，请耐心等待。
* 登录得到的 `cookie` 会有有效期，过期后需要重新登录获取新的 `cookie`。
* 不同接口对 `cookie` 的要求不同：

   * **获取房间列表**（`/library/roomids`）：不会检查 `cookie` 是否过期，所以即使 `cookie` 失效，这个接口也可能还能正常返回。
   * **查询用户是否在馆**、**获取座位信息**：会严格检查 `cookie`，如果过期会返回空数据或失败。

👉 举个例子：
你调用 `/library/roomids` 得到了房间号，但在 `/library/seatinfo` 里查不到座位信息，这很可能是因为 `cookie` 已经过期了。可以尝试重新登录获取新的 `cookie`。

---

