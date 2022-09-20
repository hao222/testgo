package person

// 类型 Person 被明确的导出，但是它的字段没有被导出
// 提供 getter() 和 setter() 方法

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

// 模拟子类继承父类的效果 将父类型放在子类型中来实现亚型
type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func (c *Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}
