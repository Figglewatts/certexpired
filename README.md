# certexpired
Simple CLI utility for checking if TLS certificates are expired.

## Usage
`certexpired [options] ADDRESS...`

`certexpired` accepts TLS endpoints (`ADDRESS`, `host:port` i.e. `lsdrevamped.net:443`) that it will connect to 
and check the `NotAfter` field of the certificate to see if it is due to expire within a given threshold. 
The exit code of the program indicates whether any TLS endpoint certificates given are within the expiry threshold. 
If any have expired, the addresses will be printed (separated by `\n`).

### Exit codes
- `0`: Successful execution, no addresses certificates within expiry threshold.
- `1`: Successful execution, printed addresses certificates within expiry threshold.
- `2`: Usage error.
- `3`: General error.

### Options
```
Usage of certexpired:
  -threshold duration
    	certificate expiry threshold (default 720h0m0s)
  -verbose
    	verbose output
```

### Executable
#### Arguments
```
$ certexpired [options] ADDRESS...
i.e.
$ certexpired lsdrevamped.net:443
```
`ADDRESS` is an address in the form accepted by [`net.Dial`](https://pkg.go.dev/net#Dial).
#### Pipe
```
$ cat << EOF | bin/certexpired
lsdrevamped.net:443
EOF
```
You can also pipe `ADDRESS` inputs, one on each line.

### Docker
```
$ docker run ghcr.io/figglewatts/certexpired [options] ADDRESS...
```