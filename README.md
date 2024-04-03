# Ptechgithub/CloudflareScanner

[![Go Version](https://img.shields.io/github/go-mod/go-version/Ptechgithub/CloudflareScanner.svg?style=flat-square&label=Go&color=00ADD8&logo=go)](https://github.com/Ptechgithub/CloudflareScanner/)[![Release Version](https://img.shields.io/github/v/release/Ptechgithub/CloudflareScanner.svg?style=flat-square&label=Release&color=00ADD8&logo=github)](https://github.com/Ptechgithub/CloudflareScanner/releases/latest)[![GitHub license](https://img.shields.io/github/license/Ptechgithub/CloudflareScanner.svg?style=flat-square&label=License&color=00ADD8&logo=github)](https://github.com/Ptechgithub/CloudflareScanner/)[![GitHub Star](https://img.shields.io/github/stars/Ptechgithub/CloudflareScanner.svg?style=flat-square&label=Star&color=00ADD8&logo=github)](https://github.com/Ptechgithub/CloudflareScanner/)[![GitHub Fork](https://img.shields.io/github/forks/Ptechgithub/CloudflareScanner.svg?style=flat-square&label=Fork&color=00ADD8&logo=github)](https://github.com/Ptechgithub/CloudflareScanner/)

Many foreign websites are using Cloudflare CDN, but the IP assigned to visitors from mainland China is not friendly (high latency, high packet loss, and slow speed).  
While Cloudflare exposes all[IP segment](https://www.cloudflare.com/zh-cn/ips/), but trying to find one that suits you among so many IPs would be exhausting, so I came up with this software.

**"Self-selected preferred IP" tests Cloudflare CDN latency and speed to obtain the fastest IP (IPv4+IPv6)**! If it's useful**Click`⭐`Give me some encouragement~**

* * *

## #Quick to use

### Download and run

1.  Download the compiled executable file ([Github Releases](https://github.com/Ptechgithub/CloudflareScanner/releases)/[Lan Zuoyun](https://pan.lanpw.com/b0742hkxe)) and unzip it.
2.  Double click to run`CloudflareScanner.exe`File (Windows system), waiting for the speed test to complete...

<details>
<summary><code><strong>「 点击查看 Linux 系统下的使用示例 」</strong></code></summary>

* * *

The following commands are only examples. For the version number and file name, please go to[**Releases**](https://github.com/Ptechgithub/CloudflareScanner/releases)Check.

```yaml
# 如果是第一次使用，则建议创建新文件夹（后续更新时，跳过该步骤）
mkdir CloudflareScanner

# 进入文件夹（后续更新，只需要从这里重复下面的下载、解压命令即可）
cd CloudflareScanner

# 下载 CloudflareScanner 压缩包（自行根据需求替换 URL 中 [版本号] 和 [文件名]）
wget -N https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareScanner_linux_amd64.zip
# 如果你是在国内网络环境中下载，那么请使用下面这几个镜像加速之一：
# wget -N https://download.scholar.rr.nu/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareScanner_linux_amd64.zip
# wget -N https://ghproxy.cc/https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareScanner_linux_amd64.zip
# wget -N https://ghproxy.net/https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareScanner_linux_amd64.zip
# wget -N https://gh-proxy.com/https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareScanner_linux_amd64.zip
# wget -N https://mirror.ghproxy.com/https://github.com/Ptechgithub/CloudflareScanner/releases/download/v2.2.5/CloudflareScanner_linux_amd64.zip
# 如果下载失败的话，尝试删除 -N 参数（如果是为了更新，则记得提前删除旧压缩包 rm CloudflareScanner_linux_amd64.zip ）

# 解压（不需要删除旧文件，会直接覆盖，自行根据需求替换 文件名）
tar -zxf CloudflareScanner_linux_amd64.zip

# 赋予执行权限
chmod +x CloudflareScanner

# 运行（不带参数）
./CloudflareScanner

# 运行（带参数示例）
./CloudflareScanner -dd -tll 90
```

> If flat**Very low latency**(such as 0.xx), then CloudflareScanner**The agent left during the speed test**, please close the proxy software before testing the speed.  
> if in**router**It is recommended to turn off the proxy in the router first (or exclude it), otherwise the speed test results may be**Inaccurate/unusable**。

</details>

* * *

> _exist**cell phone**A simple tutorial to run the CloudflareScanner speed test independently on:**[Android](https://github.com/Ptechgithub/CloudflareScanner/discussions/61)、[Android APP](https://github.com/xianshenglu/cloudflare-ip-tester-app)、[IOS](https://github.com/Ptechgithub/CloudflareScanner/discussions/321)**_

> Notice! This software only works on websites,**Preferred IP for Cloudflare WARP is not supported**, see for details:[#392](https://github.com/Ptechgithub/CloudflareScanner/discussions/392)

### Example of results

After the speed test is completed, it will be displayed by default**Fastest 10 IPs**, example:

```bash
IP 地址           已发送  已接收  丢包率  平均延迟  下载速度 (MB/s)
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

# 如果平均延迟非常低（如 0.xx），则说明 CloudflareScanner 测速时走了代理，请先关闭代理软件后再测速。
# 如果在路由器上运行，请先关闭路由器内的代理（或将其排除），否则测速结果可能会不准确/无法使用。

# 因为每次测速都是在每个 IP 段中随机 IP，所以每次的测速结果都不可能相同，这是正常的！

# 注意！我发现电脑开机后第一次测速延迟会明显偏高（手动 TCPing 也一样），后续测速都正常
# 因此建议大家开机后第一次正式测速前，先随便测几个 IP（无需等待延迟测速完成，只要进度条动了就可以直接关了）

# 软件在 默认参数 下的整个流程大概步骤：
# 1. 延迟测速（默认 TCPing 模式，HTTPing 模式需要手动加上参数）
# 2. 延迟排序（延迟 从低到高 排序并按条件过滤，不同丢包率会分开排序，因此可能会有一些延迟低但丢包的 IP 排到后面）
# 3. 下载测速（从延迟最低的 IP 开始依次下载测速，默认测够 10 个就会停止）
# 4. 速度排序（速度从高到低排序）
# 5. 输出结果（通过参数控制是否输出到命令行(-p 0)或输出到文件(-o "")）

# 注意：输出的结果文件 result.csv 通过微软 Excel 表格打开会中文乱码，这是正常的，其他表格软件/记事本都显示正常
```

The first line of the speed test result is**The fastest IP with the fastest download speed and lowest average latency**！

The complete results are saved in the current directory`result.csv`file, use**Notepad/sheet software**Open, the format is as follows:

    IP 地址, 已发送, 已接收, 丢包率, 平均延迟, 下载速度 (MB/s)
    104.27.200.69,4,4,0.00,146.23,28.64

> _You can view the complete results according to your own needs**Further filtering**, or take a look at the advanced usage**Specify filter criteria**！_

* * *

## #Advanced use

The default parameters are used for direct operation. If you want the speed measurement results to be more comprehensive and more in line with your own requirements, you can customize the parameters.

```css
C:\>CloudflareScanner.exe -h

CloudflareScanner vX.X.X
测试 Cloudflare CDN 所有 IP 的延迟和速度，获取最快 IP (IPv4+IPv6)！
https://github.com/Ptechgithub/CloudflareScanner

参数：
    -n 200
        延迟测速线程；越多延迟测速越快，性能弱的设备 (如路由器) 请勿太高；(默认 200 最多 1000)
    -t 4
        延迟测速次数；单个 IP 延迟测速的次数；(默认 4 次)
    -dn 10
        下载测速数量；延迟测速并排序后，从最低延迟起下载测速的数量；(默认 10 个)
    -dt 10
        下载测速时间；单个 IP 下载测速最长时间，不能太短；(默认 10 秒)
    -tp 443
        指定测速端口；延迟测速/下载测速时使用的端口；(默认 443 端口)
    -url https://cf.xiu2.xyz/url
        指定测速地址；延迟测速(HTTPing)/下载测速时使用的地址，默认地址不保证可用性，建议自建；

    -httping
        切换测速模式；延迟测速模式改为 HTTP 协议，所用测试地址为 [-url] 参数；(默认 TCPing)
        注意：HTTPing 本质上也算一种 网络扫描 行为，因此如果你在服务器上面运行，需要降低并发(-n)，否则可能会被一些严格的商家暂停服务。
        如果你遇到 HTTPing 首次测速可用 IP 数量正常，后续测速越来越少甚至直接为 0，但停一段时间后又恢复了的情况，那么也可能是被 运营商、Cloudflare CDN 认为你在网络扫描而 触发临时限制机制，因此才会过一会儿就恢复了，建议降低并发(-n)减少这种情况的发生。
    -httping-code 200
        有效状态代码；HTTPing 延迟测速时网页返回的有效 HTTP 状态码，仅限一个；(默认 200 301 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        匹配指定地区；地区名为当地机场三字码，英文逗号分隔，支持小写，支持 Cloudflare、AWS CloudFront，仅 HTTPing 模式可用；(默认 所有地区)

    -tl 200
        平均延迟上限；只输出低于指定平均延迟的 IP，各上下限条件可搭配使用；(默认 9999 ms)
    -tll 40
        平均延迟下限；只输出高于指定平均延迟的 IP；(默认 0 ms)
    -tlr 0.2
        丢包几率上限；只输出低于/等于指定丢包率的 IP，范围 0.00~1.00，0 过滤掉任何丢包的 IP；(默认 1.00)
    -sl 5
        下载速度下限；只输出高于指定下载速度的 IP，凑够指定数量 [-dn] 才会停止测速；(默认 0.00 MB/s)

    -p 10
        显示结果数量；测速后直接显示指定数量的结果，为 0 时不显示结果直接退出；(默认 10 个)
    -f ip.txt
        IP段数据文件；如路径含有空格请加上引号；支持其他 CDN IP段；(默认 ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        指定IP段数据；直接通过参数指定要测速的 IP 段数据，英文逗号分隔；(默认 空)
    -o result.csv
        写入结果文件；如路径含有空格请加上引号；值为空时不写入文件 [-o ""]；(默认 result.csv)

    -dd
        禁用下载测速；禁用后测速结果会按延迟排序 (默认按下载速度排序)；(默认 启用)
    -allip
        测速全部的IP；对 IP 段中的每个 IP (仅支持 IPv4) 进行测速；(默认 每个 /24 段随机测速一个 IP)

    -v
        打印程序版本 + 检查版本更新
    -h
        打印帮助说明
```

### Interface explanation

In order to prevent everyone from being concerned about the speed measurement process,**The output content is misunderstood (available, queue and other numbers, the download speed test is "interrupted" halfway?)**, I explain specifically.

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

> This example adds all common parameters, which is:`-ttl 40 -tl 150 -sl 1 -dn 5`, the final output result is as follows:

```bash
# Ptechgithub/CloudflareScanner vX.X.X

开始延迟测速（模式：TCP, 端口：443, 范围：40 ~ 150 ms, 丢包：1.00)
321 / 321 [----------------------------------------------------------------------------------] 可用: 30
开始下载测速（下限：1.00 MB/s, 数量：5, 队列：10）
3 / 5 [---------------------------------------------------------↗---------------------------]
IP 地址           已发送  已接收  丢包率  平均延迟  下载速度 (MB/s)
XXX.XXX.XXX.XXX   4       4      0.00    83.32    3.66
XXX.XXX.XXX.XXX   4       4      0.00    107.81   2.49
XXX.XXX.XXX.XXX   4       3      0.25    149.59   1.04

完整测速结果已写入 result.csv 文件，可使用记事本/表格软件查看。
按下 回车键 或 Ctrl+C 退出。
```

* * *

> People who are new to CloudflareScanner may be confused**There are obviously 30 IPs available for delay speed testing, so why are there only 3 left in the end?**  
> What does the queue in the download speed test mean? Do I still have to queue up when I download the speed test?

CloudflareScanner will first delay the speed test, during which the number of available IPs will be displayed in real time on the right side of the progress bar (`可用: 30`), but note that the available quantity refers to**Number of IPs that passed the test without timeout**, has nothing to do with the upper and lower limits of delay and packet loss conditions. When the delayed speed test is completed, because it is also specified**Delay upper and lower limits, packet loss**conditions, so after filtering according to the conditions, only`10`Already (that is, waiting for the download speed test)`队列：10`）。

That is, in the above example,`321`After the IP delay speed test is completed, only`30`Each IP test passes without timeout, and then according to the upper and lower limit range of delay:`40 ~ 150 ms`and packet loss upper limit conditions, only`10`There is an IP that meets the requirements. if you`-dd`If the download speed test is disabled, this will be output directly.`10`IP. Of course, this example is not disabled, so the software will continue to`10`IP for download speed test (`队列：10`）。

> Because the download speed test is a single-threaded one-by-one IP queue test, so the number of IPs waiting for the download speed test is called`队列`。

* * *

> You may have noticed,**It was clearly designated to find 5 IPs that meet the download speed conditions, but why was it "interrupted" after only 3?**

Download speed test progress bar`3 / 5`, the former refers to finding`3`IPs that meet the lower limit of download speed (that is, the download speed is higher than`1 MB/s`),the latter`5`Refers to your request to find`5`IPs that meet the lower limit of download speed (`-dn 5`）。

> Also, a reminder, if you specify`-dn`is larger than the download speed test queue. For example, after you delay the speed test, only`4`IP, then the numbers behind the download speed test progress bar will be the same as the download speed test queue.`4`one, not you`-dn`Specified`5`Already.

The software is finishing the speed test`10`After IPs, only`3`download speeds higher than`1 MB/s`of IPs, and the rest`7`Each IP is "failed".

Therefore, this is not`“每次测速都不到 5 就中断了”`, but all IPs have downloaded and tested the speed, but only found`3`one that meets the conditions.

* * *

If you don’t want to encounter this situation where few of the speed tests meet the conditions, then you have to**Lower the download speed limit parameter`-sl`**, or remove.

Because as long as you specify`-sl`parameters, then as long as there are not enough`-dn`number (default 10), the speed test will continue until enough or all speed tests are completed. Remove`-sl`and add`-dn 20`parameters, so that only the top 20 IPs with the lowest speed and delay are measured, and the speed measurement is stopped to save time.

* * *

In addition, if all queue IPs have been tested for speed, but there is no IP that meets the download speed conditions, then it will**Directly output the download speed test results of all queue IPs**, so that you can see the download speeds of these IPs, and you can know them in your mind, and then**Adjust it appropriately`-sl`try again**。

Similarly, in terms of delayed speed measurement,`可用: 30`、`队列：10`These two values ​​​​can also let you know whether the delay conditions you set are too harsh for you. If there are a lot of available IPs, but only 2 or 3 are left after conditional filtering, it goes without saying that you need**Adjust down expected latency/packet loss conditions**.

Of these two mechanisms, one tells you**Delayed packet loss conditions**Whether it is appropriate or not, one will tell you**Download speed conditions**Is it appropriate.

</details>

* * *

### Usage example

Windows To specify parameters you need to run in CMD, or add parameters to the shortcut target.

> **Notice**: Each parameter has**default value**, parameters using default values ​​can be omitted (**Choose as needed**),parameter**In no particular order**。  
> **hint**：Windows**PowerShell**Just put in the following command`CloudflareScanner.exe`Change to`.\CloudflareScanner.exe`That’s it.  
> **hint**: Linux system only needs to change the following command`CloudflareScanner.exe`Change to`./CloudflareScanner`That’s it.

* * *

#### #CMD runs CloudflareScanner with parameters

People who are not familiar with command line programs may not know how to run them with parameters, so I will briefly explain.

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

Many people open CMD to**absolute path**Running CloudflareScanner will report an error. This is because the default`-f ip.txt`The parameter is a relative path, and you need to specify the absolute path of ip.txt. However, this is too troublesome, so it is recommended to enter the CloudflareScanner program directory to**relative path**Run as:

**method one**：

1.  Open the directory where the CloudflareScanner program is located
2.  Press on blank space<kbd>Shift + right mouse button</kbd>Show right-click menu
3.  choose**\[Open command window here]**to open the CMD window, which will be located in the current directory by default.
4.  Enter a command with parameters, such as:`CloudflareScanner.exe -tll 50 -tl 200`Ready to run

**Method 2**：

1.  Open the directory where the CloudflareScanner program is located
2.  Select all directly in the folder address bar and enter`cmd`Press Enter to open the CMD window, which will be in the current directory by default.
3.  Enter a command with parameters, such as:`CloudflareScanner.exe -tll 50 -tl 200`Ready to run

> Of course, you can also open a CMD window at will and enter such as`cd /d "D:\Program Files\CloudflareScanner"`to enter the program directory

> **hint**: If using**PowerShell**Just put in the command`CloudflareScanner.exe`Change to`.\CloudflareScanner.exe`That’s it.

</details>

* * *

#### #Windows shortcut to run CloudflareScanner with parameters

If you don’t frequently modify the running parameters (for example, you usually just double-click to run), it is recommended to use the shortcut, which is more convenient.

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

Right click`CloudflareScanner.exe`document -**\[Create Shortcut]**, then right-click the shortcut -**\[Attributes]**, modify its**Target**：

```bash
# 如果要不输出结果文件，那么请加上 -o " "，引号里的是空格（没有空格会导致该参数被省略）。
D:\ABC\CloudflareScanner\CloudflareScanner.exe -n 500 -t 4 -dn 20 -dt 5 -o " "

# 如果文件路径包含引号，则需要把启动参数放在引号外面，记得引号和 - 之间有空格。
"D:\Program Files\CloudflareScanner\CloudflareScanner.exe" -n 500 -t 4 -dn 20 -dt 5 -o " "

# 注意！快捷方式 - 起始位置 不能是空的，否则就会因为绝对路径而找不到 ip.txt 文件
```

</details>

* * *

#### #IPv4/IPv6

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 指定自带的 IPv4 数据文件可测速这些 IPv4 地址（-f 默认值就是 ip.txt，所以该参数可省略）
CloudflareScanner.exe -f ip.txt

# 指定自带的 IPv6 数据文件可测速这些 IPv6 地址
# 另外，v2.1.0 版本后支持 IPv4+IPv6 混合测速并移除了 -ipv6 参数，因此一个文件内可以同时包含 IPv4+IPv6 地址
CloudflareScanner.exe -f ipv6.txt

# 也可以直接通过参数指定要测速的 IP
CloudflareScanner.exe -ip 1.1.1.1,2606:4700::/32
```

> When measuring IPv6 speed, you may notice that the number of speed measurements is different each time. Learn the reason:[#120](https://github.com/Ptechgithub/CloudflareScanner/issues/120)  
> Because there are too many IPv6s (in billions), and most of the IP segments are not enabled at all, I only scanned some of the available IPv6 segments and wrote`ipv6.txt`In the file, those who are interested can scan for additions and deletions by themselves. The ASN data source comes from:[bgp.he.net](https://bgp.he.net/AS13335#_prefixes6)

</details>

* * *

#### #HTTPing

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

There are currently two delayed speed measurement modes, namely**TCP protocol, HTTP protocol**。  
The TCP protocol takes less time, consumes less resources, and has a timeout of 1 second. This protocol is the default mode.  
The HTTP protocol is suitable for quickly testing whether a domain name is accessible when pointing to a certain IP. The timeout is 2 seconds.  
For the same IP, the delay obtained by pinging various protocols is generally:**ICMP &lt; TCP &lt; HTTP**, the further to the right, the more sensitive it is to network fluctuations such as packet loss.

> Note: HTTPing is essentially a**Network scan**behavior, so if you are running on a server, you need**Reduce concurrency**(`-n`), otherwise the service may be suspended by some strict merchants. If you encounter a situation where the number of available IPs in the first HTTPing speed test is normal, and subsequent speed tests are getting less and less or even directly to 0, but then recover after stopping for a period of time, it may be because the operator or Cloudflare CDN thinks you are scanning the network.**Trigger temporary restriction mechanism**, so it will recover after a while, it is recommended**Reduce concurrency**(`-n`) to reduce the occurrence of this situation.

> In addition, this software HTTPing only obtains**response headers**, and does not obtain the body content (that is, the URL file size does not affect the HTTPing test, but if you want to download the speed test, you still need a large file), similar to the curl -i function.

```bash
# 只需加上 -httping 参数即可切换到 HTTP 协议延迟测速模式
CloudflareScanner.exe -httping

# 软件会根据访问时网页返回的有效 HTTP 状态码来判断可用性（当然超时也算），默认对返回 200 301 302 这三个 HTTP 状态码的视为有效，可以手动指定认为有效的 HTTP 状态码，但只能指定一个（你需要提前确定测试地址正常情况下会返回哪个状态码）
CloudflareScanner.exe -httping -httping-code 200

# 通过 -url 参数来指定 HTTPing 测试地址（可以是任意网页 URL，不局限于具体文件地址）
CloudflareScanner.exe -httping -url https://cf.xiu2.xyz/url
# 如果你要 HTTPing 测试其他网站/CDN，那么指定一个该网站/使用该 CDN 的地址（因为软件默认地址是 Cloudflare 的，只能用于测试 Cloudflare 的 IP）

# 注意：如果测速地址为 HTTP 协议，记得加上 -tp 80（这个参数会影响 延迟测速/下载测速 时使用的端口）
# 同理，如果要测速 80 端口，那么也需要加上 -url 参数来指定一个 http:// 协议的地址才行（且该地址不会强制重定向至 HTTPS），如果是非 80 443 端口，那么需要确定该下载测速地址是否支持通过该端口访问。
CloudflareScanner.exe -httping -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

* * *

#### #Match the specified area (colo airport three-character code)

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 该功能支持 Cloudflare CDN、AWS CloudFront CDN，且这两个 CDN 的机场三字码是通用的
# 注意：如果你要用于筛选 AWS CloudFront CDN 地区，那么要通过 -url 参数指定一个使用该 CDN 的地址（因为软件默认地址是 Cloudflare 的）

# 指定地区名后，延迟测速后得到的结果就都是指定地区的 IP 了（也可以继续进行下载测速）
# 节点地区名为当地 机场三字码，指定多个时用英文逗号分隔，v2.2.3 版本后支持小写

CloudflareScanner.exe -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD

# 注意，该参数只有在 HTTPing 延迟测速模式下才可用（因为要访问网页来获得）
```

> The two CDN airport three-character codes are common, so each region name is visible:<https://www.CloudflareScanneratus.com/>

</details>

* * *

#### #File relative/absolute path

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 指定 IPv4 数据文件，不显示结果直接退出，输出结果到文件（-p 值为 0）
CloudflareScanner.exe -f 1.txt -p 0 -dd

# 指定 IPv4 数据文件，不输出结果到文件，直接显示结果（-p 值为 10 条，-o 值为空但引号不能少）
CloudflareScanner.exe -f 2.txt -o "" -p 10 -dd

# 指定 IPv4 数据文件 及 输出结果到文件（相对路径，即当前目录下，如含空格请加上引号）
CloudflareScanner.exe -f 3.txt -o result.txt -dd


# 指定 IPv4 数据文件 及 输出结果到文件（相对路径，即当前目录内的 abc 文件夹下，如含空格请加上引号）
# Linux（CloudflareScanner 程序所在目录内的 abc 文件夹下）
./CloudflareScanner -f abc/3.txt -o abc/result.txt -dd

# Windows（注意是反斜杠）
CloudflareScanner.exe -f abc\3.txt -o abc\result.txt -dd


# 指定 IPv4 数据文件 及 输出结果到文件（绝对路径，即 C:\abc\ 目录下，如含空格请加上引号）
# Linux（/abc/ 目录下）
./CloudflareScanner -f /abc/4.txt -o /abc/result.csv -dd

# Windows（注意是反斜杠）
CloudflareScanner.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd


# 如果要以【绝对路径】运行 CloudflareScanner，那么 -f / -o 参数中的文件名也必须是【绝对路径】，否则会报错找不到文件！
# Linux（/abc/ 目录下）
/abc/CloudflareScanner -f /abc/4.txt -o /abc/result.csv -dd

# Windows（注意是反斜杠）
C:\abc\CloudflareScanner.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd
```

</details>

* * *

#### #Test speed of other ports

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 如果你想要测速非默认 443 的其他端口，则需要通过 -tp 参数指定（该参数会影响 延迟测速/下载测速 时使用的端口）

# 如果要延迟测速 80 端口+下载测速（如果 -dd 禁用了下载测速则不需要），那么还需要指定 http:// 协议的下载测速地址才行（且该地址不会强制重定向至 HTTPS，因为那样就变成 443 端口了）
CloudflareScanner.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm

# 如果是非 80 443 的其他端口，那么需要确定你使用的下载测速地址是否支持通过该非标端口访问。
```

</details>

* * *

#### #Custom speed test address

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

```bash
# 该参数适用于下载测速 及 HTTP 协议的延迟测速，对于后者该地址可以是任意网页 URL（不局限于具体文件地址）

# 地址要求：可以直接下载、文件大小超过 200MB、用的是 Cloudflare CDN
CloudflareScanner.exe -url https://cf.xiu2.xyz/url

# 注意：如果测速地址为 HTTP 协议（该地址不能强制重定向至 HTTPS），记得加上 -tp 80（这个参数会影响 延迟测速/下载测速 时使用的端口），如果是非 80 443 端口，那么需要确定下载测速地址是否支持通过该端口访问。
CloudflareScanner.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

* * *

#### #Custom speed test conditions (specify the target range of latency/packet loss/download speed)

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

> Note: The delay speed test progress bar on the right**Quantity Available**, only refers to the delayed speed measurement process**Number of IPs that have not timed out**, regardless of the delay upper and lower bound conditions.

-   Specify only**[Average latency cap]**condition

```bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s
# 即找到平均延迟低于 200 ms 的 IP，然后再按延迟从低到高进行 10 次下载测速
CloudflareScanner.exe -tl 200
```

> if**No satisfying delay found**conditional IP, nothing will be output.

* * *

-   Specify only**[Average latency cap]**conditions, and**Only delay the speed test, do not download the speed test**

```bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s，数量：不知道多少 个
# 即只输出低于 200ms 的 IP，且不再下载测速（因为不再下载测速，所以 -dn 参数就无效了）
CloudflareScanner.exe -tl 200 -dd
```

-   Specify only**[Maximum packet loss probability]**condition

```bash
# 丢包几率上限：0.25
# 即找到丢包率低于等于 0.25 的 IP，范围 0.00~1.00，如果 -tlr 0 则代表过滤掉任何丢包的 IP
CloudflareScanner.exe -tlr 0.25
```

* * *

-   Specify only**[Download speed minimum]**condition

```bash
# 平均延迟上限：9999 ms，下载速度下限：5 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 9999 ms 且下载速度高于 5 MB/s 的 IP 才会停止测速
CloudflareScanner.exe -sl 5 -dn 10
```

> if**Didn't find one that met the speed**conditional IP, then it will**Ignore conditions and output all IP speed test results**(This will help you adjust the conditions next time you test your speed).

> When the upper limit of average delay is not specified, if it has been**Not enough**The number of IPs that meet the conditions will**Always measuring speed**Go down.  
> So it is recommended**Specify at the same time[Download speed minimum]+[Average latency cap]**, so that the speed test will be terminated before the specified delay upper limit is reached before sufficient numbers are collected.

* * *

-   Specify at the same time**[Average latency cap]+[Download speed minimum]**condition

```bash
# 平均延迟上限、下载速度下限均支持小数（如 -sl 0.5）
# 平均延迟上限：200 ms，下载速度下限：5.6 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 200 ms 且下载速度高于 5 .6MB/s 的 IP 才会停止测速
CloudflareScanner.exe -tl 200 -sl 5.6 -dn 10
```

> if**No satisfying delay found**conditional IP, nothing will be output.  
> if**Didn't find one that met the speed**IP of the condition, then all IP speed test results will be output regardless of the condition (convenient for you to adjust the conditions the next time you test the speed).  
> Therefore, it is recommended to test the speed without specifying conditions first to see what the average delay and download speed are, and avoid specifying conditions.**too low/too high**！

> Because the IP range exposed by Cloudflare is**Back-to-origin IP + Anycast IP**,and**Back to source IP**is unavailable, so the download speed test is 0.00.  
> You can add it at runtime`-sl 0.01`(lower limit of download speed), filter out**Back to source IP**(Results of download speed test below 0.01MB/s).

</details>

* * *

#### #Test the speed of one or multiple IPs individually

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

**method one**:
Directly specify the IP segment data to be measured through parameters.

```bash
# 先进入 CloudflareScanner 所在目录，然后运行：
# Windows 系统（在 CMD 中运行）
CloudflareScanner.exe -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32

# Linux 系统
./CloudflareScanner -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
```

* * *

**Method 2**:
Or write these IPs into any text file in the following format, for example:`1.txt`

    1.1.1.1
    1.1.1.200
    1.0.0.1/24
    2606:4700::/32

> You can omit it for a single IP`/32`subnet masked (i.e.`1.1.1.1`Equivalent to`1.1.1.1/32`）。  
> subnet mask`/24`Refers to the last segment of this IP, that is`1.0.0.1~1.0.0.255`。

Then add startup parameters when running CloudflareScanner`-f 1.txt`to specify the IP segment data file.

```bash
# 先进入 CloudflareScanner 所在目录，然后运行：
# Windows 系统（在 CMD 中运行）
CloudflareScanner.exe -f 1.txt

# Linux 系统
./CloudflareScanner -f 1.txt

# 对于 1.0.0.1/24 这样的 IP 段只会随机最后一段（1.0.0.1~255），如果要测速该 IP 段中的所有 IP，请加上 -allip 参数。
```

</details>

* * *

#### #Speed ​​up all websites using Cloudflare CDN once and for all (no need to add domain names to Hosts one by one)

I have said before that the purpose of developing this software project is to pass**Change the Hosts method to speed up access to websites using Cloudflare CDN**。

But just like[**#8**](https://github.com/Ptechgithub/CloudflareScanner/issues/8)As mentioned, adding domain names to Hosts one by one is actually**Too troublesome**, so I found one**Once and for all**way! You can watch this[**Still adding Hosts one by one? The perfect local acceleration method for all websites using Cloudflare CDN is here!**](https://github.com/Ptechgithub/CloudflareScanner/discussions/71)and another[Rely on the local DNS service to modify the domain name resolution IP to a self-selected IP](https://github.com/Ptechgithub/CloudflareScanner/discussions/317)tutorial.

* * *

#### #Automatically update Hosts

Considering that many people need to replace the IP in the Hosts file after getting the fastest Cloudflare CDN IP.

You can watch this[**Issues**](https://github.com/Ptechgithub/CloudflareScanner/discussions/312)Obtain**Windows/Linux automatic update Hosts script**！

* * *

## feedback

If you encounter any problems, you can go first[**Issues**](https://github.com/Ptechgithub/CloudflareScanner/issues)、[Discussions](https://github.com/Ptechgithub/CloudflareScanner/discussions)Check here to see if anyone else has asked (remember to check it out)[**Closed**](https://github.com/Ptechgithub/CloudflareScanner/issues?q=is%3Aissue+is%3Aclosed)of).  
If no similar question is found, please open a new one[**Issues**](https://github.com/Ptechgithub/CloudflareScanner/issues/new)Come tell me!

> **Notice**！_and`反馈问题、功能建议`If it is irrelevant, please go to the project internal forum for discussion (above`💬 Discussions`_

* * *

## Manual compilation

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

For convenience, I write the version number into the version variable in the code when compiling, so when you compile manually, you need to add it as follows:`go build`Add after the command`-ldflags`Parameter to specify the version number:

```bash
go build -ldflags "-s -w -X main.version=v2.3.3"
# 在 CloudflareScanner 目录中通过命令行（例如 CMD、Bat 脚本）运行该命令，即可编译一个可在和当前设备同样系统、位数、架构的环境下运行的二进制程序（Go 会自动检测你的系统位数、架构）且版本号为 v2.3.3
```

If you want to compile under Windows 64-bit system**Other systems, architectures, bits**, then you need to specify**GOOS**and**GOARCH**variable.

For example, under Windows system, compile a program suitable for**Linux system amd architecture 64 bit**Binary program for:

```bat
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w -X main.version=v2.3.3"
```

For example, under Linux system, compile a program suitable for**Windows system amd architecture 32 bit**Binary program for:

```bash
GOOS=windows
GOARCH=386
go build -ldflags "-s -w -X main.version=v2.3.3"
```

> Can run`go tool dist list`to see which combinations the current Go version supports compiling.

* * *

Of course, in order to facilitate batch compilation, I will specifically specify a variable as the version number, and subsequent compilations can directly call the version number variable.  
At the same time, if you want to compile in batches, you need to put them in different folders (or the file names are different). You need to add`-o`Parameters specified.

```bat
:: Windows 系统下是这样：
SET version=v2.3.3
SET GOOS=linux
SET GOARCH=amd64
go build -o Releases\CloudflareScanner_linux_amd64\CloudflareScanner -ldflags "-s -w -X main.version=%version%"
```

```bash
# Linux 系统下是这样：
version=v2.3.3
GOOS=windows
GOARCH=386
go build -o Releases/CloudflareScanner_windows_386/CloudflareScanner.exe -ldflags "-s -w -X main.version=${version}"
```

</details>

* * *

## License

The GPL-3.0 License.
