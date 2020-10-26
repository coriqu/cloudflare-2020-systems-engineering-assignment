---
typora-root-url: ./screenshot/coriq-links-10.png
---

# Systems Assignment

## Introduction

- **About**: This CLI tool is a follow-on to the [General Assignment](https://github.com/cloudflare-hiring/cloudflare-2020-general-engineering-assignment) I completed first, which can handle the HTTP request by taking an argument of the URL and an argument of the number of requests.
- **Language**: Go

- **Build tool:** Makefile

## Instructions to run this program

Build program

```
$ cd my-app
$ make all
```

Run program

```
./hello --url=(the url) --profile=(the number of requests)
//run the URL to my General Assignment
./hello --url=https://proud-moon-53cf.coriq.workers.dev/links --profile=1000
```

Use help for more details

```
$ ./hello --help
```

## Test Results

1. My site links

   ![](/../coriq-links-10.png)

   ![](/../coriq-links-100.png)

   ![coriq-links-1000](/../coriq-links-1000.png)

2. My site

   ![](/../coriq-1000.png)

   ![coriq-10000](/../coriq-10000.png)

3. Cloudflare

   ![](/../cloudflare-10.png)

   ![cloudflare-100](/../cloudflare-100.png)

   ![cloudflare-1000](/../cloudflare-1000.png)

4. Google

   ![](/../google-10.png)

   ![google-100](/../google-100.png)

   ![google-1000](/../google-1000.png)

5. Instagram

   ![](/../instagram-10.png)

   ![instagram-100](/../instagram-100.png)

   ![instagram-1000](/../instagram-1000.png)

## Findings

1. The requests with error codes returned always happens from the **245**th request. The error is:

```
Get "<url>": dial tcp: lookup <url>: no such host
```

2. The size in bytes of the response is **fixed** when the request URL is a **static page.** 

3. The more requests I make, the fastest time is faster, the slowest time is slower, the mean time is faster and the median is slower.

   

