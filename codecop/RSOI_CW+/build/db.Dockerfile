FROM ubuntu
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && apt-get install -y postgresql
USER postgres

COPY build/scripts.sql scripts.sql

RUN service postgresql start &&\
    psql -U postgres -c "ALTER USER postgres PASSWORD 'password';" &&\
    psql -U postgres -c 'CREATE DATABASE "rsoi";' &&\
    psql -U postgres -d rsoi -a -f scripts.sql &&\
    service postgresql stop

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root
EXPOSE 5432
CMD service postgresql start && tail -f /dev/null
