FROM mysql:5.7

WORKDIR /
COPY . .
EXPOSE 3306

RUN mkdir /var/log/mysql
RUN touch /var/log/mysql/mysql-error.log
RUN touch /var/log/mysql/mysql-slow.log
RUN touch /var/log/mysql/mysql-query.log

CMD ["mysqld"]