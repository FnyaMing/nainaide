#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/nainaided/${BINARY:-nainaided}
ID=${ID:-0}
LOG=${LOG:-nainaided.log}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'nainaided' E.g.: -e BINARY=nainaided_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Run binary with all parameters
##
export NAINAIDEDHOME="/nainaided/node${ID}/nainaided"

if [ -d "`dirname ${NAINAIDEDHOME}/${LOG}`" ]; then
  "$BINARY" --home "$NAINAIDEDHOME" "$@" | tee "${NAINAIDEDHOME}/${LOG}"
else
  "$BINARY" --home "$NAINAIDEDHOME" "$@"
fi

chmod 777 -R /nainaided

