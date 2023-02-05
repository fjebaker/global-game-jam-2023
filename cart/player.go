package cart

import "cart/tic80"

const (
	PLAYER_DELTA_X = 4
	PLAYER_DELTA_Y = 4

	// Keep the player centered on screen
	PLAYER_OFFSET_X int32 = 120 - PLAYER_DELTA_X
	PLAYER_OFFSET_Y       = 67 - PLAYER_DELTA_Y

	PLAYER_START_POSITION_X = 97 * 8
	PLAYER_START_POSITION_Y = 14 * 8

	DIRT_EAT_TIME_DELTA = 10
	ROOT_EAT_TIME_DELTA = 25
)
const (
	PLAYER_MAIN_FRAME = 256
	PLAYER_MAIN_SIZE  = 1
	PLAYER_DEAD_FRAME = 261

	// DEBUG FRAME
	// PLAYER_MAIN_FRAME = 336

	PLAYER_MOVE_SFX        = 61
	PLAYER_MOVE_DURATION   = 8
	PLAYER_EAT_SFX         = 63
	PLAYER_EAT_DURATION    = 30
	PLAYER_SOUND_CHANNEL   = 3
	EATING_PARTICLE_COLOUR = 4
)

type Player struct {
	X, Y         int32
	Speed        int32
	Frame        int32
	Sprite       tic80.Sprite
	Move_fx      tic80.SoundEffect
	Eat_fx       tic80.SoundEffect
	EatStartTime int32
	Dead         bool
	Digging      bool
	Moving       bool
	HasItem      bool
	Eating       bool
	EatTimeDelta int32
}

func NewPlayer(worldX, worldY int32) Player {
	sprite := tic80.SquareSprite(PLAYER_MAIN_FRAME+2, PLAYER_MAIN_SIZE)
	sprite.Rotate = tic80.ROTATE_RIGHT
	move_fx := tic80.NewSoundEffect(
		PLAYER_MOVE_SFX, PLAYER_SOUND_CHANNEL, PLAYER_MOVE_DURATION,
	)
	eat_fx := tic80.NewSoundEffect(
		PLAYER_EAT_SFX, PLAYER_SOUND_CHANNEL, PLAYER_EAT_DURATION,
	)

	return Player{
		worldX, worldY,
		4,
		0, sprite,
		move_fx, eat_fx,
		0,
		false, false, false, false, false,
		0,
	}
}

///////////////////////////////////////////////////////////////////////////////
// Methods

func (player *Player) Draw() {
	player.Sprite.Draw(PLAYER_OFFSET_X, PLAYER_OFFSET_Y)
	if player.Eating {
		player.drawEatingParticles()
	}
}

func (player *Player) GetInfront() (int32, int32) {
	var x, y int32
	switch player.Sprite.Rotate {
	case tic80.ROTATE_NONE:
		y = player.Y - 8 + PLAYER_DELTA_Y - 1
		x = player.X
	case tic80.ROTATE_DOWN:
		y = player.Y + 8 - PLAYER_DELTA_Y
		x = player.X
	case tic80.ROTATE_RIGHT:
		y = player.Y
		x = player.X + 8 - PLAYER_DELTA_X
	case tic80.ROTATE_LEFT:
		y = player.Y
		x = player.X - 8 + PLAYER_DELTA_Y - 1
	}

	return x + PLAYER_DELTA_X, y + PLAYER_DELTA_Y
}

func (player *Player) HandleInteraction(t int32) {
	// We always check for digging
	player.Digging = tic80.BUTTON_B.IsPressed()

	if tic80.BUTTON_A.IsPressed() {
		player.Speed = 1
		player.Move_fx.Note = 2
	} else {
		player.Speed = 3
		player.Move_fx.Note = 0
	}

	if tic80.BUTTON_UP.IsPressed() {
		player.Sprite.Rotate = tic80.ROTATE_NONE
		player.Moving = true
		return
	}
	if tic80.BUTTON_RIGHT.IsPressed() {
		player.Sprite.Rotate = tic80.ROTATE_RIGHT
		player.Moving = true
		return
	}
	if tic80.BUTTON_LEFT.IsPressed() {
		player.Sprite.Rotate = tic80.ROTATE_LEFT
		player.Moving = true
		return
	}
	if tic80.BUTTON_DOWN.IsPressed() {
		player.Sprite.Rotate = tic80.ROTATE_DOWN
		player.Moving = true
		return
	}
	player.Moving = false
	player.Move_fx.Stop()
}

