#!/bin/bash

curl \
    -X GET \
    -H "traceparent: 00-a0892f3577b34da6a3ce929d0e0e4736-f03067aa0ba902b7-01" \
    http://localhost:8080/
