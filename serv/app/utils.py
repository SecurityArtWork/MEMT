# -*- coding: utf-8 -*-
from __future__ import print_function, absolute_import

import re
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
