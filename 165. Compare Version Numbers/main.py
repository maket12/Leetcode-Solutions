class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        temp1, temp2 = "", ""
        i, j = 0, 0
        while i < len(version1) or j < len(version2):
            temp1, temp2 = "", ""
            while i < len(version1) and version1[i] != ".":
                temp1 += version1[i]
                i += 1
            while j < len(version2) and version2[j] != ".":
                temp2 += version2[j]
                j += 1

            i += 1
            j += 1
            
            val1 = int(temp1) if len(temp1) > 0 else 0
            val2 = int(temp2) if len(temp2) > 0 else 0

            if val1 < val2:
                return -1
            elif val1 > val2:
                return 1
        return 0
