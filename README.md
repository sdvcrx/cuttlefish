# cuttlefish proxy

A simple HTTP forward proxy that support multiple parent proxies.

[![Build Status](https://travis-ci.org/sdvcrx/cuttlefish.svg?branch=master)](https://travis-ci.org/sdvcrx/cuttlefish)
[![Drone Build Status](https://drone.sdvcrx.com/api/badges/sdvcrx/cuttlefish/status.svg?branch=master)](https://drone.sdvcrx.com/sdvcrx/cuttlefish)

## Features

* Full HTTP proxy implementation, including HTTPS through `CONNECT`
* Proxy basic authentication support
* Redirect request to multiple parent proxies (like squid)
* toml configuration

## Install

Download binary file from [release](https://github.com/sdvcrx/cuttlefish/releases).

Or compile it by yourself:

``` shell
go get github.com/sdvcrx/cuttlefish
```

## TODO

[TODOs.md](https://github.com/sdvcrx/cuttlefish/blob/master/TODOs.md)

## LICENSE

MIT
