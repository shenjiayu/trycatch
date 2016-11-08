package trycatch

// simple implementation of try catch in golang
// this is a dev branch
// resolve conflicts
func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}
