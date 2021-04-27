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

### Configuration
The application can be configured via the conf.json file. A typical conf.json looks like this:

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

### Waking a machine
To wake a machine you only need to send a GET request to the http endpoint. On linux you could use `wget` or `curl`.

`wget 127.0.0.1:81/wake/00:80:41:ae:fd:7e`

**Status codes**
|Status code|Message|
|---|---|
|200|none|
|400|Error message of gowol err|
