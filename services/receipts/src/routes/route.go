package routes

import (
	"net/http"
	"regexp"
)

// the Request struct wraps the http.Request struct, providing a slice of
// strings representing the positional arguments found in a url pattern, and a
// map[string]string called kwargs representing the named parameters captured
// in url parsing.
type Request struct {
	*http.Request
	Args   []string
	Kwargs map[string]string
}

// HandlerFunc is nearly the same as http.HandlerFunc, it simply takes a
// routes.Request object instead of an http.Request object.
type HandlerFunc func(http.ResponseWriter, *Request)

// struct Router implements http.Handler, so that it may be used with the
// default http library.  It keeps a registry mapping regexes to functions for
// easier url parsing.
type Router struct {
	routes []*entry
}

type entry struct {
	route   *Route
	handler HandlerFunc
}

// this doesn't really do anything yet, but it probably will in the future.
func NewRouter() *Router {
	return &Router{}
}

// implements the http.Handler interface, so that we may use our router with
// the default http package.
func (r *Router) ServeHTTP(w http.ResponseWriter, raw *http.Request) {
	req, fn := r.match(raw)
	if req == nil {
		http.NotFound(w, raw)
		return
	}
	fn(w, req)
}

// checks an incoming http request against our list of known routes.  If the
// request matches one of the routes, the request is transformed into a
// routes.Request, and its Args and Kwargs fields are filled in based on the
// url.  If no match is found, returns (nil, nil)
func (r *Router) match(req *http.Request) (*Request, HandlerFunc) {
	for _, e := range r.routes {
		if match := e.route.Match(req.URL.Path); match != nil {
			return &Request{req, match.Args, match.Kwargs}, e.handler
		}
	}
	return nil, nil
}

// adds a regex-based route in the normal human fashion.
func (router *Router) AddRoute(pattern string, fn HandlerFunc) {
	router.routes = append(router.routes, &entry{route: NewRoute(pattern), handler: fn})
}

// right now, just embeds a regex.  A "name" field should also be added here.
type Route struct {
	*regexp.Regexp
}

type RouteMatch struct {
	Args   []string
	Kwargs map[string]string
}

func NewRoute(pattern string) *Route {
	return &Route{regexp.MustCompile(pattern)}
}

func (r *Route) Match(target string) *RouteMatch {
	submatches := r.FindStringSubmatch(target)
	if submatches == nil {
		return nil
	}

	if len(submatches) == 1 {
		return new(RouteMatch)
	}

	m := new(RouteMatch)
	submatches = submatches[1:]
	for i, name := range r.SubexpNames()[1:] {
		if name == "" {
			m.Args = append(m.Args, submatches[i])
		} else {
			if m.Kwargs == nil {
				m.Kwargs = make(map[string]string)
			}
			m.Kwargs[name] = submatches[i]
		}
	}
	return m
}
