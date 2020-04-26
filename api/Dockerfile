ARG GO_VERSION=1.13

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /src

WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -o /app .

# Final stage: the running container.
FROM scratch AS final

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

ENTRYPOINT ["/app"]