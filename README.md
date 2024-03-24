# 👗 outfit-picker API  

## 📄Add Item

- Endpoint
  - /api/items
- Method
    - POST
  - Description
      - 프론트엔드에서 사용자의 옷장에 아이템을 추가합니다.
- Request
    - POST 객체가 전달됨
      - POST 객체
        - itemName : string (추가할 아이템의 이름)
        - category : string (아이템의 분류)
        - image : string (아이템의 사진 URL)
          
          ```json
          { 
              "itemName": "기모후드",
              "category":"아우터",
               "image": "https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png"
          }
          ```
- Response
  - 상태 코드 200 OK
  - 상태 코드 400 Bad Request
    ```json
      {
      "status": "error",
      "message": "Invalid request. Please provide valid data for clothing registration."
      }
      ```

## 📄Delete Item

- Endpoint
  - /api/items/:id
- Method
  - DELETE
- Description
  - 사용자의 옷장에서 선택한 아이템을 제거합니다.
- Request
  - 제거할 아이템 아이디
- Response
  - 상태 코드 200 OK
  - 상태 코드 400 Bad Request

      ```json
      {
        "status": "error",
        "message": "Invalid request. Please provide valid data for clothing registration."
      }
      ```

## 📄Get Item

- Endpoint
  - /api/items
- Method
  - GET
- Description
  - 자신의 옷장에 추가한 전체 의류 아이템을 확인합니다.
- Request
  - 없음
- Response
  - POST 객체가 배열로 전달됨
  - POST 객체
    - itemId : ing
    - itemName : string
    - category : string
    - image : string

      ```json
      [
          {
          "itemId": 1,
          "itemName": "기모후드",
          "category": "아우터",
              "image": "https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png"
        }
      ]
      
## 📄User Register

- Endpoint
  - /api/users
- Method
  - POST
- Description
  - 사용자가 회원 가입을 요청합니다.
- Request
  - POST 객체가 전달됨
  - POST 객체
    - id : string (유저 아이디) **필수**
    - password : string (유저 비밀번호) **필수**
    - name : string (유저 이름) **필수**
    - birthday : string (유저 생일) **필수**
    - phoneNumber : string (유저 전화번호) **필수**
    - gender : int (유저 성별)
      - 0 = male
      - 1 = female

        ```json
        {  
             "userId" : "aba1740", 
             "password" : "dbwls1234", 
             "name" : "박유진", 
             "gender" : false
        }
        ```

- Response
  - 상태 코드 200 OK
  - 상태 코드 400 Bad Request
    - 필수 입력값 누락

        ```json
        {
          "status":  "error",
          "message": "필수 입력값을 입력해주세요."
        }
        ```

    - ID 중복 

      ```json
      {
        "status":  "error",
				"message": "id가 중복되었습니다."
      }
      ```
  - 상태 코드 500 Internal Server Error
    ```json
    {
      "status":  "error",
      "message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
    }
    ```

## 📄Clothing Wear Log

- Endpoint
  - /api/coordis
- Method
  - POST
- Description
  - 사용자가 착용한 옷 사진을 업로드하고 이를 날짜,기온,날씨와 함께 기록합니다.
- Request
  - POST 객체가 전달됨
  - POST 객체
    - date : string (날짜) **필수**
    - photo : string (옷 착용 사진) **필수**
    - temperature : int (기온)
    - weather : int (날씨) **필수**
      - 0 : 맑음
      - 1 : 흐림
      - 2 : 비
      - 3 : 눈

        ```json
        {  
             "date" : "2024-03-24", 
             "photo" : "https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png", 
             "temperature" : "16", 
             "weather" : 1
        }
        ```

- Response
  - 상태 코드 200 OK
  - 상태 코드 400 Bad Request
    - 필수 입력값 누락

      ```json
      {
        "status":  "error",
				"message": "필수 입력값을 입력해주세요."
      }
      ```
  - 상태 코드 500 Internal Server Error
    ```json
    {
      "status":  "error",
      "message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
    }
    ```


## 📄Login

- Endpoint
  - /api/login
- Method
  - POST
- Description
  - 사용자 인증을 위해 로그인을 수행합니다.
- Request
  - POST 객체가 전달됨
  - POST 객체
    - id : string (유저 아이디) **필수**
    - password : string (유저 비밀번호) **필수**
        ```json
        {  
             "id" : "aba1740", 
             "password" : "dbwls1234", 
        }
        ```

- Response
  - 상태 코드 200 OK
  - 상태 코드 400 Bad Request
    - 필수 입력값 누락

      ```json
      {
        "status":  "error",
				"message": "잘못된 요청입니다. 올바른 데이터를 제공해주세요."
      }
      ```
      
    - ID 불일치

      ```json
      {
        "status":  "error",
				"message": "잘못된 로그인 정보입니다. 다시 시도해주세요."
      }
      ```
    - password 불일치

      ```json
      {
        "status":  "error",
				"message": "잘못된 로그인 정보입니다. 다시 시도해주세요."
      }
      ```

## 📄Get Clothing Wear Log List

- Endpoint
  - /api/coordis
- Method
  - GET
- Description
  - 사용자가 원하는 연도와 월에 대해 착용한 코디 목록을 조회할 수 있습니다. 연도와 월을 함께 지정하여 *해당 월에 착용한 코디 목록만*을 반환합니다. 
- Request
  - 쿼리 매개변수
    - month: 정보를 검색할 월 (예: 1월은 "01", 12월은 "12")
    - year: 정보를 검색할 연도 (예: "2024", "2025")
    ```HTTP
    http://localhost:8080/api/coordis?year=2024&month=03
    ```
- Response
  - 상태 코드 200 OK
    - POST 객체가 배열로 전달됨
    - POST 객체
      - id : int (목록 번호)
      - date : string (날짜) 
      - photo : string (코디 사진)
      - temperature : int (기온)
      - weather : int (날씨)
        ```json
        [
            {
            "id": 1,
            "date": "2024-03-18",
            "photo": "https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png",
            "temperature" : 10,
            "weather" : 1
          },
          {
            "id": 6,
            "date": "2024-03-24",
            "photo": "https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png",
            "temperature" : 15,
            "weather" : 0
          }
        ]
  - 상태 코드 404 StatusNotFound
     - 조건에 해당하는 행이 없을 경우
      ```json
      {
        "status":  "error",
        "message": "해당하는 날짜의 코디가 없습니다.",
      }
      ```
  - 상태 코드 500 Internal Server Error
    ```json
    {
      "status":  "error",
      "message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
    }
    ```
    
## 📄Delete Clothing Wear Log

- Endpoint
  - /api/coordis/:id
- Method
  - DELETE
- Description
  - 사용자의 코디 기록에서 해당하는 정보를 삭제합니다. 
- Request
  - 제거할 코디 아이디
- Response
  - 상태 코드 200 OK
  - 상태 코드 404 StatusNotFound
      ```json
      {
        "status":  "error",
				"message": "해당하는 ID를 찾을 수 없습니다."
      }
      ```
  - 상태 코드 500 Internal Server Error
    ```json
    {
      "status":  "error",
      "message": "서버에서 문제가 발생했습니다. 잠시 후에 다시 시도해주세요.",
    }
    ```