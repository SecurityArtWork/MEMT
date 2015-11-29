# -*- coding: utf-8 -*-
"""Graph API module.
This module define every API related to maps.
"""
from __future__ import print_function, absolute_import

from . import __api_prefix__

from flask import jsonify
from flask import request
from flask import redirect
from flask import url_for
from flask import abort
from flask import flash

from flask.ext.classy import FlaskView, route

from app.extensions import mongo

class SearchView(FlaskView):
    route_prefix = __api_prefix__

    @route('/<hash>', methods=['GET'])
    @route('/<hash>/', methods=['GET'], defaults={'type': 'json'})
    def search(self, hash, type=None):
        if hash:
            assets = mongo.db.assets
            query = assets.find({"$or": [{"ssdeep": {"$eq": hash}}, {"md5": {"$eq": hash}}, {"sha1": {"$eq": hash}}, {"sha256": {"$eq": hash}},{"sha512": {"$eq": hash}}]})
            if query.count() == 1:
                for q in query:
                    if type == "json":
                        return jsonify(ecode=200, info="Found", data=get_data(q))
                    else:
                        return jsonify(ecode=200, info="Found", data=get_data(q))
            if type == "json":
                return jsonify(ecode=404, info="Not Found", data={})
            else:
                return jsonify(ecode=404, info="Not Found", data={})

def get_data(malware):
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
    obj["arch"] = malware["arch"]
    obj["strain"] = malware["strain"]
    obj["mutations"] = []
    obj["siblings"] = []
    if obj["strain"] == "":  # This is a strain
        obj["mutations"] = malware["mutations"]
    else:  # This is a mutation
        assets = mongo.db.assets
        strains = assets.find({"sha256": obj["strain"]})
        for strain in strains:
            if strain["mutations"] != obj["sha256"]:
                obj["siblings"].append(strain["mutations"])
        obj["siblings"].remove(obj["sha256"])
    return obj
