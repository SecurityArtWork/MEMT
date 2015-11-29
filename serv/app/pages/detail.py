# -*- coding: utf-8 -*-
"""Module that controls the details view.
"""
from __future__ import absolute_import

import os

from flask import current_app as app
from flask import Blueprint
from flask import render_template
from flask import abort
from flask import flash

from app.extensions import mongo

from app.common import get_img_to_b64

bp = Blueprint('detail', __name__, url_prefix='/detail')


@bp.route('/<sha256:hash>', methods=['GET'])
def index(hash):
    assets = mongo.db.assets
    malwares = assets.find({"sha256": hash})
    obj = {}
    if malwares:
        for malware in malwares:
            print(malware["ssdeep"])

            # This loop is intended to run just once
            obj["ssdeep"] = malware["ssdeep"]
            obj["md5"] = malware["md5"]
            obj["sha1"] = malware["sha1"]
            obj["sha256"] = malware["sha256"]
            obj["sha256"] = malware["sha256"]
            obj["sha512"] = malware["sha512"]
            obj["strain"] = malware["strain"]
            obj["strain"] = malware["strain"]
            obj["format"] = malware["format"]
            obj["format"] = malware["format"]
            obj["symbols"] = malware["symbols"]
            obj["symbols"] = malware["symbols"]
            obj["imports"] = malware["imports"]
            obj["imports"] = malware["imports"]
            obj["sections"] = malware["sections"]
            obj["mutations"] = malware["mutations"]
            obj["image"] = get_img_to_b64(malware["imageDir"])
            obj["arch"] = malware["arch"]
            obj["arch"] = malware["arch"]
            obj["ipmeta"] = malware["ipMeta"]
            del obj["ipmeta"][0]["ip"]
            return render_template('detail/index.html', info=obj)
    else:
        flash("", "")
        return abort(404)
