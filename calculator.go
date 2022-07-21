package main

import "fmt"

func main() {

  fmt.Print("***WELCOME MY FIRST PROGRAMMER IN LANGUAGE PROGRAMMING 'GO'***")

  fmt.Println("\n\nplease enter you name")
  var name string
  fmt.Scanln(&name)
  
  fmt.Println("\nprogrammer please enter your name.")
  var name_two string
  fmt.Scanln(&name_two)
  
  fmt.Printf("\n\nhi, %s! i´m %s new programmer language Go!", name, name_two)

  fmt.Print("\n\n**************MULTIPLICATION*********")
  
  fmt.Println("\n\nplease enter one number")
  var number int
  fmt.Scanln(&number)

  fmt.Println("\n\nplease enter more one number")
  var number_two int
  fmt.Scanln(&number_two)
  
  fmt.Println("\n\n", number, number_two)

  var multiplication int

  multiplication = number * number_two

  fmt.Printf("\n\n%d", multiplication)

  fmt.Print("\n\n***************DIVISION*********")

  fmt.Println("\n\nplease enter one number")
  var number_div int
  fmt.Scanln(&number_div)

  fmt.Println("\n\nplease enter more one number")
  var number_div_two int
  fmt.Scanln(&number_div_two)

  fmt.Print("\n\n", number_div, number_div_two)

  var division int
  
  division = number_div / number_div_two

  fmt.Println("\n\ntotal value is", division)




  

}