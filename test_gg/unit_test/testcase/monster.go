package testcase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 给monster绑定方法， 可将monster变量，序列化保存到文件
func (this *Monster) Store() bool {

	data, err := json.Marshal(this)
	if err != nil {
		fmt.Printf("marshal err %v", err)
		return false
	}
	filePath := "C:\\localuse\\py_test\\go_test\\testgo\\test_gg\\unit_test\\testcase\\monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Printf("write err %v", err)
		return false
	}
	return true

}

func (this *Monster) ReStore() bool {
	//// restore 反序列化对象
	filePath := "C:\\localuse\\py_test\\go_test\\testgo\\test_gg\\unit_test\\testcase\\monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read err %v", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Printf("unmarshal err %v", err)
		return false
	}
	return true

}
