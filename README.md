# MyGram

My Final Project for Scalable Web Service With Golang by Hactive8

# API Endpoint
## User

### Register User
Endpoint : {url}/user/register

Method : POST

Payload (JSON/Form) :
- username: string
- email: string
- password: string
- age: integer

### Login User
Endpoint : {url}/user/login

Method : POST

Payload (JSON/Form) :
- username: string
- email: string
- password: string

## Photo

### Create Photo
Endpoint : {url}/photo

Method : POST

Payload (JSON/Form) :
- title: string
- caption: string
- photo_url: string

*Authentication Required*

### Get All Photo
Endpoint : {url}/photo

Method : GET

*Authentication Required*

### Get One Photo
Endpoint : {url}/photo/:photo_id

Method : GET

*Authentication Required*

### Update Photo
Endpoint : {url}/photo/:photo_id

Method : POST

Payload (JSON/Form) :
- title: string
- caption: string
- photo_url: string

*Authorization Required*

### Delete Photo
Endpoint : {url}/photo/:photo_id

Method : DELETE

*Authorization Required*

## Comment

### Create Comment
Endpoint : {url}/comment/:photo_id

Method : POST

Payload (JSON/Form) :
- message: string

*Authentication Required*

### Get All Comment
Endpoint : {url}/comment/photo/:photo_id

Method : GET

*Authentication Required*

### Get One Comment
Endpoint : {url}/comment/:photo_id

Method : GET

*Authentication Required*

### Update Comment
Endpoint : {url}/comment/:comment_id

Method : POST

Payload (JSON/Form) :
- message: string

*Authorization Required*

### Delete Comment
Endpoint : {url}/comment/:comment_id

Method : DELETE

*Authorization Required*

## Social Media

### Create Social Media
Endpoint : {url}/socialmedia

Method : POST

Payload (JSON/Form) :
- name: string
- social_media_url: string

*Authentication Required*

### Get All Social Media
Endpoint : {url}/socialmedia

Method : GET

*Authentication Required*

### Get One Social Media
Endpoint : {url}/socialmedia/:socialmedia_id

Method : GET

*Authentication Required*

### Update Social Media
Endpoint : {url}/socialmedia/:socialmedia_id

Method : POST

Payload (JSON/Form) :
- name: string
- social_media_url: string

*Authorization Required*

### Delete Social Media
Endpoint : {url}/socialmedia/:socialmedia_id

Method : DELETE

*Authorization Required*