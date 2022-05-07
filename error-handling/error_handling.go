package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var retry bool = true
	var resource Resource
	for retry {
		resource, err = opener()
		// if getting TransientError then keep re-trying, otherwise return the error
		if err != nil {
			switch err.(type) {
			case TransientError:
				continue
			default:
				return err
			}
		}
		retry = false
	}
	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(r.(FrobError).defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()

	// this one may panic with FrobError (or another type of error)
	// if FrobError call Defrog() using FrobError's .defrobTag as input to Defrob()
	// Use should return this error
	resource.Frob(input)
	return err
}
