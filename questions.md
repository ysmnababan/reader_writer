## List of Questions

1. does io.read means like pop?
  yes, because it is a stream
2. what is readAll, why it can be dangerous? what is the alternative
3. why interface closer is important. to free? what to free, why some implementation doesn't have closer
4. what is io.copy and why is it important
5. what is io.buffer
6. how to now it is has closer or not
7. relation between []bytes, buffer, bytes.buffer
8. why using buffer bufio, it is because performance? how difference it is?
9. why buffer is used with flush?
10. what is io.pipe
12. what does this mean? "stream data without buffering to memory"
13. what is syscall?
14. why copying this?
```go
// handler
defer r.Body.Close()
out, _ := os.Create("out.dat")
defer out.Close()
io.Copy(out, r.Body)
```


## golang exercise