import "net/http"

// START OMIT
type myHttpHandler func(http.ResponseWriter, *http.Request) error

// END OMIT
