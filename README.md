# 👗 outfit-picker API  

## 📄Add-item

- Endpoint
  - /items
- Method
    - POST
  - Description
      - 프론트엔드에서 사용자의 옷장에 아이템을 추가하기 위한 API
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

## 📄Delete-item

- Endpoint
  - /items/:id
- Method
  - DELETE
- Description
  - 사용자의 옷장에서 선택한 아이템을 제거하기 위한 API
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
## 📄Get-item

- Endpoint
  - /items
- Method
  - GET
- Description
  - 자신의 옷장에 추가한 전체 의류 아이템을 확인하기 위한 API
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
          "itemId": 123,
          "itemName": "기모후드",
          "category": "아우터",
              "image": "https://img.icons8.com/?size=80&id=mw8n5jxdoKlM&format=png"
        }
      ]
      ```
      
## 📄User-register

- Endpoint
  - /users
- Method
  - POST
- Description
  - 회원 가입 API
- Request
  - POST 객체가 전달됨
  - POST 객체
    - userId : string (유저 아이디)
    - password : string (유저 비밀번호)
    - userName : string (유저 이름)
    - gender : bool (유저 성별)
      - ture = male
      - flase = female

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
  - **상태 코드 400 Bad Request**
      ```json
      {
        "status": "error",
        "message": "Invalid request. Please provide valid data for clothing registration."
      }
      ```


