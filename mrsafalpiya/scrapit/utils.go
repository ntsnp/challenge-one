package scrapit

import (
	"net/url"
	"strings"
)

// Functions
// ---------

/*
 * Extracts the URL parameter from the CSS value.
 *
 * Example: `url(//foo)` -> `foo`
 */
func urlFromCSSVal(v string) string {
	i := strings.Index(v, "(")
	if i >= 0 {
		j := strings.Index(v, ")")
		if j >= 0 {
			/* the link may have leading `//` */
			k := 0
			for ; k < len(v)-i; k++ {
				if v[i+k+1] != '/' {
					break
				}
			}
			return v[i+k+1 : j]
		}
	}
	return ""
}

func cleanUrl(inputUrl string, protocol string, host string) string {
	leadSlashCount := 0
	for i := range inputUrl {
		if inputUrl[i] != '/' {
			break
		}
		leadSlashCount++
	}
	inputUrl = inputUrl[leadSlashCount:]

	_, err := url.ParseRequestURI(inputUrl)
	if err == nil {
		return inputUrl
	}

	splits := strings.Split(inputUrl[leadSlashCount:], "/")
	if !strings.Contains(splits[0], ".") {
		inputUrl, _ = url.JoinPath(protocol, host, inputUrl)
	} else {
		inputUrl, _ = url.JoinPath(protocol, inputUrl)
	}

	return inputUrl
}
