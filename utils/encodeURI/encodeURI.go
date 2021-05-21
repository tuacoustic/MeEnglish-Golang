package encodeURI

import "net/url"

// UrlEncoded encodes a string like Javascript's encodeURIComponent()
func UrlEncoded(str string) string {
	u, _ := url.Parse(str)
	// if err != nil {
	// 	return "", err
	// }
	return u.String()
}
