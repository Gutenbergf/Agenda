package main

import (
    "fmt"
    "log"
    "net/rpc"
    "os/exec"
    "os"
    "io/ioutil"
    "time"
)


type Contact struct{
    Name string
    Tel, Id,Pos int

} 

var( // Global variables
    idContact int
    agenda []Contact  // Receives response from RPC call  
    contact Contact   // Receives response from RPC call

)


func main() {
    idContact = 1   
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    } 
    for{
        menu(client,err)                 
    }
}



func insert(client *rpc.Client, err error){
  var contact Contact
    fmt.Printf("Enter the name: ")
    fmt.Scan(&contact.Name)    
    fmt.Printf("Enter the number: ")
    fmt.Scan(&contact.Tel)     
    if(verifyExistingNumber(contact.Tel)==true){
        fmt.Println("Number already registered!")
        showContact(contact)  
                        
    }else{
        contact.Id=idContact                     
        err = client.Call("Contacts.Add",contact,&agenda)
        if err != nil {
            throwError(err)
        }
        fmt.Println("Contact added!")
        idContact=idContact+1   
    }   


}

func insertPos(pos int)Contact{
    var contact Contact
    fmt.Printf("Enter the name:  ")
    fmt.Scan(&contact.Name)    
    fmt.Printf("Enter the number:  ")
    fmt.Scan(&contact.Tel)  
    contact.Pos = pos - 1
    return contact 


}

func search(client *rpc.Client, err error){
    if(idContact>1){
        var option int
        menuSearch()
        fmt.Scan(&option)
        switch {
        case option==1:
            searchName(client,err)

        case option==2:
            searchNumber(client,err)

        case option>2 || option<1:
            fmt.Println("Option invalid!")   

        } 
    }else{
        fmt.Println("No contacts registered!")
    }
   

   
    

}

func searchName(client *rpc.Client, err error){
     
        var name string
        fmt.Printf("Enter the name: ")
        fmt.Scan(&name)
        err = client.Call("Contacts.Search",name,&contact)
        if err != nil {
            throwError(err)                
        }else{
            if(contact.Tel!=0){
                fmt.Println("The searched contact number is: ",contact.Tel)                                    
            }
            if(contact.Tel==0){
                fmt.Println("Contact not registered!")                
            }          

        }
    
    
}

func searchNumber(client *rpc.Client, err error){     
        var tel int
        fmt.Printf("Enter the number: ")
        fmt.Scan(&tel)
        err = client.Call("Contacts.SearchNumber",tel,&contact)
        if err != nil {
               throwError(err)
        }else{
            if(contact.Tel!=0){
                fmt.Println("The contact name you searched for is: ",contact.Name)                                  
            }
            if(contact.Tel==0){
                fmt.Println("Contact not registered!")                
            }                       

        }
    
    
}

func remove(client *rpc.Client, err error) {
    if(idContact>1){
        show(client,err)
        var pos int
        fmt.Printf("Which contact do you want to delete?(ID) ")
        fmt.Scan(&pos) 
        if((pos-1)>=idContact-1 || (pos-1)<0){
            fmt.Println("Contact invalid!")
        }else{                     
                err = client.Call("Contacts.Remove",pos-1,&agenda)
                checkError(err)
                fmt.Println("Contact Removed!")                
                idContact = idContact -1               
                
        }
       
    }else{
        fmt.Println("No contacts registered!")
    }
        
    
}


func altera(client *rpc.Client, err error){
    if(idContact>1){
        show(client,err)
        var pos int
        fmt.Printf("Which contact do you want to alter?(ID)")
        fmt.Scan(&pos) 
        if((pos-1)>=idContact-1 || (pos-1)<0){
            fmt.Println("Contact invalid!")

        }else{
            c:=insertPos(pos) 
            if(verifyExistingNumber(c.Tel) ==true){
                fmt.Println("Number already registered!")
            }else{
                err = client.Call("Contacts.Alter",c,&agenda)
                checkError(err)
                fmt.Println("Contact changed!")
            }
            
        }
    }else{
        fmt.Println("No contacts registered!")
    }
       
    
}

func verifyExisting(name string)bool{
    for i := 0; i < len(agenda); i++{
       if(agenda[i].Name==name){
           return true
       }      
        
    }
    return false

}

func verifyExistingNumber(tel int)bool{
    for i := 0; i < len(agenda); i++{
       if(agenda[i].Tel==tel){
           return true
       }      
        
    }
    return false

}

func show(client *rpc.Client, err error){
    if(idContact==1){
        fmt.Println("No contacts registered!")
    }else{
        err = client.Call("Contacts.Show",1,&agenda) // GET ON CURRENT LIST (KEEP CONSISTENCY)
        checkError(err)

        fmt.Println("ID     |   NAME    |   NUMBER  ")
        for i := 0; i <idContact-1; i++{        
            fmt.Println(agenda[i].Id,"       ",agenda[i].Name,"      ",agenda[i].Tel)             
            
        }
    }
    
    
}

func showContact(c Contact){
    fmt.Println("NAME    |   NUMBER  ")  
    fmt.Println(c.Name,"      ",c.Tel) 
    

}

//FUNCTIONS TO SHOW MENU'S

func menuSearch(){
    dat, err := ioutil.ReadFile("menuSearch.txt")
    check(err)
    fmt.Println(string(dat))
}

func headerAdd(){
    dat, err := ioutil.ReadFile("headerAdd.txt")
    check(err)
    fmt.Println(string(dat))
}

func headerRemove(){
    dat, err := ioutil.ReadFile("headerRemove.txt")
    check(err)
    fmt.Println(string(dat))
}

func headerSearch(){
    dat, err := ioutil.ReadFile("headerSearch.txt")
    check(err)
    fmt.Println(string(dat))
}

func headerAgenda(){
    dat, err := ioutil.ReadFile("headerAgenda.txt")
    check(err)
    fmt.Println(string(dat))
}

func headerAlter(){
    dat, err := ioutil.ReadFile("headerAlter.txt")
    check(err)
    fmt.Println(string(dat))
}

func menu(client *rpc.Client, err error){
    var opcao int
    dat, err := ioutil.ReadFile("menu.txt")
    check(err)
    fmt.Println(string(dat))
    fmt.Scan(&opcao)
        switch {
        case opcao == 1:
            clear()
            headerAdd()
            insert(client,err)             
            time.Sleep(1 * time.Second)            
            clear()

        case opcao==2:
            headerAgenda()
            show(client,err)
           

        case opcao==3:
            clear()
            headerSearch()
            search(client,err)
            time.Sleep(1 * time.Second)  
            clear()        

        case opcao==4:
            clear()
            headerAlter()
            altera(client,err)
            time.Sleep(1 * time.Second)
            clear()
                       
        case opcao == 5:
            clear()
            headerRemove()
            remove(client,err) 
            time.Sleep(1 * time.Second)   
            clear()       

        case opcao == 6:
            os.Exit(1)  
        

       }
    


}

//FUNCOES OF SYSTEM

func pause(){ // Function to pause system
    cmd := exec.Command("cmd", "/c", "pause") 
    cmd.Stdout = os.Stdout
    cmd.Run() 
 }

 func clear(){ // Function to clear system
    cmd := exec.Command("cmd", "/c", "cls") 
    cmd.Stdout = os.Stdout
    cmd.Run()
 }

//FUNCTION'S TO CHECK ERROR

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkError(err error){
    if err != nil {        
        log.Fatal("Contact error:", err)
    }
}

func throwError(err error){    
    log.Fatal("Contact error:", err)
}
