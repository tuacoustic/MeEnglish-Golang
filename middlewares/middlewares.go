package middlewares

import (
	"me-english/auth"
	"me-english/utils/console"
	"me-english/utils/errorcode"
	"me-english/utils/resp"
	"net/http"
	"strings"
)

type CheckTokenStruct struct {
	WebBackEnd string
	IOS        string
	Android    string
	Cms        string
	Meatshop   string
}

type HeaderRequestStruct struct {
	ClientID       string
	VersionOS      string
	VersionApp     string
	DeviceName     string
	Hashcode       string
	AcceptLanguage string
	Token          string
}

var (
	CheckToken = CheckTokenStruct{
		WebBackEnd: "288d73810e5c7231e115d656de7fea9d",
		IOS:        "e23c79b3e195b0bff7801b7029ceae60",
		Android:    "2e56021a61e0d368106b0f1b38308e5b",
		Cms:        "c3bdaac59f247064b69047d7cd7c45c1",
		Meatshop:   "ead43219efbab1d95bc8b8afd7c6147c",
	}
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		console.Info(r.Proto, "-", r.Method, "-", r.Host+""+r.RequestURI)
		next(w, r)
	}
}
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func SetMiddlewareHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerData := HeaderRequestStruct{
			ClientID:       r.Header.Get("clientid"),
			VersionOS:      r.Header.Get("versionos"),
			VersionApp:     r.Header.Get("versionapp"),
			DeviceName:     r.Header.Get("devicename"),
			Hashcode:       r.Header.Get("hashcode"),
			AcceptLanguage: r.Header.Get("accept-language"),
		}
		switch strings.ToLower(headerData.ClientID) {
		case "0":
			if headerData.Hashcode == CheckToken.WebBackEnd {
				next(w, r)
				return
			}
		case "1":
			if headerData.Hashcode == CheckToken.IOS {
				next(w, r)
				return
			}
		case "2":
			if headerData.Hashcode == CheckToken.Android {
				next(w, r)
				return
			}
		case "3":
			if headerData.Hashcode == CheckToken.Cms {
				next(w, r)
				return
			}
		case "4":
			if headerData.Hashcode == CheckToken.Meatshop {
				next(w, r)
				return
			}
		default:
			resp.Failed(w, http.StatusForbidden, errorcode.GeneralErr.ERR_403)
			return
		}
		resp.Failed(w, http.StatusForbidden, errorcode.GeneralErr.ERR_403)
		return
	}
}

func SetMiddlewareVerifyToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := auth.VerifyToken(r)
		if err != nil {
			resp.Failed(w, http.StatusForbidden, errorcode.GeneralErr.ERR_403)
			return
		}
		r.Header.Set("id", userID)
		next(w, r)
		return
	}
}
