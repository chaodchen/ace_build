includes = \
	-I$(PWD)

main:
	g++ -o sigtool4 sigmd5.c rsa_enc.c rsa_dec.c
	./sigtool4 00000000000000000000000000000000
	./sigtool4 77777777777777777777777777777777
	./sigtool4 ffffffffffffffffffffffffffffffff

clean:
	rm sigtool4 *.o