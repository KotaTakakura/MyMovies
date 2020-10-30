#!/bin/bash
migrate -database "mysql://${DBUser}:${DBPass}@${DBProtocol}/mypipe" -path ./Migrations up
/go/src/MyPIPE/MyPIPE