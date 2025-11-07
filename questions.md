## List of Questions

1. does io.read means like pop?
  yes, because it is a stream
2. what is readAll, why it can be dangerous? what is the alternative
  because it can overload the memory. the ReadAll stores the data in memory directly, so you have to be careful. The alternative is you can store it first to tmp file, and process it further.
3. why interface closer is important. to free? what to free, why some implementation doesn't have closer
4. what is io.copy and why is it important
5. what is bytes.buffer
   it is a struct that has stream, buffer. Unlike `bytes.Reader`, it has `write` capability
6. how to now it is has closer or not
   There are multiple interface in `io` packages. so you have to read the doc or see the func to know if it implements the specific interface
7. relation between []bytes, buffer, bytes.buffer
   `[]bytes` is a data type that store data in memory.
   `buffer` is a concept in computer science for a temporary storage or memory.
   `bytes.buffer` is a struct form `bytes` that implement `buffer` concept while providing `stream` capability.
8. why using buffer bufio, it is because performance? how difference it is?
9. why buffer is used with flush?
10. what is io.pipe
12. what does this mean? "stream data without buffering to memory"
  It means that the stream is like a conveyor, it doesn't stored in memory, only consume it one by one as the data flows. It is different with the buffer where the buffer stores it to memory
  for faster access but can be dangerous if the memory is overloaded.
13. what is syscall?
14. why copying this?
```go
// handler
defer r.Body.Close()
out, _ := os.Create("out.dat")
defer out.Close()
io.Copy(out, r.Body)
```
  `io.Copy` is a function to copy bytes to `writer`. In this case, it copy to a file, means it saving it to local.

14. Why using `bytes.NewReader(buf)` instead of directly using `buf`?
    `var []byte buf` can be used but you have to access it directly. so passing it to a new reader, can make use these interfaces `[io.Reader], [io.ReaderAt], [io.WriterTo], [io.Seeker],[io.ByteScanner], and [io.RuneScanner]` for more capabilities.
    bytes.NewReader(buf) is both a buffer and a stream.
    Buffer → because the data lives in memory.
    Stream → because it implements io.Reader (and optionally Seeker) so libraries can read it as a “file”.

15. What is the difference between `bytes.Buffer` and `bytes.Reader`?
* **`bytes.Buffer`** = memory + read/write + stream
* **`bytes.Reader`** = memory + read-only stream
* Both implement `io.Reader` → can be used wherever Go expects a **stream**.
* Buffers are especially useful when:

  1. You need to **accumulate data** dynamically
  2. You need a **re-readable stream**
  3. You need both **io.Reader** + **io.Writer**

## golang exercise