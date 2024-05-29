FROM debian:trixie-slim

RUN apt-get update &&\
	apt-get install -y yt-dlp &&\
	apt-get autoremove &&\
	rm -rf /var/lib/{apt,dpkg,cache,log}/

WORKDIR /var/www/

COPY ./main ./main
COPY ./static ./static

ENTRYPOINT ./main
