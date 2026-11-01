package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSuccessGetAPIKey(t *testing.T) {
	header := http.Header{}
	want := "jddjdjadlkjdoiwdoiwioeu987498489hdhiooioiooi90u09"
	header.Set("Authorization", "ApiKey " + want)
	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("error : %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}


func TestFailGetAPIKey(t *testing.T) {
	header := http.Header{}
	want := "jddjdjadlkjdoiwdoiwioeu987498489hdhiooioiooi90u09"
	Sended := "khdkhsdhsdshdksjhdksjdhksjdhjksdhkjsdhskdhskdhs"
	header.Set("Authorization","ApiKey " + Sended)
	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
	
	if reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestMissingHeader(t *testing.T) {
	header := http.Header{}
	_, err := GetAPIKey(header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got: %v", err)
	}
}

func TestWrongScheme(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "Bearer somekey123")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected error for wrong scheme, got nil")
	}
}

func TestNoKeyAfterPrefix(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected error for missing key, got nil")
	}
}

func TestLowercaseScheme(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "apikey somekey123")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected error for lowercase scheme, got nil")
	}
} 