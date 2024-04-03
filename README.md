# CloudflareScanner

**Many foreign websites use Cloudflare CDN, but the IPs allocated to visitors in mainland China are not user-friendly (high latency, high packet loss, slow speed). Although Cloudflare has publicly disclosed all [IP ranges](https://www.cloudflare.com/ips/), finding the most suitable one among so many IPs can be exhausting. Hence, this software was developed.**

**"Select the Best IP Yourself" tests the latency and speed of Cloudflare CDN, obtaining the fastest IP (IPv4+IPv6)! If you find it useful, give it a star to encourage us~**

****
## \# Quick Start

### Download and Run

1. Download the precompiled executable file ([Github Releases](https://github.com/Ptechgithub/CloudflareScanner/releases) / [蓝奏云](https://pan.lanpw.com/b0742hkxe)) and unzip it.
2. Double click on the `CloudflareST.exe` file (for Windows), and wait for the speed test to complete...

<details>
<summary><code><strong>「 Click to view usage example on Linux system 」</strong></code></summary>

****

The following commands are for demonstration purposes only. Please visit [**Releases**](https://github.com/Ptechgithub/CloudflareScanner/releases) to check the version number and file names.

``` yaml
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

> If the **average latency is very low** (e.g., 0.xx), it indicates that CloudflareST is **using a proxy** during the speed test. Please disable the proxy software before testing again.
> If running on a **router**, it's recommended to disable any proxy settings within the router (or exclude CloudflareST from them), otherwise the speed test results may be **inaccurate/unusable**.

</details>

****

> _A simple tutorial on how to independently run CloudflareST speed test on **mobile devices**: **[Android](https://github.com/Ptechgithub/CloudflareScanner/discussions/61)**, **[Android APP](https://github.com/xianshenglu/cloudflare-ip-tester-app)**, **[iOS](https://github.com/Ptechgithub/CloudflareScanner/discussions/321)**_

> Note! This software is only applicable to websites and **does not support selecting Cloudflare WARP preferred IPs**. For details, see: [#392](https://github.com/Ptechgithub/CloudflareScanner/discussions/392)

### Example Results

After the speed test is completed, the **top 10 fastest IPs** will be displayed by default, for example:

``` bash
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

The first line of the speed test result is the **fastest IP with both the fastest download speed and the lowest average latency**!

The complete results are saved in the `result.csv` file in the current directory. Open it with **Notepad/spreadsheet software**. The format is as follows:

```
IP Address, Sent, Received, Loss Rate, Avg Latency, Download Speed (MB/s)
104.27.200.69, 4, 4, 0.00, 146.23, 28.64
```

> _Everyone can further filter and process the complete results according to their own needs, or explore advanced usage to specify filtering criteria!_

****
## \# Advanced Usage

Running the tool directly uses default parameters. If you want more comprehensive and customized speed test results, you can customize the parameters.

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

### Interface Explanation

To avoid misunderstanding the **output content during the speed test process (available, queued numbers, download speed "interrupted" halfway)**, I will explain it here.

<details>
<summary><code><strong>「 Click to expand to view content 」</strong></code></summary>

****

> The example adds commonly used parameters, which are: `-ttl 40 -tl 150 -sl 1 -dn 5`, and the final output result is as follows:

``` bash
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

****

> For those who are new to CloudflareST, they might be confused: **"There were originally 30 usable IPs for latency testing, why are there only 3 left now?"** What does the queue in the download speed test mean? Do I have to wait in line for the download speed test as well?

CloudflareST first conducts latency testing. During this process, the progress bar on the right side will display the real-time number of available IPs (`Available: 30`). However, note that this available number refers to the **number of IPs that have passed the test without timing out**, regardless of the latency upper and lower limits or packet loss conditions. After the latency testing is completed, because specific conditions for latency upper and lower limits and packet loss were specified, only `10` IPs remain after filtering (which indicates a download speed test queue of `10`).

In the example above, out of `321` IPs, only `30` IPs passed the latency test without timing out. Then, after filtering based on the latency upper and lower limits (`40 ~ 150 ms`) and the packet loss limit, only `10` IPs meeting the requirements remain. If you disable download speed testing with `-dd`, then these `10` IPs will be output directly. However, in this example, download speed testing is not disabled, so the software will continue to perform download speed testing on these `10` IPs (`Queue: 10`).

> Since download speed testing is single-threaded and tests IPs one by one in sequence, the number of IPs waiting for download speed testing is referred to as the `Queue`.

****

> You may have noticed, **even though you specified to find 5 IPs that meet the download speed conditions, why did the process "terminate" with only 3?**

In the download speed test progress bar, `3 / 5` indicates that `3` IPs meeting the download speed lower limit condition (i.e., download speed exceeding `1 MB/s`) have been found, while `5` indicates that you requested to find `5` IPs meeting this condition (`-dn 5`).

> Additionally, please note that if you specify `-dn` to be greater than the download speed test queue, for example, if only `4` IPs remain after latency testing, then the number following the `/` in the download speed test progress bar will be `4`, just like the download speed test queue, rather than the `5` you specified with `-dn`.

After testing the download speed of these `10` IPs, only `3` IPs were found to have a download speed exceeding `1 MB/s`, while the remaining `7` IPs did not meet the criteria.

Therefore, it's not that `“the test terminates every time without reaching 5”`, but rather that all IPs were tested for download speed, but only `3` met the criteria.

****

If you don't want to encounter situations where only a few IPs meet the criteria after testing all of them, you can **lower the download speed upper limit parameter `-sl`** or remove it.

Because as long as the `-sl` parameter is specified, testing will continue until the `-dn` quantity (default 10 IPs) is reached, or all IPs are tested. Removing `-sl` and adding the `-dn 20` parameter will only test the latency of the top 20 IPs with the lowest latency and then stop, saving time.

****

Furthermore, if all IPs in the queue have been tested for download speed but none meet the download speed criteria, then **the download speed test results for all IPs in the queue will be output directly**. This way, you can see the download speeds of these IPs and have an idea of ​​how they perform. Then, **try lowering `-sl` appropriately and try again**.

Similarly, regarding latency testing, the two values `Available: 30` and `Queue: 10` can also help you determine whether the latency conditions you set are too strict. If you have plenty of available IPs, but after filtering conditions, only 2 or 3 remain, then it's clear that you need to **lower your expected latency/packet loss conditions**.

These two mechanisms, one informing you about **latency and packet loss conditions**, and the other about **download speed conditions**, help you determine whether your settings are appropriate.

</details>

****

### Usage Example

On Windows, to specify parameters, you need to run them in CMD, or add the parameters to the target of a shortcut.

> **Note**: All parameters have **default values**, and using parameters with default values can be omitted (**choose as needed**), parameters can be specified in any order.  
> **Tip**: For Windows **PowerShell**, simply change `CloudflareST.exe` in the commands below to `.\CloudflareST.exe`.  
> **Tip**: For Linux systems, simply change `CloudflareST.exe` in the commands below to `./CloudflareST`.

****

#### Running CloudflareST with Parameters in CMD

For those unfamiliar with command-line programs, you may not know how to run them with parameters, so I'll explain briefly.

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

****

Many people encounter errors when running CloudflareST with **absolute paths** in CMD. This is because the default `-f ip.txt` parameter uses a relative path, so you need to specify the absolute path to `ip.txt`. However, this can be cumbersome, so it's recommended to run CloudflareST in its program directory using **relative paths**:

**Method 1**:
1. Open the directory where CloudflareST is located.
2. Right-click on a blank space and select **\[Open command window here\]** from the context menu by pressing <kbd>Shift + Right Click</kbd>, which will open CMD with the current directory.
3. Enter the command with parameters, such as: `CloudflareST.exe -tll 50 -tl 200` to run the program.

**Method 2**:
1. Open the directory where CloudflareST is located.
2. Directly type `cmd` in the address bar of the folder and press Enter to open CMD with the current directory.
3. Enter the command with parameters, such as: `CloudflareST.exe -tll 50 -tl 200` to run the program.

> Of course, you can also open any CMD window and then enter a command like `cd /d "D:\Program Files\CloudflareST"` to enter the program directory.

> **Tip**: If you're using **PowerShell**, simply change `CloudflareST.exe` in the command to `.\CloudflareST.exe`.

</details>

****

#### Running CloudflareST with Parameters via Windows Shortcut

For those who don't frequently modify runtime parameters (such as usually double-clicking to run), using a shortcut is more convenient.

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

****

Right-click the `CloudflareST.exe` file - **\[Create Shortcut\]**, then right-click the created shortcut - **\[Properties\]**, and modify its **Target**:

```bash
# If you don't want to output a result file, please add -o " ", where the quotes contain a space (omitting the space will cause this parameter to be omitted).
D:\ABC\CloudflareST\CloudflareST.exe -n 500 -t 4 -dn 20 -dt 5 -o " "

# If the file path contains quotes, then the startup parameters need to be placed outside the quotes. Remember to have a space between the quotes and the -.
"D:\Program Files\CloudflareST\CloudflareST.exe" -n 500 -t 4 -dn 20 -dt 5 -o " "

# Note! The starting location of the shortcut cannot be empty, otherwise, it will not find the ip.txt file due to the absolute path.
```

</details>

****

#### \# IPv4/IPv6

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

****
```bash
# Specify the built-in IPv4 data file to test these IPv4 addresses (the -f parameter defaults to ip.txt, so it can be omitted)
CloudflareST.exe -f ip.txt

# Specify the built-in IPv6 data file to test these IPv6 addresses
# Additionally, starting from version 2.1.0, CloudflareST supports testing IPv4+IPv6 addresses together and removes the -ipv6 parameter. Therefore, one file can contain both IPv4+IPv6 addresses.
CloudflareST.exe -f ipv6.txt

# You can also directly specify the IPs to be tested through parameters
CloudflareST.exe -ip 1.1.1.1,2606:4700::/32
```

When testing IPv6, you may notice that the number of tests varies each time. The reason is explained here: [#120](https://github.com/Ptechgithub/CloudflareScanner/issues/120)  
Because there are too many IPv6 addresses (in the billions), and the vast majority of IP ranges are not in use, I only scanned a portion of the available IPv6 ranges and wrote them to the `ipv6.txt` file. If you're interested, you can scan and modify them yourself. The ASN data source is from: [bgp.he.net](https://bgp.he.net/AS13335#_prefixes6)

****

#### \# HTTPing

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

****

There are currently two latency testing modes, **TCP protocol and HTTP protocol**.  
The TCP protocol is faster and consumes fewer resources, with a timeout of 1 second, and this is the default mode.  
The HTTP protocol is suitable for quickly testing whether a domain points to a certain IP and whether it can be accessed. The timeout is set to 2 seconds.  
For the same IP, the latencies obtained by each protocol generally follow this order: **ICMP < TCP < HTTP**, with HTTP being more sensitive to network fluctuations such as packet loss.

> Note: HTTPing is essentially a form of **network scanning** behavior. Therefore, if you are running it on a server, you need to **reduce concurrency** (`-n`), otherwise, you may be suspended by some strict providers. If you encounter a situation where the number of available IPs decreases during HTTPing, or even becomes 0, but then recovers after a while, it may be due to **temporary restriction mechanisms** triggered by the ISP or Cloudflare CDN, perceiving your activity as network scanning. In this case, it usually recovers after a while. It is recommended to **reduce concurrency** (`-n`) to reduce the occurrence of such situations.

> Additionally, this software's HTTPing only retrieves **response headers** and does not retrieve the body content (meaning the URL file size does not affect HTTPing testing, but if you also want to perform download speed testing, you will need a large file). It is similar to the curl -i functionality.

``` bash
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

****

#### \# 匹配指定地区(colo 机场三字码)

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

``` bash
# 该功能支持 Cloudflare CDN、AWS CloudFront CDN，且这两个 CDN 的机场三字码是通用的
# 注意：如果你要用于筛选 AWS CloudFront CDN 地区，那么要通过 -url 参数指定一个使用该 CDN 的地址（因为软件默认地址是 Cloudflare 的）

# 指定地区名后，延迟测速后得到的结果就都是指定地区的 IP 了（也可以继续进行下载测速）
# 节点地区名为当地 机场三字码，指定多个时用英文逗号分隔，v2.2.3 版本后支持小写

CloudflareST.exe -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD

# 注意，该参数只有在 HTTPing 延迟测速模式下才可用（因为要访问网页来获得）
```

> 两个 CDN 机场三字码通用，因此各地区名可见：https://www.cloudflarestatus.com/

</details>

****

#### \# 文件相对/绝对路径

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

``` bash
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

****

#### \# 测速其他端口

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

``` bash
# 如果你想要测速非默认 443 的其他端口，则需要通过 -tp 参数指定（该参数会影响 延迟测速/下载测速 时使用的端口）

# 如果要延迟测速 80 端口+下载测速（如果 -dd 禁用了下载测速则不需要），那么还需要指定 http:// 协议的下载测速地址才行（且该地址不会强制重定向至 HTTPS，因为那样就变成 443 端口了）
CloudflareST.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm

# 如果是非 80 443 的其他端口，那么需要确定你使用的下载测速地址是否支持通过该非标端口访问。
```

</details>

****

#### \# 自定义测速地址

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

``` bash
# 该参数适用于下载测速 及 HTTP 协议的延迟测速，对于后者该地址可以是任意网页 URL（不局限于具体文件地址）

# 地址要求：可以直接下载、文件大小超过 200MB、用的是 Cloudflare CDN
CloudflareST.exe -url https://cf.xiu2.xyz/url

# 注意：如果测速地址为 HTTP 协议（该地址不能强制重定向至 HTTPS），记得加上 -tp 80（这个参数会影响 延迟测速/下载测速 时使用的端口），如果是非 80 443 端口，那么需要确定下载测速地址是否支持通过该端口访问。
CloudflareST.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

****

#### \# 自定义测速条件（指定 延迟/丢包/下载速度 的目标范围）

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

> 注意：延迟测速进度条右边的**可用数量**，仅指延迟测速过程中**未超时的 IP 数量**，和延迟上下限条件无关。

- 仅指定 **[平均延迟上限]** 条件

``` bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s
# 即找到平均延迟低于 200 ms 的 IP，然后再按延迟从低到高进行 10 次下载测速
CloudflareST.exe -tl 200
```

> 如果**没有找到一个满足延迟**条件的 IP，那么不会输出任何内容。

****

- 仅指定 **[平均延迟上限]** 条件，且**只延迟测速，不下载测速**

``` bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s，数量：不知道多少 个
# 即只输出低于 200ms 的 IP，且不再下载测速（因为不再下载测速，所以 -dn 参数就无效了）
CloudflareST.exe -tl 200 -dd
```

- 仅指定 **[丢包几率上限]** 条件

``` bash
# 丢包几率上限：0.25
# 即找到丢包率低于等于 0.25 的 IP，范围 0.00~1.00，如果 -tlr 0 则代表过滤掉任何丢包的 IP
CloudflareST.exe -tlr 0.25
```

****

- 仅指定 **[下载速度下限]** 条件

``` bash
# 平均延迟上限：9999 ms，下载速度下限：5 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 9999 ms 且下载速度高于 5 MB/s 的 IP 才会停止测速
CloudflareST.exe -sl 5 -dn 10
```

> 如果**没有找到一个满足速度**条件的 IP，那么会**忽略条件输出所有 IP 测速结果**（方便你下次测速时调整条件）。

> 没有指定平均延迟上限时，如果一直**凑不够**满足条件的 IP 数量，就会**一直测速**下去。  
> 所以建议**同时指定 [下载速度下限] + [平均延迟上限]**，这样测速到指定延迟上限还没凑够数量，就会终止测速。

****

- 同时指定 **[平均延迟上限] + [下载速度下限]** 条件

``` bash
# 平均延迟上限、下载速度下限均支持小数（如 -sl 0.5）
# 平均延迟上限：200 ms，下载速度下限：5.6 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 200 ms 且下载速度高于 5 .6MB/s 的 IP 才会停止测速
CloudflareST.exe -tl 200 -sl 5.6 -dn 10
```

> 如果**没有找到一个满足延迟**条件的 IP，那么不会输出任何内容。  
> 如果**没有找到一个满足速度**条件的 IP，那么会忽略条件输出所有 IP 测速结果（方便你下次测速时调整条件）。  
> 所以建议先不指定条件测速一遍，看看平均延迟和下载速度大概在什么范围，避免指定条件**过低/过高**！

> 因为 Cloudflare 公开的 IP 段是**回源 IP+任播 IP**，而**回源 IP**是无法使用的，所以下载测速是 0.00。  
> 运行时可以加上 `-sl 0.01`（下载速度下限），过滤掉**回源 IP**（下载测速低于 0.01MB/s 的结果）。

</details>

****

#### \# 单独对一个或多个 IP 测速

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

****

**方式 一**：
直接通过参数指定要测速的 IP 段数据。
``` bash
# 先进入 CloudflareST 所在目录，然后运行：
# Windows 系统（在 CMD 中运行）
CloudflareST.exe -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32

# Linux 系统
./CloudflareST -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
```

****

**方式 二**：
或者把这些 IP 按如下格式写入到任意文本文件中，例如：`1.txt`

```
1.1.1.1
1.1.1.200
1.0.0.1/24
2606:4700::/32
```

> 单个 IP 的话可以省略 `/32` 子网掩码了（即 `1.1.1.1`等同于 `1.1.1.1/32`）。  
> 子网掩码 `/24` 指的是这个 IP 最后一段，即 `1.0.0.1~1.0.0.255`。


然后运行 CloudflareST 时加上启动参数 `-f 1.txt` 来指定 IP 段数据文件。

``` bash
# 先进入 CloudflareST 所在目录，然后运行：
# Windows 系统（在 CMD 中运行）
CloudflareST.exe -f 1.txt

# Linux 系统
./CloudflareST -f 1.txt

# 对于 1.0.0.1/24 这样的 IP 段只会随机最后一段（1.0.0.1~255），如果要测速该 IP 段中的所有 IP，请加上 -allip 参数。
```

</details>

****

#### \# 一劳永逸加速所有使用 Cloudflare CDN 的网站（不需要再一个个添加域名到 Hosts 了）

我以前说过，开发该软件项目的目的就是为了通过**改 Hosts 的方式来加速访问使用 Cloudflare CDN 的网站**。

但就如 [**#8**](https://github.com/Ptechgithub/CloudflareScanner/issues/8) 所说，一个个添加域名到 Hosts 实在**太麻烦**了，于是我就找到了个**一劳永逸**的办法！可以看这个 [**还在一个个添加 Hosts？完美本地加速所有使用 Cloudflare CDN 的网站方法来了！**](https://github.com/Ptechgithub/CloudflareScanner/discussions/71) 和另一个[依靠本地 DNS 服务来修改域名解析 IP 为自选 IP](https://github.com/Ptechgithub/CloudflareScanner/discussions/317) 的教程。

****
## Manual Compilation

<details>
<summary><code><strong>「 Click to expand to view content 」</strong></code></summary>

****

为了方便，我是在编译的时候将版本号写入代码中的 version 变量，因此你手动编译时，需要像下面这样在 `go build` 命令后面加上 `-ldflags` 参数来指定版本号：

```bash
go build -ldflags "-s -w -X main.version=v2.3.3"
# 在 CloudflareScanner 目录中通过命令行（例如 CMD、Bat 脚本）运行该命令，即可编译一个可在和当前设备同样系统、位数、架构的环境下运行的二进制程序（Go 会自动检测你的系统位数、架构）且版本号为 v2.3.3
```

如果想要在 Windows 64位系统下编译**其他系统、架构、位数**，那么需要指定 **GOOS** 和 **GOARCH** 变量。

例如在 Windows 系统下编译一个适用于 **Linux 系统 amd 架构 64 位**的二进制程序：

```bat
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w -X main.version=v2.3.3"
```

例如在 Linux 系统下编译一个适用于 **Windows 系统 amd 架构 32 位**的二进制程序：

```bash
GOOS=windows
GOARCH=386
go build -ldflags "-s -w -X main.version=v2.3.3"
```

> 可以运行 `go tool dist list` 来查看当前 Go 版本支持编译哪些组合。

****

当然，为了方便批量编译，我会专门指定一个变量为版本号，后续编译直接调用该版本号变量即可。  
同时，批量编译的话，还需要分开放到不同文件夹才行（或者文件名不同），需要加上 `-o` 参数指定。

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

****

## License

The GPL-3.0 License.