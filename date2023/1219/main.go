package main
func readChan(chanName <-chan int){
	<-chanName
}
func writeChan(chanName chan<- int){
	chanName <-1
}
func main(){
	mychan:=make(chan int,10)
	writeChan(mychan)
	readChan(mychan)
}
