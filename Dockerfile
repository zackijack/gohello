# Build stage
FROM golang:1.13-buster as build

LABEL app="gohello"
LABEL REPO="https://github.com/zackijack/gohello"

ADD . /code/gohello

WORKDIR /code/gohello

RUN make build

# Final stage
FROM zackijack/debian-base-image:buster

LABEL maintainer="m.zackky@gmail.com"

ARG GIT_COMMIT
ARG VERSION

LABEL app="gohello"
LABEL REPO="https://github.com/zackijack/gohello"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

ENV PATH=$PATH:/opt/gohello/bin

WORKDIR /opt/gohello/bin

COPY --from=build /code/gohello/bin/gohello /opt/gohello/bin

RUN chmod +x /opt/gohello/bin/gohello

# Create appuser
RUN adduser --disabled-password --gecos '' gohello
USER gohello

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/gohello/bin/gohello"]
