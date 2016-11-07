package trycatch

// simple implementation of try catch in golang
<<<<<<< HEAD
// this is the master branch
=======
// this is a dev branch
>>>>>>> dev
func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}
