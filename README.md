# gotest

testing: Golang Postgress Docker Nginx React.

init go.mod & go.sum
```
cd backend
go mod init gotest
go mod tidy
```

development
```
docker-compose -f docker-compose.dev.yaml up --build
```

production
```
docker-compose up -d
```

need to improve:
- csrf
- jwt include user data