FROM golang:1.13-alpine as builder

ADD ../ $GOPATH/src/github.com/soxueren/greenplum-node

WORKDIR $GOPATH/src/github.com/soxueren/greenplum-node

RUN CGO_ENABLED=0 GOOS=linux go build -o gpdb-node cmd/main.go

ROM centos:7

RUN yum install -y wget

RUN wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo

RUN mkdir /init

# greenplum-db-6.0.0-rhel7-x86_64.rpm Dependent packages
RUN yum install -y apr apr-util bzip2 krb5-devel libedit libevent libyaml openssh openssl perl rsync zip net-tools less openssh-clients which iproute openssh-server 
# other Dependent packages
#RUN yum -y install unzip ed cmake gcc gcc-c++ make automake autoconf libtool coreutils glib2 lrzsz sysstat e4fsprogs xfsprogs  readline-devel zlib-devel

RUN groupadd gpadmin && \
    useradd gpadmin -g gpadmin && \       
    echo gpadmin | passwd --stdin gpadmin

ADD init/*.rpm /init

RUN rpm -ihv /init/greenplum-db-6.0.0-rhel7-x86_64.rpm

COPY --from=0 /go/src/github.com/soxueren/greenplum-node/gpdb-node  .

RUN chmod 0777 /usr/local

ADD init/*.sh /init

RUN chmod +x /init/*.sh && \
    /init/pre_install.sh

RUN yum clean all && rm -rf /init

USER gpadmin

RUN  ssh-keygen -t rsa -P ''  &&  ssh-keygen -t rsa -f .ssh/id_rsa -P '' 

CMD ["sleep","3600"]
