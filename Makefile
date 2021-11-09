SRC = examples/main.go
TARGET = examples/main.out
DEBUG = -gcflags "-N -l" -race

.PHONY:clean

debug: $(SRC)
	go build -o $(TARGET) $(DEBUG) $^

build: $(SRC)
	go build -o $(TARGET) $^

clean:
	rm examples/*.out
