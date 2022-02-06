package main

import "fmt"

func main() {
    vals := []int{0,1,4,5,6,7,8,9,10,14,16,19,41}
    res := search(vals, len(vals)-1, 0, 4)
    fmt.Println(res)
}


func search(values []int, high int, low int, target int) (int) {
    mid := low + (high - low) / 2  + 1
    if(high >= low) {
        if(values[mid] == target) {
            return mid
        } 
        if(target > values[mid]) {
            return search(values, high, mid, target)
        }
        return search(values, mid, low, target);
    }
    return -1
}
