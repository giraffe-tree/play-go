package main

import "log"

func main() {
	protect(g)

}
func protect(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		// 总结：panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止。
		// recover 只能在 defer 修饰的函数（参见 6.4 节）中使用：用于取得 panic 调用中传递过来的错误值，
		// 如果是正常执行，调用 recover 会返回 nil，且没有其它效果。
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g() //   possible runtime-error
}
func g() {
	panic("A severe error occurred: stopping the program!")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//func errorHandler(fn fType1) fType1 {
//	return func(a type1, b type2) {
//		defer func() {
//			if err, ok := recover().(error); ok {
//				log.Printf("run time panic: %v", err)
//			}
//		}()
//		fn(a, b)
//	}
//}
