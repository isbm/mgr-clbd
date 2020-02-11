#!/usr/bin/env python

"""
Main CLI script
"""
from mgrclbd.mgradm import ClusterAdmin
import argparse

def main():
    p = argparse.ArgumentParser()

    general = p.add_argument_group("General")
    general.add_argument("-u", "--director-url", help="Cluster Director URL")

    info = p.add_argument_group("Info")
    info.add_argument("-z", "--list-zones", help="List all cluster zones", action="store_true")
    info.add_argument("-l", "--list-nodes", help="List all cluster nodes", action="store_true")

    funcs = p.add_argument_group("API functions")
    funcs.add_argument("-f", "--list-api", help="List all API functions", action="store_true")
    funcs.add_argument("-d", "--help-api", help="Get help on an API function")
    funcs.add_argument("-c", "--command", help="Call an API endpoint with the JSON input command", action="store_true")
    funcs.add_argument("-i", "--input", help="Construct a JSON input command")

    try:
        ClusterAdmin(p).run()
    except Exception as exc:
        print("Error:", exc)
        raise

if __name__ == "__main__":
    main()
