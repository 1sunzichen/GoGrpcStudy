package handler
const HandleName="handle/HServer"
type HelloService struct{}
func (h *HelloService)Hello(req string,reply *string)error{
	*reply="h"+req
	return nil
}