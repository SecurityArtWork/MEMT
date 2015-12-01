# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import
import time

from bson.json_util import dumps

from flask import request, session
from flask_socketio import emit

from app import socketio

from app.common import get_latest_geo
from app.common import rt_map_namespace


@socketio.on('connect', namespace=rt_map_namespace)
def connect():
    data = get_latest_geo()
    emit("connect", dumps(data), namespace=rt_map_namespace)
    #rt_feed.delay()
    #keep_updating()


def background_thread():
    while True:
        time.sleep(10)
        data = get_latest_geo()
        emit('update', dumps(data), namespace=rt_map_namespace)

def keep_updating():
    global thread
    if thread is None:
        thread = Thread(target=background_thread)
        thread.daemon = True
        thread.start()
