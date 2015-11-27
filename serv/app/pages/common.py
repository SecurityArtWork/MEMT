# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import geoip2.database

from flask import current_app as app

from app.extensions import mongo


def get_common_info():
    assets = mongo.db.assets
    feeds = mongo.db.feeds
    total_assets = assets.count()
    total_strains = assets.find({"strain": ""}).count()
    last_countries = []
    query = assets.find({}, {"ipMeta.iso_code": 1, "ipMeta.country": 1, "ipMeta.city": 1, "ipMeta.geo": 1}).limit(app.config["RT_LAST_COUNTRIES"])
    for last_country in query:
        last_countries.append(last_country)
    last_news = []
    query = feeds.find().sort([("$natural", 1)]).limit(app.config["FEED_LAST_NEWS"])
    for last_new in query:
        last_news.append(last_new)
    return {"total_assets": total_assets,
            "total_strains": total_strains,
            "last_countries": last_countries,
            "last_news": last_news}


def get_geo(addr):
    reader = geoip2.database.Reader(app.config["MAXMAIN_DB_CITIES"])
    response = reader.city(addr)
    lat = response.location.latitude
    lon = response.location.longitude
    city = response.city.name
    country = response.country.name
    iso_code = response.country.iso_code
    response.close()
    return {"lat": lat,
            "long": lon,
            "city": city,
            "country": country,
            "iso_code": iso_code}
