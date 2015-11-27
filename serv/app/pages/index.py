# -*- coding: utf-8 -*-
"""Module that controls the index view.
"""
from __future__ import absolute_import

from flask import current_app as app
from flask import Blueprint
from flask import render_template


from app.extensions import mongo


bp = Blueprint('index', __name__)


@bp.route('/', methods=['GET'])
def index():
    assets = mongo.db.assets
    feeds = mongo.db.feeds
    total_assets = assets.count()
    total_strain = assets.find({"strain": ""}).count()
    last_countries = []
    for last_country in assets.find().limit(app.config["RT_LAST_COUNTRIES"]):
        last_countries.append(last_country)
    last_news = []
    for last_new in feeds.find().sort([("$natural", 1)]).limit(app.config["FEED_LAST_NEWS"]):
        last_news.append(last_new)
    print(total_assets, total_strain, last_countries, last_news)
    return render_template('index/index.html',
                           totalAssets=total_assets,
                           totalStrain=total_strain,
                           lastNews=last_news,
                           lastCountries=last_countries)
