package nodechoice

import (
    "testing"
    "fmt"
)

func TestUpdateNode1(t *testing.T) {
   head := CreateNode("", 0)
   UpdateList(head, "127.0.0.1:6666", 23)
   PrintList(head)
   UpdateList(head, "127.0.0.1:7777", 55)
   PrintList(head)
   UpdateList(head, "127.0.0.1:888", 3)
   PrintList(head)
   UpdateList(head, "127.0.0.1:6666", 77)
   PrintList(head)
   node := Front(head)
   fmt.Printf("The front: %d\n", node.Idle)
}

func TestUpdateNode2(t *testing.T) {
   head := CreateNode("", 0)
   var node *Node
   UpdateList(head, "127.0.0.1:6666", 23)
   PrintList(head)
   UpdateList(head, "127.0.0.1:6666", 55)
   PrintList(head)
   node = Front(head)
   fmt.Printf("The front: %d\n", node.Idle)
}

func TestUpdateNode3(t *testing.T) {
   head := CreateNode("", 0)
   var node *Node
   UpdateList(head, "127.0.0.1:6666", 23)
   PrintList(head)
   UpdateList(head, "127.0.0.1:5555", 55)
   PrintList(head)
   UpdateList(head, "127.0.0.1:5555", 66)
   PrintList(head)
   node = Front(head)
   fmt.Printf("The front: %d\n", node.Idle)
}

//func TestUpdateNode4(t *testing.T) {
//   head := CreateNode("", 0)
//   var node *Node
//   UpdateList(head, "xxx.0.0.1:6666", -1)
//   PrintList(head)
//   UpdateList(head, "127.0.0.1:aa", 55)
//   PrintList(head)
//   UpdateList(head, "127.0.0:5555", cc)
//   PrintList(head)
//   node = Front(head)
//   fmt.Printf("The front: %d\n", node.Idle)
//}
