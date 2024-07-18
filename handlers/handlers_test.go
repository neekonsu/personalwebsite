package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestServeHTML(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "test_static")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test index.html file
	indexContent := "<html><body>Test Index</body></html>"
	err = os.WriteFile(filepath.Join(tempDir, "index.html"), []byte(indexContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test page
	pageContent := "<html><body>Test Page</body></html>"
	err = os.MkdirAll(filepath.Join(tempDir, "pages"), 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(tempDir, "pages", "testpage.html"), []byte(pageContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Override the default static directory for testing
	originalStaticDir := "static"
	os.Setenv("STATIC_DIR", tempDir)
	defer os.Setenv("STATIC_DIR", originalStaticDir)

	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{"Home Page", "/", http.StatusOK, indexContent},
		{"Existing Page", "/testpage", http.StatusOK, pageContent},
		{"Non-existent Page", "/nonexistent", http.StatusNotFound, "404 page not found\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ServeHTML)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
