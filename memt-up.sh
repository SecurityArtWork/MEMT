#!/bin/bash

CWD="$MEMT_VIRTUAL_CWD"

rabbitmq=$(which rabbitmq-server)
mongodb=$(which mongod)
celery="celery worker -A %s --beat %s"
gunicorn="gunicorn --worker-class eventlet --workers %s --bind 127.0.0.1:5000 %s wsgi"

args=`getopt abo: $*`
# you should not use `getopt abo: "$@"` since that would parse
# the arguments differently from what the set command below does.
if [ $? != 0 ]
then
       echo 'Usage: ...'
       exit 2
fi
set -- $args
# You cannot use the set command with a backquoted getopt directly,
# since the exit code from getopt would be shadowed by those of set,
# which is zero by definition.
for i
do
       case "$i"
       in
               -a|-b)
                       echo flag $i set; sflags="${i#-} $sflags";
                       shift;;
               -o)
                       echo oarg is "'"$2"'"; oarg="$2"; shift;
                       shift;;
               --)
                       shift; break;;
       esac
done
echo single-char flags: "'"$sflags"'"
echo oarg is "'"$oarg"'"
