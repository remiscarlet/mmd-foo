import struct
from collections import namedtuple
import sys

if (sys.version_info < (3, 0)):
  sys.stderr.write("Please run with Py3.x. Unicode and bytes in py2 is a nightmare.\n")
  exit(1)

# This is literally just me screwing around with figuring out the file structure of a .vmd file.

# Why the fuck is this in CP932. The answer is "because Windows" but why the fuck.
ENCODING = "cp932"
FILENAME = "motion.vmd"

"""
possible_ENCODINGs = [
    "cp932",
    "euc_jp",
    "euc_jis_2004",
    "euc_jisx0213",
    "iso2022_jp",
    "iso2022_jp_1",
    "iso2022_jp_2",
    "iso2022_jp_2004",
    "iso2022_jp_3",
    "iso2022_jp_ext",
    "shift_jis",
    "shift_jis_2004",
    "shift_jisx0213",
]
"""

BoneFrame = namedtuple('BoneFrame', [
    'bone_name',
    'frame_id',
    'x',
    'y',
    'z',
    'x_rq',
    'y_rq',
    'z_rq',
    'w_rq',
    'frame_interpo',
])

FILE_STRUCTURE = [
    ("version", 30, str),
    ("model_name", 20, str),
    ("bone_frames_len", 4, int),
]
"""
    [
        (BoneFrame),
        ("bone_name", 15, str),
        ("frame_id", 4, int),
        ("x", 4, float),
        ("y", 4, float),
        ("z", 4, float),
        ("x_rq", 4, float),
        ("y_rq", 4, float),
        ("z_rq", 4, float),
        ("w_rq", 4, float),
        ("frame_interpo", 64, str),
    ],
]
"""

def convertByteArrayToType(b_arr, t):
    cleaned_arr = bytearray()
    #for b in b_arr:
    #    if b != b'\x00':
    #        cleaned_arr.add(b)

    if t is str:
        b_converted = b_arr.strip(b'\x00').decode(ENCODING)
    elif t is int:
        b_converted = int.from_bytes(b_arr, byteorder='big', signed=False)
    elif t is float:
       b_converted = struct.unpack('>f', b_arr)[0]

    return b_converted

def grabBytes(contents, offset, length):
    content_b = contents[offset:offset+length]
    rtn_b = bytearray()
    print("GRABBING BYTES")
    for b in content_b:
        print(b)
        if b != b'\x00':
            rtn_b += bytearray(b)
    print("RETURNING BYTES")
    return rtn_b

with open(FILENAME, "rb") as f:
    contents = bytearray(f.read())

    DATA = {}
    OFFSET = 0
    PREV_SECTION_NAME = None
    for section in FILE_STRUCTURE:
        structure_type = type(section)

        if structure_type == tuple:
            sec_name, sec_len, sec_type = section
            #print("Looking at %s of len %d and converting to type %s" % section)
            b_content = grabBytes(contents, OFFSET, sec_len)
            b_converted = convertByteArrayToType(b_content, sec_type)
            print("%s: %s - %s - %s" % (sec_name, b_converted, sec_type, b_content) )
            OFFSET += sec_len
        else:
            struct_type = section[0]
            print("Using namedtuple:")
            print(struct_type)
            eval_string = struct_type.__name__+"("
            for sub_section in section[1:]:
                print(sub_section)
                sec_name, sec_len, sec_type = sub_section
                print("Looking at %s of len %d and converting to type %s" % sub_section)

                b_content = grabBytes(contents, OFFSET, sec_len)
                b_converted = convertByteArrayToType(b_content, sec_type)

                print("+++", b_converted, repr(b_converted))

                if sec_type == str:
                    eval_string += "%s = \"%s\"," % (sec_name, b_converted)
                else:
                    eval_string += "%s = %s," % (sec_name, b_converted)
                OFFSET += sec_len
            eval_string = eval_string[:-1] + ")"
            print(eval_string)
            #print("\x".join("{:02x}".format(ord(c)) for c in eval_string))
            frame = eval(eval_string)
            print(frame)
    ver = contents[:30]
    model = contents[30:50]
    both = contents[:50]

    model.decode(ENCODING)
    print(ver.decode(ENCODING), model.decode(ENCODING))
    print(both.decode(ENCODING))

    kf_count = contents[50:54]
    kf_int = int.from_bytes(kf_count, byteorder='big', signed=False)
    print(kf_int)

    offset_start = 54

    # (dataname, bytelen, type)
    frame_struct = [
        ("bone_name", 15, str),
        ("frame_id", 4, int),
        ("x", 4, float),
        ("y", 4, float),
        ("z", 4, float),
        ("x_rq", 4, float),
        ("y_rq", 4, float),
        ("z_rq", 4, float),
        ("w_rq", 4, float),
        ("frame_interpo", 64, str),
    ]
    frame_struct_byte_len = sum(map(lambda d: d[1], frame_struct))
    print(frame_struct_byte_len)

    for i in range(kf_int):
        print(i)
        frame_data = {}

        inner_offset = 0
        for field in frame_struct:
            print(field)
            f_name, f_len, f_type = field
            print("Grabbing data from %d to %d" % (offset_start+inner_offset, offset_start+inner_offset+f_len))
            b_data = contents[offset_start+inner_offset:offset_start+inner_offset+f_len]

            frame_data[f_name] = convertByteArrayToType(b_data, f_type)
            inner_offset += f_len

        print(frame_data)
        break
