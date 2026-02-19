FROM golang AS build

WORKDIR /go/src/app
COPY . .

ENV DEBIAN_FRONTEND=noninteractive

# RUN apt-get update && \
#     apt-get install -y \
#     tesseract-ocr \
#     libtesseract-dev \
#     build-essential \
#     gcc \
#     libmupdf-dev \
#     tesseract-ocr-deu && \
#     rm -rf /var/lib/apt/lists/*

RUN go mod download
RUN cd ./cmd/lqmonitor && go vet -v
RUN cd ./cmd/lqmonitor && go test -v

RUN CGO_ENABLED=1 go build -o /go/bin/lqmonitor ./cmd/lqmonitor/lqmonitor.go

FROM gcr.io/distroless/cc-debian13

LABEL org.opencontainers.image.source=https://github.com/ashcycling/qrc-gen

COPY --from=build /go/bin/lqmonitor /
CMD ["/lqmonitor"]