#!/usr/bin/env python
# coding: utf-8

""" a hook for python importer """

import sys
import os
import imp
import inspect
from types import ModuleType

class SolaFinder(object):
    """ define a finder for .sola file """
    def __init__(self, *args, **kwargs):
        self.ext_name = '.sola'

    def find_module(self, module_name, path=None):
        """ return the a loader obj """
        file_path = self.get_filepath(module_name)
        if os.path.isfile("%s%s" % (file_path, self.ext_name)):
            return SolaLoader(file_path, self.ext_name)
        else:
            return None

    def get_parent_path(self, path, level=1):
        """ get parent path """
        for _ in range(level):
            path = os.path.dirname(path)
        return path

    def get_filepath(self, module_full_name):
        """ get the file path """
        # who is calling import
        # current: 0, find_module: 1
        importer_frame = sys._getframe(2)
        importer_file = inspect.getfile(importer_frame)
        package_dir = os.path.dirname(importer_file)
        if '__name__' in importer_frame.f_globals:
            if '.' in importer_frame.f_globals['__name__']:
                total_upper_level = len(importer_frame.f_globals['__name__'].split('.'))
                # get the actual path when we import like import a.b.c
                package_dir = self.get_parent_path(package_dir, total_upper_level - 1)
        if '.' in module_full_name:
            return os.path.join(*([package_dir] + module_full_name.split('.')))
        return os.path.join(package_dir, module_full_name)


class SolaLoader(object):
    """ a loader """
    def __init__(self, file_path, ext_name):
        self.file_path = file_path
        self.ext_name = ext_name
        self.full_file_name = "%s%s" % (self.file_path, self.ext_name)

    def load_module(self, full_name):
        """ load module """
        if full_name in sys.modules:
            return sys.modules[full_name]
        mod = None
        print(self.full_file_name)
        with open(self.full_file_name, "rb") as fp:
            # mod = imp.load_module(full_name, fp, self.full_file_name, (self.ext_name, "rb", 1))
            mod = ModuleType(full_name)
            compiled = compile(fp.read(), '', 'exec')
            exec(compiled, mod.__dict__)
        if mod is not None:
            mod.__file__ = self.full_file_name
            sys.modules.setdefault(full_name, mod)
        return mod


if sys.version_info >= (3, 0, 0) or sys.version_info < (2, 3, 0):
    print "warning: this import hook do not support python version 3 or lower than 2.3"
else:
    sys.meta_path.append(SolaFinder())
