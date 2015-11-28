# -*- coding: utf-8 -*-
"""Graph API module.
This module define every API related to maps.
"""
from __future__ import print_function, absolute_import

from flask import jsonify
from flask import request
from flask import redirect
from flask import url_for
from flask import abort
from flask import flash

from flask.ext.classy import FlaskView, route

from app.extensions import mongo

class SearchView(FlaskView):

    @route('/<hash>', methods=['GET','POST'])
    def search(self, hash=None):
        print("BUSCANDO")
        if request.method == 'POST':
            print("HOLA")
            hash = request.data["hash"]
        if hash:
            print("ADIOS")
            assets = mongo.db.assets
            query = assets.find({"$or": [{"ssdeep": {"$eq": hash}}, {"md5": {"$eq": hash}}, {"sha1": {"$eq": hash}}, {"sha256": {"$eq": hash}},{"sha512": {"$eq": hash}}]})
            if query.count() == 0:
                flash("404 - Hash not found :(", "danger")
                return abort(404)
            elif query.count() == 1:
                for q in query:
                    return redirect(url_for("detail.index", hash=q['sha256']))
            else:
                flash("500 - Hash not found :(", "danger")
                return abort(500)
        return redirect(url_for("index.index"))
