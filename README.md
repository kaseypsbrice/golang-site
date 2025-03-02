Start the webserver:
```
go run ./main.go
```

To find what's running on a specific port:
```
lsof -i :<port_number>
```

Kill the process running on a port:
```
kill -9 <PID>
```
