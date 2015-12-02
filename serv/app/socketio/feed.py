# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import
import time

from bson.json_util import dumps

from flask import request, session
from flask_socketio import emit

from app import socketio
from app.common import get_latest_feeds
from app.common import rt_feed_namespace


@socketio.on('connect', namespace=rt_feed_namespace)
def connect():
    print("FEED - Connected")
    data = get_latest_feeds()
    #emit("connect", dumps(data), namespace=rt_feed_namespace)
