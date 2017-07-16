FROM golang
LABEL maintainer="ariel.gerardo.rios@gmail.com"

RUN ["mkdir", "-p", "/srv/xy"]
COPY ["./api", "/srv/xy"]

EXPOSE 8080

WORKDIR "/srv/xy"
ENTRYPOINT ["./bin/start-dev.sh"]
