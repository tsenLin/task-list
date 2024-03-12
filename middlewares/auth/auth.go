package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"task-list/config"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidAuth(c *gin.Context) {
	xDate := c.GetHeader("X-Date")
	if xDate == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if isRequestExpired(xDate) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	auth, ok := c.Request.Header["Authorization"]
	if !ok || auth[0] == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authMaps := parseAuthorizationHeader(auth[0])
	if _, ok := authMaps["signature"]; !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Get the HTTP method, request path, and protocol version
	// Construct the HTTP request line string
	requestLine := generateRequestLine(c.Request.Method, c.Request.URL.Path, c.Request.Proto)

	byteBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Check digest
	var digest string
	if len(byteBody) > 0 {
		digest = generateDigest(byteBody)
		if digest != c.GetHeader("Digest") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}

	// Generate signature
	realSign := generateSignature(xDate, requestLine, digest)

	if authMaps["signature"] != realSign {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))
}

func isRequestExpired(xdate string) bool {
	xdateTime, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", xdate)
	now := time.Now()
	if xdateTime.After(now) {
		return true
	}

	validPeriod := 1 * time.Minute
	diff := now.Sub(xdateTime)
	return diff > validPeriod
}

func parseAuthorizationHeader(authHeader string) map[string]string {
	authMap := make(map[string]string)

	// Manually trim leading and trailing spaces
    authHeader = strings.TrimSpace(authHeader)
	
	// Split the Authorization header string by commas
    elements := strings.Split(authHeader, ",")

    re := regexp.MustCompile(`\s*([^=]+)="([^"]+)"\s*`)

    // Loop through each element and parse key-value pairs
    for _, element := range elements {
        matches := re.FindStringSubmatch(element)
        if len(matches) == 3 {
            // Store key-value pairs in the map
            authMap[strings.TrimSpace(matches[1])] = strings.TrimSpace(matches[2])
        }
    }
    return authMap
}

func generateRequestLine(method string, path string, protocol string) string {
	return fmt.Sprintf("%s %s %s", method, path, protocol)
}

// GenerateDigest generates the SHA-256 digest of the given byte body.
// It takes a byte slice as input and returns a string.
func generateDigest(byteBody []byte) string{
	if len(byteBody) == 0 {
		return ""
	}
	
	h := sha256.New()
	h.Write(byteBody)

	bodyHash := h.Sum(nil)
	digest := base64.StdEncoding.EncodeToString(bodyHash)

	return fmt.Sprintf("SHA-256=%s", digest)
}

func prepareSignStr(xdate string, requestLine string, digest string) string {
	// if the digest is empty, just return the x-date and request line
	if len(digest) == 0 {
		return fmt.Sprintf("x-date: %s\n%s", xdate, requestLine)
	} else {
		return fmt.Sprintf("x-date: %s\n%s\ndigest: %s", xdate, requestLine, digest)
	}
}

/*
generateSignature generates a signature using HMAC-SHA256 encryption.
sig_str = x-date request-line digest

Parameters:
- xdate: the x-date string
- requestLine: the request line string
- digest: the digest string

Returns:
- the base64-encoded signature string
*/
func generateSignature(xdate string, requestLine string, digest string) string {
	secret := []byte(config.Conf.GetString("HMAC_SECRET"))
	sigStr := prepareSignStr(xdate, requestLine, digest)

	// create a new HMAC by defining the hash type and the key
	hmac := hmac.New(sha256.New, secret)

	// write Data to it
	hmac.Write([]byte(sigStr))

	dataHmac := hmac.Sum(nil)

	// encode bytes to base64 string
	return base64.StdEncoding.EncodeToString(dataHmac)
}