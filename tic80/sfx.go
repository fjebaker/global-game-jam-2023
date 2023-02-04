package tic80

type SoundEffect struct {
	Id, Note, Octave, Duration, Channel, Volume, Speed int32
	T_start                                            int32
}

func NewSoundEffect(id, channel int32) SoundEffect {
	return SoundEffect{id, 0, 5, 30, channel, 10, 0, 0}
}

func (sfx *SoundEffect) PlayRecordTime(t int32) {
	sfx.T_start = t
	sfx.Play()
}

func (sfx *SoundEffect) IsPlaying(t int32, tModulo int32) bool {
	if t < sfx.T_start {
		return sfx.Duration >= (t + tModulo - sfx.T_start)
	} else {
		return sfx.Duration >= (t - sfx.T_start)
	}
}

func (sfx *SoundEffect) Play() {
	sfx.playId(sfx.Id)
}

func (sfx *SoundEffect) Stop() {
	sfx.playId(-1)
}

func (sfx *SoundEffect) playId(id int32) {
	_sfx(id, sfx.Note, sfx.Octave, sfx.Duration, sfx.Channel, sfx.Volume, sfx.Volume, sfx.Speed)
}
