## Alice - Cynical Night Plan
シニカルナイトプラン

#### Stage Export Stuff
- Coords: (2621, 1729) <=> (2738, 1824) (May be too small - not enough trees on house's left side)
	- Trying - (2591, 1699) <=> (2768, 1854)
- Floor: 30, Ceil: 100
- Must separate all blocks with transparency and load them in separately due to mmd model draw order weirdness. This is super annoying. Consider learning to render in blender with cam movements in mmd? (but then whole different rendering environment to relearn)
	- Leaves
	- Flowers
	- Fences??? (Normal face dir seems to be screwy?)

#### Stage Import Stuff...
It's ridiculous I need to even write something like this...

Import order/model draw order:
- alice_house_full_minus_vegetation_and_ironbars
- alice_house_grass_and_tallgrass_only
- alice_house_acacia_leaves_only
- alice_house_dark_oak_leaves_only
- alice_house_jungle_leaves_only
- alice_house_oak_leaves_only
- alice_house_iron_bars_only
- alice_house_birch_leaves_only

#### MMD Stuff
- Fug, I think Alice is missing a bone the motion needs

## History
- 2021/07/17 - Basic project setup. Lighting, materials, dummy bones. Might consider scaling models up/down. Use `scalecontrol.pmd`?
- 2021/07/15-16 - Create stage. MMD model render order is a pita.
