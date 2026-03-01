package checker

import (
	"fmt"
	"net/http"
)

func CheckIfUp(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println(url, "is up! (200 OK)")
	} else {
		fmt.Printf("%s returned status: %d\n", url, resp.StatusCode)
	}
}
