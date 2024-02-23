.PHONY: pb
pb: pb-clear
	buf generate

.PHONY: pb-clear
pb-clear: 
	rm -f *.pb.go