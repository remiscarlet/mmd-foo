## Alice - Cynical Night Plan
シニカルナイトプラン

### Stage Export Stuff
- Coords: (2621, 1729) <=> (2738, 1824) (May be too small - not enough trees on house's left side)
	- Trying - (2591, 1699) <=> (2768, 1854)
- Floor: 30, Ceil: 100
- Must separate all blocks with transparency and load them in separately due to mmd model draw order weirdness. This is super annoying. Consider learning to render in blender with cam movements in mmd? (but then whole different rendering environment to relearn)
	- Leaves
	- Flowers
	- Fences??? (Normal face dir seems to be screwy?)

### Stage Import Stuff...
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

### MMD Notes
- ~~Fug, I think Alice is missing a bone the motion needs~~
	- Lol just using a different model.

### Weird UV Mapping Thing on Transparent Textures
- Most obvious on `tall_grass.png` tetxures where it seems the very top line of pixels being rendered is actually wrapping and displaying the very bottom row of pixels.
	- This is even more apparent when observing the two halves of double tall grass. You can see the pattern of "wrapped" pixels at the top of each half matches the grass pattern from the bottom of the same block.
- For workaround, upscale the textures x2 and blank out the very bottom row of pixels. Seems to be invisible enough on tall grass esp if one just shifts the vegetation heights down by like 0.01 Y pos.

### Section Planning
- Intro (4m)(0:00-0:08) - Closeups/slightly fast zoom away?
- Instrumental (8m)(0:08-0:24) - Add credits with 4 scenes 2 measures each. Alice stands to left or right sides with credit text on opposite side
- SectA (8m)(0:24-0:40)
- SectA' (9m)(0:40-1:00) (end has the stair walk up/down)
- Chorus (8m)(1:00-1:16)
- Instrumental (8m)(1:16-1:32) - Semi fancy cams
- Bridge(8m)(1:32-1:48) - Use high contrast fx? Lines only? Mono? Etc, play into the "radio voice" effect
- SectB (8m)(1:48-2:04) 
- Intro x2 (4m x2 + 1m)(2:04-2:22) - Focus on finger walk at 2:08
- Chorus (8m)(2:22-2:38)
- Chorus' (8m)(2:38-2:54)
- Instrumental (8m)(2:54-3:10)
- Instrumental/Outro (4m)(3:10-3:18/end)

## History
- 2021/07/18
	- render.1 - Intro first pass up to ~0:24. More material/lighting tweaks.
	- render.2 - Cam motion up to ~0:50. Add fade fx. 
	- render.3 - Cam motion up to ~1:29. Add HW post filter. More lighting adjusts. DoF adjusts.
	- render.3-a - Adjust interpolation and framing particularly in Chorus. Turn diffusion down a notch. Add cheaplens.
	- render.3-b - PostMovie??? Maybe a bit too grainy
	- render.3-c - Fix 2-sided faces being transparent on one side (vegetation, glass, iron fence, etc). Add back in non-grass vegetation... how did I miss this? Add emissive mats to light emitting blocks (lantern, glowstone, lamphead, fireplace). Adjust cams where newly visible vegation blocks LoS. Hacky workaround for transparent texture bug... UV mapping float imprecision? See section further up. Upscale most vegetation textures 2x to try work around alpha bleeds.
- 2021/07/17 - Basic project setup. Lighting, materials, dummy bones. Might consider scaling models up/down. Use `scalecontrol.pmd`?
- 2021/07/15-16 - Create stage. MMD model render order is a pita.
