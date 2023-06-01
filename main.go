package main
// docker 		  run image <cmd> <params>
// go run main.go run 		<cmd> <params>

func main()  {
	switch os.Args {
	case "run":
		run()
	default:
		panic("bad command")
	}


}

func run() {
	fmt.printf("running %v\n", os.Args[2:])

	cmd := exec.command(os.Args[2], os.Args[3:]...)
	cmd.stdin	= os stdin
	cmd.stdout	= os stdout
	cmd.stderr	= os stderr
	cmd.run()

}
