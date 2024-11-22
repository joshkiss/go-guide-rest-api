`GET     /events`
`GET     /events/<id>`
`POST    /events`         [auth]
`PUT     /events/<id>`    [auth: only by creator]
`DELETE  /events/<id>`    [auth: only by creator]

`POST    /signup`
`POST    /login `      {auth token (JWT)}

`POST    /events/<id>/register`   [auth]
`DELETE  /events/<id>/register`   [auth]
