all: fork hello sleep ping

fork: fork.c
	gcc -static fork.c -o fork

hello: hello.c
	gcc -static hello.c -o hello

sleep: sleep.c
	gcc -static sleep.c -o sleep

ping: ping.c
	gcc -static ping.c -o ping

clean:
	rm fork hello sleep ping
