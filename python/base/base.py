from sortedcontainers import SortedDict, SortedList


class NameDic(SortedDict):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def add(self,key,val):
        self[key]=val


n=NameDic()
n.add("1","2")
n.add("12","2")
n.add("11","2")
print(n)