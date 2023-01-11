# Group 4
- Fajar Muhammad Hamka
- Dave Russel
- Richard William
- Indra

## Techstacks Backend
- Golang
- Clean Architecture
- Gin
- JWT
- UUID
- godotoenv
- GORM
- PostgreSQL

## Techstacks Frontend
- React
- React Context
- Next js
- Eslint
- Bootstrap
- React-loader-spinner
- React-toastify
- Cookies-next

## How to use it
- clone this repo
- Go to backend folder & terminal
- create file .env and set up (you can see the env.example)
- You can use database docker:
  - set up docker (optional, you can use by default also (see the password))
  - run `docker compose up` for database
- You can also use your database local (without docker if u want)
- create database mygallery (should be same with env)
- run `go install`
- run `go mod tidy`
- run `go run .`
- Go folder frontend folder & terminal
- Create file .env.local and set up (you can see the env.local.example)
- Run `yarn dev` or `npm run dev`
- Open Link http://localhost:3000 for accessing frontend