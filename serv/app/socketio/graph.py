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
