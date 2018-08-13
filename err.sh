#!/usr/bin/env bash

# ------------------------------------------------------------------------------
# sail
#
#   sail invoker
#
#   Usage:
#     ./sail funcName [Options]  # if invoked from the file
#     # Line to add to .bashrc if desired: source $PATH_TO_FILE/sail.sh --source-only
# ------------------------------------------------------------------------------

export PROJECT_DIR=$(dirname ${BASH_SOURCE[0]})
source $PROJECT_DIR/.env

function err {
    help_text=$(cat << EOF
Useful Email Read Reciepts functions

Usage: ./err <function> [options]

Functions:
$(
for f in $(find $PROJECT_DIR/scripts -type f -maxdepth 1); do
    func=$(basename $f)
    echo "  ${func%.*}"
done
)
EOF
)

    func="$1"
    args="${@:2}"
    code=0

    if [[ -z $func ]] || [[ $func == "help" ]]; then
        echo "$help_text"
    elif [[ -f "$PROJECT_DIR/scripts/$func.sh" ]]; then
        source "$PROJECT_DIR/scripts/$func.sh"
        $func $args
        unset -f $func
    elif [[ -f "$PROJECT_DIR/scripts/$func.py" ]]; then
        python3 "$PROJECT_DIR/scripts/$func.py" $args
    else
        echo "Unknown function: $func"
        code=1
    fi

    unset -f $PROJECT_DIR
    return $code
}

if [ "${1}" != "--source-only" ]; then
    err "${@}"
fi