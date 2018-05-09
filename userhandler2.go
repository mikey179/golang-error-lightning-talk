import (
	"fmt"
	"net/http"
)

func deleteStream(s Streams, u Upstream) myHttpHandler {
	return func(user *User, w *responseWriter, r *http.Request) (reterr error) {
		data := map[string]interface{}
		stream, err := s.FindByPath(r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data["title"] = "Internal Server Error"
			data["error"] = "stream_lookup_error"
			reterr = fmt.Errorf("could not load stream %s from database: %v", r.URL.Path, err)
		} else if stream == nil {
			w.WriteHeader(http.StatusNotFound)
			data["title"] = "Not Found"
			data["error"] = "stream_does_not_exist"
		} else if stream.Owner.Username != user.Name {
			w.WriteHeader(http.StatusForbidden)
			data["title"] = "Forbidden"
			data["error"] = "not_allowed_to_delete_stream"
// START OMIT
		} else {
			data["title"] = stream.Title
			nowStreaming, err := u.NowStreaming(stream.Path)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				data["error"] = "stream_availability_error_not_sure_if_safe_to_delete"
				reterr = fmt.Errorf("failed to check availability for stream %q: %v", stream.Path, err)
			} else if nowStreaming {
				w.WriteHeader(http.StatusLocked)
				data["error"] = "stream_can_not_be_deleted_while_streaming"
			} else {
				err := s.Delete(stream)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					data["error"] = "could_not_delete_stream"
					reterr = fmt.Errorf("failed to delete stream %q: %v", stream.Path, err)
				}
			}
		}

		w.Render("deleted.gtmpl", data)
		return
	}
}

// END OMIT
