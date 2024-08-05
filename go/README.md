# GDSC Backend API

Python APP converted to Go.

## Requirements

- Go 1.16 or later

## Installation


   ```sh
   git clone https://github.com/jashvakharia/gdsc_backend
   cd gdsc_backend/go
   go mod init modules
   go get github.com/gin-gonic/gin
   go get gorm.io/gorm
   go get github.com/glebarez/sqlite
   go run main.go
   ```

## API Reference

#### Get the alive message.

```http
  GET /test
```


#### Get all the data stored.

```http
  GET /data
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `x-api-key` | `header` | **Required**. Your API key |


#### Post new data.

```http
  POST /post
```

| Example String |
| :-------- |
| `{"name":"gdsc-jash vakharia","value":"J069"}` 

| Content-type |
| :-------- |
| `application/json` 
