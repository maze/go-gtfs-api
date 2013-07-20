#!/bin/bash
curl -XDELETE 'http://localhost:9200/gtfs_agency';
curl -XPUT 'http://localhost:9200/gtfs_agency' -d @agency_settings.json
curl -XPUT 'http://localhost:9200/_river/gtfs_agency/_meta' -d @agency_river.json
