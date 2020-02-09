#!/usr/bin/python3

import argparse
import requests

class ClusterAdmin:
    def __init__(self, parser: argparse.ArgumentParser):
        self.__parser = parser
        self.__args = self.__parser.parse_args()
        self.__director_url = self.__args.director_url
        self.api_root = self.__director_url + "/api/v1"

        self._dummy_token = {"token": "0"}

    def run(self):
        called = False
        if self.__args.list_nodes:
            self.list_nodes()
            called = True

        if self.__args.list_zones:
            self.list_zones()
            called = True

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
        data.update(self._dummy_token)
        resp = requests.post(self.api_root + url, data=data)
        ret = resp.json()
        ret["errcode"] = resp.status_code
        return ret

    def list_zones(self):
        """
        list_zones -- list cluster zones.
        """
        ret = self._post("/nodes/list", {})
        print(ret)

    def list_nodes(self):
        """
        list_nodes -- list cluster nodes.
        """

    def list_systems(self, mid: str):
        """
        list_systems -- list all registered systems to a cluster node.

        :param mid: machine-id of a registered cluster node.
        :type mid: string
        """

def main():
    p = argparse.ArgumentParser()
    p.add_argument("-z", "--list-zones", help="List all cluster zones", action="store_true")
    p.add_argument("-l", "--list-nodes", help="List all cluster nodes", action="store_true")
    p.add_argument("-u", "--director-url", help="Cluster Director URL", required=True)
    ClusterAdmin(p).run()

if __name__ == "__main__":
    main()
