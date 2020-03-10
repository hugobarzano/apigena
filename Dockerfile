# Dockerfile2 from https://github.com/hugobarzano/LProductosSoftware.git
FROM    httpd:2.4  
MAINTAINER    hugobarzano  
RUN    sed 's/^Listen 80/Listen 6666/g' /usr/local/apache2/conf/httpd.conf > httpd.new  
RUN    mv httpd.new /usr/local/apache2/conf/httpd.conf  
COPY    html/ /usr/local/apache2/htdocs/  
EXPOSE    6666  
