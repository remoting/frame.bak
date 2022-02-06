package util

func TryCatch(try func(), catch ...func(exception error)) {
	defer func() {
		if exception := recover(); exception != nil && len(catch) > 0 {
			if v, ok := exception.(error); ok {
				catch[0](v)
			}
		}
	}()
	try()
}
func Throw(exception interface{}) {
	if exception != nil {
		panic(exception)
	}
}
