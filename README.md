# zipwhip-bulk-message-send

A .csv driven bulk message send console application.

### Input Parameters

- **session:** The `session` for the account being sent from.
- **fileName:** The name of the file in the same directory as the application.
- **threads:** The number of 'threads' to use for delivery, defaults to 5.


### Install Project

```
go get github.com/acapps/zipwhip-bulk-message-send
```

### Build Project

```
cd $GOPATH/src/github.com/acapps/zipwhip-bulk-message-send
go build

```

### Sample Request

```
./zipwhip-bulk-message-send \
    -session='{{session}}' \
    -fileName=sample.csv`
```

### Sample Output

The output running the application provides the response from sending the message.
It also includes two time stamps; the first, when the request to send began and the second, when the request finished.
The final output value is the time in milliseconds to send the message.

```
2017/06/21 17:59:55 Starting to Send Messages:
2017/06/21 17:59:58 '{"response":{"id":877692606781038592,"status":"queued"},"success":true}','1498093195124660372','1498093198112592589','2987'
2017/06/21 17:59:58 '{"response":{"id":877692607346868224,"status":"queued"},"success":true}','1498093195124719283','1498093198252136443','3127'
2017/06/21 17:59:58 '{"response":{"id":877692607896186880,"status":"queued"},"success":true}','1498093195124649792','1498093198390950306','3266'
2017/06/21 17:59:58 '{"response":{"id":877692608500301824,"status":"queued"},"success":true}','1498093195124618311','1498093198545974822','3421'
2017/06/21 17:59:58 '{"response":{"id":877692608978317312,"status":"queued"},"success":true}','1498093198112766691','1498093198638354765','525'
2017/06/21 17:59:58 '{"response":{"id":877692609137836032,"status":"queued"},"success":true}','1498093195124644598','1498093198678437765','3553'
2017/06/21 17:59:58 '{"response":{"id":877692609200750592,"status":"queued"},"success":true}','1498093198252265303','1498093198691247751','438'
2017/06/21 17:59:58 '{"response":{"id":877692609209139200,"status":"queued"},"success":true}','1498093198391162285','1498093198694567943','303'
2017/06/21 17:59:58 Finished Sending Messages!
```