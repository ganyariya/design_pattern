package particle

import "fmt"

/*
intrinsic な情報（色や画像）は別に取り除いた MovingParticle
extrinsic 可変的な情報を MovingParticle が持つ
*/
type MovingParticle struct {
	x        int
	y        int
	particle *Particle
}

func NewMovingParticle(x, y int, color, image string) *MovingParticle {
	particle := ParticleFactoryInstance.GetParticle(color, image)
	return &MovingParticle{x: x, y: y, particle: particle}
}

func (mp *MovingParticle) Move(dx, dy int) {
	mp.x += dx
	mp.y += dy
}

func (mp *MovingParticle) Draw() string {
	return fmt.Sprintf("%d %d %s %s", mp.x, mp.y, mp.particle.GetColor(), mp.particle.GetImage())
}
