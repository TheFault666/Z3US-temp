package modules

import (
	"fmt"
	"net/http"
	"strings"
)

// DetectBrokenAccessControl checks for potential broken access control issues on the given website URL.
func BrokenAccess(url string) {
	// Parse the input URL
	parsedURL, err := url.Parse(url)
	if err != nil {
		fmt.Println("Error parsing input URL:", err)
		return
	}

	// Taking Paths to test for Broken Access from user as input
	fmt.Print("Enter paths to test (comma-separated): ")
	scanner.Scan()
	pathsInput := scanner.Text()
	pathsToTest := strings.Split(pathsInput, ",")

	// Trim spaces from each path
	for i, path := range pathsToTest {
		pathsToTest[i] = strings.TrimSpace(path)
	}

	for _, path := range pathsToTest {
		// Making the test URL by combining the base URL and the path
		testURL := fmt.Sprintf("%s%s", parsedURL.String(), path)

		// Sending a GET request to the test URL
		response, err := http.Get(testURL)
		if err != nil {
			fmt.Println("Error sending GET request:", err)
			return
		}
		defer response.Body.Close()

		// Checking if the response indicates potential broken access control
		if response.StatusCode == http.StatusOK {
			fmt.Printf("Access to %s is possible. Potential broken access control detected.\n", testURL)
		} else {
			fmt.Printf("Access to %s is not allowed. Access control seems intact.\n", testURL)
		}
	}
}
