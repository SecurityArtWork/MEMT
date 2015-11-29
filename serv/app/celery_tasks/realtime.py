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
    feeds_collection = mongo.db.feeds
    now = datetime.utcnow()
    from_ = now - timedelta(minutes=now.minute % 5 + app.config["FEED_REFRESH"], seconds=now.second, microseconds=now.microsecond)
    feeds = feeds_collection.find({"date": {"$gte": from_}})
    socketio.emit("update", dumps(feeds), namespace=rt_feed_namespace)

@celery.task(name="memt.rt.map", bind=True)
def rt_feed(self):
    #socketio.emit("", namespace=rt_map_namespace)
    pass

