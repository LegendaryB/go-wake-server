﻿﻿<h1 align="center">go-wake-server</h1><div align="center">

[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![GitHub license](https://img.shields.io/github/license/LegendaryB/go-wake-server.svg?longCache=true&style=flat-square)](https://github.com/LegendaryB/go-wake-server/blob/master/LICENSE.md)

Simple HTTP to Wake-on-LAN bridge.
<br>
<br>
<sub>Built with ❤︎ by Daniel Belz</sub>
</div><br>

## Configuration
The application can be configured via the conf.json file. A typical configuration looks like this:

```json
{
    "port": "81",
    "mac_regex_pattern": "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$",
    "broadcast": {
        "address": "255.255.255.255",
        "port": "9"
    }
}
```

|Property|Description|
|---|---|
|port|Port on which the application should listen for requests.|
|mac_regex_pattern|The regex pattern is used to check if the given request parameter is valid.|
|broadcast.address|The broadcast address to which the magic packet should be send.|
|broadcast.port|The broadcast port on which the magic packet should be send.|

### Command-line arguments
At the moment you can only specify a custom configuration file via the command-line options:

`sudo ./go-wake-server -c myconf.json`

## Sending a magic packet
You can trigger sending a Wake-on-LAN packet by sending a GET request to the http endpoint.
On Unix based systems you can use `wget` or `curl` for that purpose.

**curl**

`curl 127.0.0.1:81/wake/00:80:41:ae:fd:7e`

**wget**

`wget 127.0.0.1:81/wake/00:80:41:ae:fd:7e`
