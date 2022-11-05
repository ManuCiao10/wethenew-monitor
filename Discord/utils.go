package discord

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ManuCiao10/wethenew-monitor/data"
)

var (
	mu sync.Mutex
)

func GetProxy() string {
	mu.Lock()
	file, err := os.Open("data/proxies2.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	_ = file.Close()
	if len(txtlines) == 0 {
		panic("Please add proxies to proxies.txt")
	}
	index := rand.Intn(len(txtlines))
	mu.Unlock()
	proxy := strings.Split(txtlines[index], ":")
	proxy_url := "http://" + proxy[2] + ":" + proxy[3] + "@" + proxy[0] + ":" + proxy[1]
	return proxy_url
}

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("|%v|\n", time.Since(start))
	}
}

func MonitorPid(pid int) {
	process, err := os.FindProcess(int(pid))
	if err != nil {
		fmt.Printf("Failed to find process: %s\n", err)
	} else {
		err := process.Signal(syscall.Signal(0))
		fmt.Printf("process.Signal on pid %d returned: %v\n", pid, err)
	}
}

func SaveSlice(class data.Info) []int {
	var slice []int

	for _, v := range class.Results {
		slice = append(slice, v.ID)
		
	}
	return slice
}

func SaveSliceTest(class data.Info) []int {
	var slice []int

	for _, v := range class.Results {
		if v.ID != 275 {
			slice = append(slice, v.ID)
		}
	}
	return slice
}

func Contains(s []int, id int) bool {
	for _, v := range s {
		if v == id {
			return true
		}
	}
	return false
}
