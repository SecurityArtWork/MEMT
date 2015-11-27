# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import geoip2.database

from flask import current_app as app

from app.extensions import mongo


def get_common_info():
    """This function has to be called from inside a  request, beacause it uses a app object.
    """
    assets = mongo.db.assets
    feeds = mongo.db.feeds
    total_assets = assets.count()
    total_strains = assets.find({"strain": ""}).count()
    last_countries = get_latest_geo(app.config["RT_LAST_COUNTRIES"])
    last_news = get_latest_feeds(app.config["FEED_LAST_NEWS"])
    return {"total_assets": total_assets,
            "total_strains": total_strains,
            "last_countries": last_countries,
            "last_news": last_news}


def get_latest_feeds(limit=5):
    feeds = mongo.db.feeds
    last_news = []
    query = feeds.find().sort([("$natural", 1)]).limit(limit)
    for last_new in query:
        last_news.append(last_new)
    return last_news

def get_latest_geo(limit=100):
    assets = mongo.db.assets
    last_countries = []
    query = assets.find({}, {"ipMeta.iso_code": 1,
                             "ipMeta.country": 1,
                             "ipMeta.city": 1,
                             "ipMeta.geo": 1}).limit(limit)
    for last_country in query:
        last_countries.append(last_country)
    return last_countries

def get_geo_from_ip(addr):
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
