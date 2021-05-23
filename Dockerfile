FROM golang
WORKDIR /
COPY . .
RUN go build -o /build/oapigenclient .

ENTRYPOINT [ "/build/oapigenclient" ]
EXPOSE 8080