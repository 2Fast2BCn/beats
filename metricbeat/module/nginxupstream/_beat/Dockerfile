FROM centos:centos7

MAINTAINER "Kobe Nys" <kobe.nys@allnice.be>

ENV NGINX_VERSION         1.12.1

# Install packages
RUN yum update -y
RUN yum install gcc gcc-c++ make zlib-devel pcre-devel openssl-devel zip wget tar hostname git patch curl -y && yum clean all

RUN useradd --no-create-home nginx

RUN cd /tmp && git clone https://github.com/2Fast2BCn/nginx_upstream_check_module.git

# Download and build nginx
RUN cd /tmp && wget -q -O - http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz | tar zxf -
RUN cd /tmp/nginx-${NGINX_VERSION} \
        && patch -p0 < /tmp/nginx_upstream_check_module/check_1.11.5+.patch \
        && ./configure --prefix=/etc/nginx/ --sbin-path=/usr/sbin/nginx --add-module=/tmp/nginx_upstream_check_module --pid-path=/var/run/nginx.pid \
        && make install

# Cleanup
RUN rm -rf /tmp/nginx-${NGINX_VERSION}
RUN rm -rf /tmp/nginx_upstream_check_module

EXPOSE 80

COPY ./nginx.conf /etc/nginx/

WORKDIR /etc/nginx/

CMD ["/usr/sbin/nginx"]

