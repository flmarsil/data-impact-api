# FROM golang:1.17-alpine3.15 AS build
FROM golang@sha256:4918412049183afe42f1ecaf8f5c2a88917c2eab153ce5ecf4bf2d55c1507b74 AS build

# Create appuser.
# ENV USER=appuser
# ENV UID=10001 

# # See https://stackoverflow.com/a/55757473/12429735RUN 
# RUN apk update \
#     && adduser \    
#     --disabled-password \    
#     --gecos "" \    
#     --home "/nonexistent" \    
#     --shell "/sbin/nologin" \    
#     --no-create-home \    
#     --uid "${UID}" \    
#     "${USER}" 

WORKDIR /go/src/app

COPY . ./

# Get dependancies
RUN go mod download \
    && go mod verify

# Remove debug informations and compile only for linux target and disabling cross compilation 
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/webserver ./cmd/main.go

# FROM alpine:3.15 
FROM alpine@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300

# Import the user and group files from the builder.
# COPY --from=build /etc/passwd /etc/passwd
# COPY --from=build /etc/group /etc/group

# Import go binary
COPY --from=build /go/src/app/bin /go/bin

# Use an unprivileged user.
# USER appuser:appuser

CMD ["/go/bin/webserver"]

# https://docs.docker.com/engine/reference/commandline/pull/#pull-an-image-by-digest-immutable-identifier