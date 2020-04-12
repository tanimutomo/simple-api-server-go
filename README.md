# Simple API Server by GO
Features
- Signup and Login
- Post and get a article
- Add tags to articles and get them
- ...


## Start running server
Base endpoint: `http://localhost:8080`

```
go run main.go
```

## Sample Requst JSON
#### Header
```
Content-Type : application/json
```

### Singup and Login
- Signup
```
# Endpoint
/signup

# JSON
{
  "Username": "taroyamada",
  "Password": "taro-pass",
  "Email": "taro@example.com"
}
```

- Login
```
# Endpoint
/login

# JSON
{
  "Username": "taroyamada",
  "Password": "taro-pass",
}
```

### Operation to articles and tags
#### Header
```
Authorization : {token}
- token is included in the response of `/login`
```

- Get all articles
```
# Endpoint
GET) /users/{username}/articles
- username = taroyamada
```

- Post a new article
```
# Endpoint
POST) /users/{username}/articles
- username = taroyamada

# JSON
{
  "Title": "title-a",
  "Content": "content-a",
  "Tags": [{"Name": "tag-a"}]
}
```

- Get all tags
```
# Endpoint
GET) /users/{username}/tags
- username = taroyamada
```

- Add a new tag to the existing article
```
# Endpoint
POST) /users/{username}/articles/{articleID}
- username = taroyamada
- articleID = 1

# JSON
{
  "Name": "tag-b"
}
```


## Refereces
[tanimutomo/gin-api-server-ja-tutorials](https://github.com/tanimutomo/gin-api-server-ja-tutorials)