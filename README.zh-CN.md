# Cloudflare扫描仪

**很多国外网站都使用Cloudflare CDN，但分配给中国大陆访问者的IP不方便（高延迟、高丢包、速度慢）。尽管 Cloudflare 已经公开披露了所有[IP范围](https://www.cloudflare.com/ips/)，在这么多的IP中找到最合适的一个是一件很累人的事。于是，开发了这个软件。**

**“自己选择最佳 IP”测试 Cloudflare CDN 的延迟和速度，获得最快的 IP（IPv4+IPv6）！如果觉得有用的话，给个star鼓励一下吧~**

* * *

## #快速开始

### 下载并运行

1.  下载预编译的可执行文件（[Github 发布](https://github.com/Ptechgithub/CloudflareScanner/releases)/[蓝奏云](https://pan.lanpw.com/b0742hkxe)）并解压它。
2.  双击`CloudflareST.exe`文件（适用于 Windows），然后等待速度测试完成...

<details>
<summary><code><strong>「 Click to view usage example on Linux system 」</strong></code></summary>

* * *

以下命令仅用于演示目的。请拜访[**发布**](https://github.com/Ptechgithub/CloudflareScanner/releases)检查版本号和文件名。

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

> 如果**平均延迟非常低**（例如 0.xx），表示 CloudflareST 是**使用代理**在速度测试期间。请先禁用代理软件再进行测试。
> 如果运行在**路由器**，建议禁用路由器内的任何代理设置（或从中排除 CloudflareST），否则速度测试结果可能会**不准确/无法使用**.

</details>

* * *

> _关于如何独立运行 CloudflareST 速度测试的简单教程**移动设备**:**[安卓](https://github.com/Ptechgithub/CloudflareScanner/discussions/61)**,**[安卓应用程序](https://github.com/xianshenglu/cloudflare-ip-tester-app)**,**[iOS系统](https://github.com/Ptechgithub/CloudflareScanner/discussions/321)**_

> 笔记！本软件仅适用于网站和**不支持选择 Cloudflare WARP 首选 IP**。详细信息请参见：[#392](https://github.com/Ptechgithub/CloudflareScanner/discussions/392)

### 结果示例

速度测试完成后，**最快的 10 个 IP**默认会显示，例如：

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

速度测试结果的第一行是**最快的 IP，具有最快的下载速度和最低的平均延迟**!

完整的结果保存在`result.csv`当前目录中的文件。打开它**记事本/电子表格软件**。格式如下：

    IP Address, Sent, Received, Loss Rate, Avg Latency, Download Speed (MB/s)
    104.27.200.69, 4, 4, 0.00, 146.23, 28.64

> _大家可以根据自己的需要对完整的结果进行进一步的过滤和处理，或者探索高级用法来指定过滤条件！_

* * *

## #高级用法

直接运行该工具使用默认参数。如果您想要更全面、定制化的速度测试结果，可以自定义参数。

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

### 接口说明

为了避免误解**测速过程中的输出内容（可用、排队号码、下载速度中途“中断”）**，我在这里解释一下。

<details>
<summary><code><strong>「 Click to expand to view content 」</strong></code></summary>

* * *

> 示例添加了常用参数，分别是：`-ttl 40 -tl 150 -sl 1 -dn 5`，最终输出结果如下：

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

> 对于那些刚接触 CloudflareST 的人来说，他们可能会感到困惑：**“原本有30个可用IP用于延迟测试，为什么现在只剩下3个了？”**下载速度测试中的队列是什么意思？我也需要排队等待下载速度测试吗？

CloudflareST 首先进行延迟测试。在此过程中，右侧进度条会显示实时可用IP数量（`Available: 30`）。但请注意，此可用数字指的是**已通过测试且未超时的 IP 数量**，无论延迟上限和下限或丢包情况如何。时延测试完成后，由于规定了时延上下限和丢包的具体条件，因此只需`10`过滤后保留IP（表示下载速度测试队列为`10`).

在上面的例子中，从`321`仅 IP`30`IP 通过了延迟测试，没有超时。然后，根据延迟上下限进行过滤后（`40 ~ 150 ms`）和丢包限制，仅`10`符合要求的 IP 保留。如果您禁用下载速度测试`-dd`，那么这些`10`IP将直接输出。但本例中并没有禁用下载速度测试，因此软件会继续对这些进行下载速度测试。`10`IP（`Queue: 10`).

> 由于下载速度测试是单线程的，并且是逐个依次测试IP，因此等待下载速度测试的IP数量称为`Queue`.

* * *

> 您可能已经注意到，**即使您指定查找 5 个满足下载速度条件的 IP，为什么该过程“终止”只有 3 个？**

在下载速度测试进度条中，`3 / 5`表明`3`满足下载速度下限条件（即下载速度超过`1 MB/s`）已被发现，同时`5`表示您请求查找`5`满足此条件的 IP（`-dn 5`).

> 此外，请注意，如果您指定`-dn`大于下载速度测试队列，例如，如果只有`4`延迟测试后保留 IP，然后是后面的数字`/`在下载速度测试进度条中将会`4`，就像下载速度测试队列一样，而不是`5`你指定了`-dn`.

测试完这些的下载速度后`10`仅 IP`3`发现IP下载速度超过`1 MB/s`，而剩余的`7`IP 不符合标准。

因此，并不是这样的`“the test terminates every time without reaching 5”`，而是对所有 IP 进行下载速度测试，但仅`3`符合标准。

* * *

如果您不想遇到测试完所有IP后只有少数IP符合条件的情况，您可以**降低下载速度上限参数`-sl`**或将其删除。

因为只要`-sl`指定参数后，测试将继续，直到`-dn`数量（默认10个IP）已达到，或测试所有IP。去除`-sl`并添加`-dn 20`参数只会测试延迟最低的前 20 个 IP 的延迟，然后停止，以节省时间。

* * *

此外，如果队列中的所有 IP 都已进行下载速度测试，但没有一个符合下载速度标准，则**直接输出队列中所有IP的下载速度测试结果**。这样，您就可以看到这些 IP 的下载速度并了解它们的表现。然后，**尝试降低`-sl`适当地再试一次**.

同样，关于延迟测试，这两个值`Available: 30`和`Queue: 10`还可以帮助您确定您设置的延迟条件是否过于严格。如果你有很多可用的IP，但是经过过滤条件后，只剩下2个或3个，那么很明显你需要**降低预期的延迟/丢包情况**.

这两种机制，其中一种通知您**延迟和丢包情况**，另一个关于**下载速度条件**，帮助您确定您的设置是否合适。

</details>

* * *

### 使用示例

在Windows上，要指定参数，您需要在CMD中运行它们，或者将参数添加到快捷方式的目标中。

> **笔记**：所有参数都有**默认值**，并且使用具有默认值的参数可以省略（**根据需要选择**)，参数可以按任意顺序指定。  
> **提示**：对于 Windows**电源外壳**，简单地改变`CloudflareST.exe`在下面的命令中`.\CloudflareST.exe`.  
> **提示**：对于Linux系统，只需更改`CloudflareST.exe`在下面的命令中`./CloudflareST`.

* * *

#### 在 CMD 中使用参数运行 CloudflareST

对于那些不熟悉命令行程序的人来说，你可能不知道如何带参数运行它们，所以我将简要解释一下。

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

许多人在运行 CloudflareST 时遇到错误**绝对路径**在命令中。这是因为默认`-f ip.txt`参数使用的是相对路径，所以需要指定绝对路径`ip.txt`。但是，这可能很麻烦，因此建议使用以下命令在其程序目录中运行 CloudflareST**相对路径**:

**方法一**:

1.  打开CloudflareST所在目录。
2.  右键单击空白处并选择**\[在这里打开命令窗口]**从上下文菜单中按<kbd>Shift + 右键单击</kbd>，这将使用当前目录打开 CMD。
3.  输入带有参数的命令，例如：`CloudflareST.exe -tll 50 -tl 200`运行程序。

**方法2**:

1.  打开CloudflareST所在目录。
2.  直接输入`cmd`在文件夹的地址栏中输入并按 Enter 键，以当前目录打开 CMD。
3.  输入带有参数的命令，例如：`CloudflareST.exe -tll 50 -tl 200`运行程序。

> 当然，你也可以打开任意CMD窗口，然后输入如下命令`cd /d "D:\Program Files\CloudflareST"`进入程序目录。

> **提示**: 如果你正在使用**电源外壳**，简单地改变`CloudflareST.exe`在命令中`.\CloudflareST.exe`.

</details>

* * *

#### 通过 Windows 快捷方式使用参数运行 CloudflareST

对于不经常修改运行参数（比如平时双击运行）的人来说，使用快捷方式更方便。

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

右键单击`CloudflareST.exe`文件 -**\[创建快捷方式]**，然后右键单击创建的快捷方式 -**\[特性]**，并修改其**目标**:

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

测试 IPv6 时，您可能会注意到每次测试的数量都不同。原因解释如下：[#120](https://github.com/Ptechgithub/CloudflareScanner/issues/120)  
因为 IPv6 地址太多（数十亿），而且绝大多数 IP 范围都没有使用，所以我只扫描了一部分可用的 IPv6 范围并将其写入`ipv6.txt`文件。如果您有兴趣，可以自行扫描并修改。 ASN数据来源来自：[BGP.和.net](https://bgp.he.net/AS13335#_prefixes6)

* * *

#### #HTTP处理

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

目前有两种延迟测试模式，**TCP协议和HTTP协议**.  
TCP协议速度更快，消耗的资源更少，超时时间为1秒，这是默认模式。  
HTTP协议适合快速测试某个域名是否指向某个IP以及是否可以访问。超时设置为 2 秒。  
对于同一个IP，各个协议得到的时延一般遵循这样的顺序：**ICMP &lt; TCP &lt; HTTP**，HTTP 对网络波动（例如丢包）更加敏感。

> 注意：HTTPing 本质上是一种形式**网络扫描**行为。因此，如果您在服务器上运行它，您需要**减少并发**(`-n`），否则，您可能会被一些严格的提供商暂停。如果您遇到 HTTPing 过程中可用 IP 数量减少，甚至变成 0，但过一段时间又恢复的情况，可能是由于**临时限制机制**由 ISP 或 Cloudflare CDN 触发，将您的活动视为网络扫描。在这种情况下，通常会在一段时间后恢复。建议**减少并发**(`-n`）以减少此类情况的发生。

> 此外，该软件的 HTTPing 仅检索**响应头**并且不检索正文内容（意味着URL文件大小不会影响HTTPing测试，但如果您还想进行下载速度测试，则需要一个大文件）。它类似于curl -i 功能。

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

#### #匹配指定地区(colo 机场三字码)

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

> 两个 CDN 机场三字码通用，因此各地区名可见：[HTTPS://呜呜呜.cloud flare status.com/](https://www.cloudflarestatus.com/)

</details>

* * *

#### #文件相对/绝对路径

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

#### #测速其他端口

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

#### #自定义测速地址

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

#### #自定义测速条件（指定 延迟/丢包/下载速度 的目标范围）

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

> 注意：延迟测速进度条右边的**可用数量**，仅指延迟测速过程中**未超时的 IP 数量**，和延迟上下限条件无关。

-   仅指定**[平均延迟上限]**条件

```bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s
# 即找到平均延迟低于 200 ms 的 IP，然后再按延迟从低到高进行 10 次下载测速
CloudflareST.exe -tl 200
```

> 如果**没有找到一个满足延迟**条件的 IP，那么不会输出任何内容。

* * *

-   仅指定**[平均延迟上限]**条件，且**只延迟测速，不下载测速**

```bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s，数量：不知道多少 个
# 即只输出低于 200ms 的 IP，且不再下载测速（因为不再下载测速，所以 -dn 参数就无效了）
CloudflareST.exe -tl 200 -dd
```

-   仅指定**[丢包几率上限]**条件

```bash
# 丢包几率上限：0.25
# 即找到丢包率低于等于 0.25 的 IP，范围 0.00~1.00，如果 -tlr 0 则代表过滤掉任何丢包的 IP
CloudflareST.exe -tlr 0.25
```

* * *

-   仅指定**[下载速度下限]**条件

```bash
# 平均延迟上限：9999 ms，下载速度下限：5 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 9999 ms 且下载速度高于 5 MB/s 的 IP 才会停止测速
CloudflareST.exe -sl 5 -dn 10
```

> 如果**没有找到一个满足速度**条件的 IP，那么会**忽略条件输出所有 IP 测速结果**（方便你下次测速时调整条件）。

> 没有指定平均延迟上限时，如果一直**凑不够**满足条件的 IP 数量，就会**一直测速**下去。  
> 所以建议**同时指定[下载速度下限]+[平均延迟上限]**，这样测速到指定延迟上限还没凑够数量，就会终止测速。

* * *

-   同时指定**[平均延迟上限]+[下载速度下限]**条件

```bash
# 平均延迟上限、下载速度下限均支持小数（如 -sl 0.5）
# 平均延迟上限：200 ms，下载速度下限：5.6 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 200 ms 且下载速度高于 5 .6MB/s 的 IP 才会停止测速
CloudflareST.exe -tl 200 -sl 5.6 -dn 10
```

> 如果**没有找到一个满足延迟**条件的 IP，那么不会输出任何内容。  
> 如果**没有找到一个满足速度**条件的 IP，那么会忽略条件输出所有 IP 测速结果（方便你下次测速时调整条件）。  
> 所以建议先不指定条件测速一遍，看看平均延迟和下载速度大概在什么范围，避免指定条件**过低/过高**！

> 因为 Cloudflare 公开的 IP 段是**回源 IP+任播 IP**，而**回源 IP**是无法使用的，所以下载测速是 0.00。  
> 运行时可以加上`-sl 0.01`（下载速度下限），过滤掉**回源 IP**（下载测速低于 0.01MB/s 的结果）。

</details>

* * *

#### #单独对一个或多个 IP 测速

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

**方式 一**：
直接通过参数指定要测速的 IP 段数据。

```bash
# 先进入 CloudflareST 所在目录，然后运行：
# Windows 系统（在 CMD 中运行）
CloudflareST.exe -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32

# Linux 系统
./CloudflareST -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
```

* * *

**方式 二**：
或者把这些 IP 按如下格式写入到任意文本文件中，例如：`1.txt`

    1.1.1.1
    1.1.1.200
    1.0.0.1/24
    2606:4700::/32

> 单个 IP 的话可以省略`/32`子网掩码了（即`1.1.1.1`等同于`1.1.1.1/32`）。  
> 子网掩码`/24`指的是这个 IP 最后一段，即`1.0.0.1~1.0.0.255`。

然后运行 CloudflareST 时加上启动参数`-f 1.txt`来指定 IP 段数据文件。

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

#### #一劳永逸加速所有使用 Cloudflare CDN 的网站（不需要再一个个添加域名到 Hosts 了）

我以前说过，开发该软件项目的目的就是为了通过**改 Hosts 的方式来加速访问使用 Cloudflare CDN 的网站**。

但就如[**#8**](https://github.com/Ptechgithub/CloudflareScanner/issues/8)所说，一个个添加域名到 Hosts 实在**太麻烦**了，于是我就找到了个**一劳永逸**的办法！可以看这个[**还在一个个添加 Hosts？完美本地加速所有使用 Cloudflare CDN 的网站方法来了！**](https://github.com/Ptechgithub/CloudflareScanner/discussions/71)和另一个[依靠本地 DNS 服务来修改域名解析 IP 为自选 IP](https://github.com/Ptechgithub/CloudflareScanner/discussions/317)的教程。

* * *

## 手动编译

<details>
<summary><code><strong>「 Click to expand to view content 」</strong></code></summary>

* * *

为了方便，我是在编译的时候将版本号写入代码中的 version 变量，因此你手动编译时，需要像下面这样在`go build`命令后面加上`-ldflags`参数来指定版本号：

```bash
go build -ldflags "-s -w -X main.version=v2.3.3"
# 在 CloudflareScanner 目录中通过命令行（例如 CMD、Bat 脚本）运行该命令，即可编译一个可在和当前设备同样系统、位数、架构的环境下运行的二进制程序（Go 会自动检测你的系统位数、架构）且版本号为 v2.3.3
```

如果想要在 Windows 64位系统下编译**其他系统、架构、位数**，那么需要指定**GOOS**和**GOARCH**变量。

例如在 Windows 系统下编译一个适用于**Linux 系统 amd 架构 64 位**的二进制程序：

```bat
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w -X main.version=v2.3.3"
```

例如在 Linux 系统下编译一个适用于**Windows 系统 amd 架构 32 位**的二进制程序：

```bash
GOOS=windows
GOARCH=386
go build -ldflags "-s -w -X main.version=v2.3.3"
```

> 可以运行`go tool dist list`来查看当前 Go 版本支持编译哪些组合。

* * *

当然，为了方便批量编译，我会专门指定一个变量为版本号，后续编译直接调用该版本号变量即可。  
同时，批量编译的话，还需要分开放到不同文件夹才行（或者文件名不同），需要加上`-o`参数指定。

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

## 执照

GPL-3.0 许可证。
