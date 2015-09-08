#!/usr/bin/python -u

def convert_to_str(st):
    """ use this function to convert all strings to str"""
    if isinstance(st, unicode):
        return st.encode('utf-8')
    return str(st)
 

if __name__ == "__main__":

    val1 = "aaé"
    print val1        # aaé
    print type(val1)  # <type 'str'>
     
    print val1.encode('utf-8')
    """
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    UnicodeDecodeError: 'ascii' codec can't decode byte 0xc3 in position 0: ordinal not in range(128)
    """
     
    print val1.encode('ascii')
    """
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    UnicodeDecodeError: 'ascii' codec can't decode byte 0xc3 in position 0: ordinal not in range(128)
    """
     
    val2 = u"aaé"
    print val2                  # aaé
    print type(val2)            # <type 'unicode'>
    print val2.encode('utf-8')  # aaé
     
    print val2.encode('ascii')
    """
    Traceback (most recent call last):
      File "<stdin>", line 1, in <module>
    UnicodeEncodeError: 'ascii' codec can't encode character u'\xe9' in position 0: ordinal not in range(128)
    """
     
     
    print val2.encode('ascii', 'ignore') # aa
    # é is missing
     
    import unicodedata
    unicodedata.normalize('NFKD', val2).encode('ascii','ignore')
    # aae
    # é got converted to e
 

    val1 = "ébc"
    val2 = u"ébc"
     
    print val1, type(val1), convert_to_str(val1), type(convert_to_str(val1))
    # ébc <type 'str'> ébc <type 'str'>
     
    print val2, type(val2), convert_to_str(val2), type(convert_to_str(val2))
    # ébc <type 'unicode'> ébc <type 'str'>
