### Install
- Windows:
1. run `SETX GO111MODULES on`
2. open new console && cd to project
3. Pick package 
```
go get -u github.com/sirupsen/logrus
go get -u github.com/labstack/echo/...
go get -u github.com/labstack/echo
go get -u -v github.com/kat-co/vala
```

Code hot reload [air](https://github.com/cosmtrek/air)

4. now check go.mod values 
5. run magic command to download source and create vendor:
```
go mod download
go mod vendor
```

### IF not using vendor mode GO module in different folder

copy onsite folder to GOPATH

`go get -v`

### Logs

### Updates

[] upgrade to v4 or v3.3.10-dev !!


