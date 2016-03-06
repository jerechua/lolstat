package gong

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Scope struct {
	req *http.Request
	res http.ResponseWriter
}

// Write writes to the body of the response.
func (s *Scope) Write(out []byte) {
	s.res.Write(out)
}

// BodyJSON reads the body of the request and parses it into JSON.
func (s *Scope) BodyJSON(in interface{}) error {
	decoder := json.NewDecoder(s.req.Body)
	decoder.UseNumber()
	if err := decoder.Decode(in); err != nil {
		return fmt.Errorf("unable to decode post payload")
	}
	return nil
}

// Vars returns the map of URL variables to value.
func (s *Scope) Vars() map[string]string {
	return mux.Vars(s.req)
}

// Var returns the URL variable to value.
func (s *Scope) Var(v string) (string, error) {
	val, ok := mux.Vars(s.req)[v]
	if !ok {
		return "", fmt.Errorf("could not find var: %q", v)
	}
	return val, nil
}
