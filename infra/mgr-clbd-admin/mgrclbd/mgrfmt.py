"""
API formatters.
"""

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
        out = []
        if jsondata.get("errcode", 0) != 200:
            out.append("Error: " + jsondata.get("error", ""))
        else:
            out.append("Info: " + jsondata.get("msg", "N/A") or "")
            out.append(str(jsondata.get("data", {})))

        return "\n".join(out)
