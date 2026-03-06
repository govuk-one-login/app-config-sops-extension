FROM golang:1.25-alpine AS build
WORKDIR /build
COPY go.mod go.sum ./
COPY main.go .

RUN go build -tags lambda.norpc -o main main.go

FROM public.ecr.aws/lambda/provided:al2023

WORKDIR /lambda

#RUN curl -LO https://github.com/getsops/sops/releases/download/v3.11.0/sops-v3.11.0.linux.arm64 \
# && mv sops-v3.11.0.linux.arm64 /usr/local/bin/sops \
# && chmod +x /usr/local/bin/sops

COPY --from=build /build/main /lambda/main

ENTRYPOINT [ "/lambda/main" ]