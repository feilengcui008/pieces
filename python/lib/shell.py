#! /usr/bin/env python
# -*- encoding:utf-8 -*-

""" some shell utilities """

import subprocess


class ShellUtility(object):
    """ some utility functions """
    @staticmethod
    def getcodeoutput(cmd, shell=True, env=None):
        """ get return code and output """
        proc = subprocess.Popen(cmd, stdin=subprocess.PIPE,
                                stdout=subprocess.PIPE,
                                stderr=subprocess.PIPE,
                                shell=shell,
                                env=env)
        out = proc.communicate()[0]
        err = proc.communicate()[1]
        code = proc.returncode
        return (code, out, err)
