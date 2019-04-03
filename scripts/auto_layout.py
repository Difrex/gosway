#!/usr/bin/env python

# This is a simple i3blocks script

import json
import os


LAYOUTS = {
    "manual": "[M]",
    "spiral": "[S]",
    "failed": "[DOWN]"
}


def swaymgr_cmd(cmd):
    """sends a command to the swaymgr"""
    return os.popen("swaymgr -s '" + cmd + "'").read()


def get_current_layout():
    """get the currently active workspace layout configuration"""
    try:
        layout = json.loads(swaymgr_cmd("get layout"))
    except Exception:
        print(LAYOUTS["failed"])
        return
    if not layout["managed"]:
        print(LAYOUTS["manual"])
    else:
        print(LAYOUTS[layout["layout"]])


if __name__ == "__main__":
    get_current_layout()
