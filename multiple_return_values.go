package main


import "fmt"

func vals() (int,int){
	return 3,7
}

func main(){
	a,b:=vals()
	fmt.Println(a)
	fmt.Println(b)

	f, _:=vals()  //either we can use  f, _ or _ , f but not _ , _ for assigning the two numbers of function returns values
	fmt.Println(f)
}