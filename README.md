# adjust
Tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

## requirements
* Golang 1.17

## Options
* `-parallel` defines the number of parallel jobs

## Args
Different http address, seperated by whitespace.

If address has no schema, `http` will be used by default 

## Local run example
``
go run main.go http://google.com
``

## Docker example 
```
> docker build -t myhttp .
> docker run myhttp calcmd5 www.google.com http://yandex.com
```