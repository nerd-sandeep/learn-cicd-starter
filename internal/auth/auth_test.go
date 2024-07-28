package auth

import (
    "reflect"
    "testing"
	"net/http"
	"fmt"

)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization","ApiKey ValidKey")
	apiKey, err := GetAPIKey(headers)
    got := apiKey
    want := "ValidKey1"
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("API Key:", apiKey)
	}

    if !reflect.DeepEqual(want, got) {
         t.Fatalf("expected: %v, got: %v", want, got)
    }

    // headers1 := http.Header{}
	// headers1.Set("Authorization","ApiKey ValidKey1")
	// apiKey1, err1 := GetAPIKey(headers1)
    // got1 := apiKey1
    // want1 := "ValidKey1"
	// if err1 != nil {
	// 	fmt.Println("Error:", err1)
	// } else {
	// 	fmt.Println("API Key:", apiKey1)
	// }

    // if !reflect.DeepEqual(want1, got1) {
    //      t.Fatalf("expected: %v, got: %v", want1, got1)
    // }

}