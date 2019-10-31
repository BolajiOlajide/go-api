# People API (GoLang)

Here's me trying to do some weird stuff with GoLang. Tried revamping the output of a tutorial and make use of:

- Go Modules
- Go Packages

## Setup

- Clone the project (obviously, lol)

- Make a copy of the `.env.example` and name it `.env`

- Fill in the content of the `.env` file with the appropriate variables

- Start the application with the command `go run main.go`

### Endpoints

| EndPoint                                 | Functionality                 | Public Access|
| -----------------------------------------|:-----------------------------:|-------------:|
| **GET** /person/{id}                     | Fetch a person by ID          |    TRUE      |
| **GET** /people                          | Fetch every person in the DB  |    TRUE      |
| **POST** /person                         | Create a person               |    TRUE      |

Gracias!
