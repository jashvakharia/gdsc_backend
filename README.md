# Google DSC Executives Task for WebDev backend.

A backend task given by GDSC team from MPSTME.


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






## Authors

- [@jashvakharia](https://www.github.com/jashvakharia)


## Installation

Install python3

```bash
  pip install Flask SQLAlchemy flask_sqlalchemy
  cd gdsc_backend
```
    
## Run Locally

Clone the project

```bash
  git clone https://github.com/jashvakharia/gdsc_backend
```

Go to the project directory

```bash
  cd gdsc_backend
```

Install dependencies

```bash
  Refer Installation.
```

Start the server

```bash
  python app.py
```

