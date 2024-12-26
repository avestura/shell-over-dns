# Shell over DNS

Tiny experiment to respond to DNS queries by shell executing them on host.

> [!CAUTION]
> This is obviously not secure. Do not use it.

## Example:

```
$> dig "@localhost" -p 9953 "hostname"

; <<>> DiG 9.16.22 <<>> @localhost -p 9953 hostname
; (2 servers found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 9825
;; flags: qr rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;hostname.                      IN      A

;; ANSWER SECTION:
shell.                  3600    IN      TXT     "VMForTest\013\010"

;; Query time: 22 msec
;; SERVER: 127.0.0.1#9953(127.0.0.1)
;; WHEN: Thu Dec 26 23:17:32 Iran Standard Time 2024
;; MSG SIZE  rcvd: 53
```