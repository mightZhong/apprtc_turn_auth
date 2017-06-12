package nodechoice

import (
    "fmt"
    //"sync"
)

const fail int = -1
const success int = 0
const lost_threshold = 5
//var lock *sync.Mutex
type Node struct {
    Addr string
    Idle uint32
    Count uint
    next *Node
}

func isSameNode(newNode *Node, node *Node) bool {
    if newNode == nil || node == nil {
        fmt.Printf("Involid parameters in isSameNode\n")
        return false
    }
    return newNode.Addr == node.Addr
}

func isGreater(newNode *Node, node *Node) bool {
    if  newNode == nil || node == nil {
        fmt.Printf("Involid parameters in UpdataNode\n")
        return false
    }
    if newNode.Idle > node.Idle {
        return true
    } else {
        return false
    }
}

func CreateNode(addr string, idle uint32) *Node {
        //TODO check parameter
        var node *Node
        node = new(Node)
        //TODO
        if node == nil {
            fmt.Printf("create node failed\n")
            return nil
        }
        node.Addr, node.Idle, node.Count, node.next = addr, idle, lost_threshold, nil
        //lock = &sync.Mutex{}
        return node
}

func isNodeExist(node *Node, addr string) bool {
    if node == nil {
        fmt.Printf("Invalid parameter in isNodeExist\n")
    }

    if node.Addr == addr {
        return true
    } else {
        return false
    }
}

func getNode(head *Node, addr string, idle uint32) *Node {
    if head == nil {
        fmt.Printf("Involid parameters in getNode\n")
        return nil
    }
    //lock.Lock()
    for node:=head; node.next!=nil; node=node.next {
        if isNodeExist(node.next, addr) == true {
            takeoutNode := node.next
            takeoutNode.Idle = idle
            takeoutNode.next = nil
            takeoutNode.Count++
            node.next = node.next.next
            return takeoutNode
        }
    }
    //lock.Unlock()
    return CreateNode(addr, idle)
}

func insertNode(head *Node, newNode *Node) int {
       if head == nil || newNode == nil {
            fmt.Printf("Involid parameters in insertNode\n")
            return fail
        }
        //insert in Front
       if head.next == nil {
            fmt.Printf("Insert in Front\n")
            head.next = newNode
            return success
        }
        var node *Node
        //lock.Lock()
        for node=head; node.next!=nil; node = node.next {
            if isGreater(newNode, node.next) {
                node.next, newNode.next = newNode, node.next
                break
            }
        }
        //lock.Unlock()
        //insert in back
        if node.next == nil {
            fmt.Printf("Insert in Back\n")
            node.next = newNode
        }
        return success
}

func UpdateList(head *Node, addr string, idle uint32) int {
        //TODO parameter checking addr idle
        if head == nil {
            fmt.Printf("Involid parameters in UpdataList\n")
            return fail
        }
        fmt.Printf("--> %d -- %s\n", idle, addr);

        node := getNode(head, addr, idle)

        return insertNode(head, node)
}

func PrintList(head *Node) int {
        if head == nil {
            fmt.Printf("Invalid parameter in PrintNodes\n")
            return fail
        }
        fmt.Printf("Result: ")
        for node:=head.next; node!=nil; node = node.next {
               fmt.Printf("%d ", node.Idle)
        }
        fmt.Printf("\n\n")
        return success
}

func Front(head *Node) *Node {
    if head == nil {
        fmt.Printf("Invalid parameter in Front\n")
        return nil
    }
    return head.next
}

func CleanLostNodes(head *Node) int {
    if head == nil {
        fmt.Printf("Invalid parameter in CleanLossNode\n")
        return fail
    }
    //lock.Lock()
    for node:=head; node.next!=nil; node = node.next {
       if node.next.Count < lost_threshold {
           fmt.Printf("TURN server %s is lost count %d\n", node.next.Addr, node.next.Count)
           //takeoutNode := node.next
           node.next = node.next.next
           //TODO using GC?
           //free(takeoutNode)
       }
       if node.next != nil {
           node.next.Count = 0
       } else {
           break
       }
    }
    //lock.Unlock()
    return success
}

func GetBestServer(head *Node) string {
    server := ""
    node := Front(head)
    if node == nil {
        fmt.Printf("Sorry, there is no vaild TRUN sever\n")
        return server
    }
    fmt.Printf("choose node: %s %d", node.Addr, node.Idle)
    node.Idle = node.Idle - 3
    UpdateList(head, node.Addr, node.Idle)
    PrintList(head)
    return node.Addr
}
