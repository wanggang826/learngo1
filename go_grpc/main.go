package main

import "go_grpc/pb/person"

func main(){
	var p person.Person
	one := p.TestOneOf.(*person.Person_One)
	one.One = "123"  //赋值


}
