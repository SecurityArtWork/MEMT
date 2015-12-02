# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import geoip2.database
import base64

from flask import current_app as app

from app.extensions import mongo

rt_feed_namespace = "/feed"
rt_map_namespace = "/rtmap"
celery_namespace = "/celery"

messages = {
    "feed_title": "New analysis",
    "feed_msg": "",
    "feed_title_strain": "Strain Found!",
    "feed_msg_strain": "New malware found!! Checkout this link and get more info."
}


def get_common_info():
    """This function has to be called from inside a  request, beacause it uses a app object.
    """
    assets = mongo.db.assets
    total_assets = assets.count()
    total_strains = assets.find({"strain": ""}).count()
    last_countries = get_latest_geo(app.config["RT_LAST_COUNTRIES"])
    last_news = get_latest_feeds(app.config["FEED_LAST_NEWS"])
    return {"total_assets": total_assets,
            "total_strains": total_strains,
            "last_countries": last_countries,
            "last_news": last_news}


def get_latest_feeds(limit=5):
    last_news = []
    feeds = mongo.db.feed
    query = feeds.find().sort([("$natural", -1)]).limit(limit)
    for last_new in query:
        last_news.append(last_new)
    return last_news

def get_latest_geo(limit=100):
    last_countries = []
    assets = mongo.db.assets
    query = assets.find({"ipmeta.country": {"$ne": "unknown"}},
                        {"ipmeta.iso_code": 1,
                         "ipmeta.country": 1,
                         "ipmeta.city": 1,
                         "ipmeta.geo": 1}).sort([("$natural", -1)])\
                                          .limit(limit)
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

def get_img_to_b64(img):
    print("IMG LOC: {}".format(img))
    try:
        with open(img, "rb") as image_file:
            encoded_string = base64.b64encode(image_file.read())
    except (IOError):
        return "data:image/png;base64,"
    return "data:image/png;base64," + encoded_string
