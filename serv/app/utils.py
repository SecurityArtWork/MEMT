# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import re
import json

from datetime import datetime
from time import mktime

from werkzeug.routing import BaseConverter
from werkzeug.routing import ValidationError


class sha256Converter(BaseConverter):

    regex = "[A-Fa-f0-9]{64}"

    def __init__(self, map):
        BaseConverter.__init__(self, map)
        self.fixed_digits = 64


    def to_python(self, value):
        if len(value) != self.fixed_digits:
            raise ValidationError()
        return str(value)

    def to_url(self, value):
        return str(value)




class MEMTEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, datetime):
            return {
                '__type__': '__datetime__',
                'epoch': int(mktime(obj.timetuple()))
            }
        else:
            return json.JSONEncoder.default(self, obj)

def memt_decoder(obj):
    if '__type__' in obj:
        if obj['__type__'] == '__datetime__':
            return datetime.fromtimestamp(obj['epoch'])
    return obj

# Encoder function
def memt_dumps(obj):
    return json.dumps(obj, cls=MEMTEncoder)

# Decoder function
def memt_loads(obj):
    return json.loads(obj, object_hook=memt_decoder)
