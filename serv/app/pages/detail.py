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

from config import BASEDIR

from app.extensions import mongo
from app.common import get_img_to_b64

bp = Blueprint('detail', __name__, url_prefix='/detail')


@bp.route('/<sha256:hash>', methods=['GET'])
def index(hash):
    assets = mongo.db.assets
    malwares = assets.find({"sha256": hash})
    if malwares.count():
        for malware in malwares:
            print(malware)
            # This loop is intended to run just once
            obj = get_object(malware)
            return render_template('detail/index.html', info=obj)
        flash("Internal error")
        return abort(500)
    else:
        flash("Not found, sorry", "")
        return abort(404)


def get_object(malware):
    print("PARSING MALWARE {}".malware)
    obj = {}
    obj["ssdeep"] = malware["ssdeep"]
    obj["md5"] = malware["md5"]
    obj["sha1"] = malware["sha1"]
    obj["sha256"] = malware["sha256"]
    obj["sha512"] = malware["sha512"]
    obj["format"] = malware["format"]
    obj["symbols"] = malware["symbols"]
    obj["imports"] = malware["imports"]
    obj["sections"] = malware["sections"]
    obj["image"] = get_img_to_b64(os.path.join(BASEDIR, "..", "aux", malware["imagedir"]))
    obj["arch"] = malware["arch"]
    obj["ipmeta"] = malware["ipmeta"]
    obj["strain"] = malware["strain"]

    if obj["strain"] == "":  # This is a strain
        obj["mutations"] = malware["mutations"]
    else:  # This is a mutation
        assets = mongo.db.assets
        obj["siblings"] = []
        strains = assets.find({"sha256": obj["strain"]})
        for strain in strains:
            if strain["mutations"] != obj["sha256"]:
                obj["siblings"].append(strain["mutations"])
        obj["siblings"].remove(obj["sha256"])

    del obj["ipmeta"][0]["ip"]
    return obj
