﻿﻿<h1 align="center">go-wake-server</h1><div align="center">

[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![GitHub license](https://img.shields.io/github/license/LegendaryB/go-wake-server.svg?longCache=true&style=flat-square)](https://github.com/LegendaryB/go-wake-server/blob/master/LICENSE.md)

Simple http server which sends a wake on lan packet to the specific mac address.
<br>
<br>
<sub>Built with ❤︎ by Daniel Belz</sub>
</div><br>

## Getting started

### Command-line arguments
The application can be configured via command-line arguments. The table shows all possible values. You can also show all command-line arguments when starting the application with the `-h` flag.

|Argument|Description|Default|
|---|---|---|
|port|The port on which the application should listen for http requests.|81|
|use-static-mac|Flag to indicate if the static 'mac-addr' value should be used when the http resource is hit.|false|
|mac-addr|MAC address which is used in case the 'use-static-mac' flag is set to true.|none|
|broadcast-addr|Address to which the generated magic packet will be send.|255.255.255.255|

### Command-line arguments usage
