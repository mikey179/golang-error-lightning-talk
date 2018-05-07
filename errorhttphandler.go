import "net/http"

// START OMIT
func (fn myHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		mylogger.Println(err)
	}
}

http.Handle("/delete", appHandler(deleteStream(s, u)))
// END OMIT
