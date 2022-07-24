#!/bin/bash
UN_FMT_FILES="$(gofmt -l .)"
[[ -z $UN_FMT_FILES ]] || {
  echo "file un formated"
  echo $UN_FMT_FILES
}
