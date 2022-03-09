package main

import (
	"go/ast"
)

//ValueSpec 不仅可以用来表示常量声明还可以用来表示变量声明
//对应一行声明语句
type ValueSpec struct {
	Doc     *ast.CommentGroup //注释
	Names   []*ast.Ident      //常量名
	Type    ast.Expr          //常量类型
	Values  []ast.Expr        //常量初始值
	Comment *ast.CommentGroup //行注释
}

func main() {
	//编译器构造常量伪代码
	//var ValueSpecs []ValueSpec
	//for iota, name := range ValueSpecs {
	//逐个生成常量
	//obj:=types.NewConst(name,iota)
	//...
	//}

}
