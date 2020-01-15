/*
go test -v
*/
package main
import "testing"
import "fmt"
/*
参与测试的函数必须是TestXxx的形式
*/
func TestPrint(t *testing.T){
    t.SkipNow()//必须写在第一行,暂时无法解决的问题，临时忽略
    res := Print1to20()
    fmt.Println("hey")
    if res != 210{
        //print error infos
        t.Errorf("Wrong result of Print1to20")

    }

}
/*
无法辨认无法test
*/
func testT(t *testing.T){
    fmt.Println("test")
}
/*
测试函数之间相互依赖，比如test1的输出作为test2的输入，可以使用t.Run(),保证test的顺序
这种测试方式称为subtest
*/
func TestPrintx(t *testing.T){
    t.Run("ax",func(t *testing.T){fmt.Println("a1")})
    t.Run("a2",func(t *testing.T){fmt.Println("a2")})
    t.Run("a3",func(t *testing.T){fmt.Println("a3")})
}
/*
TestMain作为初始化test，并且使用m.Run来调用其他tests可以完成一些需要初始化操作的testing,
比如文件打开，rest服务登录等
如果没有在TestMain中调用m.Run()则除了TestMain以外的其他tests都不会被调用执行
*/
func TestMain(m *testing.M){
    fmt.Println("test main first")
    m.Run()
}
/*
benchmark函数一般以Benchmark开头
一般会跑N次，而且每次执行都会如此
在执行过程中会根据实际case的执行时间是否稳定会增加b.N的此时以达到稳定
*/
//go test -bench=.
func BenchmarkAll(b *testing.B){
    //b.N会变化,直到达到稳态
    for n:=0;n<b.N;n++{
        Print1to20()
        //aaa(n)
    }
}
func aaa(n int) int{
    //达不到稳态，会一直跑
    for n > 0{
        n--
    }
    return n
}
