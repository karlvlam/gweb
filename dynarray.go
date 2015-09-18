package main

import "fmt"
import "container/list"

type Box struct {
    a int
    b int
}

func main() {
   var strSlice []string

   // Append and assign to a reference
   strSlice = append(strSlice, "One")
   fmt.Println(strSlice)

   strSlice = append(strSlice, "Two")
   fmt.Println(strSlice)

   strSlice = append(strSlice, "Three")
   fmt.Println(strSlice)

   // remove the last object of the Slice
   strSlice = strSlice[0:len(strSlice)-1]
   fmt.Println(strSlice)

   // create a list and add 3 nodes 
   var strList = list.New()
   strList.PushBack("One")
   printList(strList)
   strList.PushBack("Two")
   printList(strList)
   strList.PushBack("Three")
   printList(strList)

   // remove the last one
   strList.Remove(strList.Back())
   printList(strList)
}

func printList(l *list.List) {
    var s = ""
    for node := l.Front(); node != nil; node = node.Next() {
        s += node.Value.(string) + " "
    }
    fmt.Println(s)
}

