package slack

import "testing"

func Test_Client(t *testing.T) {
	config := &Config{APIToken: "token"}
	_, err := config.Client()
	if err != nil {
		t.Fatal(err)
	}
}