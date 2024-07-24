package main_test

import "testing"

func TestExample(t *testing.T) {
	t.Log("i couldn't use gomock and ( proto-gen-go and proto-gen-go-grpc) together because there was a dependency error ")
	i := 1
	if i != 1 {
		t.Errorf("Wow! 1 != 1")
	}

}
