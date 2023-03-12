# This A Very small tool To validate Domain
## It has A CLI to validate the Domain of an email address
```terminal
$ go run main.go
```
## Then Write The email 
```terminal
$ <EMAIL>@<DOMAIN>
```

## Example of The CLI Output
```terminal
gmail.com,true,true,v=spf1 redirect=_spf.google.com,true,v=DMARC1; p=none; sp=quarantine; 
rua=mailto:mailauth-reports@google.com
```

## You Can Also Use The "CheckDomain" function
```go
    str := CheckDomain("gmail.com")
```
## And you will get the same output as the CLI
```terminal
gmail.com,true,true,v=spf1 redirect=_spf.google.com,true,v=DMARC1; p=none; sp=quarantine; 
rua=mailto:mailauth-reports@google.com
```