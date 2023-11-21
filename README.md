# strace-from-scratch
As seen at Gophercon 2017. [Here's a walkthrough of this code](https://medium.com/@lizrice/strace-in-60-lines-of-go-b4b76e3ecd64) and [here's the slide deck](https://speakerdeck.com/lizrice/a-go-programmers-guide-to-syscalls). 

## Seccomp demo

Comment out the disallow() at main.go line 18 to demonstrate seccomp.

```
make docker && docker run -it --rm strace tail -n1 LICENSE
```

Potentially useful links:

Intercepting and Emulating Linux System Calls with Ptrace: https://nullprogram.com/blog/2018/06/23/

https://blog.rchapman.org/posts/Linux_System_Call_Table_for_x86_64/

https://refspecs.linuxfoundation.org/elf/x86_64-abi-0.99.pdf

https://www.juliensobczak.com/inspect/2021/08/10/linux-system-calls-under-the-hood.html

https://dr-knz.net/go-calling-convention-x86-64.html

https://poonai.github.io/posts/how-debuggers-works-part1/

https://medium.com/@lizrice/a-debugger-from-scratch-part-1-7f55417bc85f

https://www.pixelstech.net/article/1633245055-The-Go-Pointer-Magic



