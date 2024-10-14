package client

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"task2/client/client"
	"testing"
	"time"
)

// Ensure server is working now to run tests
var ServerUrl = "http://127.0.0.1:8080"

var strings = []string{
	"Hello World!",
	"Hello\nMultiline\nWorld!",
	"",
	"Привет, мир!",
}

func TestVersion(t *testing.T) {
	newClient := client.NewClient(ServerUrl)
	version, err := newClient.GetVersion()
	assert.Equal(
		t, nil, err, "Expected no error",
	)
	assert.Equal(
		t, version, "0.1.1", "Different version",
	)
	fmt.Println(version)
}

func TestDecode(t *testing.T) {
	newClient := client.NewClient(ServerUrl)
	for _, str := range strings {
		resultStr, err := newClient.DecodeString(base64.StdEncoding.EncodeToString([]byte(str)))
		assert.Equal(
			t, nil, err, "Expected no error",
		)
		assert.Equal(
			t, resultStr, str, "Different string",
		)
		fmt.Println(resultStr)
	}
}

// Takes a lot of time to try to check all scenarios in random test
func TestHardOp(t *testing.T) {
	newClient := client.NewClient(ServerUrl)
	for i := 0; i < 10; i++ {
		timeStart := time.Now()
		status, responseCode, seconds, err := newClient.GetHardOp()
		if err != nil {
			fmt.Println("force stopped")
			assert.Greater(t, time.Now().Sub(timeStart), time.Duration(15*time.Second),
				"Error faster than excepted")
		} else {
			assert.Contains(t, []int{500, 200}, responseCode,
				"Unexpected response code")
			assert.Greater(t, time.Now().Sub(timeStart), time.Duration(10*time.Second),
				"Too fast for hard op")
			assert.Less(t, seconds, 15,
				"Too slow to be done")
			fmt.Println(status, responseCode, seconds)
		}
	}
}
