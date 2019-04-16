#!/bin/bash

SERVER="gapp"
CMD_DIR=$PWD

function start()
{
    if [ "`pgrep ${SERVER}`" != "" ]
    then
        echo "${SERVER} already running"
        exit 1
    fi

	nohup ${CMD_DIR}/${SERVER}  &>"./var/log/nohup.$SERVER.log" &

    if [ "`pgrep ${SERVER}`" == "" ]
    then
        echo "${SERVER} start failed."
        exit 1
    fi
}



function stop()
{
    if [ "`pgrep ${SERVER}  -u ${UID}`" != "" ]
    then
		kill -9 `pgrep $SERVER -u $UID`
        exit 1
    fi

    if [[ "`pgrep ${SERVER} -u ${UID}`" != "" ]]
    then
        echo "${SERVER} stop failed"
        exit 1
    fi
}

function status()
{
	if [[ "`pgrep ${SERVER} -u ${UID}`" != "" ]];then
		echo "${SERVER} is running"
	else
		echo ${SERVER} is not running
	fi
}


case "$1" in
    'start')
        start
        ;;
    'stop')
        stop
        ;;
    'restart')
        stop && start
        ;;
    'status')
        status
        ;;
    *)
        echo "usage: $0 {start|stop|restart|status}"
        exit 1
    ;;
esac