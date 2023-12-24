package authn

import (
	"net/http"
	"net/url"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gweffectx/safedav/internal/conf"
	"github.com/gweffectx/safedav/internal/setting"
	"github.com/gweffectx/safedav/server/common"
)

func NewAuthnInstance(r *http.Request) (*webauthn.WebAuthn, error) {
	siteUrl, err := url.Parse(common.GetApiUrl(r))
	if err != nil {
		return nil, err
	}
	return webauthn.New(&webauthn.Config{
		RPDisplayName: setting.GetStr(conf.SiteTitle),
		RPID:          siteUrl.Hostname(),
		//RPOrigin:      siteUrl.String(),
		RPOrigins: []string{siteUrl.String()},
		// RPOrigin: "http://localhost:5173"
	})
}
