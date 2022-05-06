# 使用するDockerイメージ
FROM mysql:latest

# MySQL設定ファイルをイメージ内にコピー
ADD ./my.cnf /etc/mysql/conf.d/my.cnf

# docker runに実行される
CMD ["mysqld"]
