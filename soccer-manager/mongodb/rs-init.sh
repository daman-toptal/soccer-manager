#!/bin/bash


mongo -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --quiet <<EOF
var config = {
    "_id": "rs0",
    "members": [
        {
            "_id": 1,
            "host": "127.0.0.1:27017",
        }
    ]
};
rs.initiate(config);
rs.status();
EOF