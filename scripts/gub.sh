#!/usr/bin/env bash

# ------------------------------------------------------------------------------
# gub
#
#   Go utility belt
#
#   Usage:
#     gub setGopath  # Sets the #GOPATH
#     gub setup      # Setup dep and fetch dependencies
#     gub install    # Install the Go project
#     gub run        # Run the Go project
# ------------------------------------------------------------------------------

source $PROJECT_DIR/.env

#######################################
# Sets the #GOPATH
# Globals:
#    $PROJECT_DIR
# Arguments:
#    None
# Returns:
#    None
#######################################
function setGopath {
    export GOPATH=$PROJECT_DIR/services/receipts
}

#######################################
# Setup dep and fetch dependencies
# Globals:
#    $PROJECT_DIR
# Arguments:
#    None
# Returns:
#    None
#######################################
function setup {
    setGopath
    if [ ! -d $PROJECT_DIR/services/receipts/bin ]; then
        mkdir $PROJECT_DIR/services/receipts/bin
    fi

    if [ ! -f $PROJECT_DIR/services/receipts/bin/dep ]; then
        curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
    fi
    cd $PROJECT_DIR/services/receipts/src/receipts
    $PROJECT_DIR/services/receipts/bin/dep ensure
}

#######################################
# Install the Go receipts project
# Globals:
#    $PROJECT_DIR
# Arguments:
#    None
# Returns:
#    None
#######################################
function install {
    source $PROJECT_DIR/.env
    setGopath
    go install receipts
}

#######################################
# Run the Go receipts project
# Globals:
#    $PROJECT_DIR
# Arguments:
#    None
# Returns:
#    None
#######################################
function run {
    install
    echo $DB_HOST
    $PROJECT_DIR/services/receipts/bin/receipts
}

function gub {
    case "$1" in
    "set_path")
        setGopath
        ;;
    "install")
        install
        ;;
    "run")
        run
        ;;
    "setup")
        setup
        ;;
    *)
        echo "Default"
        ;;
    esac
}
