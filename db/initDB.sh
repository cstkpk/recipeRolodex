#!/bin/bash
# Drop current DBs and create new schema
mysql < schema.sql
# Seed DBs
mysql < seedingredients.sql
mysql < seedrecipes.sql
mysql < seedlink.sql