## Building blocks
- A fork of https://github.com/Massad/gin-boilerplate
- Using https://github.com/appleboy/gin-jwt instead of Redis-based session storage. 

## Configured with

* [go-gorp](github.com/go-gorp/gorp): Go Relational Persistence
* ~~[RedisStore](https://github.com/gin-gonic/contrib/tree/master/sessions): Gin middleware for session management with multi-backend support (currently cookie, Redis).~~
* Configurable JWT session management (stateless)
* Built-in **CORS Middleware**
* Feature **PostgreSQL 9.4** JSON queries
* Unit test

## Installation

1. clone && cd /clone
2. `$ go get -t -v ./...`

> Sometimes you need to get this package manually
```
$ go get github.com/bmizerany/assert
```

You will find the **database.sql** in `db/database.sql`

And you can import the postgres database using this command:
```
$ psql -U postgres -h localhost < ./db/database.sql
```

### Run it
```
$ go run main.go
```

### Build it
```
$ go build -o main main.go
```

### Test it(?)
```
$ go test -v ./tests/*
```

----

## Routing
API routes are nested under `/v1` and the web app under `/`.

Authentication uses JSON Web Tokens (JWTs), which are stateless server-side, meaning you don't have to use Redis or any other session store to keep track of logins. Authenticated routes live under `/v1/auth`. 


## Call and response. 

> Install httpie. 
> - `$ brew install httpie`

### Public
Get all articles
`http -v --json GET localhost:9000/v1/a`

Get one article
`http -v --json GET localhost:9000/v1/a/1`

### Authy
Signin
`http -v POST localhost:9000/v1/signin email=rotblauer@gmail.com password=testes`

Refresh token 
> While you're still logged in, ie. you can't refresh with/an expired token. 
`http -v --json GET localhost:9000/v1/auth/refresh_token "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjkzODYzNzcsImlkIjoiMSIsIm9yaWdfaWF0IjoxNDY5MzgyNzc3LCJ1c2VyX2lkIjoicm90YmxhdWVyQGdtYWlsLmNvbSJ9.qEBqr-XtIsplGnQTa8j_gEGKLrjrwdBG-dJ1FqVHOzg"`

> __There is no signout__. Why not? Tokens are cryptic self-contained bad-ass mutha-fuckas. For one, they have (adjustable) expiration dates. When a token expires you'll be defacto signed out. The other way to manually signout is to just remove the token from the client's memory store (local storage of any variety, or client-handled cookies). Much sexy. Oh and payloads are cool. 

You'll get a response like this when you signin. 
```json
{
    "expire": "2016-07-24T13:33:12-04:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjkzODE1OTIsImlkIjoiMiIsIm9yaWdfaWF0IjoxNDY5Mzc3OTkyfQ.A_qboyShXutO5BUiS_sru8bgryvYkYm3sQE1-vfgm6E"
}
```

Create article
`http -v --json POST localhost:9000/v1/auth/a "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJvdGJsYXVlckBnbWFpbC5jb20iLCJleHAiOjE0Njk0NzgzOTcsIm9yaWdfaWF0IjoxNDY5NDc0Nzk3LCJ1aWQiOiIxIn0.YfOfewNf7PCpy69IqlrpnZCDV_ec_DcC6I3LRluZgiQ" title=hello content=world`

Update article
`http -v --json PUT localhost:9000/v1/auth/a/5 "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJvdGJsYXVlckBnbWFpbC5jb20iLCJleHAiOjE0NjkzOTMyMTAsImlkIjoiMSIsIm9yaWdfaWF0IjoxNDY5Mzg5NjEwfQ.4F65k-x8Zlwfs9Rj9agiupshd2Xz_P1ZW-Mwi0-ykY0" title=hello content=wordies`

Delete article
`http -v --json DELETE localhost:9000/v1/auth/a/1 "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJvdGJsYXVlckBnbWFpbC5jb20iLCJleHAiOjE0NjkzOTMyMTAsImlkIjoiMSIsIm9yaWdfaWF0IjoxNDY5Mzg5NjEwfQ.4F65k-x8Zlwfs9Rj9agiupshd2Xz_P1ZW-Mwi0-ykY0"`

----

## License
(The MIT License)

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
