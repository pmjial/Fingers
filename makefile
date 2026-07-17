##
# Fingers
#
# @file
# @version 0.1

TARGET = Fingers.exe

all: $(TARGET)
	./$(TARGET)

$(TARGET): main.go db/db.go
	go build

# end
