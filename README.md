﻿﻿<h1 align="center">go-wake-server</h1><div align="center">

[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![GitHub license](https://img.shields.io/github/license/LegendaryB/go-wake-server.svg?longCache=true&style=flat-square)](https://github.com/LegendaryB/go-wake-server/blob/master/LICENSE.md)

Simple http server which sends a wake on lan packet to the specific mac address.
<br>
<br>
<sub>Built with ❤︎ by Daniel Belz</sub>
</div><br>

## Command-line arguments
The application can be configured via command-line arguments. The following table should be self explaining.

### port (default: 81)
The port on which the application should listen for http requests.

### use-static-mac (default: false)
Flag to indicate if the static 'mac-addr' value should be used when the http resource is hit.

### mac-addr (default: none)
MAC address which is used in case the 'use-static-mac' flag is set to true.

### broadcast addr (default: 255.255.255.255)
Address to which the generated magic packet will be send.

## Command-line samples
