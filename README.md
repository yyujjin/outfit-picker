
# 👗 outfit-picker API

## users

#### Endpoint
```http
/api/users
```
<details>
 <summary>-<code>POST</code> </summary>

##### Description
사용자가 회원 가입을 요청합니다.

##### Request

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |id|필수|string|유저 아이디|
> |password|필수|string|유저 비밀번호|
> |name|필수|string|유저 이름|
> |birthday|필수|string|유저 생일|
> |phoneNumber|필수|string|유저 전화번호|
> |gender|선택|int|유저 성별 ( **default** 0 = male  , 1 = female)|


##### Example JSON
```json
  {  
    "id":"yaho", 
    "password":"lululala123", 
    "name":"박유진",
    "birthday":"19961204",
    "phoneNumber":"010-1234-5678", 
    "gender":1
  }
```
> 
##### Responses

> | http code  | case|response     |
> |--------------------------|-----------------------------------|-------------------------|
> | `200`         |        | |
> | `400`         | 필수 입력값 누락|"status":  "error","message": "필수 입력값을 입력해주세요."|
> |   `400`        |   ID 중복     |  "status":  "error","message": "id가 중복되었습니다." |
> | `500`         |        | "status":  "error","message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요."|
</details>

------------------------------------------------------------------------------------------

## login

#### Endpoint
```http
/api/login
```
<details>
 <summary>-<code>POST</code> </summary>

##### Description
사용자가 인증을 위해 로그인을 수행합니다.

##### Request

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |id|필수|string|유저 아이디|
> |password|필수|string|유저 비밀번호|

##### Example JSON
```json
{  
  "id":"yaho", 
  "password":"lululala123"
}
```
> 
##### Responses

> | http code  | case|response     |
> |--------------------------|-----------------------------------|-------------------------|
> | `200`         |        | |
> | `400`         | 필수 입력값 누락|"status":  "error","message": "잘못된 요청입니다. 올바른 데이터를 제공해주세요."|
> |    `400`      |   id 불일치     |  "status":  "error","message": "잘못된 로그인 정보입니다. 다시 시도해주세요."|
> |    `400`      |   password 불일치     | "status":  "error","message": "잘못된 로그인 정보입니다. 다시 시도해주세요." |
</details>

------------------------------------------------------------------------------------------

## items

#### Endpoint
```http
/api/items
```

<details>
 <summary>-<code>GET</code> </summary>

##### Description
자신의 옷장에 추가한 전체 의류 아이템을 확인합니다.

##### Request

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |N/A|N/A|N/A|N/A|

##### Responses

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |itemId|필수|int|아이템 아이디|
> |itemName|필수|string|아이템 이름|
> |category|필수|string|아이템 분류|
> |image|필수|string|아이템 사진|

##### Example JSON
```json
[
  {
    "itemId":1,
    "itemName":"기모후드",
    "category":"아우터",
    "image":"https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png"
  }
]
```
</details>

------------------------------------------------------------------------------------------
<details>
 <summary>-<code>POST</code> </summary>

##### Description
사용자의 옷장에 아이템을 추가합니다. 

##### Request

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |itemName|필수|string|아이템 이름|
> |cagegory|필수|string|아이템 분류|
> |image|필수|string|아이템 사진|

##### Example JSON
```json
{
  "itemName":"기모후드",
  "category":"아우터",
  "image":"https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png"
}
```
> 
##### Responses

> | http code  | response     |
> |--------------------------|-----------------------------------|
> | `200`         |        |
> | `400`         | "status": "error", "message": "Invalid request. Please provide valid data for clothing registration."|



</details>

------------------------------------------------------------------------------------------

<details>
 <summary>-<code>DELETE</code> </summary>

##### Description
사용자의 옷장에서 선택한 아이템을 제거합니다.

##### Request

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |id|필수|semantics|제거할 아이템 번호|

##### Example URL
```HTTP
http://localhost:8080/api/items/4
```
> 
##### Responses

> | http code  | response     |
> |--------------------------|-----------------------------------|
> | `200`         |        |
> | `400`         | "status": "error", "message": "Invalid request. Please provide valid data for clothing registration."|

