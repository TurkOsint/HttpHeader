<p align="center">
  <img src="https://github.com/IntelSights/About/blob/main/img/IntelSights.png?raw=true" alt="IntelSights" width="250" />
</p>

<h1 align="center">HttpHeader</h1>

HttpHeader is a tool that returns headers specified as HEAD in HTTP requests and presents them to the user.

<h2 align="left">Features:</h2>

- [x] Http/s Head 
- [x] P3P
- [x] Set-Cookie 
- [x] Content-Type
- [x] Date
- [x] Server
- [x] Accept-Ranges
- [x] X-Xss-Protection
- [x] X-Cache
- [x] Expires
- [x] Cache-Control
- [x] Content-Encoding / Length
- [x] Alt-Svc
- [x] Etag:
- [x] Last-Modified


<h2 align="left">Installing:</h2>

```
git clone https://github.com/IntelSights/HttpHeader.git
```

<h2 align="left">Using:</h2>

```
cd HttpHeader/

./HttpHeader -g https://example.com/
```

<h2 align="left">Example:</h2>

```
cd HttpHeader/

./HttpHeader -g https://example.com

Output;
Accept-Ranges: bytes
Age: 431381
Date: Sun, 28 Mar 2021 22:32:51 GMT
Server: ECS (dcb/7EA2)
X-Cache: HIT
Content-Length: 648
Content-Encoding: gzip
Cache-Control: max-age=604800
Content-Type: text/html; charset=UTF-8
Etag: "3147526947"
Expires: Sun, 04 Apr 2021 22:32:51 GMT
Last-Modified: Thu, 17 Oct 2019 07:18:26 GMT

```


<h2 align="left">To-Do:</h2>

- [x] Infrastructure Will Be Changed
- [x] Modular Building Will Be Brought
- [x] Arguments Will Be Replaced

<h2 align="left">Developers:</h2>

[@EyupErgin](https://github.com/eyupergin)


