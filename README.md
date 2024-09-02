# KODE Notes App
Golang Notes REST API built on `Clean Architecure` and `Dependency Injection` principles. Authentication is implemented using JWT. Each user can add and access their own notes ONLY using a valid JSON Web Token generated on `0.0.0.0:8080/auth/sign-in` endpoint.

There is `KODE.postman_collection.json` Postman collection file in root directory for API testing.


## Startup guide

Run `make run` in root folder to apply all necessary migrations and start the project.

## Endpoints

### Authorization

- 0.0.0.0:8080/auth/sign-up &ensp; `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Create new user
- 0.0.0.0:8080/auth/sign-in &ensp; `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Log in to get access token

### API

- 0.0.0.0:8080/api/v1/notes &ensp; `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Get all notes for user
- 0.0.0.0:8080/api/v1/notes &ensp; `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Add note to user


### Other

- 0.0.0.0:8080/health &ensp; `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Check API connection

# Request and response body example

## Create User

Request
  ```
  curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"your_name","password":"your_password"}' \
  http://localhost:8080/auth/sign-up
  ```
Response

```json
{"id":1}
```

## Log in

Request
  ```
  curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"your_name","password":"your_password"}' \
  http://localhost:8080/auth/sign-in
  ```
Response 

```json
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjUwOTMxODIsImlhdCI6MTcyNTA4OTU4MiwiaWQiOjN9.T_PykTNRzx9ie2RjudyqrPYPuruAx3mo5ti2s5TZq0I"}
```
## Add note to a user (using JWT access token)
Request
  ```
curl --header "Content-Type: application/json" \
-H "Authorization: Bearer your_token" \
--request POST \
--data '{"content": "текст без ашибак"}' \
http://localhost:8080/api/v1/notes
  ```
Response

```json
{"id":1}
```

## Get notes of a user (using JWT access token)
Request
  ```
curl --header "Content-Type: application/json" \
-H "Authorization: Bearer your_token" \
--request GET \
http://localhost:8080/api/v1/notes
  ```
Response
```json
[{"content":"текст без ошибок"}]
```

## Yandex.Speller integration

<p>
<img src="https://github.com/gitkoDev/KODE-test-task/blob/main/speller.png">
</p>

<p>
<img src="https://github.com/gitkoDev/KODE-test-task/blob/main/note.png">
</p>

