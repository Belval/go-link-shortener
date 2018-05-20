# go-link-shortener
A basic link shortening service written in go. Please consider as a work-in-progress.

## How-To

1. Edit the config.json file
2. Run `./go-link-shortener -config=/path/to/your/config`
3. Shorten an url with `/api/v1/new/?url=[your.url.com]`
    - Don't forget to add & escape the https:// part if you really want it (otherwise it will default to http://)
4. Use the returned address

## A more permanent installation

If you want to run it as a service, `install_or_update.sh` will set up everything in a nice systemd service. Worry free!

1. Clone the project
2. cd go-link-shortener
3. sudo ./install_or_update.sh

## Benchmarks

On a single vCore machine with 1GB of RAM and behind a Nginx proxy using SSL (so a real scenario).

Using the Apache benchmarking tool:

`ab -n 1000 -c 50 https://belv.al/82dbd` (1000 requests, 50 concurrents):

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:       71  297  41.2    292     460
    Processing:    33   53   8.8     52      91
    Waiting:       33   52   8.6     51      87
    Total:        115  349  42.5    345     508

    Percentage of the requests served within a certain time (ms)
    50%    345
    75%    373
    100%   508 (longest request)

`ab -n 500 -c 10 https://belv.al/82dbd` (500 requests, 10 concurrents):

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:       63   89  10.8     88     138
    Processing:    19   34   8.6     32      70
    Waiting:       18   33   8.6     32      69
    Total:         96  123  13.1    123     178

    Percentage of the requests served within a certain time (ms)
    50%    123
    75%    131
    100%   178 (longest request)

I ran tests without Nginx and timings were below 70ms so it's quite fast.

## Requirements

- https://github.com/google/uuid
- https://github.com/mattn/go-sqlite3

