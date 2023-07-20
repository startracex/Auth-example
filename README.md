# Authorization Application

## Login

`POST` `/login` Allow cors some

Login action

Parameter:

 - Body: E-mail, password(SHA-1 encryption)

## Register

`POST` `/register` Allow cors some

Register action

Parameter:

 - Body: E-mail, password(SHA-1 encryption)

## Authorization

`POST` `/auth` Allow cors all

Return `authorization code`

Parameter:

 - Body: E-mail, password(SHA-1 encryption)

## Application programming interface

`GET` `/api/auth` Allow cors all

Return `temporary token`

Query: 
  + code: `authorization code`

`GET` `/api/user` Allow cors

Return `user information`

Parameter:

 - Header:

   - Token: `temporary token`
