
# What is this

idk, random stuff. Fucking with VMD content parsing.

## Other Tool Ideas
- VMD parser:
    - Parsing the content could be written as a state machine??? This might make things more straight forward
    - Or I might do this as a side project and just use `pymeshio` to get working on the actual tools...
- You know what I really hate. MME effects and how I have 100+ folders and only know what a dozen of them look like...
	- Somehow auto generate MMD screenshots of effects? Would be more aimed at postfx or general screen fx though. What category of fx are there?
- More generally for above, some kind of MME organizer/parser/etc in general would be great. There's currently so little support for organizing them.
- Long shot. Custom read minecraft world data and custom render it in unity/blender/etc and export to mmd to use as stage. Eg, with Remi and SDM.
	- World data parser
		- `pip install NBT` or `DonoA/PyAnvilEditor`(No pip?) for parsing the data
	- Less complicated:
		- export world as obj (`jMc2Obj`)
			- Texture export from resource pack is stalling. Gonna need to dig into source code to figure out why :(
		- Load into blender and export as `.x`
	- This is done via:
		- mineways for mc world to `.obj` export
			- Add custom texturefile
			- Each block 1000mm
			- Export separate types
			- Material per family
			- Do NOT export individual blocks
			- Do NOT split by block type
		- Blender for `.obj` to `.pmx` using modified `blpymeshio`
		- PMXEditor for `.pmx` to `.x` (`.pmx` crashes MMD due to size?)(Take advantage of AutoLuminous when exporting)
		- MMD for importing `.x` (Load AL4 if using)

## Notes
- I think camera motions are a special case with bone names and don't have the "camera" string assigned but rather a special byte marker.
    - `bytes: [0 0 0 0 25 1 0 0 0 0 0 0 0 0 24] - bytes2: 000000001901000000000000000018`
    - Meh. Back to `motion.vmd`
    - Welp. Yep as of 7.10 looks like they changed it up. No use camera motions.

Resources:
- https://mikumikudance.fandom.com/wiki/VMD_file_format
- https://blog.goo.ne.jp/torisu_tetosuki/e/bc9f1c4d597341b394bd02b64597499d
- http://atupdate.web.fc2.com/vmd_format.htm
Motion file from:
- https://www.nicovideo.jp/watch/sm36969890
