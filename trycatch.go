package trycatch

// simple implementation of try catch in golang
// testing the merge function
// this is a master branch
// resolve conflicts
func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}
