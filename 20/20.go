package main

import (
	"fmt"
	"../utils/argparse"
	"../utils/intutils"
	"../utils/parsers"
	"../utils/stringutils"
	"strings"
)

const (
	AXE_X = 0
	AXE_Y = 1
	AXE_Z = 2
)

type Troika struct {
	x int
	y int
	z int
}

func (t Troika) add(other Troika) Troika {
	return Troika{t.x + other.x, t.y + other.y, t.z + other.z}
}

func (t Troika) getLength() int {
	return intutils.Abs(t.x) + intutils.Abs(t.y) + intutils.Abs(t.z)
}

type Particle struct {
	position Troika
	velocity Troika
	acceleration Troika
}

func (p *Particle) nextTick() {
	p.velocity = p.velocity.add(p.acceleration)
	p.position = p.position.add(p.velocity)
}

func (p *Particle) equalsPosition(other *Particle) bool {
	return p.position == other.position
}

func main() {
	if err := execute(); err != nil {
		fmt.Println("Day 20: Particle Swarm. Parameters: /path/to/file(string)", err)
	}
}

func execute() error {
	if err := argparse.ValidateLength(2); err != nil {
		return err
	}
	matrix, err := argparse.ReadStringMatrix(1, ", ")
	if err != nil {
		return err
	}

	particles, err := readInput(matrix)
	if err != nil {
		return err
	}

	const TICKS = 1000
	duplicates := simulate(particles, TICKS)

	leftCount := len(particles) - countTrue(duplicates)

	fmt.Println("Left particles:", leftCount)

	return nil
}

func countTrue(duplicates []bool) int {
	count := 0
	for _, val := range duplicates {
		if val {
			count += 1
		}
	}
	return count
}

func simulate(particles []*Particle, ticks int) []bool {
	duplicates := make([]bool, len(particles))
	duplicateCount := 0
	for i := 0; i < ticks; i++ {
		newDuplicateCount := reduce(particles, duplicates)
		for _, particle := range particles {
			particle.nextTick()
		}

		if newDuplicateCount != 0 || i % 10000 == 0 {
			duplicateCount += newDuplicateCount
			fmt.Println("Ticks:", i, "Duplicates count:", duplicateCount)
		}
	}
	return duplicates
}

func reduce(particles []*Particle, duplicates []bool) int {
	duplicateCount := 0
	newDuplicates := make([]int, 0)
	for idx, particle := range particles {
		if duplicates[idx] {
			continue
		}
		if positionCount(particles, duplicates, particle) != 1 {
			newDuplicates = append(newDuplicates, idx)
			duplicateCount += 1
		}
	}

	for _, newDuplicate := range newDuplicates {
		duplicates[newDuplicate] = true
	}
	return duplicateCount
}

func positionCount(particles []*Particle, duplicates []bool, particle *Particle) int {
	count := 0
	for idx, particleToCheck := range particles {
		if duplicates[idx] {
			continue
		}
		if particleToCheck.equalsPosition(particle) {
			count += 1
		}
	}
	return count
}

func readInput(matrix [][]string) ([]*Particle, error) {
	particles := make([]*Particle, len(matrix))

	for particleIdx, row := range matrix {
		particle, err := readParticle(row[0], row[1], row[2])
		if err != nil {
			return nil, err
		}
		particles[particleIdx] = &particle
	}

	return particles, nil
}

func readParticle(rawPosition, rawVelocity, rawAccelerate string) (particle Particle, err error) {
	position, err := readTroika(rawPosition)
	if err != nil {
		return
	}
	velocity, err := readTroika(rawVelocity)
	if err != nil {
		return
	}
	accelerate, err := readTroika(rawAccelerate)

	particle = Particle{position:position, velocity:velocity, acceleration:accelerate}
	return
}

func readTroika(item string) (res Troika, err error) {
	rawCoordinates := stringutils.MapStrings(strings.Split(item[3 : len(item)-1], ","), func (str string) string {
		return strings.Trim(str, " ")
	})
	coordinates, err := parsers.StringsToNumbers(rawCoordinates)
	if err != nil {
		return
	}
	res = Troika{
		x: coordinates[AXE_X],
		y: coordinates[AXE_Y],
		z: coordinates[AXE_Z],
	}
	return
}