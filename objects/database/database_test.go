package database

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDatabaseCityMarshalling(t *testing.T) {
	a := &City{Area: "test_area", Region: "test_region", Important: true}
	a.ID = 1
	a.Title = "test_title"

	jsn, _ := json.Marshal(a)

	fmt.Println(string(jsn))

	tst := "{\"id\":1,\"title\":\"test_title\",\"area\":\"test_area\",\"region\":\"test_region\",\"important\":1}"

	tstDBC := &City{}

	b := json.Unmarshal([]byte(tst), tstDBC)

	fmt.Println(b, tstDBC)
}
