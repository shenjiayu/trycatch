package trycatch

// simple implementation of try catch in golang
func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}
