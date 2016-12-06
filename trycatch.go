package trycatch

// simple implementation of try catch in golang
// dev branch
// master branch
func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}
