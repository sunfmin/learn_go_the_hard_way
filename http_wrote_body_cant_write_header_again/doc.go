/*
	HTTP Headers should be before start writing HTTP Body

	If you start write body, the http package will automatically help you to write headers before you write the body.
	Normally this will start the writing of body:

		func Hello(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", "Felix")
		}

	So if you want to change http Headers, you must add that before you start writing body, for example:

		func Hello(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Set-Cookie", "...")
			fmt.Fprintf(w, "Hello, %s", "Felix")
		}

	But normally people forget about this, Especially if you write a wrapper to abstract some logic out:

		type SessionHandlerFunc func(session map[interface{}]interface{}, w http.ResponseWriter, r *http.Request)

		func WithSession(shf SessionHandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				s, err := store.Get(r, "tricky_session_name")
				shf(s.Values, w, r)
				s.Save(r, w)
			}
		}

		func UpdateSession2(session map[interface{}]interface{}, w http.ResponseWriter, r *http.Request) {
			session["update2"] = "value2"
			fmt.Fprintf(w, "Hello, %s", "Felix")
		}

	In this case, because s.Save(r, w) will set "Set-Cookie" to headers, But since it's after the fmt.Fprintf in UpdateSession2
	So it won't be write into the http response, So the browser won't get the Set-Cookie instruction, so the cookie is not saved!

*/
package main
