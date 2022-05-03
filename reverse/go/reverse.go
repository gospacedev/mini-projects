//Reverses a string
package main

import "fmt"

func Reverse(s string) (result string){
	for _,v := range s {
	  result = string(v) + result
	}
	return
}

func main(){
  a := Reverse("cool")

  fmt.Println(a)
}
