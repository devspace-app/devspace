FROM ubuntu:20.04

RUN addgroup --system xusers \
  && adduser \
			--home /home/xuser \
			--disabled-password \
			--shell /bin/bash \
			--ingroup xusers \
			--quiet \
			xuser

RUN apt-get update && apt-get install -y --no-install-recommends \
				xvfb \
				xauth \
				x11vnc \
				x11-utils \
				x11-xserver-utils \
		&& rm -rf /var/lib/apt/lists/*

RUN mkdir -p /data/Xauthority && chown -R xuser:xusers /data/Xauthority
VOLUME /data/Xauthority

# start x11vnc and expose its port
ENV DISPLAY :0.0
EXPOSE 5900
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

USER xuser
ENTRYPOINT ["/entrypoint.sh"]