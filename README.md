# Authorization Application

## Login

`GET`  `/login`

Login page

`POST/PUT`  `/login`

Login action

Parameter:

 - Body: E-mail, password(SHA-1 encryption)

## Register

`GET`  `/register`

Register page

`POST/PUT`  `/register`

Register action

Parameter:

 - Body: E-mail, password(SHA-1 encryption)

## Authorization

`GET`  `/auth`

Authorization page

`POST/PUT`  `/auth`

Return `authorization code`

Parameter:

 - Body: E-mail, password(SHA-1 encryption)

## Application programming interface

`GET`  `/api/auth` Allow cors

Return `temporary token`

Query: 
  - code: `authorization code`

`GET`  `/api/user` Allow cors

Return `user information`

Parameter:

 - Header:

   - Token: `temporary token`
