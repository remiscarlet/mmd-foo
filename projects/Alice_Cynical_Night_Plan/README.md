## Alice - Cynical Night Plan
シニカルナイトプラン

### Stage Export Stuff
- Coords: (2621, 1729) <=> (2738, 1824) (May be too small - not enough trees on house's left side)
	- Trying - (2591, 1699) <=> (2768, 1854)
- Floor: 30, Ceil: 100
- Must separate all blocks with transparency and load them in separately due to mmd model draw order weirdness. This is super annoying. Consider learning to render in blender with cam movements in mmd? (but then whole different rendering environment to relearn)
	- Leaves
	- Flowers
	- Fences??? ~~(Normal face dir seems to be screwy?)~~ (Just need to enable 2-sided in pmxeditor)

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
- ~~0:40 Don't pan horizontally again. Vertical?~~
	- Completely redid section
- ~~Intro - Y of stage isn't right. Shoes clips through ground.~~
- ~~0:27 - DoF focuses on hand~~
- ~~Fade to black isn't full black because of... some conflicting effects.~~
- The dirlight intensity change in intro is a bit too noticeable maybe

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
- Chorus (8m)(2:22-2:38) - Need to redo this section. A bit "basic" right now.
- Chorus' (8m)(2:38-2:54) - Need to make make it a bit more "variation"-y. Too copypastad from earlier still.
- Instrumental (8m)(2:54-3:10)
- Instrumental/Outro (4m)(3:10-3:18/end)

## Credits
- Music: Ayase
- Vocals: くろくも
- Model: Arlvit
- Choreo: わに
- Motion: リングイネ
- MME: rui, そぼろ, ビームマンP,  ましまし,  SANDMAN , ikeno
- Fonts: Togalite
- Stage: kaiwakap
- Resource Pack: Steven's Traditional
- Alice's House/アリスの家: remidapyon x Yukkuricraft
- Camera, Video, Pose: remidapyon

Music: Ayase
Vocals: くろくも

Model: Arivit
Choreo: わに   Motion: リングイネ
Fonts: Togalite  Skybox: Kaiwakap
MME: rui, そぼろ, ビームマンP,  ましまし,
           SANDMAN , ikeno
		   
Resource Pack: Steven's Traditional
Alice's House/アリスの家:
    remidapyon x Yukkuricraft
Camera, Video, Pose:
    remidapyon

### Description
TitleJp:「MMD」アリスのシニカルナイトプラン
TitleEng:「MMD」Alice's Cynical Night Plan

どもす。remidapyonです。

第三作めとなります。今回はカメラの動きに集中し、なんとか作ってみました。心の癒やすアリスの舞い踊り、いかがですか？

また今回は幻想郷をマイクラで作ったコミュニティ「Yukkuricraft」の「幻想郷 V3」のマップを使ったステージです。
「幻想郷 V3」のトレーラー動画はこちらです。どうぞお楽しみください：[Link Here]

どうぞ、第三作の「シニカルナイトプラン」


Sup. remidapyon here.

This marks my third work so far. I tried to focus much more on the camerawork this time around. Alice is dancing the night away - Will you join her?

Once again, I created the stage from Yukkuricraft's Genso v3 map. Yukkuricraft is the community that recreated the entirety of Gensokyo in a fully explorable public minecraft server.
If you would like, please check out the Genso v3 Trailer here: [Link Here]

Anyways, please enjoy my third work: Cynical Night Plan!


Credits:
[Music] Ayase
[Vocals] くろくも
[Model] Arivit
[Choreo] わに
[Motion] リングイネ
[MME] rui, そぼろ, ビームマンP,  ましまし,  SANDMAN , ikeno
[Fonts] Togalite
[Skydome] kaiwakap

[Resource Pack] Steven's Traditional
[Alice's House/アリスの家 Stage] remidapyon x Yukkuricraft
[Camera, Video, Pose] remidapyon


## History
- 2021/08/02
	- render.8-e - Fix a ton of hand-clipping-through-skirt bits. Reposition intro bits stage at parts. Move lighting around in intro to be a bit more interesting. Lower specularity on eyeball materials (less reflective). Create thumbnail.
- 2021/07/22
	- render.8-d - Remove "dolley off screen"s at ~0:35 and ~0:40. Fix non-true darkness on fadeouts caused by iklightrays.
- 2021/07/21
	- render.8-c - Oops didn't have `main.fx` on any of the credit pmxs. Adjust AL4 for readability on credits. Add blue petals to camera tilt in intro. Add lightrays and fogs to intro. Replace HW filter with RC filter. Increase exposure a bit in first ~30 seconds.
- 2021/07/20
	- render.8 - Add credits as 3d objects using Blender.
	- render.8-a - Lots of material adjustments for credits text. Outline on, texture as blue concrete powder for titles and lapis lazuli for names, add toon01, set reflection ~103-105 (AL4). Reposition credits text and camera for better framing.
	- render.8-b - Adjust intro cam work a bit. Adjust framing to emphasize skirt spin a bit more at ~2:28 and ~2:38. Redo "arm slams" at 2:39-2:41
- 2021/07/19
	- render.4 - Clean up/smooth out a lot of the interpolation heavy movements. Add motions for first half of bridge (~1:43)
	- render.5 - Lots of new cam motion. ~2:13
	- render.5-a - Smooth out a lot of the faster cam motions - particularly with better framing. Add a bit more motion ~2:22
	- render.5-b - Fix a ton of arm/hand clipping problems.
	- render.6 - Cam motions up till final instrumental ~2:54. Fix more arm clippings.
	- render.7 - Finished first final draft cam movements. Still need to refine a bunch of stuff.
	- render.7-a - Fix the パンチラ oops. Extend idle animation and fade out at end. Lower dirt and grass material specular vals. Increase EnvSpecDiff. Increase SSAO. Bump exposure up a _tiny_ bit. (Overexposing causes house walls to white out...)
	- render.7-b - Smooth out transition between Chorus/Chorus'. Move stage to Forest of Magic instead of front of house for intro. Smooth/fix misc minor motions.
	- render.7-c - Connect transition at ~2:02. Lighting tweaks - gamma down, contrast up, exposure up. Tweak timing of transition at 2:14-2:15
- 2021/07/18
	- render.1 - Intro first pass up to ~0:24. More material/lighting tweaks.
	- render.2 - Cam motion up to ~0:50. Add fade fx. 
	- render.3 - Cam motion up to ~1:29. Add HW post filter. More lighting adjusts. DoF adjusts.
	- render.3-a - Adjust interpolation and framing particularly in Chorus. Turn diffusion down a notch. Add cheaplens.
	- render.3-b - PostMovie??? Maybe a bit too grainy
	- render.3-c - Fix 2-sided faces being transparent on one side (vegetation, glass, iron fence, etc). Add back in non-grass vegetation... how did I miss this? Add emissive mats to light emitting blocks (lantern, glowstone, lamphead, fireplace). Adjust cams where newly visible vegation blocks LoS. Hacky workaround for transparent texture bug... UV mapping float imprecision? See section further up. Upscale most vegetation textures 2x to try work around alpha bleeds.
- 2021/07/17 - Basic project setup. Lighting, materials, dummy bones. Might consider scaling models up/down. Use `scalecontrol.pmd`?
- 2021/07/15-16 - Create stage. MMD model render order is a pita.