</details>

------------------------------------------------------------------------------------------




## coordis

#### Endpoint
```http
/api/coordis
```

<details>
 <summary>-<code>GET</code> </summary>

##### Description
사용자가 원하는 연도와 월에 대해 착용한 코디 목록을 조회할 수 있습니다. 연도와 월을 함께 지정하여 *해당 월에 착용한 코디 목록만을 반환합니다.*

##### Request

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |month|필수|query|검색할 월 (*1월은 "01",12월은 "12"*) |
> |year|필수|query|검색할 연도 (*"2024", "2025"*)|

##### Example URL
```HTTP
http://localhost:8080/api/coordis?year=2024&month=04
```

##### Responses

>| http code  | case|response     |
> |--------------------------|-----------------------------------|-------------------------|
> | `200`         |        ||

>| id (목록 번호)  |  data (날짜)    | photo  (코디 사진) | temperature (기온) | weather (날씨)|
> |----------|----------|-----------|----------|----------|
> |int| string     |  string |int   |int|

##### Example JSON
```json
[
  {
    "id":1,
    "date":"2024-03-18",
    "photo":"https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png",
    "temperature" :10,
    "weather" :1
  },
  {
    "id":6,
    "date":"2024-03-24",
    "photo":"https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png",
    "temperature":15,
    "weather":0
  }
]
```

> | http code  | case|response     |
> |--------------------------|-----------------------------------|-------------------------|
> | `404`         | 조건에 해당하는 행이 없을 경우 | "status":  "error","message": "해당하는 날짜의 코디가 없습니다."|
> | `500`         |        | "status":  "error","message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요."|


</details>

------------------------------------------------------------------------------------------

<details>
 <summary>-<code>POST</code> </summary>

##### Description
사용자가 착용한 옷 사진을 업로드하고 이를 날짜, 기온, 날씨와 함께 기록합니다.

##### Request


> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |date|필수|string|날짜|
> |image|필수|string|코디 사진|
> |temperature|선택|int|기온|
> |weather|필수|int|날씨 (*0 = 맑음, 1 = 흐림, 2 = 비, 3 = 눈*)|

##### Example JSON
```json
  {  
    "date" :"2024-03-24", 
    "photo":"https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png", 
    "temperature":"16", 
    "weather":1
  }
```
> 
##### Responses

> | http code  | case|response     |
> |--------------------------|-----------------------------------|-------------------------|
>|`200`| | |
> | `400`         | 필수 입력 값 누락 | "status":  "error","message": "필수 입력값을 입력해주세요."|
> | `500`         |        | "status":  "error","message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요."|



</details>

------------------------------------------------------------------------------------------

<details>
 <summary>-<code>DELETE</code> </summary>

##### Description
사용자의 코디 기록에서 해당하는 정보를 삭제합니다.



##### Request

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |id|필수|semantics|제거할 코디 번호 |

##### Example URL
```HTTP
http://localhost:8080/api/coordis/4
```
> 
##### Responses

> | http code  | case|response     |
> |--------------------------|-----------------------------------|-------------------------|
>|`200`| | |
> | `400`         |  | "status":  "error","message": "해당하는 ID를 찾을 수 없습니다." |
> | `500`         |        | "status":  "error","message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요."|


</details>

------------------------------------------------------------------------------------------


## categories

#### Endpoint
```http
/api/categories
```

<details>
 <summary>-<code>GET</code> </summary>

##### Description
카테고리 리스트를 전달합니다.

##### Requests

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |N/A|N/A|N/A|N/A|

##### Responses

> |name|type|data type|description|
>|---------|--------|----------|-----------|
> |id|필수|int|카데고리 번호|
> |name|필수|string|카데고리 이름|

##### Example JSON
```json
[
  {
    "id":1,
    "name":"아우터"
  },
  {
    "id":2,
    "name":"상의"
  },
  {
    "id":3,
    "name":"하의"
  }
]
```

</details>

------------------------------------------------------------------------------------------
