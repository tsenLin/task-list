package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestValidAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Test case: Valid headers
	t.Run("ValidHeaders", func(t *testing.T) {
		method := "POST"
		path := "/test"
		body := "test body"
		headers := map[string]string{
			"X-Date": time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"),
			"Digest": generateDigest([]byte(body)),
		}

		w, c := createMockTestRequest(method, path, headers, body)

		// signature
		requestLine := generateRequestLine(c.Request.Method, c.Request.URL.Path, c.Request.Proto)
		c.Request.Header.Set("Authorization", fmt.Sprintf("signature=\"%s\"", generateSignature(headers["X-Date"], requestLine, headers["Digest"])))

		ValidAuth(c)

		// Add assertions for the expected behavior based on the test cases
		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d; got %d", http.StatusOK, w.Code)
		}
	})

	// Test case: Valid headers without body
    t.Run("ValidHeadersWithoutBody", func(t *testing.T) {
		method := "GET"
		path := "/test"
		body := ""
		headers := map[string]string{
			"X-Date": time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"),
			"Digest": generateDigest([]byte(body)),
		}

		w, c := createMockTestRequest(method, path, headers, body)

		// signature
		requestLine := generateRequestLine(c.Request.Method, c.Request.URL.Path, c.Request.Proto)
		c.Request.Header.Set("Authorization", fmt.Sprintf("signature=\"%s\"", generateSignature(headers["X-Date"], requestLine, headers["Digest"])))

		ValidAuth(c)

		// Add assertions for the expected behavior based on the test cases
		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d; got %d", http.StatusOK, w.Code)
		}
	})

	// Test case: Missing X-Date header
    t.Run("MissingXDateHeader", func(t *testing.T) {
		method := "POST"
		path := "/test"
		body := "test body"
		headers := map[string]string{
			"Digest": generateDigest([]byte(body)),
		}

		w, c := createMockTestRequest(method, path, headers, body)

		// signature
		requestLine := generateRequestLine(c.Request.Method, c.Request.URL.Path, c.Request.Proto)
		c.Request.Header.Set("Authorization", fmt.Sprintf("signature=\"%s\"", generateSignature(headers["X-Date"], requestLine, headers["Digest"])))

		ValidAuth(c)

		// Add assertions for the expected behavior based on the test cases
		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status %d; got %d", http.StatusUnauthorized, w.Code)
		}
	})

	// Test case: Missing Digest header
    t.Run("MissingDigestHeader", func(t *testing.T) {
		method := "POST"
		path := "/test"
		body := "test body"
		headers := map[string]string{
			"X-Date": time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"),
		}

		w, c := createMockTestRequest(method, path, headers, body)

		// signature
		requestLine := generateRequestLine(c.Request.Method, c.Request.URL.Path, c.Request.Proto)
		c.Request.Header.Set("Authorization", fmt.Sprintf("signature=\"%s\"", generateSignature(headers["X-Date"], requestLine, headers["Digest"])))

		ValidAuth(c)

		// Add assertions for the expected behavior based on the test cases
		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status %d; got %d", http.StatusUnauthorized, w.Code)
		}
	})

	// Test case: Expired request
    t.Run("ExpiredRequest", func(t *testing.T) {
		method := "POST"
		path := "/test"
		body := "test body"
		headers := map[string]string{
			"X-Date": time.Now().Add(-2 * time.Minute).Format("Mon, 02 Jan 2006 15:04:05 MST"),
			"Digest": generateDigest([]byte(body)),
		}

		w, c := createMockTestRequest(method, path, headers, body)

		// signature
		requestLine := generateRequestLine(c.Request.Method, c.Request.URL.Path, c.Request.Proto)
		c.Request.Header.Set("Authorization", fmt.Sprintf("signature=\"%s\"", generateSignature(headers["X-Date"], requestLine, headers["Digest"])))

		ValidAuth(c)

		// Add assertions for the expected behavior based on the test cases
		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status %d; got %d", http.StatusUnauthorized, w.Code)
		}
    })

	// Test case: Future request
    t.Run("FutureRequest", func(t *testing.T) {
		method := "POST"
		path := "/test"
		body := "test body"
		headers := map[string]string{
			"X-Date": time.Now().Add(2 * time.Minute).Format("Mon, 02 Jan 2006 15:04:05 MST"),
			"Digest": generateDigest([]byte(body)),
		}

		w, c := createMockTestRequest(method, path, headers, body)

		// signature
		requestLine := generateRequestLine(c.Request.Method, c.Request.URL.Path, c.Request.Proto)
		c.Request.Header.Set("Authorization", fmt.Sprintf("signature=\"%s\"", generateSignature(headers["X-Date"], requestLine, headers["Digest"])))

		ValidAuth(c)

		// Add assertions for the expected behavior based on the test cases
		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status %d; got %d", http.StatusUnauthorized, w.Code)
		}
    })
}

func createMockTestRequest(method string, path string, headers map[string]string, body string) (*httptest.ResponseRecorder, *gin.Context){
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest(method, path, strings.NewReader(body))

	for key, value := range headers {
		c.Request.Header.Set(key, value)
	}

	return w, c
}