func (player *Player) Update(t int32, world *World, game *Game, desired *RetrievableItem, rabbit *Rabbit) {
	if rabbit.IsDead() && rabbit.PointInZone(player.X, player.Y) {
		game.ChangeState(GAME_STATE_OVER)
		player.SetDead()
		// no other actions needed
		return
	}

	// check what is infront
	x, y := player.GetInfront()
	tileIndex := world.GetMapTile(x, y)

	// stop eating if not digging
	if !player.Digging {
		player.Eating = false
	}

	finished_eating := player.Eating && TimeSince(t, player.EatStartTime) >= player.EatTimeDelta
	if finished_eating {
		player.Eating = false
		world.Dig(x, y, game)
	}

	// disallow movement when eating
	if !player.Eating && player.Moving {
		// check sfx update
		if player.Move_fx.IsPlaying(t, OVERFLOW_MODULO_TIME) == false {
			player.Move_fx.PlayRecordTime(t)
		}
		// check whether to advance location
		if ((t*15)/10)%player.Speed == 0 {
			player.move(world)
		}
	}

	if world.IsDeadly(tileIndex) {
		game.ChangeState(GAME_STATE_OVER)
		player.SetDead()
		// no other actions needed
		return
	}

	if !player.Eating && player.Digging {
		switch {
		case world.IsDirt(tileIndex):
			player.startEating(t, DIRT_EAT_TIME_DELTA)
		case world.IsTree(tileIndex):
			player.startEating(t, ROOT_EAT_TIME_DELTA)
		case world.IsItem(tileIndex):
			world.CollectItem(x, y)
		}
		// check if the item was what we wanted
		if tileIndex == desired.Id() {
			player.HasItem = true
		}
	}

	player.animate(t)
}

///////////////////////////////////////////////////////////////////////////////
// Utils

func (player *Player) startEating(t int32, eatTimeDelta int32) {
	player.Eat_fx.Play()
	player.Eating = true
	player.EatStartTime = t
	player.EatTimeDelta = eatTimeDelta
}

func (player *Player) animate(t int32) {
	var mod int32
	switch {
	case player.Eating:
		mod = 1
	case player.Moving:
		if player.Speed == 1 {
			mod = 2
		} else {
			mod = 5
		}
	default:
		mod = 12
	}
	if t%mod == 0 {
		player.incrementFrame()
	}
}

func (player *Player) drawEatingParticles() {
	x, y := player.GetInfront()
	base_x := x - player.X + PLAYER_OFFSET_X - PLAYER_DELTA_X
	base_y := y - player.Y + PLAYER_OFFSET_Y - PLAYER_DELTA_Y
	for i := 0; i < 5; i = i + 1 {
		tic80.PaintPixel(
			base_x+int32(RandInt(0, 8)),
			base_y+int32(RandInt(0, 8)),
			EATING_PARTICLE_COLOUR,
		)
	}
}

func (player *Player) incrementFrame() {
	player.Frame = (player.Frame + 1) % 5
	player.Sprite.Id = PLAYER_MAIN_FRAME + player.Frame
}

func (player *Player) move(world *World) {
	x, y := player.GetInfront()
	// if we are trying to move out of bounds
	// don't
	if !world.IsInBounds(x, y) {
		return
	}
	// What does the tile in that position contain?
	tileIndex := world.GetMapTile(x, y)

	if world.IsIndestructible(tileIndex) || world.IsDirt(tileIndex) || world.IsTree(tileIndex) {
		return
	}

	switch player.Sprite.Rotate {
	case tic80.ROTATE_NONE:
		player.Y -= 1
	case tic80.ROTATE_DOWN:
		player.Y += 1
	case tic80.ROTATE_RIGHT:
		player.X += 1
	case tic80.ROTATE_LEFT:
		player.X -= 1
	}
}

func (player *Player) SetDead() {
	player.Dead = true
	// change the tileset explicitly
	player.Sprite.Id = PLAYER_DEAD_FRAME
}
