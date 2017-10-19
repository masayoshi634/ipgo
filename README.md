Ipgo is like ip command.

## Requirements
- Linux

## Installation


### Build from source
```bash
 $ go get github.com/masayoshi634/ipgo
 $ make
 $ make install
```

## Usage
```
$ ip addr show lo
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 brd 127.255.255.255 scope global lo
       valid_lft forever preferred_lft forever
    inet 127.0.0.2/8 brd 127.255.255.255 scope global secondary lo
       valid_lft forever preferred_lft forever

$ ipgo addr show | jq 'map(select(.["Label"]=="lo"))' | jq .
[
  {
    "IP": "127.0.0.1",
    "Label": "lo",
    "Flags": 128,
    "Scope": 0,
    "Broadcast": "127.255.255.255",
    "PreferedLft": 4294967295,
    "ValidLft": 4294967295,
    "Mask": 8,
    "Peer": {
      "IP": "127.0.0.1",
      "Mask": 8
    }
  },
  {
    "IP": "127.0.0.2",
    "Label": "lo",
    "Flags": 129,
    "Scope": 0,
    "Broadcast": "127.255.255.255",
    "PreferedLft": 4294967295,
    "ValidLft": 4294967295,
    "Mask": 8,
    "Peer": {
      "IP": "127.0.0.2",
      "Mask": 8
    }
  }
]
```
