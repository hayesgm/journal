# Journal
#
# Run a Journal mirror

FROM centos
MAINTAINER Geoffrey Hayes <hayesgm@gmail.com>

RUN git clone https://github.com/hayesgm/journal /srv/journal && cd /srv/journal

ENTRYPOINT ["/srv/journal.linux"]

EXPOSE 80:2222