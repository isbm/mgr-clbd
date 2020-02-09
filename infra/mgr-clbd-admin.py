#!/usr/bin/python3

import argparse
import requests

class ClusterAdmin:
    def __init__(self, parser):
        self.__parser = parser
        self.__args = self.__parser.parse_args()
        self.__director_url = self.__args.director_url
        self.api_root = self.__director_url + "/api/v1"

        self._dummy_token = {"token": "0"}

    def run(self):
        if self.__args.list_nodes:
            self.list_nodes()
        else:
            self.__parser.print_help()


    def list_nodes(self):
        r = requests.post("{}/nodes/list".format(self.api_root), data=self._dummy_token)
        if r.status_code == 200:
            print(r.json())
        else:
            print("ERROR:", r.json().get("error", "Unknown error"))

    def list_systems(self, fp):
        pass

def main():
    p = argparse.ArgumentParser()
    p.add_argument("-l", "--list-nodes", help="List all cluster nodes", action="store_true")
    p.add_argument("-u", "--director-url", help="Cluster Director URL", required=True)
    ClusterAdmin(p).run()
 
if __name__ == "__main__":
    main()
