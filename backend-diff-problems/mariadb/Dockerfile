# 使用するDockerイメージ
FROM mariadb:10.7.3-focal

# ポート番号
EXPOSE 3308

# MySQL設定ファイルをイメージ内にコピー
ADD ./my.cnf /etc/mysql/conf.d/my.cnf

# docker runに実行される
CMD ["mysqld"]
