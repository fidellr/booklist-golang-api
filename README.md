# online-bookshelf-golang-api
This are a simple CRUD RESTful api implementation with Golang


## How to use:
1. Install mux router
#`go get -u github.com/gorilla/mux`
2. `go build `

## Endpoints

## Get All Books
`GET api/books`

## Get Single Book
`GET api/books/{id}`

## Delete Book
`DELETE api/books/{id}`

## Create Book
POST api/books

# Request sample
`` {
   "isbn":"4545454",
   "title":"Book Three",
   "author":{"firstname":"Jenkins",  "lastname":"Sholo"}
  }``

Update Book
PUT api/books/{id}

# Request sample
`` {
   "isbn":"4545454",
   "title":"Updated Title",
   "author":{"firstname":"Peter",  "lastname":"Kirk"}
  }``

