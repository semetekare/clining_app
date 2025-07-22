#!/bin/sh

chmod -R a+rx /clining

python3 manage.py collectstatic --no-input --clear
python3 manage.py makemigrations
python3 manage.py migrate

exec "$@"
