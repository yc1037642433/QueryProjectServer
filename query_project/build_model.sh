#!/bin/bash
goctl model mysql ddl -src queryuser.sql -dir .. -style go_zero
