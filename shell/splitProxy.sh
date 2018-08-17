#!/bin/bash
pattern='(https?://)?((([^:\/]+)(:([^\/]*))?@)?([^:\/?]+)(:([0-9]+))?)'

[ -n "$http_proxy" ] && HTTPPROXY=$http_proxy
[ -n "$HTTP_PROXY" ] && HTTPPROXY=$HTTP_PROXY
[ -n "$https_proxy" ] && HTTPSPROXY=$https_proxy
[ -n "$HTTPS_PROXY" ] && HTTPSPROXY=$HTTPS_PROXY

if [ -n "$HTTPPROXY" ]; then
	if [[ "$HTTPPROXY" =~ $pattern ]]; then
		[ -n "${BASH_REMATCH[4]}" ] && JAVA_OPTS="$JAVA_OPTS -Dhttp.proxyUser=${BASH_REMATCH[4]}"
		[ -n "${BASH_REMATCH[6]}" ] && JAVA_OPTS="$JAVA_OPTS -Dhttp.proxyPassword=${BASH_REMATCH[6]}"
		[ -n "${BASH_REMATCH[7]}" ] && JAVA_OPTS="$JAVA_OPTS -Dhttp.proxyHost=${BASH_REMATCH[7]}"
		[ -n "${BASH_REMATCH[9]}" ] && JAVA_OPTS="$JAVA_OPTS -Dhttp.proxyPort=${BASH_REMATCH[9]}"
	fi
fi
if [ -n "$HTTPSPROXY" ]; then
	if [[ "$HTTPSPROXY" =~ $pattern ]]; then
		[ -n "${BASH_REMATCH[4]}" ] && JAVA_OPTS="$JAVA_OPTS -Dhttps.proxyUser=${BASH_REMATCH[4]}"
		[ -n "${BASH_REMATCH[6]}" ] && JAVA_OPTS="$JAVA_OPTS -Dhttps.proxyPassword=${BASH_REMATCH[6]}"
		[ -n "${BASH_REMATCH[7]}" ] && JAVA_OPTS="$JAVA_OPTS -Dhttps.proxyHost=${BASH_REMATCH[7]}"
		[ -n "${BASH_REMATCH[9]}" ] && JAVA_OPTS="$JAVA_OPTS -Dhttps.proxyPort=${BASH_REMATCH[9]}"
	fi
fi

echo $JAVA_OPTS