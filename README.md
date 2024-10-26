# Langmal-Server
<p align="center">
    <img width="256" height="256" src="logo.jpg" alt="Logo of Langmal backend project featuring a blue chinchilla sitting on a server.">
</p>

Welcome to the backend service of the [Langmal](https://github.com/MikolajRatajczyk/Langmal-Apple) quiz app.<br>
This project is written in [Go](https://go.dev/) and covered by **unit tests**.

## Features
- Register, login, and logout user using JWT.
- Store and serve quiz results from a DB.
- Provide a health check endpoint for monitoring.

## Quick start ⚡️
```sh
git clone https://github.com/MikolajRatajczyk/Langmal-Server-Go.git
```
```sh
cd Langmal-Server
```
```sh
go run main.go
```
The server should now be running at `localhost:5001`.

## Config ⚙️
You may provide the full configuration by:
1. Switching to the release mode
```sh
export GIN_MODE=release
```
2. Setting the secret used for JWTs:
```sh
export LANGMAL_JWT_SECRET=YourSecret
```

## Running the tests ✅
```sh
go test ./...
```

---

## API 🔌
Endpoints marked with 🔐 require a JWT in `Authorization` header, example: `Bearer yourJwt`.

### Register
- endpoint: `/user/register`
- method: `POST`
- body:
  ```json
  {
    "email": "foo@bar.com",
    "password": "password"
  }
  ```
- example response:
  ```
  200
  ```
  ```json
  {
    "message": "User has been registered."
  }
  ```
### Log-in
- endpoint: `/user/login`
- method: `POST`
- body:
  ```json
  {
    "email": "foo@bar.com",
    "password": "password"
  }
  ```
- example response:
  ```
  200
  ```
  ```json
  {
    "jwt": "yourJwt"
  }
  ```
### Log-out
- endpoint: `/user/logout`
- method: `POST`
- body:
  ```json
  {
    "token": "yourJwt"
  }
  ```
- example response:
  ```
  200
  ```
  ```json
  {
    "message": "Logged-out (token has been blocked)"
  }
  ```
### Get quizzes 🔐
- endpoint: `/content/quizzes`
- method: `GET`
- example response:
  ```
  200
  ```
  ```json
  [
    {
      "title": "Space and beyond 🚀",
      "id": "5e8ef788-f305-4ee3-ad69-ba8924ca3806",
      "questions": [
        {
          "title": "Which planet is known as the \"Red Planet\" 🔴?",
          "options": [
            "Jupiter",
            "Saturn",
            "Mars"
          ],
          "answer": 2
        },
        {
          "title": "What galaxy is Earth located in 🌌?",
          "options": [
            "Andromeda",
            "Milky Way",
            "Large Magellanic Cloud"
          ],
          "answer": 1
        }
      ]
    }
  ]
  ```
### Get results 🔐
- endpoint: `/content/results`
- method: `GET`
- example response:
  ```
  200
  ```
  ```json
  [
    {
      "correct": 2,
      "wrong": 1,
      "quiz_id": "4e2778d3-57df-4fe9-83ec-afffec1ec5c",
      "created_at": 1729953525,
      "quiz_title": "Giants of the world 🌍"
    }
  ]
  ```
### Save result 🔐
- endpoint: `/content/results`
- method: `POST`
- body:
  ```json
  {
    "created_at": 1729953525,
    "wrong": 1,
    "quiz_id": "4e2778d3-57df-4fe9-83ec-afffec1ec5c",
    "correct": 2
  }
  ```
- example response:
  ```
  201
  ```
  ```json
  {
    "message": "Result saved."
  }
  ```
### Check health
- endpoint: `/health`
- method: `GET`
- example response:
  ```
  200
  ```
