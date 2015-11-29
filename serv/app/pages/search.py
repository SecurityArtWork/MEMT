# -*- coding: utf-8 -*-
"""Module that controls the search view.
"""
from __future__ import absolute_import

from flask import Blueprint
from flask import render_template
from flask import request
from flask import flash
from flask import abort
from flask import redirect
from flask import url_for

from app.extensions import mongo

bp = Blueprint('search', __name__, url_prefix='/search')


@bp.route('', methods=['POST'])
@bp.route('/<hash>', methods=['GET'])
def index():
    if request.method == 'POST':
        hash = request.form["hash"]
    if hash:
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
