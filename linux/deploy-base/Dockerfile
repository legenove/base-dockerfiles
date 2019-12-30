FROM amazonlinux:1

RUN yum -y install libxml2-devel curl-devel libpng-devel gd-devel autoconf zlib-devel gcc make openssl-devel unzip crontabs procps vi

RUN yum update -y -q --security && \
        yum install -y -q wget \
        yum install -y -q git \
        yum install -y -q gzip \
        yum install -y -q tar.x86_64 \
        yum install -y -q make \
        yum clean all

RUN curl https://bootstrap.pypa.io/get-pip.py | python
RUN pip install supervisor==3.3.5