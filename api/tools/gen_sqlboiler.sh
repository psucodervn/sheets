#!/usr/bin/env bash

sqlboiler -c tools/sqlboiler/conf.toml \
    --templates tools/sqlboiler/templates \
    --templates tools/sqlboiler/templates_test \
    --templates tools/sqlboiler/own_templates \
     psql
#    --templates tools/sqlboiler/override_psql \
#    --templates tools/sqlboiler/override_psql_test \
