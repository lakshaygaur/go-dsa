package main

import "fmt"


func canPlaceFlowers(flowerbed []int, n int) bool {
    // prevIndex:=0
    // nextIndex:=0
    if len(flowerbed) == 1 {
        if flowerbed[0] == 0 {
            n--
        }
        if n > 0 {
            return false
        } else {
            return true
        }
    }


    for i:=0; i < len(flowerbed); i++ {
        // beginning array check
        if i==0 {
		 if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
            fmt.Print("\nin beginning", n)
            flowerbed[i]=1
            n--
		 }
        } else if i== len(flowerbed)-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
            // end array check
            fmt.Print("\nin last", n)
            flowerbed[i]=1
            n--
			}
        } else {
			if i != len(flowerbed)-1 && flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0{ // mid array
            fmt.Print("\nin mid", n," i=",i)
            flowerbed[i]=1
            n--
			}
        } 
        if n == 0 {
            return true
        }
    }

    if n > 0 {
        return false
    }
    fmt.Print("ending with last")
    return false

}


func main() {
	inputArr :=[]int{0,1,0}
	n:=1
	canPlaceFlowers(inputArr, n)
}