"""
API formatters.
"""
import json

class OutputFormatter:
    """
    Generic output formatter.
    """
    def format(self, jsondata: dict) -> str:
        """
        format -- Format JSON to the stdout

        :param jsondata: dictionary to format the data
        :type jsondata: dict
        :return: formatter string for the output
        :rtype: str
        """
        return json.dumps(jsondata, indent="  ")
