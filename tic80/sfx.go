package tic80

type SoundEffect struct {
	Id, Note, Octave, Duration, Channel, Volume, Speed	int32
}

func NewSoundEffect(id, channel int32) SoundEffect {
	return SoundEffect{id, 0, 8, 30, channel, 10, 0}
}

func (sfx *SoundEffect) Play() {
	_sfx(sfx.Id, sfx.Note, sfx.Octave, sfx.Duration, sfx.Channel, sfx.Volume, sfx.Volume, sfx.Speed)
}