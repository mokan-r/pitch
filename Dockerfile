FROM golang

WORKDIR /project/

COPY server /project/server
COPY model /project/model
COPY db /project/db
COPY internal /project/internal
COPY go.mod /project

RUN go mod tidy

ENTRYPOINT ["go", "run", "github.com/mokan-r/pitch/server/cmd"]