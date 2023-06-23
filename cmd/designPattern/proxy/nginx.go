package main

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int, 0),
	}
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowd := n.checkRateLimiting(url)
	if !allowd {
		return 403, "Not Allowed"
	}

	return n.application.handleRequest(url, method)
}
