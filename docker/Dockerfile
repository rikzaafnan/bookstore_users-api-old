#################################
# STEP 1 build executable binary
#################################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

# # 2. Create appuser.
# ENV USER=appuser
# ENV UID=10001
# # See https://stackoverflow.com/a/55757473/12429735RUN
# RUN adduser \
#     --disabled-password \
#     --gecos "" \
#     --home "/nonexistent" \
#     --shell "/sbin/nologin" \
#     --no-create-home \
#     --uid "${UID}" \
#     "${USER}"

# 3. grab deps and build
WORKDIR $GOPATH/src/bookstore-users-api/
COPY ../go.mod .
#COPY ../go.sum .
RUN go mod download
RUN go mod verify

COPY .. .
# disable debug, compile to linux only, and disable cross compile
# https://stackoverflow.com/questions/52640304/standard-init-linux-go190-exec-user-process-caused-no-such-file-or-directory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /app/bookstore-users-api

#################################
# STEP 1 build scratch image
#################################
#FROM scratch
FROM alpine

#RUN apk add curl
#RUN curl -L https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz | tar -xJ
#RUN apk update && apk add --no-cache wkhtmltopdf xvfb fontconfig freetype ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family
#RUN apk update freetype freetype-dev ttf-freefont
#COPY --from=wkhtmltopdf_image /bin/wkhtmltopdf /usr/bin/

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/bookstore-users-api /app/bookstore-users-api
#ADD ../ca-certificates.crt /app/
# ADD ../template /template

# # Use an unprivileged user.
# USER appuser:appuser

# Logging
ARG LOG_DIR=/app
#RUN mkdir -p ${LOG_DIR}
ENV LOG_FILE_LOCATION=${LOG_DIR}/bookstore-users-api.log

# Env variables as needed
ARG enviro=xxx
ARG apm=xxx
ENV bookstore_users_api_ENV=$enviro


ENTRYPOINT ["/app/bookstore-users-api"]
