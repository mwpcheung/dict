package dict

func raiseGenericError(err error) {
	if err == nil {
		return
	}

	Raise(NewFileGenericError(err.Error()))
}
