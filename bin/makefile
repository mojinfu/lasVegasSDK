RM= rm -rf

all: clean release

release: lasVegas

clean:
	$(RM) ../log/* 
	$(RM)  lasVegas

lasVegas:
	go build -ldflags "-s -w" -o $@ ../src/main/*.go
