package functional

type List func() (head interface{}, tail List)

func Cons(head interface{}, tail List) List {
	return func() (interface{}, List) { return head, tail }
}

func Head(list List) interface{} {
	head, _ := list()
	return head
}

func Tail(list List) List {
	_, tail := list()
	return tail
}

func Map(fn func(interface{}) interface{}, list List) List {
	if list == nil {
		return nil
	}
	return Cons(fn(Head(list)), Map(fn, Tail(list)))
}

func Reduce(fn func(x, y interface{}) interface{}, list List) interface{} {
	if Tail(list) == nil {
		return Head(list)
	}
	return fn(Head(list), Reduce(fn, Tail(list)))
}
