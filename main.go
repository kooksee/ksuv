package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"strconv"
	"github.com/shirou/gopsutil/process"
	"github.com/fatedier/frp/utils/log"
	"os/exec"
	"bufio"
)

func printUsage(u *disk.UsageStat) {
	fmt.Println(u.Path + "\t" + strconv.FormatFloat(u.UsedPercent, 'f', 2, 64) + "% full.")
	fmt.Println("Total: " + strconv.FormatUint(u.Total / 1024 / 1024 / 1024, 10) + " GiB")
	fmt.Println("Free:  " + strconv.FormatUint(u.Free / 1024 / 1024 / 1024, 10) + " GiB")
	fmt.Println("Used:  " + strconv.FormatUint(u.Used / 1024 / 1024 / 1024, 10) + " GiB")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//runtime.GOMAXPROCS(runtime.NumCPU())

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	init_urls(r)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 100)

	t1 := time.Now().UnixNano()
	fmt.Println(t1)

	go func() {
		cmd := exec.Command("locate", "py")

		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		r := bufio.NewReader(stdout)

		p := cmd.Process
		fmt.Println(p.Pid)

		//ticker := time.NewTicker(time.Millisecond * 500)
		//defer ticker.Stop()


		for {
			line, _, _ := r.ReadLine()

			if len(line) > 0 {
				fmt.Println(string(line))
				time.Sleep(time.Microsecond)
				fmt.Println("ok")
			} else {
				fmt.Println("没有数据,睡2秒")
				time.Sleep(time.Second * 2)
			}

		}
		//for {
		//select {
		//case <-ticker.C:

		//line, _, _ := r.ReadLine()
		//fmt.Println(string(line))
		//fmt.Println("ok")

		//for i := range line {
		//	fmt.Println(string(i))
		//	fmt.Println("ok\n")
		//}

		//}
		//}

	}()

	fmt.Println(os.Getpid())

	p, err := process.NewProcess(int32(os.Getpid()))

	fmt.Println(p.Name())
	fmt.Println(p.Ppid())

	children, err := p.Children()
	for i := range children {
		fmt.Println(i)
		fmt.Println("ok\n")
	}


	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		fmt.Println("退出程序")
		log.Info("退出程序")

		//done <- true
		os.Exit(0)
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		defer ticker.Stop()
		done <- true

		for {
			select {
			case v1 := <-done:
				fmt.Println("job.....", v1)
			case <-ticker.C:
				done <- false
				fmt.Println("job.....")
			}
		}
	}()

	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	parts, err := disk.Partitions(false)
	check(err)

	var usage []*disk.UsageStat

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		check(err)
		usage = append(usage, u)
		printUsage(u)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}



