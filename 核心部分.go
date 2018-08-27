func main(){
	ln, err := net.Listen(network:"tcp", address:":9000")
	if err != nil{
		log.Println("端口监听失败," + err.Error())
	}


	log.Println(v:"running...")
	//web 服务端
	go func(){
		log.Println(v:"web server running...")
		//静态资源
		http.HandleFunc(pattern:"/", func(w http.ResponseWriter, r *http.Requset){http.ServeFile(w, r, name:"index.html")
			})
		http.Handle(pattern:"/static/", http.StripPrefix(prefix:"/static", 
			http.FileServer(http.Dir("static"))))

		//ajax请求
		http.HandleFunc(pattern:"/login", login)
		http.HandleFunc(pattern:"/logout", logout)
		http.HandleFunc(pattern:"/last", last)
		http.HandleFunc(pattern:"/history", history)
		http.HandleFunc(pattern:"/download", download)
		http.ListenAndServe(addr:":80", handler:nil)
	}()

	//处理状态上报
	for{
		conn, err := Ln.Accept()
		if err != nil {
			log.Println("accept error, " + err.Error())
		}
		go handleConn(conn)
	}
}