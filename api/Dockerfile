FROM golang
LABEL maintainer="ariel.gerardo.rios@gmail.com"

RUN ["mkdir", "-p", "/srv/xy"]
COPY [".", "/srv/xy"]

EXPOSE 8080

WORKDIR "/srv/xy"
RUN ["./bin/requirements.sh"]

ENTRYPOINT ["go", "run", "main.go"]
