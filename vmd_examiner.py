import struct

# This is literally just me screwing around with figuring out the file structure of a .vmd file.

"""
possible_encodings = [
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
# Why the fuck is this in CP932. The answer is "because Windows" but why the fuck.
encoding = "cp932"
filename = "motion.vmd"
with open(filename, "rb") as f:
    contents = bytearray(f.read())
    ver = contents[:30]
    model = contents[30:50]
    both = contents[:50]

    model.decode(encoding)
    print(ver.decode(encoding), model.decode(encoding))
    print(both.decode(encoding))

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
            if f_type is str:
                print(b_data)
                b_converted = b_data.decode(encoding)
                print(b_converted)
            elif f_type is int:
                b_converted = int.from_bytes(b_data, byteorder='big', signed=False)
            elif f_type is float:
               b_converted = struct.unpack('>f', b_data)

            frame_data[f_name] = b_converted

            inner_offset += f_len

        print(frame_data)
        break
