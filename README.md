# generateMac
MAC Address generating tool with ability to specify OUI and generate 1000 entries at once


| Flag | Description |
| ------ | ------ |
| -h | Show this help |
| -p | Specify the OUI (first six hex) separated by colon (:) or dash (-) |
| -n | Number of entries to generate (up to 1000) |

# Example Input
```sh
./generateMac -n 10
```

# Example Output
```sh
28:99:c7:5f:43:7e
28:99:c7:6c:87:ac
28:99:c7:56:db:b0
28:99:c7:c4:9f:fd
28:99:c7:43:4b:b8
28:99:c7:38:f8:99
28:99:c7:ad:0b:8f
28:99:c7:8f:f1:87
28:99:c7:3a:e4:8c
28:99:c7:be:7a:26
```

# Example Input
```sh
./generateMac -p 54-8a-ba
```

# Example Output
```sh
54:8a:ba:aa:f4:f6
```

# Example Input
```sh
./generateMac -n 4 -p 54:8a:ba
```

# Example Output
```sh
54:8a:ba:4d:89:4e
54:8a:ba:88:75:28
54:8a:ba:6a:96:3e
54:8a:ba:26:22:1c
```
