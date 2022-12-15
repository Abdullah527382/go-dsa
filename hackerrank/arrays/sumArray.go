package main

import (
    "fmt"
)

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
    // Write your code here
    var sum int32 = 0
    for i:= 0; i < len(ar); i++ {
        sum += ar[i]
    }
    return sum
}

func main() {

    var ar = []int32{1, 2, 3, 4, 5}

    result := simpleArraySum(ar)

    fmt.Printf("The result of array sum %d\n", result)

}
