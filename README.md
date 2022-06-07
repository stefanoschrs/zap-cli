# zap-cli

> Pretty print [uber-go/zap](https://github.com/uber-go/zap) logs

## Installation
```bash
make
ln -s $(pwd)/dist/zap-cli /usr/local/bin/zap-cli
```

## Usage / Examples
1.  
    ```bash
    $ go run ./examples/test.go
    {"level":"\u001b[35mDEBUG\u001b[0m","time":"2022-06-07T08:19:31+03:00","message":"Failed to fetch URL","url":"https://github.com/stefanoschrs/zap-cli","attempt":3}
    {"level":"\u001b[34mINFO\u001b[0m","time":"2022-06-07T08:19:31+03:00","message":"Failed to fetch URL"}
    ```
    ```bash
    $ go run ./examples/test.go | zap-cli 
    2022-06-07T08:19:33+03:00 DEBUG	Failed to fetch URL	{"attempt":3,"url":"https://github.com/stefanoschrs/zap-cli"}
    2022-06-07T08:19:33+03:00 INFO	Failed to fetch URL
    ```
2.  
    ```bash
    $ cat examples/test.log 
    {"level":"\u001b[35mDEBUG\u001b[0m","time":"2022-06-06T14:00:24Z","message":"Processing","CountryId":"1","CountryName":"United States"}
    {"level":"\u001b[35mDEBUG\u001b[0m","time":"2022-06-06T14:00:24Z","message":"Page 0"}
    {"level":"\u001b[35mDEBUG\u001b[0m","time":"2022-06-06T14:00:24Z","message":"fetching data"}
    {"level":"\u001b[35mDEBUG\u001b[0m","time":"2022-06-06T14:00:26Z","message":"Found 50 entries"}
    {"level":"\u001b[31mERROR\u001b[0m","time":"2022-06-06T14:00:26Z","message":"no entry was added"}
    ```
    ```bash
    $ cat examples/test.log | zap-cli
    2022-06-06T14:00:24Z DEBUG	Processing	{"CountryId":"1","CountryName":"United States"}
    2022-06-06T14:00:24Z DEBUG	Page 0
    2022-06-06T14:00:24Z DEBUG	fetching data
    2022-06-06T14:00:26Z DEBUG	Found 50 entries
    2022-06-06T14:00:26Z ERROR	no entry was added
    ```

## TODO
- [ ] Use a config/flags to map the base fields (level, time, message)
