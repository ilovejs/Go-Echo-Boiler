### Install

- on windows:

1. run
`SETX GO111MODULES on`

2. open new console && cd to project

3. Pick package 
```
go get github.com/sirupsen/logrus

//get latest
go get -u github.com/labstack/echo/...
go get -u github.com/labstack/echo
```

- Code hot reload

[air](https://github.com/cosmtrek/air)

4. now check go.mod values 

5. download source and create vendor
```
go mod download
go mod vendor
```

### Logs



