# Notes

## General thoughts
- Use focal length variation more
- Default should be with interpolation. Use linear sparingly.
    - Accelerate (x^2) for zoom in/out feel
    - Decelerate (-x^2) for "sliding into next shot" feel
    - tan(x) for 転けた feel
    - tanh(x) for "doppler effect" feel
- Use more continuous shots. Eg, keep continuous motion with direction/focal changes rather than teleporting camera.
    - Shifting speeds while keeping same direction is nice.
    - Consider angle changes with speed shift.
- If the choreo is interesting, chill out on the camera to let the motion shine.
- If the choreo is telling a "story" or symbolic/meaningful movement, try to make camera play along with it.
- Watch more videos while framestepping. 
- How about moving the models so it's not in the same place every damn scene... 
- Move camera along a single axis with tilted angle?
- Use "very slow movement" to "speed up and zoom away" kind of feel 
	- "du-wah" kind of movement if used in repetition
- "Swooping in from above" movement.
- Extreme acceleration but over more frames so feels more dramatic?
- If music is heavily syncopated, use still camera with jump cuts.
- Constant speed but with cuts changing position/angle. Keep the feel with camera moving same speed/same direction.
- Use cuts alternating between two general positions/angles so it kinda feels like you're going back and forth
	- And more generally, keep some sense of continuity between cuts. This will probably help with cuts sometimes feeling like they happen too often
- For slow/non movements on the choreo, use shots from behind.
	- Or just use shots from behind more often.
- If choreo has finger pointing/moving with the model looking in that direction, position camera where they're pointing?
- Generally take advantage of mirrored movements/"larger picture" movements.
	- Isn't this the same thing as the continuity comment?
- For faster/frequent cuts, either have the bg or the model moving, but don't move both. Moving both just feels motion sick. 
	- Good example: https://youtu.be/bQs3I6cMQUI?t=163
- I think the subtitles and credit lines in the HQ vids are just custom .x files containing the text since they're obviously 3d rendered in mmd... clever
	- .x is a standard extension 
	- For logos like for songs...
		- SS logo, upscale, clean up in PS leaving a png with alpha.
		- Import to blender, turn into 3d object and apply texture??
- Use DoF shift in short cuts to give a zoom feel?

## Focal length shift usages
- "Zoom in" feel without zooming in
- Can generally keeps a shot from looking too static (don't overdo though)

## DoF
- Start out of focus and shift into focus with camera movement

## FX
- If going for effect heavy, consider multiple renders and compositing in post (Eg, MOKA's Dramaturgy)
	- Not actually sure if this is what they do or if there's a plugin, but the color borders makes me think it's composite.

## MME
- ray-mmd
	- [Shaders](https://github.com/MikuMikuShaders)
- AutoSmoke/SoftParticleEngine
- WorldSnow (Particle fx)
- 空中に舞う埃 (Particle fx)
- CheapLens (Vignette fx)
- LensFlare_2
- Dot_v1.4 (ドット絵風)
- BodyLine (Neuroi fx...?)
- KiraKira_v1 (+KiraKira_Z)
- Spectrum (Spectrometer fx)
- LightBloom (Yeaaah this is what I was looking for)
- ましまし Diffusions
- lol conway GoL - penne sdCellularAutomata
- Heh... Shinshiwaku fx...
- HanaParticle/LooksReal fx

## Things to read up on
- UE has some nice [rendering docs](https://docs.unrealengine.com/en-US/Engine/Rendering/index.html)
- Color Grading
- AA

## Channels/Pages
### Vids/Motions
- [MAO](https://ch.nicovideo.jp/mao415xxx/blomaga/ar732620?ref=pc_watch_description) (PV kits Umm???)
- [MOKA](https://www.youtube.com/channel/UCuITEWz9SRl-k2Ftjq_1Vaw)
- [PizaCG](https://www.youtube.com/channel/UCpKgRpioJ1pMvIC9vqLzuGw)
- [ちゃーりぃ](https://www.nicovideo.jp/mylist/49944437) 
- [yurie](https://www.nicovideo.jp/user/5728901/video)

### Models
- [銀時愛P](https://seiga.nicovideo.jp/user/illust/22024414)
- [那由他](https://w.atwiki.jp/nayutaworksltd/pages/1.html)
- [極東ロデオ](http://hhp.x0.to/rodeo/3d/)
- [kiribayasi](https://ch.nicovideo.jp/kiribayasi/blomaga/ar1027449)
- [2hu models](https://w.atwiki.jp/vpvpwiki/pages/223.html)
- [deadsilvervirus](https://www.deviantart.com/o-deadsilvervirus-o/gallery/56065516/mmd-download)
- [Stages](https://www.deviantart.com/mmd-stages)
- [山田淀子](https://www.nicovideo.jp/user/41277494)(MME too)

### MME
- [神](https://www.nicovideo.jp/user/421727/video)
- [ましまし](https://www.nicovideo.jp/user/14675117/video)
- [pennennennennennenem](https://pennennennennennenem.github.io/MME/index.html)
- [針金P](https://ch.nicovideo.jp/Harigane/blomaga/ar500418)

### Tools
- [Mogg](https://sites.google.com/site/moggproject/enghome)

### Repos/GH
- [VMD-Lifting](https://github.com/errno-mmd/VMD-Lifting) Yoooooooooo this is cool.

## Random Links
- [Project Diva rips](https://www.deviantart.com/flyingspirits-p/journal/Project-Diva-Motion-Convert-MMD-Downloads-798970572)
- [BeammanP Wiki](https://w.atwiki.jp/beamman/)
- [raymmd](https://learnmmd.com/http:/learnmmd.com/using-raycast-mmd-like-a-pro/)
- [MMDステージ配布あり](https://seiga.nicovideo.jp/tag/MMD%E3%82%B9%E3%83%86%E3%83%BC%E3%82%B8%E9%85%8D%E5%B8%83%E3%81%82%E3%82%8A) (p14)
- [RemiTda :(](https://3d.nicovideo.jp/works/td26230)
- [DirectX .x File Format Specification](http://www.cgdev.net/axe/x-file.html)

## Powershell wryy
- `Get-ChildItem -Path *.zip -File | ForEach {expand-archive -path $_.fullname -destinationpath 'unzipped/'}`

## Tool ideas
- PMX auto bone translator via CLI.
