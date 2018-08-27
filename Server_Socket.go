package main  
import (  
    "fmt"  
    "net"  
    "log"  
    "os"  
)  

func main() {  
    //建立socket,监听端口
    //establish socket server 
    netListen, err := net.Listen("tcp", "9000")  
    CheckError(err)  
    defer netListen.Close()  

    Log("Waiting for clients")  
    for {  
        conn, err := netListen.Accept()  
        if err != nil {  
            continue  
        }  
        Log(conn.RemoteAddr().String(), " tcp connect success")  
        handleConnection(conn)  
    }  
}  

//处理连接  
//deal with the connection
func handleConnection(conn net.Conn) {  
    buffer := make([]byte, 2048)  
    for { 
        n, err := conn.Read(buffer)  
        if err != nil {
            Log(conn.RemoteAddr().String(), " connection error: ", err)  
            return  
        }  
        Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))  
    }
}

func CheckError(err error) {  
    if err != nil {  
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())  
        os.Exit(1)  
    }  
}   