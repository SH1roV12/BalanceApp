package customErrors

import "fmt"


type MyError struct{
	Err error
	RawError error
	Location string
}


func NewRepoAppError(err,raw error, method string)MyError{
	return MyError{
		Err: err,
		RawError: err,
		Location: fmt.Sprintf("repo:%s",method),
	}
}

func(e MyError)Unwrap()error{
	return e.Err
}

func(e MyError)Error()string{
	return fmt.Sprintf("[%s]:%s", e.Location,e.Err)
}