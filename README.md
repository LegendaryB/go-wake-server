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
    "broadcast": {
        "address": "255.255.255.255",
        "port": "9"
    }
}
```

## Sending a magic packet via http
You can trigger sending a Wake-on-LAN packet by sending a GET request to the http endpoint.
On Unix based systems you can use `wget` or `curl` for that purpose.

**curl**

`curl 127.0.0.1:81/wake/00:80:41:ae:fd:7e`

**wget**

`wget 127.0.0.1:81/wake/00:80:41:ae:fd:7e`
