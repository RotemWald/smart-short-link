# smart-short-link

> SelectMedia.asia Home Assignment

* Using Go 1.17
* Make sure you have the [`google/uuid`]((https://github.com/google/uuid)) library installed
* Server is running under port 8080 (http://localhost:8080/)

## Endpoints
* `GET /uuid/{key}` redirects by UUID key
* `GET /counter/{key}` redirects by counter key
* `POST /uuid` creates new UUID shorted link with supplied urls
* `POST /counter` creates new counter shorted link with supplied urls

### Data Types
To create new shorted link please attach the below JSON schema:
```json
[
    {
        "start_hour": 0,
        "end_hour": 13,
        "url": "http://www.google.co.il"
    },
    {
        "start_hour": 13,
        "end_hour": 23,
        "url": "http://www.ynet.co.il"
    }
]
```

As a response you will get a redirected link by the following format (depends on what key type you used):
```json
{
    "url": "http://localhost:8080/uuid/39fbbdb1-f28b-49db-8500-ef90d47d3cb4"
}
```