includes = \
	-I$(PWD)

main:
	g++ -o sigtool3 sigmd5.c rsa_enc.c rsa_dec.c
	./sigtool3 9D09CE1D31DBD03100AB41E80E35C9C2
	./sigtool3 deeb6e2059484e58242481572a6d60d1

clean:
	rm sigtool3 *.o