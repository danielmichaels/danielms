+++
date = "2018-02-28"
categories = ["python"]
tags = ["scapy", "python"]
title = "Scapy Snippets"
slug = "Scapy-Snippets"
+++

### How to get RSSI from WLAN packet

```python
from scapy.layers.dot11 import RadioTap, Dot11

def get_rssi(packet):
    if packet.haslayer(RadioTap):
        return packet.dbm_antsignal
```

**caveat: currently only tested on Ralink: RT5370 chipset**
