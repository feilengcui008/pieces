#!/usr/bin/env python
# -*- encoding:utf-8 -*-

""" this file contains http provider interface and its urllib2 and
requests implementation """

from abc import ABCMeta, abstractmethod
import urllib
import urllib2


class AbstractHttpProvider(object):
    """ abstract class """
    __metaclass__ = ABCMeta

    @abstractmethod
    def get(self, url, headers, data):
        """ get method """
        pass

    @abstractmethod
    def head(self, url, headers):
        """ head method """
        pass

    @abstractmethod
    def post(self, url, headers, data, files):
        """ post method """
        pass

    @abstractmethod
    def put(self, url, headers, data, files):
        """ put method """
        pass

    @abstractmethod
    def patch(self, url, headers, data):
        """ patch method """
        pass

    @abstractmethod
    def delete(self, url, headers):
        """ delete method """
        pass


class HttpProviderException(Exception):
    """ http provider exception """
    pass


class Urllib2HttpProvider(AbstractHttpProvider):
    """
        urllib2 implementation for http provider
        note: this implementation do not support file uploading,
        and it has no dependency packages, this is convenient in
        some cases where we do not want too many dependencies
    """

    @staticmethod
    def get(url, headers=None, data=None):
        if data is not None and isinstance(data, dict):
            url = "%s?%s" % (url, urllib.urlencode(data))
        req = urllib2.Request(url, headers=headers)
        req.get_method = lambda: "GET"
        res = None
        try:
            res = urllib2.urlopen(req)
        except urllib2.HTTPError as exp:
            return (exp.code, exp.msg, None)
        return (res.code, res.msg, res.read(), res.headers.dict)

    @staticmethod
    def head(url, headers=None):
        req = urllib2.Request(url, headers=headers)
        req.get_method = lambda: 'HEAD'
        res = None
        try:
            res = urllib2.urlopen(req)
        except urllib2.HTTPError as exp:
            return (exp.code, exp.msg, None)
        return (res.code, res.msg, None, res.headers.dict)

    @staticmethod
    def post(url, headers=None, data=None, files=None):
        if files is not None:
            raise HttpProviderException(
                "urllib2 implementation does not support file uploading")
        if data is not None and isinstance(data, dict):
            data = urllib.urlencode(data)
        req = urllib2.Request(url, headers=headers)
        req.get_method = lambda: "POST"
        opener = urllib2.build_opener()
        res = None
        try:
            res = opener.open(req, data)
        except urllib2.HTTPError as exp:
            return (exp.code, exp.msg, None)
        return (res.code, res.msg, res.read(), res.headers.dict)

    @staticmethod
    def put(url, headers=None, data=None, files=None):
        if files is not None:
            raise HttpProviderException(
                "urllib2 implementation does not support file uploading")
        if data is not None and isinstance(data, dict):
            data = urllib.urlencode(data)
        req = urllib2.Request(url, headers=headers)
        req.get_method = lambda: "PUT"
        opener = urllib2.build_opener()
        res = None
        try:
            res = opener.open(req, data)
        except urllib2.HTTPError as exp:
            return (exp.code, exp.msg, None)
        return (res.code, res.msg, res.read(), res.headers.dict)

    @staticmethod
    def patch(url, headers=None, data=None):
        if data is not None and isinstance(data, dict):
            data = urllib.urlencode(data)
        req = urllib2.Request(url, headers=headers)
        req.get_method = lambda: "PATCH"
        opener = urllib2.build_opener()
        res = None
        try:
            res = opener.open(req, data)
        except urllib2.HTTPError as exp:
            return (exp.code, exp.msg, None)
        return (res.code, res.msg, res.read(), res.headers.dict)

    @staticmethod
    def delete(url, headers=None):
        req = urllib2.Request(url, headers=headers)
        req.get_method = lambda: 'DELETE'
        res = None
        try:
            res = urllib2.urlopen(req)
        except urllib2.HTTPError as exp:
            return (exp.code, exp.msg, None)
        return (res.code, res.msg, res.read(), res.headers.dict)


class RequestsHttpProvider(AbstractHttpProvider):
    """ requests implementation of http provider """

    def __init__(self):
        try:
            import requests
            self.requests = requests
        except ImportError:
            raise HttpProviderException(
                "RequestsHttpProvider depend on requests package")

    def get(self, url, headers=None, data=None):
        """ get method """
        res = self.requests.get(url, params=data, headers=headers)
        return (res.status_code, res.reason, res.content, res.headers)

    def head(self, url, headers=None):
        """ head method """
        res = self.requests.head(url, headers=headers)
        return (res.status_code, res.reason, None, res.headers)

    def post(self, url, headers=None, data=None, files=None):
        """ post method """
        res = self.requests.post(url, data=data, headers=headers, files=files)
        return (res.status_code, res.reason, res.content, res.headers)

    def put(self, url, headers=None, data=None, files=None):
        """ put method """
        res = self.requests.put(url, data=data, headers=headers, files=files)
        return (res.status_code, res.reason, res.content, res.headers)

    def patch(self, url, headers=None, data=None):
        """ patch method """
        res = self.requests.patch(url, data=data, headers=headers)
        return (res.status_code, res.reason, res.content, res.headers)

    def delete(self, url, headers=None):
        """ delete method """
        res = self.requests.delete(url, headers=headers)
        return (res.status_code, res.reason, None, res.headers)
