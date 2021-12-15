# Echo

This is a simple Golang webserver which replies the current HTTP request including its headers.

I wrote this little app to test Microsoft Azure Application Proxy Header-based SSO.

## Usage

```
Usage of ./echo:
  -Help
        Print help
  -Port int
        Listen on this port (default 8090)
```

## AAD App Proxy Configuration

![AAD Config](aad-headersso-config.png)

## Sample Response

```
GET / HTTP/1.1 
[Headers]
Accept=[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9]
Accept-Encoding=[gzip, deflate, br]
Accept-Language=[de,en-US;q=0.9,en;q=0.8]
Cache-Control=[no-cache]
Connection=[Keep-Alive]
Dnt=[1]
Front-End-Https=[on]
Pragma=[no-cache]
Referer=[https://login.microsoftonline.com/]
Sec-Ch-Ua=[" Not A;Brand";v="99", "Chromium";v="96", "Microsoft Edge";v="96"]
Sec-Ch-Ua-Mobile=[?0]
Sec-Ch-Ua-Platform=["Windows"]
Sec-Fetch-Dest=[document]
Sec-Fetch-Mode=[navigate]
Sec-Fetch-Site=[cross-site]
Upgrade-Insecure-Requests=[1]
User-Agent=[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36 Edg/96.0.1054.43]

X-Auth-Displayname=[pri]
X-Auth-Email=[pri@*******.ch]
X-Auth-Groups=[DOMAIN\Z-ServerAdmin,DOMAIN\Z-CloudUsers,DOMAIN\Z-RDPAdmin,DOMAIN\Z-ClientAdmin]
X-Auth-Username=[pri]

X-Forwarded-For=[194.***.***.***]
X-Ms-Proxy=[AzureAD-Application-Proxy]
Host = "echo.******.ch:8090"
RemoteAddr= "***.***.***.***:62656"

```

# License
MIT License

Copyright (c) 2021 Philipp Ritter

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.