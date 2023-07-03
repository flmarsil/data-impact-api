# mongo:5.0.6
FROM mongo@sha256:9ae745b709512a09a8c105959f75bde3d8a25c3215842a8251c073e14cd2a04d

COPY ./config/mongo-init.json /mongo-init.json

CMD mongoimport                                 \
    # Host (docker-compose service name)
    --host db                                   \
    # user and pass from .env file
    --username $MONGO_INITDB_ROOT_USERNAME      \
    --password $MONGO_INITDB_ROOT_PASSWORD      \
    --authenticationDatabase admin              \
    # initial db and collection from .env file
    --db $MONGO_INITDB_DATABASE                 \
    --collection $MONGO_INITDB_COLLECTION       \
    --type json                                 \
    --file /mongo-init.json                     \
    --jsonArray                                 \
    # WARNING! drop collection before run import
    --drop                                      \
    # verbose
    -v                                          \