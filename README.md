# GoPortscanner
This is a simple TCP Port scanner written in Go-lang.

# Long description
This code scans the specified target for the specified ports. For each port, it attempts to connect using net.DialTimeout, with a timeout of one second. If the connection is successful, the port is considered "open" and the program prints a message indicating that. If the connection fails, the port is considered "closed" and the program continues scanning other ports.

Note that this code only scans TCP ports, but you can modify it to scan UDP ports as well by changing the net.DialTimeout to net.DialUDPTimeout. Also note that this code only scans a small number of ports - for a full port scan, you would need to scan all 65,535 available ports, which could take a while.

# usage
``go run scanner.go`` (and of course, you do need to have Go-lang installed on your computer in $PATH in order to be to run this.)
