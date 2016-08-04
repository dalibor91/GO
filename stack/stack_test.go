package main 

import ( 
	"fmt"
	stack "./stack" 
)


func main() {


	s := new(stack.Stack)

	s.Push("This is first line") //0
	s.Push("This is second") //1
	s.Push("Third") //2
	s.Push("Blah")
	s.Push("More ....")
	s.Push("....")

	fmt.Printf("Size of stack is %d\n-----------------------------\n", s.Len())

	s.Print()

	fmt.Printf("\n-----------------------------\n")

	index := 2
	b, err := s.Get(index)

	if err != nil {
		fmt.Println("Error !")
		return 
	}

	fmt.Printf("On index %d in stash we have %s\n", index, b.GetValue())
	fmt.Printf("Next value is %s\n", b.Next().GetValue())
	fmt.Printf("Prev value is %s\n", b.Prev().GetValue())
	fmt.Println("\n\n\n")
	fmt.Printf("Remove %d index from stash\n", index)

	s.Remove(index)

	fmt.Printf("Size of stack is %d\n-----------------------------\n", s.Len())

	s.Print()

	fmt.Printf("\n-----------------------------\n")

	b, err = s.Get(index)

	if err != nil {
		fmt.Println("Error !")
		return 
	}

	fmt.Printf("On index %d in stash we have %s\n", index, b.GetValue())
	fmt.Printf("Next value is %s\n", b.Next().GetValue())
	fmt.Printf("Prev value is %s\n", b.Prev().GetValue())

}
