# event-checksum
1.) Clone repo locally

2.) `cd cmd`

3.) `go build`

4.) There are several ways to execute the program. 
    
    4a.) To generate a checksum for an event without a binary value, simply call `./cmd`.
    4b.) To generate a checksum for an event with a binary value, call `./cmd -b`
    4c.) To generate an event with a binary value and POST it to a local instance of EdgeXFoundry, call `./cmd -b -p`

5.) Do it repeatedly and you should see that the checksum is different each time.
