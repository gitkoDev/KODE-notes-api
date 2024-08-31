# KODE Notes App
Golang Notes REST API built on `Clean Architecure` and `Dependency Injection` principles

## Endpoints

### Authorization

- 0.0.0.0:8080/auth/sign-up &ensp; `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Create new user
- 0.0.0.0:8080/auth/sign-in &ensp; `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Log in to get access token

### API

- 0.0.0.0:8080/api/v1/notes &ensp; `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Get all notes for user
- 0.0.0.0:8080/api/v1/notes &ensp; `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Add note to user


### Other

- 0.0.0.0:8080/health &ensp; `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Check API connection

## Request body example

### Create User

  ```
  curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"your_name","password":"your_password"}' \
  http://localhost:8080/auth/sign-up
  ```

### Log in

  ```
  curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"your_name","password":"your_password"}' \
  http://localhost:8080/auth/sign-in
  ```
### Add note to a user (using JWT access token)

  ```
curl --header "Content-Type: application/json" \
-H "Authorization: Bearer your_token" \
--request POST \
--data '{"content": "Your new note"}' \
http://localhost:8080/api/v1/notes
  ```

### Get notes of a user (using JWT access token)

  ```
curl --header "Content-Type: application/json" \
-H "Authorization: Bearer your_token" \
--request GET \
http://localhost:8080/api/v1/notes
  ```

## Yandex.Speller integration

<p>
<img src="https://github.com/gitkoDev/KODE-test-task/blob/main/speller.png">
</p>

<p>
<img src="https://github.com/gitkoDev/KODE-test-task/blob/main/note.png">
</p>

