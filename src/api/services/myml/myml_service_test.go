package myml

import (
	"net/http"
	"testing"
)

func BenchmarkGetUser(b *testing.B) {
	for n :=0; n < b.N; n++ {
		for n :=0; n < 100; n++ {
			http.Get("http://localhost:8081/user/152581223")
		}
	}
}

func TestGetUser(t *testing.T) {
	result, _ := GetUser(152581223)
	if result == nil {
		t.Error("ERROR")
	}
}


