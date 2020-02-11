"""
API helper. This dynamically generates info about remote API endpoinds, based on OpenAPI specs.
"""
import json
from typing import Any
import requests
import textwrap
import sys


class APIHelp:
    """
    API Help builds OpenAPI-based definitions.
    """
    def __init__(self, url: str):
        """
        __init__ -- constructor
        """
        self.__api_cache = []
        self._url = url
        self.api = []
        self._update_api()

    @staticmethod
    def type2python(typename: str) -> str:
        """
        type2python -- Convert OpenAPI types to Python types

        :param typename: name of the type
        :type typename: str
        :return: new type name
        :rtype: str
        """
        out = ""
        if typename == "string":
            out = "str"
        else:
            out = typename
        return out

    def _update_api(self):
        """
        _upate_api -- get OpenAPI definitions and format it.
        """
        self.api.clear()
        buff = requests.get(self._url + "/swagger/doc.json").json()["paths"]
        for apipath in buff:
            for method in buff[apipath]:
                apidata = buff[apipath][method]
                apiset = {
                    "_urn": apipath.replace("/api/v1", ""),
                    "_method": method.upper(),
                    "_order": [],
                    "_summary": apidata.get("description", "N/A")
                }
                for argset in apidata.get("parameters", []):
                    apiset[argset["name"]] = (self.type2python(argset["type"]), argset["description"])
                    apiset["_order"].append(argset["name"])
                self.api.append((apidata["operationId"].replace("-", "_").lower(), apiset,))

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
                inp["method"] = apidata["_method"]
                inp["urn"] = apidata["_urn"]
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
            "arg": apiargs,
            "urn": "/endpoint/function",
            "method": "<POST or GET>"
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
