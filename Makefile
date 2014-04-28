BIN_DIR=bin
CC=clang

all: bind

bind: main.go
	CC=${CC} go build -o ${BIN_DIR}/main main.go

clean:
	rm -f ${BIN_DIR}/main