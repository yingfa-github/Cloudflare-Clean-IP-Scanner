package main

import (
        "encoding/json"
        "flag"
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
        "runtime"
        "time"

        "github.com/Ptechgithub/CloudflareScanner/task"
        "github.com/Ptechgithub/CloudflareScanner/utils"
)

var (
        version, versionNew string
)

func init() {
        var printVersion bool
        var help = `
CloudflareSpeedTest ` + version + `
Test the latency and speed of all IP addresses of Cloudflare CDN, and get the fastest IP (IPv4+IPv6)!
https://github.com/Ptechgithub/CloudflareScanner

Parameters:
    -n 200
        Latency test threads; the more threads, the faster the latency test. Devices with weak performance (such as routers) should not be too high; (default 200, maximum 1000)
    -t 4
        Latency test times; the number of times a single IP is tested for latency; (default 4 times)
    -dn 10
        Download test quantity; after latency test and sorting, the quantity of download test from the lowest latency; (default 10)
    -dt 10
        Download test time; the maximum time for downloading test of a single IP, cannot be too short; (default 10 seconds)
    -tp 443
        Specify test port; port used for latency test/download test; (default port 443)
    -url https://speed.cloudflare.com/__down?bytes=52428800
        Specify test address; address used for latency test (HTTPing)/download test, default address does not guarantee availability, self-built is recommended;

    -httping
        Switch test mode; switch latency test mode to HTTP protocol, test address used is [-url] parameter; (default TCPing)
    -httping-code 200
        Valid status code; valid HTTP status code returned when HTTPing latency test is performed, only one; (default 200 301 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        Match specified region; the region name is the local airport three-letter code, separated by English commas, only available in HTTPing mode; (default all regions)

    -tl 200
        Upper limit of average latency; only output IP addresses with average latency below the specified upper limit, each upper and lower limit condition can be used in combination; (default 9999 ms)
    -tll 40
        Lower limit of average latency; only output IP addresses with average latency above the specified lower limit; (default 0 ms)
    -tlr 0.2
        Upper limit of packet loss rate; only output IP addresses with packet loss rate lower than or equal to the specified limit, range 0.00~1.00, 0 filters out any packet loss IP addresses; (default 1.00)
    -sl 5
        Lower limit of download speed; only output IP addresses with download speed higher than the specified limit, stop testing only when the specified quantity [-dn] is reached; (default 0.00 MB/s)

    -p 10
        Display result quantity; directly display the specified quantity of results after testing, when 0, do not display results and exit directly; (default 10)
    -f ip.txt
        IP segment data file; if the path contains spaces, please add quotes; support other CDN IP segments; (default ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        Specify IP segment data; specify the IP segment data to be tested directly through parameters, separated by English commas; (default empty)
    -o result.csv
        Write result file; if the path contains spaces, please add quotes; when empty, do not write to file [-o ""]; (default result.csv)

    -dd
        Disable download test; after disabling, test results will be sorted by latency (default sorted by download speed); (default enabled)
    -allip
        Test all IPs; test each IP in the IP segment (IPv4 only); (default test one IP randomly in each /24 segment)

    -v
        Print program version + check for updates
    -h
        Print help message
`
        var minDelay, maxDelay, downloadTime int
        var maxLossRate float64
        flag.IntVar(&task.Routines, "n", 200, "Latency test threads")
        flag.IntVar(&task.PingTimes, "t", 4, "Latency test times")
        flag.IntVar(&task.TestCount, "dn", 10, "Download test quantity")
        flag.IntVar(&downloadTime, "dt", 10, "Download test time")
        flag.IntVar(&task.TCPPort, "tp", 443, "Specify test port")
        flag.StringVar(&task.URL, "url", "https://speed.cloudflare.com/__down?bytes=52428800", "Specify test address")

        flag.BoolVar(&task.Httping, "httping", false, "Switch test mode")
        flag.IntVar(&task.HttpingStatusCode, "httping-code", 0, "Valid status code")
        flag.StringVar(&task.HttpingCFColo, "cfcolo", "", "Match specified region")

        flag.IntVar(&maxDelay, "tl", 9999, "Upper limit of average latency")
        flag.IntVar(&minDelay, "tll", 0, "Lower limit of average latency")
        flag.Float64Var(&maxLossRate, "tlr", 1, "Upper limit of packet loss rate")
        flag.Float64Var(&task.MinSpeed, "sl", 0, "Lower limit of download speed")

        flag.IntVar(&utils.PrintNum, "p", 10, "Display result quantity")
        flag.StringVar(&task.IPFile, "f", "ip.txt", "IP segment data file")
        flag.StringVar(&task.IPText, "ip", "", "Specify IP segment data")
        flag.StringVar(&utils.Output, "o", "result.csv", "Output result file")

        flag.BoolVar(&task.Disable, "dd", false, "Disable download test")
        flag.BoolVar(&task.TestAll, "allip", false, "Test all IP")

        flag.BoolVar(&printVersion, "v", false, "Print program version")
        flag.Usage = func() { fmt.Print(help) }
        flag.Parse()

        if task.MinSpeed > 0 && time.Duration(maxDelay)*time.Millisecond == utils.InputMaxDelay {
                fmt.Println("[Tip] When using [-sl] parameter, it is recommended to use [-tl] parameter to avoid continuous testing due to insufficient [-dn] quantity...")
        }
        utils.InputMaxDelay = time.Duration(maxDelay) * time.Millisecond
        utils.InputMinDelay = time.Duration(minDelay) * time.Millisecond
        utils.InputMaxLossRate = float32(maxLossRate)
        task.Timeout = time.Duration(downloadTime) * time.Second
        task.HttpingCFColomap = task.MapColoMap()

        if printVersion {
                println(version)
                fmt.Println("Checking for updates...")
                checkUpdate()
                if versionNew != "" {
                        fmt.Printf("*** Found new version [%s]! Please go to [https://github.com/Ptechgithub/CloudflareScanner] to update! ***", versionNew)
                } else {
                        fmt.Println("Current version is the latest [" + version + "]!")
                }
                os.Exit(0)
        }
}

func main() {
        task.InitRandSeed() // Set random seed

        fmt.Printf("# XIU2/CloudflareSpeedTest %s \n\n", version)

        // Start latency testing + filter delay/loss
        pingData := task.NewPing().Run().FilterDelay().FilterLossRate()
        // Start download speed testing
        speedData := task.TestDownloadSpeed(pingData)
        utils.ExportCsv(speedData) // Export to file
        speedData.Print()          // Print results

        if versionNew != "" {
                fmt.Printf("\n*** Found new version [%s]! Please go to [https://github.com/Ptechgithub/CloudflareScanner] to update! ***\n", versionNew)
        }
        endPrint()
}

func endPrint() {
        if utils.NoPrintResult() {
                return
        }
        if runtime.GOOS == "windows" { // If Windows, need to press Enter or Ctrl+C to exit (avoids closing after completion when run by double-clicking)
                fmt.Printf("Press Enter or Ctrl+C to exit.")
                fmt.Scanln()
        }
}

// Check for updates
func checkUpdate() {
        timeout := 10 * time.Second
        client := http.Client{Timeout: timeout}
        res, err := client.Get("https://api.github.com/repos/Ptechgithub/CloudflareScanner/releases/latest")
        if err != nil {
                return
        }
        defer res.Body.Close()

        var release struct {
                TagName string `json:"tag_name"`
        }
        err = json.NewDecoder(res.Body).Decode(&release)
        if err != nil {
                return
        }

        if release.TagName != version {
                versionNew = release.TagName
        }
}