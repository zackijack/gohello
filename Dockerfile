FROM debian:buster-slim

ENV TZ=Asia/Jakarta

RUN apt-get update \
    && apt-get install -y \
        bash \
        ca-certificates \
        dumb-init \
        openssl \
        tzdata

RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && \
	echo $TZ > /etc/timezone

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/pushm0v/gohello"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/gohello

WORKDIR /opt/gohello

COPY ./bin/gohello /opt/gohello/

RUN chmod +x /opt/gohello/gohello

# Create appuser
RUN adduser --disabled-password --gecos '' gohello
USER gohello

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/gohello/gohello"]
