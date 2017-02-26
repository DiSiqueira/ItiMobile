package itimobile

import (
	"github.com/mssola/user_agent"
	"net/http"
)

// MobileMatcher Store the schema to match with the request Path
type MobileMatcher struct {
	isMobile bool
}

// New is the constructor to ItiMobile
func New(isMobile bool) *MobileMatcher {
	return &MobileMatcher{
		isMobile: isMobile,
	}
}

// Match returns if the request can be handled by this Route.
func (b *MobileMatcher) Match(req *http.Request) bool {
	ua := user_agent.New(req.UserAgent())
	if ua.Mobile() != b.isMobile {
		return false
	}
	return true
}
