# reshare-service

[![.github/workflows/ci.yml](https://github.com/jhu-library-applications/reshare-service/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/jhu-library-applications/reshare-service/actions/workflows/ci.yml)

This is a webservice that checks ReShare's VuFind instance for 
a loanable copy of an item using an ISN or a title/author search. 

If a loanable copy is found the URL for that item is returned. 

# Running

1. Download and install go:

`https://golang.google.cn/dl/`

2. Checkout the code from GitHub

3. Build and run the server:

`go build . && ./reshare-service`

## Examples

http://localhost:5050/request?title=becoming&author=Michelle%20obama

http://localhost:5050/request?isn=1524763144

http://localhost:5050/request?title=becoming&author=Michelle%20obama&isn=123
