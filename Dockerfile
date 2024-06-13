FROM ubuntu:20.04

RUN apt-get update && apt-get install -y wget lsof git build-essential uuid-dev bpfcc-tools

RUN wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz

# Install Golang
# f: this must be the last flag of the command, and the tar file must be immediately after. It tells tar the name and path of the compressed file.
# z: tells tar to decompress the archive using gzip
# x: tar can collect files or extract them. x does the latter.
# v: makes tar talk a lot. Verbose output shows you all the files being extracted.
RUN tar -xvzf go1.22.4.linux-amd64.tar.gz -C /usr/local
# ln -s /path/to/file /path/to/symlink
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

RUN mkdir /tmp/21

# Install valgrind-demo
RUN git clone https://github.com/evgeniy-scherbina/valgrind
WORKDIR /valgrind
# RUN CGO_LDFLAGS="-luuid" GOARCH=arm64 CGO_ENABLED=1 go install .
# RUN ./build.sh
# ln -s /path/to/file /path/to/symlink
# RUN ln -s /root/go/bin/linux_arm64/cgo-tutorial-v3 /usr/local/bin/cgo-tutorial-3

CMD sleep infinity
