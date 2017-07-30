package main

import (
    
    "fmt"
    "net"
    "os"
    "net/rpc"
    
)


type Contact struct{
    Name string
    Tel, Id,Pos int
} 


type Contacts Contact

var( // Global variables
    myAgenda []Contact
    id  int
)


func(t *Contacts) Add(c Contact,agenda *[]Contact) error{   
    myAgenda[id-1] = c 
    *agenda = myAgenda  
    id=id+1
    return nil

}


func(t *Contacts) Search(name string,contact *Contact) error{   
    for i:=0;i<len(myAgenda);i++{
        if(name==myAgenda[i].Name){
            *contact = myAgenda[i]
        }
    }
    return nil

}

func(t *Contacts) SearchNumber(tel int,contact *Contact) error{   
    for i:=0;i<len(myAgenda);i++{
        if(tel==myAgenda[i].Tel){
            *contact = myAgenda[i]
        }
    }
    return nil

}

func(t *Contacts) Remove(position int,agenda *[]Contact) error{  
    for i:=position+1;i<len(myAgenda);i++{ // Decrement ID
        myAgenda[i].Id = myAgenda[i].Id -1                 
        
    }  
    myAgenda = append(myAgenda[:position], myAgenda[position+1:]...) // Remove    
    *agenda = myAgenda
     id=id-1

        
    return nil

}

func(t *Contacts) Alter(c Contact,agenda *[]Contact) error{ 
    for i:=0;i<len(myAgenda);i++{
        if(i==c.Pos){
            myAgenda[i] = c   
        }
    }     
      
    *agenda = myAgenda        
    return nil

}


func(t *Contacts) Show(i int,agenda *[]Contact) error{       
    *agenda = myAgenda        
    return nil

}



func main() {   
    id =1  
    myAgenda = make([]Contact, 10)      
    contact := new(Contacts)
    rpc.Register(contact)    
    rpc.HandleHTTP()   
   

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        rpc.ServeConn(conn)
    }

}


func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}