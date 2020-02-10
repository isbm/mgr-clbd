#!/usr/bin/python3

import argparse
import requests
import sys
import json
import textwrap
from typing import Any

class APIHelp:
    api = [
        ("add_zone", {
            "name": ("str", "Name of the Cluster Zone"),
            "description": ("str", "Short summary of the Zone"),
            "_order": ["name", "description"],
            "_summary": "Create a new Cluster Zone. Cluster Zones cannot be deleted, if at least one Cluster Node is in it."
        })
    ]

    def cast(self, value: str, tname: str) -> Any:
        """
        cast -- cast a data to a type.

        :param value: any value in string
        :type tname: str
        :param tname: type name (Python conventions)
        :type tname: str
        :return: typed valued
        :rtype: Any
        """
        out = None
        if tname == "int":
            out = int(value)
        else:
            out = str(value)
        return out

    def create_json_input(self, name: str) -> bool:
        """
        create_json_input -- interactively construct JSON input.

        :param name: name of the API call
        :type name: str
        """
        inp = {"api": name, "arg": {}}
        for descr in self.api:
            apiname, apidata = descr
            if name == apiname:
                for pname in apidata["_order"]:
                    ptype, pdescr = apidata[pname]
                    sys.stderr.write("Define '{}' ({}): ".format(pname, ptype))
                    pdata = input()
                    inp["arg"][pname] = self.cast(pdata, ptype)
        print(json.dumps(inp, indent=4))
        return True

    def list_api(self) -> bool:
        """
        list_api -- list supported API.
        """
        print("Supported API:")
        for descr in self.api:
            print("  -",descr[0])
        return True

    def help_json_sample(self, apidata: dict) -> bool:
        """
        help_json_sample -- generate a JSON input sample to STDOUT.

        :param apidata: api data from the help
        :type apidata: dict
        :return: True
        :rtype: bool
        """
        apiname, apimeta = apidata

        apiargs = {}
        for argname in apimeta["_order"]:
            argtype, argdescr = apimeta[argname]
            apiargs[argname] = "<some {}>".format(argtype)

        sample = {
            "api": apiname,
            "arg": apiargs
        }

        print(json.dumps(sample, sort_keys=True, indent=3))


    def help_on_api(self, name: str) -> bool:
        wrapper = textwrap.TextWrapper(initial_indent="   ", subsequent_indent="   ")
        found = False
        for descr in self.api:
            apiname, apidata = descr
            if name == apiname:
                found = True
                print("Help on '{}' API call:".format(name))
                print("\nWhat it does:")
                print("\n".join(wrapper.wrap(apidata["_summary"])))
                print("\nParameters:")
                for pname in apidata["_order"]:
                    ptype, phelp = apidata[pname]
                    print("  - {} ({})".format(pname, ptype))
                    print("      ", phelp)
                print("\nJSON Sample:\n")
                self.help_json_sample(descr)
                print()
        if not found:
            print("Cannot find '{}' API call".format(name))
        return True


class ClusterAdmin:
    def __init__(self, parser: argparse.ArgumentParser):
        self.__parser = parser
        self.__args = self.__parser.parse_args()
        self.__director_url = self.__args.director_url or ""
        self.api_root = self.__director_url + "/api/v1"
        self.apihelp = APIHelp()

        self._dummy_token = {"token": "0"}

    def run(self):
        called = False
        if self.__args.list_nodes:
            called = self.list_nodes()
        elif self.__args.list_zones:
            called = self.list_zones()
        elif self.__args.list_api:
            called = self.apihelp.list_api()
        elif self.__args.help_api:
            called = self.apihelp.help_on_api(self.__args.help_api)
        elif self.__args.input:
            called = self.apihelp.create_json_input(self.__args.input)

        if not called:
            self.__parser.print_help()

    def _post(self, url: str, data: dict):
        """
        _post -- call request POST.

        :param url: URI to the API endpoint
        :type url: str
        :param data: key/value form payload
        :type data: dict
        """
        if self.__director_url == "":
            raise Exception("Cluster URL is not defined.")

        data.update(self._dummy_token)
        resp = requests.post(self.api_root + url, data=data)
        ret = resp.json()
        ret["errcode"] = resp.status_code
        return ret

    def list_zones(self) -> bool:
        """
        list_zones -- list cluster zones.
        """
        ret = self._post("/nodes/list", {})
        print(ret)
        return True

    def add_zone(self) -> bool:
        """
        add_zone -- create a zone in the cluster.
        """
        return True


    def list_nodes(self) -> bool:
        """
        list_nodes -- list cluster nodes.
        """
        return True

    def list_systems(self, mid: str):
        """
        list_systems -- list all registered systems to a cluster node.

        :param mid: machine-id of a registered cluster node.
        :type mid: string
        """

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

    ClusterAdmin(p).run()

if __name__ == "__main__":
    main()
