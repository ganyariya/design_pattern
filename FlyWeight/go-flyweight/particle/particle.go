package particle

/*
intrinsic 本質的な情報（変更されない）
FlyWeight オブジェクトである

FlyWeight = フライ級（軽量級）を指しているが、実態としては重量級のメモリを持っている

外部から変更できないように実装する
*/
type Particle struct {
	color string
	image string
}

func newParticle(color string, image string) *Particle {
	return &Particle{color: color, image: image}
}

func (p *Particle) GetColor() string {
	return p.color
}
func (p *Particle) GetImage() string {
	return p.image
}
