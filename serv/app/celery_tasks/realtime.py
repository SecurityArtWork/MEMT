# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import requests
import json

from datetime import datetime
from datetime import timedelta

from bson.json_util import dumps
from flask import current_app as app

from app import socketio
from app import celery

from app.extensions import mongo

from app.common import rt_map_namespace
from app.common import rt_feed_namespace


@celery.task(name="memt.rt.feed", bind=True)
def rt_feed(self):
    feeds_collection = mongo.db.feed
    now = datetime.utcnow()
    from_ = now - timedelta(minutes=now.minute % 5 + app.config["FEED_REFRESH"], seconds=now.second, microseconds=now.microsecond)
    feeds = feeds_collection.find()
    print(dumps(feeds))
    socketio.emit("update", dumps(feeds), namespace=rt_feed_namespace)


@celery.task(name="memt.rt.map", bind=True)
def rt_map(self):
    assets_collection = mongo.db.assets
    now = datetime.utcnow()
    from_ = now - timedelta(minutes=now.minute % 5 + app.config["RTMAP_REFRESH"], seconds=now.second, microseconds=now.microsecond)
    assets = assets_collection.find({"date": {"$gte": from_}}, {"ipmeta.iso_code": 1, "ipmeta.city": 1, "ipmeta.country": 1, "ipmeta.geo": 1})
    socketio.emit("update", dumps(assets), namespace=rt_map_namespace)


