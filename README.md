# strace-from-scratch
As seen at Gophercon 2017. [Here's a walkthrough of this code](https://medium.com/@lizrice/strace-in-60-lines-of-go-b4b76e3ecd64) and [here's the slide deck](https://speakerdeck.com/lizrice/a-go-programmers-guide-to-syscalls). 

## Seccomp demo

Comment out the disallow() at main.go line 18 to demonstrate seccomp.
