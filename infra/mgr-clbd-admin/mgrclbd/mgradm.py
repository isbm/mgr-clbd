"""
Cluster Admin class.
"""
import sys
import argparse
import json
import requests
from mgrclbd.mgrhlp import APIHelp
from mgrclbd.mgrfmt import OutputFormatter


class ClusterAdmin:
    def __init__(self, parser: argparse.ArgumentParser):
        self.__parser = parser
        self.__args = self.__parser.parse_args()
        self.__director_url = self.__args.director_url or ""
        self.api_root = self.__director_url + "/api/v1"
        self.apihelp = APIHelp(self.__director_url)

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
        elif self.__args.command:
            sys.stdout.write(OutputFormatter().format(self.json_call()))
            called = True

        if not called:
            self.__parser.print_help()

    def _call(self, url: str, data: dict, method="POST"):
        """
        _call -- call request POST.

        :param url: URI to the API endpoint
        :type url: str
        :param data: key/value form payload
        :type data: dict
        """
        if self.__director_url == "":
            raise Exception("Cluster URL is not defined.")

        ret = {}
        data.update(self._dummy_token)
        if method in ["POST", "DELETE"]:
            resp = getattr(requests, method.lower())(self.api_root + url, data=data)
        elif method == "GET":
            resp = requests.get(self.api_root + url, params=data)
        else:
            raise Exception("Method {} not supported".format(method))

        if resp.status_code != 200:
            ret["error"] = resp.text
        else:
            ret = resp.json()
        ret["errcode"] = resp.status_code
        return ret

    def json_call(self) -> dict:
        """
        json_call -- call a remote API endpoint on JSON input.

        :return: JSON dictionary result
        :rtype: dict
        """
        args = json.loads(sys.stdin.read())
        ret = self._call(args["urn"], data=args["arg"], method=args["method"])

        return ret

    def list_zones(self) -> bool:
        """
        list_zones -- list cluster zones.
        """
        ret = self._call("/zones/list", data={}, method="GET")
        if ret.get("errcode", 0) == 200:
            print("Cluster zones:")
            for zone in ret.get("data", []):
                print("  - ", zone["name"])
                print("    ", zone["description"])
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
