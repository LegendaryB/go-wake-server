﻿﻿<h1 align="center">go-wake-server</h1><div align="center">

[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![GitHub license](https://img.shields.io/github/license/LegendaryB/go-wake-server.svg?longCache=true&style=flat-square)](https://github.com/LegendaryB/go-wake-server/blob/master/LICENSE.md)

Simple HTTP to Wake-on-LAN bridge.
<br>
<br>
<sub>Built with ❤︎ by Daniel Belz</sub>
</div><br>

## Getting started

### Configuration
The application can be configured via the conf.json file. A typical configuration looks like this:

```json
{
    "port": "81",
    "allow_any_mac": false,
    "mac_address": "38:ea:a7:a1:07:5b",
    "broadcast": {
        "address": "255.255.255.255",
        "port": "9"
    }
}
```
**Note**: When `allow_any_mac` is set to true the application allows GET requests to `yourip:yourport/wake/yourMAC` and `yourip:yourport/wake/`. 

When set to false only GET requests to `yourip:yourport/wake/` are accepted. The value from property `mac_address` is used.

**Wake-on-LAN to MAC defined in conf.json**

`curl 127.0.0.1:81/wake/`

**Wake-on-LAN with custom MAC**

`curl 127.0.0.1:81/wake/00:80:41:ae:fd:7e`
