package error

type RouterExistError struct {
	Url string
}


func (error RouterExistError) Error() string{
	return "Url alreday exist. Url:" + error.Url
}