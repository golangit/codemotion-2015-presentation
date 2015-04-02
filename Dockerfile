
FROM google/golang:stable
# Godep for vendoring
RUN go get github.com/tools/godep
# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

MAINTAINER liuggio@gmail.com
ENV APP_DIR $GOPATH/src/github.com/golangit/codemotion-2015-presentation
 
# Set the entrypoint 
ENTRYPOINT ["/opt/app/codemotion-2015-presentation"]
ADD . $APP_DIR

# Compile the binary and statically link
RUN mkdir /opt/app
RUN cd $APP_DIR && godep restore
RUN cd $APP_DIR && CGO_ENABLED=0 go build -o /opt/app/codemotion-2015-presentation -ldflags '-d -w -s'

EXPOSE 80
