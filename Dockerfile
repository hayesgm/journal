# Journal
#
# Run a Journal mirror

FROM centos
MAINTAINER Geoffrey Hayes <hayesgm@gmail.com>

RUN yum install git-core -y

RUN git clone https://github.com/hayesgm/journal /srv/journal && cd /srv/journal

CMD cd /srv/journal && ./journal.linux

EXPOSE 80:2222