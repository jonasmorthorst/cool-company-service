ARG GO_VERSION=1.13

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

#ARG app_env
#ENV APP_ENV $app_env

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

# Expose both 443 and 80 to our application
#EXPOSE 443
EXPOSE 8080

# Mount the certificate cache directory as a volume, so it remains even after
# we deploy a new version
#VOLUME ["/cert-cache-test"]

# Run the compiled binary.

# RUN go get github.com/pilu/fresh

ENTRYPOINT ["/app"]

## before RUN go build -o main .

#RUN git clone https://github.com/vishnubob/wait-for-it.git

#RUN echo "$PWD"
#RUN echo "$LS"

#WORKDIR /app
#COPY wait-for-it /

# beforeRUN go get github.com/pilu/fresh

#CMD if [ ${APP_ENV} = production ]; \
#	then \
	#"/app/main" \
#	else \
#	go get github.com/pilu/fresh && \
	#fresh; \
#	fi

# before EXPOSE 8080

#before CMD ["/app/main"]
