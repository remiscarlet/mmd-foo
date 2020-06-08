
# What is this

idk, random stuff. Fucking with VMD content parsing.

## Notes
- I think camera motions are a special case with bone names and don't have the "camera" string assigned but rather a special byte marker.
    - `bytes: [0 0 0 0 25 1 0 0 0 0 0 0 0 0 24] - bytes2: 000000001901000000000000000018`
    - Meh. Back to `motion.vmd`
    - Welp. Yep as of 7.10 looks like they changed it up. No use camera motions.

Resources:
- https://mikumikudance.fandom.com/wiki/VMD_file_format
- https://blog.goo.ne.jp/torisu_tetosuki/e/bc9f1c4d597341b394bd02b64597499d
Motion file from:
- https://www.nicovideo.jp/watch/sm36969890
