FROM golang
LABEL maintainer="ariel.gerardo.rios@gmail.com"

RUN ["mkdir", "-p", "/go/src/github.com/ariel17/xy/api"]
WORKDIR "/go/src/github.com/ariel17/xy/api"
COPY . .

RUN ["go", "get", "-u", "github.com/golang/dep/cmd/dep"]
RUN ["dep", "ensure"]

EXPOSE 8080
ENTRYPOINT ["go", "run", "main.go", "8080"]
