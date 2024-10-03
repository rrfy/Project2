package main

import "fmt"

type Attacker  interface {
	Attack()
}

type Moveable interface{
	Move()
}


type GameUnit struct {
	id int
	Name   string
    X int
	Y int
}

func (gu GameUnit) int GetId {
	return gu.id
}

func (gu GameUnit) string GetName {
	return gu.Name
}

func (gu GameUnit) int GetX {
	return gu.X
}

func (gu GameUnit) int GetY {
	return gu.Y
}


type Unit struct {
	base GameUnit
	Hp int
	attc int
}

func (u *Unit) bool IsAlive {
	if (u.Hp > 0){ return True}
	else { return False}
}

func (u *Unit) int gethp{
	return u.Hp
}

func (u *Unit) ReceiveDamage(int damage){
	u.Hp -= damage
}

func (u *Unit) Attack(at Unit) {
	at.ReceiveDamage(u.attc)
}

func (u *Unit) Move(){
	u.X += 1
	u.Y += 1
}


type Structure struct {
	base GameUnit
	Built bool
}

func (s *Structure) int IsBuilt{
	return Built
}


type Fort struct {
	base Structure
	attc int
}

func (f *Fort) Attack(at Unit) {
	at.ReceiveDamage(f.attc)
}


type MobileHouse struct {
	base Structure
	XMove int
	YMove int
}

func (mh *MobileHouse) Move() {
	u.X += XMove
	u.Y += YMove
}

func main() {
    
}