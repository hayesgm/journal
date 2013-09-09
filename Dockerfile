# Journal
#
# Run a Journal mirror

FROM centos
MAINTAINER Geoffrey Hayes <hayesgm@gmail.com>

RUN mkdir -p /srv/journal
RUN wget https://github.com/hayesgm/journal/releases/download/v0.0.1pre/journal.linux -O /srv/journal/journal
RUN chmod +x /srv/journal/journal

ENTRYPOINT ["/srv/journal/journal"]

EXPOSE 80:2222