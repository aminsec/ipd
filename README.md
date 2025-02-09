# IP-Details

Install:

```
go install -v github.com/aminsec/ipd@latest
```

Sample usage:
```
echo 131.123.11.1 | ipd
```

Result: 
```
{
  "IP": "131.123.11.1",
  "ASN": {
    "number": 11050,
    "name": "KENT-STATE",
    "desc": "Kent State University"
  },
  "CIDR": {
    "range": "131.123.0.0/16",
    "desc": "Kent State University"
  }
}
```
