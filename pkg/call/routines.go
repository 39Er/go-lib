package call

func GoSafe(fn func()) {
	go RunSafe(fn)
}
