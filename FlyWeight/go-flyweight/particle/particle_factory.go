package particle

var ParticleFactoryInstance particleFactory = particleFactory{pool: map[[2]string]*Particle{}}

/*
Particle を管理するファクトリ
同じ Particle は複数生成しない

Singleton のような設計にしてある
*/
type particleFactory struct {
	pool map[[2]string]*Particle
}

func (p *particleFactory) GetParticle(color, image string) *Particle {
	if particle, exist := p.pool[[2]string{color, image}]; exist {
		return particle
	}
	particle := newParticle(color, image)
	p.pool[[2]string{color, image}] = particle
	return particle
}
