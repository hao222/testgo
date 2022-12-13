package main

// 运算符是一元或二元函数，它返回一个新对象而不修改其参数
// Go 语言并不支持运算符重载：为了克服该限制，运算符必须由函数来模拟

// 1 函数作为运算符 switch断言中测试类型作为包的私有函数，并暴露单一的 Add() 函数作为公共 API
func Add(a Matrix, b Matrix) Matrix {
	switch a.(type) {
	case sparseMatrix:
		switch b.(type) {
		case sparseMatrix:
			return addSparseToSparse(a.(sparseMatrix), b.(sparseMatrix))
		case denseMatrix:
			return addSparseToDense(a.(sparseMatrix), b.(denseMatrix))
			…
		}
	default:
		// 不支持的参数
		…
	}
}

// 2 方法作为运算符  根据接收者类型不同，可以区分不同的方法
func (a *sparseMatrix) Add(b Matrix) Matrix {
	switch b.(type) {
	case sparseMatrix:
		return addSparseToSparse(a.(sparseMatrix), b.(sparseMatrix))
	case denseMatrix:
		return addSparseToDense(a.(sparseMatrix), b.(denseMatrix))
		…
	default:
		// 不支持的参数
		…
	}
}

// 3 使用接口，当在不同类型上 （todo 相似的类型）执行相同的方法时，创建一个通用化的接口以实现多态
type Algebraic interface {
	Add(b Algebraic) Algebraic
	Min(b Algebraic) Algebraic
	Mult(b Algebraic) Algebraic
	Elements()
}

func (a *denseMatrix) Add(b Algebraic) Algebraic {
	switch b.(type) {
	case sparseMatrix:
		return addDenseToSparse(a, b.(sparseMatrix))
		…
	default:
		for x
		in range
		b.Elements() …
	}
}
