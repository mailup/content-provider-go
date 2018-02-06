package client

import (
	"testing"
)

func Test_Authorize(t *testing.T) {
	token, err := Authorize("fb77e991-61a4-4d8f-a685-1cd6a004edb7", "4e0513da-1ad9-412b-8876-a43dfb678696")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	t.Logf("Token: %v", token)
}
