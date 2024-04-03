# CloudflareScanner

**بسیاری از وب سایت های خارجی از Cloudflare CDN استفاده می کنند، اما IP های اختصاص داده شده به بازدیدکنندگان در سرزمین اصلی چین کاربر پسند نیستند (تأخیر بالا، از دست دادن بسته های بالا، سرعت پایین). اگرچه Cloudflare همه را به طور عمومی افشا کرده است[محدوده IP](https://www.cloudflare.com/ips/)، یافتن مناسب ترین مورد در میان بسیاری از IP ها می تواند خسته کننده باشد. از این رو این نرم افزار توسعه یافت.**

**"Select the Best IP Yourself" تأخیر و سرعت Cloudflare CDN را آزمایش می کند و سریع ترین IP (IPv4+IPv6) را به دست می آورد! اگر آن را مفید می دانید، برای تشویق ما به آن ستاره بدهید**

* * *

## #شروع سریع

### دانلود و اجرا کنید

1.  دانلود فایل اجرایی از پیش کامپایل شده ([انتشارات Github](https://github.com/Ptechgithub/CloudflareScanner/releases)/[لان زویون](https://pan.lanpw.com/b0742hkxe)) و آن را از حالت فشرده خارج کنید.
2.  روی آن دابل کلیک کنید`CloudflareST.exe`فایل (برای ویندوز) و منتظر بمانید تا تست سرعت کامل شود...

<details>
<summary><code><strong>「 Click to view usage example on Linux system 」</strong></code></summary>

* * *

دستورات زیر فقط برای اهداف نمایشی هستند. لطفا ببینید[**منتشر شده**](https://github.com/Ptechgithub/CloudflareScanner/releases)برای بررسی شماره نسخه و نام فایل ها.

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

> اگر**متوسط ​​تأخیر بسیار کم است**(به عنوان مثال، 0.xx)، نشان می دهد که CloudflareST است**با استفاده از یک پروکسی**در طول تست سرعت لطفاً قبل از آزمایش مجدد، نرم افزار پروکسی را غیرفعال کنید.
> اگر در حال اجرا بر روی a**روتر**، توصیه می شود تنظیمات پروکسی را در روتر غیرفعال کنید (یا CloudflareST را از آنها حذف کنید)، در غیر این صورت ممکن است نتایج تست سرعت**نادرست/غیرقابل استفاده**.

</details>

* * *

> _یک آموزش ساده در مورد نحوه اجرای مستقل تست سرعت CloudflareST روی**دستگاه های تلفن همراه**:**[اندروید](https://github.com/Ptechgithub/CloudflareScanner/discussions/61)**,**[برنامه اندروید](https://github.com/xianshenglu/cloudflare-ip-tester-app)**,**[iOS](https://github.com/Ptechgithub/CloudflareScanner/discussions/321)**_

> توجه داشته باشید! این نرم افزار فقط برای وب سایت ها و**از انتخاب IP های ترجیحی Cloudflare WARP پشتیبانی نمی کند**. برای جزئیات، نگاه کنید به:[#392](https://github.com/Ptechgithub/CloudflareScanner/discussions/392)

### نتایج نمونه

پس از اتمام تست سرعت،**10 سریعترین IP برتر**به طور پیش فرض نمایش داده می شود، به عنوان مثال:

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

خط اول نتیجه تست سرعت عبارت است از**سریعترین IP با سریعترین سرعت دانلود و کمترین میانگین تأخیر**!

نتایج کامل در قسمت ذخیره می شود`result.csv`فایل در دایرکتوری فعلی آن را با**نرم افزار دفترچه یادداشت/صفحه گسترده**. قالب به شرح زیر است:

    IP Address, Sent, Received, Loss Rate, Avg Latency, Download Speed (MB/s)
    104.27.200.69, 4, 4, 0.00, 146.23, 28.64

> _هرکسی می‌تواند نتایج کامل را بر اساس نیازهای خود فیلتر و پردازش کند، یا استفاده پیشرفته را برای تعیین معیارهای فیلتر جستجو کند!_

* * *

## #استفاده پیشرفته

اجرای مستقیم ابزار از پارامترهای پیش فرض استفاده می کند. اگر می‌خواهید نتایج تست سرعت جامع‌تر و سفارشی‌تر شود، می‌توانید پارامترها را سفارشی کنید.

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

### توضیح رابط

برای جلوگیری از سوء تفاهم**محتوای خروجی در طول فرآیند تست سرعت (موجود، اعداد در صف، سرعت دانلود "قطع" در نیمه راه)**، اینجا توضیح می دهم.

<details>
<summary><code><strong>「 Click to expand to view content 」</strong></code></summary>

* * *

> این مثال پارامترهای رایج را اضافه می کند که عبارتند از:`-ttl 40 -tl 150 -sl 1 -dn 5`و نتیجه خروجی نهایی به شرح زیر است:

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

> برای کسانی که در CloudflareST تازه کار هستند، ممکن است گیج شوند:**"در ابتدا 30 IP قابل استفاده برای آزمایش تاخیر وجود داشت، چرا اکنون فقط 3 IP باقی مانده است؟"**صف در تست سرعت دانلود به چه معناست؟ آیا برای تست سرعت دانلود هم باید در صف منتظر بمانم؟

CloudflareST ابتدا آزمایش تأخیر را انجام می دهد. در طول این فرآیند، نوار پیشرفت در سمت راست تعداد IP های موجود را نمایش می دهد (`Available: 30`). با این حال، توجه داشته باشید که این شماره موجود به**تعداد IP هایی که تست را بدون مهلت گذرانده اند**، صرف نظر از محدودیت های بالا و پایین تاخیر یا شرایط از دست دادن بسته. پس از اتمام تست تأخیر، به دلیل اینکه شرایط خاصی برای حدود بالا و پایین تأخیر و از دست دادن بسته مشخص شده بود، فقط`10`IP ها پس از فیلتر کردن باقی می مانند (که نشان دهنده صف تست سرعت دانلود است`10`).

در مثال بالا، خارج از`321`فقط آی پی ها`30`IP ها تست تاخیر را بدون مهلت سپری کردند. سپس پس از فیلتر کردن بر اساس حدود بالا و پایین تاخیر (`40 ~ 150 ms`) و فقط محدودیت از دست دادن بسته`10`IP های مطابق با الزامات باقی می مانند. اگر تست سرعت دانلود را با`-dd`، سپس اینها`10`آی پی ها مستقیماً خروجی خواهند شد. با این حال، در این مثال، تست سرعت دانلود غیرفعال نیست، بنابراین نرم افزار به انجام تست سرعت دانلود در این موارد ادامه می دهد.`10`آی پی ها (`Queue: 10`).

> از آنجایی که تست سرعت دانلود به صورت تک رشته ای است و IP ها را یک به یک به ترتیب تست می کند، به تعداد IP هایی که منتظر تست سرعت دانلود هستند، می گویند:`Queue`.

* * *

> شاید متوجه شده باشید،**با وجود اینکه مشخص کردید 5 آی پی را پیدا کنید که شرایط سرعت دانلود را داشته باشد، چرا فرآیند فقط با 3 "خاتمه" یافت؟**

در نوار پیشرفت تست سرعت دانلود،`3 / 5`نشان میدهد که`3`IPهایی که شرایط حد پایین‌تر سرعت دانلود را دارند (یعنی بیش از سرعت دانلود`1 MB/s`) پیدا شده اند، در حالی که`5`نشان می دهد که شما درخواست پیدا کرده اید`5`IP هایی که این شرایط را دارند (`-dn 5`).

> علاوه بر این، لطفاً توجه داشته باشید که اگر مشخص کنید`-dn`به عنوان مثال، اگر فقط از صف تست سرعت دانلود بیشتر باشد`4`IP ها پس از تست تأخیر باقی می مانند، سپس تعداد زیر`/`در نوار پیشرفت تست سرعت دانلود قرار خواهد گرفت`4`، درست مانند صف تست سرعت دانلود، به جای`5`شما مشخص کردید با`-dn`.

بعد از تست سرعت دانلود اینها`10`فقط آی پی ها`3`سرعت دانلود IPها بیش از حد معمول بود`1 MB/s`، در حالی که باقی مانده است`7`IP ها معیارها را برآورده نمی کنند.

بنابراین، این نیست`“the test terminates every time without reaching 5”`، بلکه همه IP ها برای سرعت دانلود تست شده اند، اما فقط`3`معیارها را برآورده کرد.

* * *

اگر نمی خواهید با شرایطی مواجه شوید که تنها چند آی پی پس از آزمایش همه آنها معیارها را دارند، می توانید**پارامتر حد بالایی سرعت دانلود را کاهش دهید`-sl`**یا آن را حذف کنید.

زیرا تا زمانی که`-sl`پارامتر مشخص شده است، آزمایش تا زمانی ادامه خواهد یافت`-dn`مقدار (پیش‌فرض 10 IP) به دست می‌آید، یا همه IP‌ها آزمایش می‌شوند. در حال برداشتن`-sl`و اضافه کردن`-dn 20`پارامتر فقط تأخیر 20 IP برتر را با کمترین تأخیر آزمایش می کند و سپس متوقف می شود و در زمان صرفه جویی می کند.

* * *

علاوه بر این، اگر تمام IP های موجود در صف برای سرعت دانلود تست شده باشند اما هیچ کدام معیارهای سرعت دانلود را نداشته باشند، پس**نتایج تست سرعت دانلود برای تمام IP های موجود در صف مستقیماً خروجی خواهد شد**. به این ترتیب می توانید سرعت دانلود این آی پی ها را مشاهده کنید و از عملکرد آنها ایده بگیرید. سپس،**سعی کنید پایین بیاورید`-sl`مناسب و دوباره امتحان کنید**.

به طور مشابه، در مورد آزمایش تأخیر، این دو مقدار`Available: 30`و`Queue: 10`همچنین می تواند به شما کمک کند تعیین کنید که آیا شرایط تأخیر تعیین شده خیلی سخت است یا خیر. اگر تعداد زیادی IP در دسترس دارید، اما پس از شرایط فیلتر کردن، تنها 2 یا 3 IP باقی مانده است، پس واضح است که باید**شرایط تأخیر/از دست دادن بسته مورد انتظار خود را کاهش دهید**.

این دو مکانیسم، یکی به شما اطلاع می دهد**تأخیر و شرایط از دست دادن بسته**، و دیگری در مورد**شرایط سرعت دانلود**، به شما کمک می کند تعیین کنید که آیا تنظیمات شما مناسب هستند یا خیر.

</details>

* * *

### مثال استفاده

در ویندوز، برای تعیین پارامترها، باید آنها را در CMD اجرا کنید، یا پارامترها را به هدف یک میانبر اضافه کنید.

> **توجه داشته باشید**: همه پارامترها دارند**مقادیر پیش فرض**، و استفاده از پارامترهای با مقادیر پیش فرض را می توان حذف کرد (**در صورت نیاز انتخاب کنید**، پارامترها را می توان به هر ترتیبی مشخص کرد.  
> **نکته**: برای ویندوز**پاورشل**، به سادگی تغییر دهید`CloudflareST.exe`در دستورات زیر به`.\CloudflareST.exe`.  
> **نکته**: برای سیستم های لینوکس، به سادگی تغییر دهید`CloudflareST.exe`در دستورات زیر به`./CloudflareST`.

* * *

#### اجرای CloudflareST با پارامترها در CMD

برای کسانی که با برنامه های خط فرمان آشنا نیستند، ممکن است ندانید که چگونه آنها را با پارامترها اجرا کنید، بنابراین به طور خلاصه توضیح می دهم.

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

بسیاری از مردم هنگام اجرای CloudflareST با خطا مواجه می شوند**مسیرهای مطلق**در CMD این به دلیل پیش فرض است`-f ip.txt`پارامتر از یک مسیر نسبی استفاده می کند، بنابراین باید مسیر مطلق را مشخص کنید`ip.txt`. با این حال، این می تواند دست و پا گیر باشد، بنابراین توصیه می شود CloudflareST را در فهرست برنامه های آن با استفاده از**مسیرهای نسبی**:

**روش 1**:

1.  دایرکتوری را که CloudflareST در آن قرار دارد باز کنید.
2.  روی یک فضای خالی کلیک راست کرده و انتخاب کنید**\[پنجره فرمان را در اینجا باز کنید]**از منوی زمینه با فشار دادن<kbd>Shift + کلیک راست</kbd>، که CMD را با دایرکتوری فعلی باز می کند.
3.  دستور را با پارامترهایی وارد کنید، مانند:`CloudflareST.exe -tll 50 -tl 200`برای اجرای برنامه

**روش 2**:

1.  دایرکتوری را که CloudflareST در آن قرار دارد باز کنید.
2.  مستقیم تایپ کنید`cmd`در نوار آدرس پوشه و Enter فشار دهید تا CMD با دایرکتوری فعلی باز شود.
3.  دستور را با پارامترهایی وارد کنید، مانند:`CloudflareST.exe -tll 50 -tl 200`برای اجرای برنامه

> البته می توانید هر پنجره CMD را نیز باز کنید و سپس دستوری مانند آن را وارد کنید`cd /d "D:\Program Files\CloudflareST"`برای ورود به دایرکتوری برنامه

> **نکته**: اگر استفاده می کنید**پاورشل**، به سادگی تغییر دهید`CloudflareST.exe`در فرمان به`.\CloudflareST.exe`.

</details>

* * *

#### اجرای CloudflareST با پارامترها از طریق میانبر ویندوز

برای کسانی که اغلب پارامترهای زمان اجرا را تغییر نمی دهند (مانند معمولاً دوبار کلیک کردن برای اجرا)، استفاده از میانبر راحت تر است.

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

کلیک راست کنید`CloudflareST.exe`فایل -**\[ایجاد میانبر]**، سپس روی میانبر ایجاد شده راست کلیک کنید -**\[خواص]**، و آن را اصلاح کنید**هدف**:

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

هنگام آزمایش IPv6، ممکن است متوجه شوید که تعداد تست ها در هر بار متفاوت است. دلیل آن در اینجا توضیح داده شده است:[#120](https://github.com/Ptechgithub/CloudflareScanner/issues/120)  
از آنجایی که آدرس‌های IPv6 بسیار زیادی وجود دارد (میلیاردها)، و اکثریت قریب به اتفاق محدوده‌های IP مورد استفاده قرار نمی‌گیرند، من فقط بخشی از محدوده IPv6 موجود را اسکن کردم و آن‌ها را در ایمیل نوشتم.`ipv6.txt`فایل. اگر علاقه مند هستید، می توانید خودتان آنها را اسکن و اصلاح کنید. منبع داده ASN از:[bgp.he.net](https://bgp.he.net/AS13335#_prefixes6)

* * *

#### #HTTPing

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

* * *

در حال حاضر دو حالت تست تاخیر وجود دارد،**پروتکل TCP و پروتکل HTTP**.  
پروتکل TCP سریعتر است و منابع کمتری را مصرف می کند، با تایم اوت 1 ثانیه، و این حالت پیش فرض است.  
پروتکل HTTP برای آزمایش سریع اینکه آیا دامنه به IP خاصی اشاره می کند و اینکه آیا می توان به آن دسترسی داشت یا خیر، مناسب است. تایم اوت روی 2 ثانیه تنظیم شده است.  
برای همان IP، تاخیرهای به دست آمده توسط هر پروتکل معمولاً از این ترتیب پیروی می کنند:**ICMP &lt; TCP &lt; HTTP**، HTTP نسبت به نوسانات شبکه مانند از دست دادن بسته حساس تر است.

> توجه: HTTPing اساساً نوعی از است**اسکن شبکه**رفتار - اخلاق. بنابراین، اگر آن را روی سرور اجرا می کنید، باید**کاهش همزمانی**(`-n`)، در غیر این صورت، ممکن است توسط برخی از ارائه دهندگان سختگیر تعلیق شوید. اگر با وضعیتی مواجه شدید که تعداد IP های موجود در طول HTTP کاهش می یابد یا حتی 0 می شود، اما پس از مدتی بازیابی می شود، ممکن است به دلیل**مکانیسم های محدودیت موقت**توسط ISP یا Cloudflare CDN فعال می شود و فعالیت شما را به عنوان اسکن شبکه درک می کند. در این حالت معمولا پس از مدتی بهبود می یابد. توصیه می شود به**کاهش همزمانی**(`-n`) برای کاهش وقوع چنین شرایطی.

> علاوه بر این، HTTPing این نرم افزار فقط بازیابی می شود**سرصفحه های پاسخ**و محتوای بدنه را بازیابی نمی کند (یعنی اندازه فایل URL روی تست HTTP تاثیر نمی گذارد، اما اگر می خواهید تست سرعت دانلود را نیز انجام دهید، به یک فایل بزرگ نیاز دارید). این شبیه به عملکرد curl -i است.

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

#### #منطقه مشخص شده را مطابقت دهید (کد سه نویسه فرودگاه کولو)

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

> دو کد سه کاراکتری فرودگاه CDN مشترک هستند، بنابراین نام هر منطقه قابل مشاهده است:<https://www.cloudflarestatus.com/>

</details>

* * *

#### #مسیر نسبی/مطلق فایل

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

#### #تست سرعت سایر پورت ها

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

#### #آدرس تست سرعت سفارشی

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

#### #شرایط تست سرعت سفارشی (محدوده هدف تاخیر/از دست دادن بسته/سرعت دانلود را مشخص کنید)

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

> توجه: نوار پیشرفت تست سرعت تاخیر در سمت راست**مقدار موجود**، فقط به فرآیند اندازه گیری سرعت تاخیری اشاره دارد**تعداد IP هایی که زمان آن تمام نشده است**، بدون توجه به تاخیر شرایط کران بالا و پایین.

-   فقط مشخص کنید**[سقف تاخیر متوسط]**وضعیت

```bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s
# 即找到平均延迟低于 200 ms 的 IP，然后再按延迟从低到高进行 10 次下载测速
CloudflareST.exe -tl 200
```

> اگر**تاخیر رضایت بخشی پیدا نشد**IP مشروط، هیچ چیزی خروجی نخواهد شد.

* * *

-   فقط مشخص کنید**[سقف تاخیر متوسط]**شرایط، و**فقط تست سرعت را به تاخیر بیندازید، تست سرعت را دانلود نکنید**

```bash
# 平均延迟上限：200 ms，下载速度下限：0 MB/s，数量：不知道多少 个
# 即只输出低于 200ms 的 IP，且不再下载测速（因为不再下载测速，所以 -dn 参数就无效了）
CloudflareST.exe -tl 200 -dd
```

-   فقط مشخص کنید**[حداکثر احتمال از دست دادن بسته]**وضعیت

```bash
# 丢包几率上限：0.25
# 即找到丢包率低于等于 0.25 的 IP，范围 0.00~1.00，如果 -tlr 0 则代表过滤掉任何丢包的 IP
CloudflareST.exe -tlr 0.25
```

* * *

-   فقط مشخص کنید**[حداقل سرعت دانلود]**وضعیت

```bash
# 平均延迟上限：9999 ms，下载速度下限：5 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 9999 ms 且下载速度高于 5 MB/s 的 IP 才会停止测速
CloudflareST.exe -sl 5 -dn 10
```

> اگر**موردی را پیدا نکردم که سرعتش را داشته باشد**IP مشروط، سپس آن را انجام خواهد داد**شرایط را نادیده بگیرید و تمام نتایج تست سرعت IP را خروجی بگیرید**(این به شما کمک می کند دفعه بعد که سرعت خود را آزمایش می کنید شرایط را تنظیم کنید).

> زمانی که حد بالایی میانگین تاخیر مشخص نشده باشد، اگر بوده است**کافی نیست**تعداد IP هایی که شرایط را برآورده می کنند**همیشه در حال اندازه گیری سرعت**برو پایین.  
> پس توصیه می شود**در همان زمان مشخص کنید[حداقل سرعت دانلود]+[سقف تاخیر متوسط]**، به طوری که تست سرعت قبل از رسیدن به حد بالایی تاخیر مشخص شده قبل از جمع آوری تعداد کافی خاتمه می یابد.

* * *

-   در همان زمان مشخص کنید**[سقف تاخیر متوسط]+[حداقل سرعت دانلود]**وضعیت

```bash
# 平均延迟上限、下载速度下限均支持小数（如 -sl 0.5）
# 平均延迟上限：200 ms，下载速度下限：5.6 MB/s，数量：10 个（可选）
# 即需要找到 10 个平均延迟低于 200 ms 且下载速度高于 5 .6MB/s 的 IP 才会停止测速
CloudflareST.exe -tl 200 -sl 5.6 -dn 10
```

> اگر**تاخیر رضایت بخشی پیدا نشد**IP مشروط، هیچ چیزی خروجی نخواهد شد.  
> اگر**موردی را پیدا نکردم که سرعتش را داشته باشد**IP شرط، سپس تمام نتایج تست سرعت IP بدون در نظر گرفتن شرایط خروجی خواهد شد (برای شما راحت است که دفعه بعد سرعت را آزمایش کنید شرایط را تنظیم کنید).  
> بنابراین توصیه می شود ابتدا سرعت را بدون تعیین شرایط تست کنید تا ببینید میانگین تاخیر و سرعت دانلود چقدر است و از تعیین شرایط خودداری کنید.**خیلی کم/خیلی زیاد**！

> زیرا محدوده IP در معرض نمایش Cloudflare است**IP بازگشت به مبدا + IP Anycast**، و**بازگشت به IP منبع**در دسترس نیست، بنابراین تست سرعت دانلود 0.00 است.  
> می توانید آن را در زمان اجرا اضافه کنید`-sl 0.01`(حد پایین سرعت دانلود)، فیلتر کردن**بازگشت به IP منبع**(نتایج تست سرعت دانلود زیر 0.01 مگابایت بر ثانیه).

</details>

* * *

#### #سرعت یک یا چند آی پی را به صورت جداگانه تست کنید

<details>
<summary><code><strong>「 点击展开 查看内容 」</strong></code></summary>

* * *

**روش یک**:
مستقیماً داده های بخش IP را برای اندازه گیری از طریق پارامترها مشخص کنید.

```bash
# 先进入 CloudflareST 所在目录，然后运行：
# Windows 系统（在 CMD 中运行）
CloudflareST.exe -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32

# Linux 系统
./CloudflareST -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
```

* * *

**روش 2**:
یا این IP ها را در هر فایل متنی با فرمت زیر بنویسید، به عنوان مثال:`1.txt`

    1.1.1.1
    1.1.1.200
    1.0.0.1/24
    2606:4700::/32

> برای یک IP می توانید آن را حذف کنید`/32`زیر شبکه پوشانده شده (یعنی`1.1.1.1`معادل با`1.1.1.1/32`）。  
> پوشش زیر شبکه`/24`به آخرین سگمنت این IP یعنی`1.0.0.1~1.0.0.255`。

سپس هنگام اجرای CloudflareST پارامترهای راه اندازی را اضافه کنید`-f 1.txt`برای تعیین فایل داده بخش IP.

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

#### #سرعت تمام وب سایت ها را با استفاده از Cloudflare CDN یکبار برای همیشه (بدون نیاز به اضافه کردن نام دامنه به هاست ها یک به یک)

قبلا هم گفته بودم که هدف از توسعه این پروژه نرم افزاری پاس است**برای افزایش سرعت دسترسی به وب سایت ها با استفاده از Cloudflare CDN، روش Hosts را تغییر دهید**。

اما درست مثل[**#8**](https://github.com/Ptechgithub/CloudflareScanner/issues/8)همانطور که گفته شد، افزودن نام دامنه به هاست ها یک به یک در واقع می باشد**خیلی دردسرسازه**، بنابراین من یکی را پیدا کردم**یک بار برای همیشه**مسیر! می توانید این را تماشا کنید[**هنوز هاست ها را یکی یکی اضافه می کنید؟ روش شتاب دهی محلی عالی برای همه وب سایت هایی که از Cloudflare CDN استفاده می کنند اینجاست!**](https://github.com/Ptechgithub/CloudflareScanner/discussions/71)و یکی دیگر[برای تغییر IP رزولوشن نام دامنه به یک IP منتخب، به سرویس DNS محلی تکیه کنید](https://github.com/Ptechgithub/CloudflareScanner/discussions/317)آموزش

* * *

## تدوین دستی

<details>
<summary><code><strong>「 Click to expand to view content 」</strong></code></summary>

* * *

برای راحتی کار، هنگام کامپایل، شماره نسخه را در متغیر نسخه در کد می نویسم، بنابراین وقتی به صورت دستی کامپایل می کنید، باید آن را به صورت زیر اضافه کنید:`go build`بعد از دستور اضافه کنید`-ldflags`پارامتر برای تعیین شماره نسخه:

```bash
go build -ldflags "-s -w -X main.version=v2.3.3"
# 在 CloudflareScanner 目录中通过命令行（例如 CMD、Bat 脚本）运行该命令，即可编译一个可在和当前设备同样系统、位数、架构的环境下运行的二进制程序（Go 会自动检测你的系统位数、架构）且版本号为 v2.3.3
```

اگر می خواهید تحت سیستم ویندوز 64 بیتی کامپایل کنید**سایر سیستم ها، معماری ها، بیت ها**، سپس باید مشخص کنید**GOOS**و**GOARCH**متغیر.

به عنوان مثال، در سیستم ویندوز، یک برنامه مناسب را کامپایل کنید**سیستم لینوکس و معماری 64 بیتی**برنامه باینری برای:

```bat
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w -X main.version=v2.3.3"
```

به عنوان مثال، در سیستم لینوکس، یک برنامه مناسب برای**سیستم ویندوز AMD معماری 32 بیتی**برنامه باینری برای:

```bash
GOOS=windows
GOARCH=386
go build -ldflags "-s -w -X main.version=v2.3.3"
```

> می تواند اجرا کند`go tool dist list`برای دیدن اینکه نسخه فعلی Go از کامپایل کدام ترکیب ها پشتیبانی می کند.

* * *

البته، به منظور تسهیل در کامپایل دسته ای، من به طور خاص یک متغیر را به عنوان شماره نسخه مشخص می کنم و کامپایل های بعدی می توانند مستقیماً متغیر شماره نسخه را فراخوانی کنند.  
در عین حال، اگر می خواهید به صورت دسته ای کامپایل کنید، باید آنها را در پوشه های مختلف قرار دهید (یا نام فایل ها متفاوت است) باید اضافه کنید.`-o`پارامترهای مشخص شده

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

## مجوز

مجوز GPL-3.0.
