package main

import ("fmt" ; "math" ; "math/rand" )  //to import multiple packages
                                        // if you remove ; and write packages in different line
                                        //it will work fine
                                        // import("fmt"
                                        //          "math"
                                        //          "math/rand")   
                                        // it will work correctly    
                                        //to import specific  sub package use example "math/rand"



func main() {

fmt.Println("The square root of 81 is ",math.Sqrt(81))
fmt.Println("A Number from 1 to 100 ", rand.Intn(100))
}
	
