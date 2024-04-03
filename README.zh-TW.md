# Cloudflare掃描儀

**許多國外網站都使用Cloudflare CDN，但分配給中國大陸訪客的IP不方便（高延遲、高丟包、速度慢）。儘管 Cloudflare 已經公開披露了所有[IP範圍](https://www.cloudflare.com/ips/)，在這麼多的IP中找到最合適的一個是一件很累人的事。於是，開發了這個軟體。**

**「自己選擇最佳 IP」測試 Cloudflare CDN 的延遲和速度，獲得最快的 IP（IPv4+IPv6）！如果覺得有用的話，給個star鼓勵一下吧~**

* * *

## #快速開始

### 下載並運行

1.  下載預編譯的可執行檔（[Github 發布](https://github.com/Ptechgithub/CloudflareScanner/releases)/[藍奏雲](https://pan.lanpw.com/b0742hkxe)）並解壓它。
2.  按兩下`CloudflareST.exe`文件（適用於 Windows），然後等待速度測試完成...

<details>
<summary><code><strong>「 Click to view usage example on Linux system 」</strong></code></summary>

* * *

以下命令僅用於演示目的。請拜訪[**發布**](https://github.com/Ptechgithub/CloudflareScanner/releases)檢查版本號碼和檔案名稱。

```yaml
# If this is your first time using, it's recommended to create a new folder (skip this step for future updates)
mkdir CloudflareST

# Enter the folder (for future updates, just repeat the following download and extraction commands from here)
cd CloudflareST

# Download the CloudflareST compressed file (replace [version number] and [file name] in the URL according to your needs)
wget -N https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# If you're downloading in a Chinese network environment, consider using one of these mirrors for acceleration:
# wget -N https://download.scholar.rr.nu/XIU2/CloudflareScanner/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# wget -N https://ghproxy.cc/https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# wget -N https://ghproxy.net/https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# wget -N https://gh-proxy.com/https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# wget -N https://mirror.ghproxy.com/https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# If the download fails, try removing the -N parameter (if it's for an update, remember to delete the old compressed file beforehand: rm CloudflareST_linux_amd64.tar.gz)

# Extract (no need to delete old files, they will be overwritten directly; replace the file name according to your needs)
tar -zxf CloudflareST_linux_amd64.tar.gz

# Grant execute permission
chmod +x CloudflareST

# Run (without parameters)
./CloudflareST

# Run (example with parameters)
./CloudflareST -dd -tll 90
```

> 如果**平均延遲非常低**（例如 0.xx），表示 CloudflareST 是**使用代理**在速度測試期間。請先停用代理軟體再進行測試。
> 如果運行在**路由器**，建議停用路由器內的任何代理設定（或從中排除 CloudflareST），否則速度測試結果可能會**不準確/無法使用**.

</details>

* * *

> _關於如何獨立運行 CloudflareST 速度測試的簡單教程**行動裝置**:**[安卓](https://github.com/Ptechgithub/CloudflareScanner/discussions/61)**,**[安卓應用程式](https://github.com/xianshenglu/cloudflare-ip-tester-app)**,**[iOS系統](https://github.com/Ptechgithub/CloudflareScanner/discussions/321)**_

> 筆記！本軟體僅適用於網站和**不支援選擇 Cloudflare WARP 首選 IP**。詳細資訊請參閱：[#392](https://github.com/Ptechgithub/CloudflareScanner/discussions/392)

### 結果範例

速度測試完成後，**最快的 10 個 IP**預設會顯示，例如：

```bash
IP Address      Sent    Received    Loss-Rate    Avg-Latency    Download Speed (MB/s)
104.27.200.69     4       4       0.00    146.23    28.64
172.67.60.78      4       4       0.00    139.82    15.02
104.25.140.153    4       4       0.00    146.49    14.90
104.27.192.65     4       4       0.00    140.28    14.07
172.67.62.214     4       4       0.00    139.29    12.71
104.27.207.5      4       4       0.00    145.92    11.95
172.67.54.193     4       4       0.00    146.71    11.55
104.22.66.8       4       4       0.00    147.42    11.11
104.27.197.63     4       4       0.00    131.29    10.26
172.67.58.91      4       4       0.00    140.19    9.14
...

# If the average latency is very low (e.g., 0.xx), it indicates that CloudflareST is using a proxy during the speed test. Please disable the proxy software before testing again.
# If running on a router, it's recommended to disable any proxy settings within the router (or exclude CloudflareST from them), otherwise the speed test results may be inaccurate/unusable.

# Because each speed test randomly selects an IP within each IP range, the results of each test cannot be the same, which is normal!

# Note! I found that the latency is significantly higher the first time I run the test after booting up my computer (the same with manual TCPing). Subsequent tests are normal.
# Therefore, it is recommended that before running the first official speed test after booting up, randomly test a few IPs (you don't need to wait for the latency test to complete, just close it as soon as the progress bar moves).

# The entire process of the software under the default parameters is roughly as follows:
# 1. Latency test (default TCPing mode, HTTPing mode requires manual addition of parameters)
# 2. Latency sorting (sort by latency from low to high and filter by conditions; different loss rates will be sorted separately, so there may be some IPs with low latency but high loss rate sorted to the end)
# 3. Download speed test (start downloading speed test from the IP with the lowest latency in sequence, and stop after testing 10 IPs by default)
# 4. Speed sorting (sort by speed from high to low)
# 5. Output results (control whether to output to the command line (-p 0) or to a file (-o ""))

# Note: The result file result.csv will display Chinese garbled characters when opened with Microsoft Excel, which is normal. It displays normally in other spreadsheet software/notepad.
```

速度測試結果的第一行是**最快的 IP，具有最快的下載速度和最低的平均延遲**!

完整的結果保存在`result.csv`目前目錄中的檔案。打開它**記事本/試算軟體**。格式如下：

    IP Address, Sent, Received, Loss Rate, Avg Latency, Download Speed (MB/s)
    104.27.200.69, 4, 4, 0.00, 146.23, 28.64

> _大家可以依照自己的需求對完整的結果進行進一步的過濾和處理，或是探索進階用法來指定篩選條件！_

* * *

## #進階用法

直接運行該工具使用預設參數。如果您想要更全面、客製化的速度測試結果，可以自訂參數。

```css
C:\>CloudflareST.exe -h

CloudflareScanner vX.X.X
Test the latency and speed of all IPs of Cloudflare CDN, and obtain the fastest IP (IPv4+IPv6)!
https://github.com/Ptechgithub/CloudflareScanner

Options:
    -n 200
        Number of latency test threads; more threads result in faster testing, but devices with weak performance (e.g., routers) should not set it too high; (default is 200, maximum is 1000)
    -t 4
        Number of latency test iterations; number of latency tests per single IP; (default is 4 times)
    -dn 10
        Number of download speed tests; after latency testing and sorting, the number of IPs to be speed tested from the lowest latency; (default is 10)
    -dt 10
        Download speed test duration; maximum time for download speed test per single IP, cannot be too short; (default is 10 seconds)
    -tp 443
        Specify the test port; port used for latency and download speed tests; (default is port 443)
    -url https://cf.xiu2.xyz/url
        Specify the test URL; address used for latency (HTTPing) and download speed tests, default address is not guaranteed to be available, self-hosting is recommended;

    -httping
        Switch test mode; change latency test mode to HTTP protocol, and set the test address to the one specified by [-url] parameter; (default is TCPing)
        Note: HTTPing is essentially a form of network scanning behavior, so if you're running it on a server, you need to reduce concurrency (-n), otherwise some strict vendors may suspend service.
        If you encounter a situation where the number of available IPs for HTTPing decreases over time or even becomes 0, but then recovers after a while, it may also be because the operator or Cloudflare CDN thinks you are performing network scanning and triggers a temporary restriction mechanism, so it will recover after a while. It is recommended to reduce concurrency (-n) to reduce the occurrence of this situation.
    -httping-code 200
        Effective status code; the valid HTTP status code returned by the web page during HTTPing latency test, only one; (default is 200 301 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        Match specified regions; region names are local airport three-letter codes, separated by commas, lowercase, support Cloudflare, AWS CloudFront, only available in HTTPing mode; (default is all regions)

    -tl 200
        Upper limit of average latency; only output IPs with average latency lower than the specified value, upper and lower limit conditions can be used together; (default is 9999 ms)
    -tll 40
        Lower limit of average latency; only output IPs with average latency higher than the specified value; (default is 0 ms)
    -tlr 0.2
        Upper limit of loss rate; only output IPs with loss rate lower than/equal to the specified value, range is 0.00~1.00, 0 filters out any IP with loss; (default is 1.00)
    -sl 5
        Lower limit of download speed; only output IPs with download speed higher than the specified value, testing will stop after reaching the specified number [-dn]; (default is 0.00 MB/s)

    -p 10
        Number of results to display; display the specified number of results directly after testing, 0 will exit without displaying results; (default is 10)
    -f ip.txt
        IP range data file; if the path contains spaces, please enclose it in quotes; support other CDN IP ranges; (default is ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        Specify IP range data; specify the IP range data to be tested directly through parameters, separated by commas; (default is empty)
    -o result.csv
        Write result to file; if the path contains spaces, please enclose it in quotes; empty value will not write to file [-o ""]; (default is result.csv)

    -dd
        Disable download speed test; after disabling, test results will be sorted by latency (default is sorted by download speed); (default is enabled)
    -allip
        Test all IPs; test each IP in IP ranges (IPv4 only); (default is randomly test one IP per /24 range)

    -v
        Print program version + check for updates
    -h
        Print help information
```

### 介面說明

為了避免誤解**測速過程中的輸出內容（可用、排隊號碼、下載速度中途「中斷」）**，我在這裡解釋一下。

<details>
<summary><code><strong>「 Click to expand to view content 」</strong></code></summary>

* * *

> 範例添加了常用參數，分別是：`-ttl 40 -tl 150 -sl 1 -dn 5`，最終輸出結果如下：

```bash
# XIU2/CloudflareScanner vX.X.X

Delay testing started (mode: TCP, port: 443, range: 40 ~ 150 ms, loss: 1.00)
321 / 321 [----------------------------------------------------------------------------------] Available: 30
Download speed testing started (lower limit: 1.00 MB/s, quantity: 5, queue: 10)
3 / 5 [---------------------------------------------------------↗---------------------------]
IP Address        Sent    Received    Loss Rate    Avg Latency    Download Speed (MB/s)
XXX.XXX.XXX.XXX   4       4           0.00         83.32          3.66
XXX.XXX.XXX.XXX   4       4           0.00         107.81         2.49
XXX.XXX.XXX.XXX   4       3           0.25         149.59         1.04

Complete test results have been written to the result.csv file, which can be viewed using Notepad/spreadsheet software.
Press Enter or Ctrl+C to exit.
```

* * *

> 對於那些剛接觸 CloudflareST 的人來說，他們可能會感到困惑：**“原本有30個可用IP用於延遲測試，為什麼現在只剩下3個了？”**下載速度測試中的佇列是什麼意思？我也需要排隊等待下載速度測試嗎？

CloudflareST 首先進行延遲測試。在此過程中，右側進度條會顯示即時可用IP數量（`Available: 30`）。但請注意，此可用數字指的是**已通過測試且未逾時的 IP 數量**，無論延遲上限和下限或丟包情況如何。時延測試完成後，由於規定了時延上下限和丟包的具體條件，因此只需`10`過濾後保留IP（表示下載速度測試隊列為`10`).

在上面的例子中，從`321`僅 IP`30`IP 通過了延遲測試，沒有逾時。然後，根據延遲上下限進行過濾後（`40 ~ 150 ms`）和丟包限制，僅`10`符合要求的 IP 保留。如果您停用下載速度測試`-dd`，那麼這些`10`IP將直接輸出。但本例中並沒有停用下載速度測試，因此軟體會繼續對這些進行下載速度測試。`10`IP（`Queue: 10`).

> 由於下載速度測試是單線程的，並且是逐個依次測試IP，因此等待下載速度測試的IP數量稱為`Queue`.

* * *

> 您可能已經註意到，**即使您指定查找 5 個符合下載速度條件的 IP，為什麼該過程「終止」只有 3 個？**

在下載速度測試進度條中，`3 / 5`表明`3`滿足下載速度下限條件（即下載速度超過`1 MB/s`）已被發現，同時`5`表示您請求尋找`5`滿足此條件的 IP（`-dn 5`).

> 此外，請注意，如果您指定`-dn`大於下載速度測試佇列，例如，如果只有`4`延遲測試後保留 IP，然後是後面的數字`/`在下載速度測試進度條中將會`4`，就像下載速度測試隊列一樣，而不是`5`你指定了`-dn`.

測試完這些的下載速度後`10`僅 IP`3`發現IP下載速度超過`1 MB/s`，而剩餘的`7`IP 不符合標準。

因此，並不是這樣的`“the test terminates every time without reaching 5”`，而是對所有 IP 進行下載速度測試，但僅`3` met the criteria.

* * *

如果您不想遇到測試完所有IP後只有少數IP符合條件的情況，您可以**降低下載速度上限參數`-sl`**或將其刪除。

因為只要`-sl`指定參數後，測試將繼續，直到`-dn`數量（預設10個IP）已達到，或測試所有IP。去除`-sl`並添加`-dn 20`參數只會測試延遲最低的前 20 個 IP 的延遲，然後停止，以節省時間。

* * *

此外，如果佇列中的所有 IP 都已進行下載速度測試，但沒有一個符合下載速度標準，則**直接輸出佇列中所有IP的下載速度測試結果**。這樣，您就可以看到這些 IP 的下載速度並了解它們的表現。然後，**嘗試降低`-sl`適當地再試一次**.

同樣，關於延遲測試，這兩個值`Available: 30`和`Queue: 10`還可以幫助您確定您設定的延遲條件是否過於嚴格。如果你有很多可用的IP，但經過過濾條件後，只剩下2個或3個，那麼很明顯你需要**降低預期的延遲/丟包情況**.

這兩種機制，其中一種通知您**延遲和丟包情況**，另一個關於**下載速度條件**，幫助您確定您的設定是否合適。

</details>

* * *

### 使用範例

在Windows上，要指定參數，您需要在CMD中執行它們，或將參數新增至捷徑的目標。

> **筆記**：所有參數都有**預設值**，並且使用具有預設值的參數可以省略（**根據需要選擇**)，參數可以任意順序指定。  
> **提示**：對於 Windows**電源外殼**，簡單地改變`CloudflareST.exe`在下面的命令中`.\CloudflareST.exe`.  
> **提示**：對於Linux系統，只需更改`CloudflareST.exe`在下面的命令中`./CloudflareST`.

* * *

#### 在 CMD 中使用參數來運行 CloudflareST

對於那些不熟悉命令列程式的人來說，你可能不知道如何帶參數運行它們，所以我將簡要解釋一下。

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

許多人在運行 CloudflareST 時遇到錯誤**絕對路徑**在命令中。這是因為預設`-f ip.txt`參數使用的是相對路徑，所以需要指定絕對路徑`ip.txt`。但是，這可能很麻煩，因此建議使用以下命令在其程式目錄中執行 CloudflareST**相對路徑**:

**方法一**:

1.  開啟CloudflareST所在目錄。
2.  右鍵單擊空白處並選擇**\[在這裡打開命令視窗]**從上下文選單中按<kbd>Shift + 右鍵單擊</kbd>，這將使用目前目錄開啟 CMD。
3.  輸入帶有參數的命令，例如：`CloudflareST.exe -tll 50 -tl 200`運行程式。

**方法二**:

1.  開啟CloudflareST所在目錄。
2.  直接輸入`cmd`在資料夾的網址列中輸入並按 Enter 鍵，以目前目錄開啟 CMD。
3.  輸入帶有參數的命令，例如：`CloudflareST.exe -tll 50 -tl 200`運行程式。

> 當然，你也可以開啟任意CMD窗口，然後輸入以下指令`cd /d "D:\Program Files\CloudflareST"`進入程序目錄。

> **提示**: 如果你正在使用**電源外殼**，簡單地改變`CloudflareST.exe`在命令中`.\CloudflareST.exe`.

</details>

* * *

#### 透過 Windows 快捷方式使用參數來執行 CloudflareST

對於不經常修改運行參數（例如平常雙擊運行）的人來說，使用快捷方式更方便。

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

右鍵單擊`CloudflareST.exe`文件 -**\[建立快捷方式]**，然後右鍵單擊創建的快捷方式 -**\[特性]**，並修改其**目標**:

```bash
# If you don't want to output a result file, please add -o " ", where the quotes contain a space (omitting the space will cause this parameter to be omitted).
D:\ABC\CloudflareST\CloudflareST.exe -n 500 -t 4 -dn 20 -dt 5 -o " "

# If the file path contains quotes, then the startup parameters need to be placed outside the quotes. Remember to have a space between the quotes and the -.
"D:\Program Files\CloudflareST\CloudflareST.exe" -n 500 -t 4 -dn 20 -dt 5 -o " "

# Note! The starting location of the shortcut cannot be empty, otherwise, it will not find the ip.txt file due to the absolute path.
```

</details>

* * *

#### #IPv4/IPv6

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

```bash
# Specify the built-in IPv4 data file to test these IPv4 addresses (the -f parameter defaults to ip.txt, so it can be omitted)
CloudflareST.exe -f ip.txt

# Specify the built-in IPv6 data file to test these IPv6 addresses
# Additionally, starting from version 2.1.0, CloudflareST supports testing IPv4+IPv6 addresses together and removes the -ipv6 parameter. Therefore, one file can contain both IPv4+IPv6 addresses.
CloudflareST.exe -f ipv6.txt

# You can also directly specify the IPs to be tested through parameters
CloudflareST.exe -ip 1.1.1.1,2606:4700::/32
```

測試 IPv6 時，您可能會注意到每次測試的數量都不同。原因解釋如下：[#120](https://github.com/Ptechgithub/CloudflareScanner/issues/120)  
因為 IPv6 位址太多（數十億），絕大多數 IP 範圍都沒有使用，所以我只掃描了一部分可用的 IPv6 範圍並將它們寫入`ipv6.txt`文件。如果您有興趣，可以自行掃描並修改。 ASN資料來源來自：[bgp.he.net](https://bgp.he.net/AS13335#_prefixes6)

* * *

#### #HTTP處理

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

目前有兩種延遲測試模式，**TCP協定和HTTP協定**.  
TCP協定速度更快，消耗的資源更少，超時時間為1秒，這是預設模式。  
HTTP協定適合快速測試某個網域是否指向某個IP以及是否可以存取。超時設定為 2 秒。  
對於同一個IP，各個協定所得到的時延一般都遵循這樣的順序：**ICMP &lt; TCP &lt; HTTP**，HTTP 對網路波動（例如丟包）更為敏感。

> 注意：HTTPing 本質上是一種形式**網路掃描**行為。因此，如果您在伺服器上運行它，您需要**減少並行**(`-n`），否則，您可能會被一些嚴格的提供者暫停。如果您遇到 HTTPing 過程中可用 IP 數量減少，甚至變成 0，但過一段時間又恢復的情況，可能是由於**臨時限制機制**由 ISP 或 Cloudflare CDN 觸發，將您的活動視為網路掃描。在這種情況下，通常會在一段時間後恢復。建議**減少並行**(`-n`）以減少此類情況的發生。

> 此外，該軟體的 HTTPing 僅檢索**回應頭**並且不檢索正文內容（意味著URL檔案大小不會影響HTTPing測試，但如果您還想進行下載速度測試，則需要一個大檔案）。它類似於curl -i 功能。

```bash
# Simply add the -httping parameter to switch to HTTP protocol latency testing mode
CloudflareST.exe -httping

# The software determines availability based on the effective HTTP status code returned when accessing the webpage (of course, timeouts are also counted). By default, the software considers 200, 301, and 302 HTTP status codes as valid. You can manually specify the HTTP status code you consider valid, but only one can be specified (you need to confirm in advance which status code the test address will return under normal circumstances).
CloudflareST.exe -httping -httping-code 200

# Use the -url parameter to specify the HTTPing test address (it can be any webpage URL, not limited to specific file addresses)
CloudflareST.exe -httping -url https://cf.xiu2.xyz/url
# If you want to HTTPing test other websites/CDNs, specify an address that uses that website/CDN (because the software defaults to Cloudflare's address, it can only be used to test Cloudflare's IPs)

# Note: If the test address is using the HTTP protocol, remember to add -tp 80 (this parameter will affect the port used during latency testing/download speed testing)
# Similarly, if you want to test port 80, you also need to add the -url parameter to specify an address with the http:// protocol (and this address should not be forcibly redirected to HTTPS). If it is a port other than 80 or 443, you need to make sure that the download speed test address supports access through that port.
CloudflareST.exe -httping -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

* * *

#### #符合指定地區(colo 機場三字碼)

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 该功能支持 Cloudflare CDN、AWS CloudFront CDN，且这两个 CDN 的机场三字码是通用的
# 注意：如果你要用于筛选 AWS CloudFront CDN 地区，那么要通过 -url 参数指定一个使用该 CDN 的地址（因为软件默认地址是 Cloudflare 的）

# 指定地区名后，延迟测速后得到的结果就都是指定地区的 IP 了（也可以继续进行下载测速）
# 节点地区名为当地 机场三字码，指定多个时用英文逗号分隔，v2.2.3 版本后支持小写

CloudflareST.exe -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD

# 注意，该参数只有在 HTTPing 延迟测速模式下才可用（因为要访问网页来获得）
```

> 兩個 CDN 機場三字碼通用，因此各地區名可見：<https://www.cloudflarestatus.com/>

</details>

* * *

#### #文件相對/絕對路徑

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 指定 IPv4 数据文件，不显示结果直接退出，输出结果到文件（-p 值为 0）
CloudflareST.exe -f 1.txt -p 0 -dd

# 指定 IPv4 数据文件，不输出结果到文件，直接显示结果（-p 值为 10 条，-o 值为空但引号不能少）
CloudflareST.exe -f 2.txt -o "" -p 10 -dd

# 指定 IPv4 数据文件 及 输出结果到文件（相对路径，即当前目录下，如含空格请加上引号）
CloudflareST.exe -f 3.txt -o result.txt -dd


# 指定 IPv4 数据文件 及 输出结果到文件（相对路径，即当前目录内的 abc 文件夹下，如含空格请加上引号）
# Linux（CloudflareST 程序所在目录内的 abc 文件夹下）
./CloudflareST -f abc/3.txt -o abc/result.txt -dd

# Windows（注意是反斜杠）
CloudflareST.exe -f abc\3.txt -o abc\result.txt -dd


# 指定 IPv4 数据文件 及 输出结果到文件（绝对路径，即 C:\abc\ 目录下，如含空格请加上引号）
# Linux（/abc/ 目录下）
./CloudflareST -f /abc/4.txt -o /abc/result.csv -dd

# Windows（注意是反斜杠）
CloudflareST.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd


# 如果要以【绝对路径】运行 CloudflareST，那么 -f / -o 参数中的文件名也必须是【绝对路径】，否则会报错找不到文件！
# Linux（/abc/ 目录下）
/abc/CloudflareST -f /abc/4.txt -o /abc/result.csv -dd

# Windows（注意是反斜杠）
C:\abc\CloudflareST.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd
```

</details>

* * *

#### #測速其他端口

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 如果你想要测速非默认 443 的其他端口，则需要通过 -tp 参数指定（该参数会影响 延迟测速/下载测速 时使用的端口）

# 如果要延迟测速 80 端口+下载测速（如果 -dd 禁用了下载测速则不需要），那么还需要指定 http:// 协议的下载测速地址才行（且该地址不会强制重定向至 HTTPS，因为那样就变成 443 端口了）
CloudflareST.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm

# 如果是非 80 443 的其他端口，那么需要确定你使用的下载测速地址是否支持通过该非标端口访问。
```

</details>

* * *

#### #自訂測速地址

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 该参数适用于下载测速 及 HTTP 协议的延迟测速，对于后者该地址可以是任意网页 URL（不局限于具体文件地址）

# 地址要求：可以直接下载、文件大小超过 200MB、用的是 Cloudflare CDN
CloudflareST.exe -url https://cf.xiu2.xyz/url

# 注意：如果测速地址为 HTTP 协议（该地址不能强制重定向至 HTTPS），记得加上 -tp 80（这个参数会影响 延迟测速/下载测速 时使用的端口），如果是非 80 443 端口，那么需要确定下载测速地址是否支持通过该端口访问。
CloudflareST.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

* * *

#### #自訂測速條件（指定 延遲/丟包/下載速度 的目標範圍）

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

> 注意：延遲測速進度條右邊的**可用數量**，僅指延遲測速過程中**未超時的 IP 數量**，和延遲上下限條件無關。

-   僅指定**[平均延遲上限]**條件

```bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s
# 即找到平均延迟低于 200 ms 的 IP，然后再按延迟从低到高进行 10 次下载测速
CloudflareST.exe -tl 200
```

> 如果**沒有找到一個滿足延遲**條件的 IP，那麼不會輸出任何內容。

* * *

-   僅指定**[平均延遲上限]**條件，且**只延遲測速，不下載測速**

```bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s，数量：不知道多少 个
# 即只输出低于 200ms 的 IP，且不再下载测速（因为不再下载测速，所以 -dn 参数就无效了）
CloudflareST.exe -tl 200 -dd
```

-   僅指定**[丟包幾率上限]**條件

```bash
# 丢包几率上限：0.25
# 即找到丢包率低于等于 0.25 的 IP，范围 0.00~1.00，如果 -tlr 0 则代表过滤掉任何丢包的 IP
CloudflareST.exe -tlr 0.25
```

* * *

-   僅指定**[下載速度下限]**條件

```bash
# 平均延迟上限：9999 ms，下载速度下限：5 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 9999 ms 且下载速度高于 5 MB/s 的 IP 才会停止测速
CloudflareST.exe -sl 5 -dn 10
```

> 如果**沒有找到一個滿足速度**條件的 IP，那麼會**忽略條件輸出所有 IP 測速結果**（方便你下次測速時調整條件）。

> 沒有指定平均延遲上限時，如果一直**湊不夠**滿足條件的 IP 數量，就會**一直測速**下去。  
> 所以建議**同時指定[下載速度下限]+[平均延遲上限]**，這樣測速到指定延遲上限還沒湊夠數量，就會終止測速。

* * *

-   同時指定**[平均延遲上限]+[下載速度下限]**條件

```bash
# 平均延迟上限、下载速度下限均支持小数（如 -sl 0.5）
# 平均延迟上限：200 ms，下载速度下限：5.6 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 200 ms 且下载速度高于 5 .6MB/s 的 IP 才会停止测速
CloudflareST.exe -tl 200 -sl 5.6 -dn 10
```

> 如果**沒有找到一個滿足延遲**條件的 IP，那麼不會輸出任何內容。  
> 如果**沒有找到一個滿足速度**條件的 IP，那麼就會忽略條件輸出所有 IP 測速結果（方便你下次測速時調整條件）。  
> 所以建議先不指定條件測速一遍，看看平均延遲和下載速度大概在什麼範圍，避免指定條件**過低/過高**！

> 因為 Cloudflare 公開的 IP 段是**回源 IP+任播 IP**，而**回源 IP**是無法使用的，所以下載測速是 0.00。  
> 運行時可以加上`-sl 0.01`（下載速度下限），過濾掉**回源 IP**（下載測速低於 0.01MB/s 的結果）。

</details>

* * *

#### #單獨對一個或多個 IP 測速

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

**方式 一**：
直接透過參數指定要測速的 IP 段資料。

```bash
# 先进入 CloudflareST 所在目录，然后运行：
# Windows 系统（在 CMD 中运行）
CloudflareST.exe -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32

# Linux 系统
./CloudflareST -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
```

* * *

**方式 二**：
或把這些 IP 依下列格式寫入到任意文字檔中，例如：`1.txt`

    1.1.1.1
    1.1.1.200
    1.0.0.1/24
    2606:4700::/32

> 單一 IP 的話可以省略`/32`子網路遮罩了（即`1.1.1.1`等同於`1.1.1.1/32`）。  
> 子網路遮罩`/24`指的是這個 IP 最後一段，即`1.0.0.1~1.0.0.255`。

然後執行 CloudflareST 時加上啟動參數`-f 1.txt`來指定 IP 段資料檔。

```bash
# 先进入 CloudflareST 所在目录，然后运行：
# Windows 系统（在 CMD 中运行）
CloudflareST.exe -f 1.txt

# Linux 系统
./CloudflareST -f 1.txt

# 对于 1.0.0.1/24 这样的 IP 段只会随机最后一段（1.0.0.1~255），如果要测速该 IP 段中的所有 IP，请加上 -allip 参数。
```

</details>

* * *

#### #一勞永逸加速所有使用 Cloudflare CDN 的網站（不需要再一個個新增網域到 Hosts 了）

我以前說過，開發該軟體專案的目的就是為了透過**改變 Hosts 的方式來加速訪問使用 Cloudflare CDN 的網站**。

但就如[**#8**](https://github.com/Ptechgithub/CloudflareScanner/issues/8)所說，一個個添加網域到 Hosts 實在**太麻煩**了，於是我就找到了個**一勞永逸**的辦法！可以看這個[**還在一個個加入 Hosts？完美本地加速所有使用 Cloudflare CDN 的網站方法來了！**](https://github.com/Ptechgithub/CloudflareScanner/discussions/71)和另一個[依賴本機 DNS 服務來修改網域解析 IP 為自選 IP](https://github.com/Ptechgithub/CloudflareScanner/discussions/317)的教程。

* * *

## 手動編譯

<details>
<summary><code><strong>「 Click to expand to view content 」</strong></code></summary>

* * *

為了方便，我是在編譯的時候將版本號寫入程式碼中的 version 變量，因此你手動編譯時，需要像下面這樣在`go build`命令後面加上`-ldflags`參數來指定版本號：

```bash
go build -ldflags "-s -w -X main.version=v2.3.3"
# 在 CloudflareScanner 目录中通过命令行（例如 CMD、Bat 脚本）运行该命令，即可编译一个可在和当前设备同样系统、位数、架构的环境下运行的二进制程序（Go 会自动检测你的系统位数、架构）且版本号为 v2.3.3
```

如果想要在 Windows 64位元系統下編譯**其他系統、架構、位數**，那麼需要指定**GOOS**和**GOARCH**變數。

例如在 Windows 系統下編譯一個適用於**Linux 系統 amd 架構 64 位元**的二進位程式：

```bat
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w -X main.version=v2.3.3"
```

例如在 Linux 系統下編譯一個適用於**Windows 系統 amd 架構 32 位**的二進位程式：

```bash
GOOS=windows
GOARCH=386
go build -ldflags "-s -w -X main.version=v2.3.3"
```

> 可以運行`go tool dist list`來查看目前 Go 版本支援編譯哪些組合。

* * *

當然，為了方便批次編譯，我會專門指定一個變數為版本號，後續編譯直接呼叫該版本號變數即可。  
同時，批次編譯的話，還需要分開到不同資料夾才行（或檔案名稱不同），需要加上`-o`參數指定。

```bat
:: Windows 系统下是这样：
SET version=v2.3.3
SET GOOS=linux
SET GOARCH=amd64
go build -o Releases\CloudflareST_linux_amd64\CloudflareST -ldflags "-s -w -X main.version=%version%"
```

```bash
# Linux 系统下是这样：
version=v2.3.3
GOOS=windows
GOARCH=386
go build -o Releases/CloudflareST_windows_386/CloudflareST.exe -ldflags "-s -w -X main.version=${version}"
```

</details>

* * *

## 執照

GPL-3.0 授權。
