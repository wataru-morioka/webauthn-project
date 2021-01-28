package config

type RestMethod int

const (
    GET RestMethod = iota
    POST
    PUT
    DELETE
)

func (r RestMethod) String() string {
    switch r {
    case GET:
        return "GET"
    case POST:
        return "POST"
    case PUT:
        return "PUT"
    case DELETE:
        return "DELETE"
    default:
        return "Unknown"
    }
}