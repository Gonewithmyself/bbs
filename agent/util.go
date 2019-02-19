package agent

func panicError(err error) {
	if nil != err {
		panic(err)
	}
}
