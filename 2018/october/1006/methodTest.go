package main

import "fmt"

/**
https://github.com/golang/go/wiki/CodeReviewComments#receiver-names
Receiver Names
The name of a method's receiver should be a reflection of its identity;
often a one or two letter abbreviation of its type suffices (such as "c" or "cl" for "Client").
Don't use generic names such as "me", "this" or "self",
identifiers typical of object-oriented languages that place more emphasis on methods as opposed to functions.
The name need not be as descriptive as that of a method argument,
as its role is obvious and serves no documentary purpose.
It can be very short as it will appear on almost every line of every method of the type;
familiarity admits brevity. Be consistent, too:
if you call the receiver "c" in one method,don't call it "cl" in another.
 */
func main() {
	i := Apple{100, 100}
	fmt.Println(&i)
	fmt.Println(i)
	fmt.Println(i.sum())
}

type Apple struct {
	size int
	flag int
}

func (a *Apple) sum() int {
	fmt.Println(a)
	return a.size * a.flag
}

/**
cannot define new methods on non-local type int
类型在其他的，或是非本地的包里定义，在它上面定义方法都会得到和上面同样的错误。
 */
//func (a int) test() int{
//	return 0
//}
