CREATE DATABASE IF NOT EXISTS mah_jong;
CREATE DATABASE IF NOT EXISTS mah_jong_test;

CREATE USER IF NOT EXISTS 'app'@'%' IDENTIFIED BY 'hoge';
GRANT ALL ON *.* TO 'app'@'%';
