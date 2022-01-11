package main

import "testing"

func TestAddUpper(t *testing.T) {

	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addupper10 报错， 实际值 %v", res)
	}
	t.Logf("addupper 10 success")
}
