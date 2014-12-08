package main

import "fmt"

func main() {

    s := make([]string, 5)
    fmt.Println("s:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    s[3] = "d"
    s[4] = "e"

    fmt.Println(s)
    fmt.Println("get 4th :", s[3])

    fmt.Println("len:", len(s))
    s = append(s, "f")
    s = append(s, "g")
    fmt.Println("len:", len(s))

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"h", "i", "j"}
    fmt.Println("t :", t)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}