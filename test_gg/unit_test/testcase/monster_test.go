package testcase

import (
	"testing"
)

func TestMonster_Store(t *testing.T) {
	monster := Monster{
		Name:  "xiaoming",
		Age:   10,
		Skill: "eating",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误 %v ", res)
	}
	t.Logf("monster.Store success!!!")

}

func TestMonster_ReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.Restore() 错误 %v ", res)
	}
	if monster.Name != "xiaoming" {
		t.Fatalf("monster.Restore() 错误 %v ", res)
	}
	t.Logf("monster.Restore test success!!!")

}
