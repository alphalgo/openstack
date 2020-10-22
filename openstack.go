package openstack

//
// OPENSTACK JUST AN EXPANDED STACK DATA STRUCTURE SIMULATION WITHOUT FULLY TEST,
// IN SOME WAY OPENSTACK IS A SUPPOSE WITH MY IDEAS AND UNVALIDATE IN REAL WORK.
//

import (
	log "github.com/sirupsen/logrus"
)

// defines target value and expect value
const (
	VA = 1 << 10
	VB = 1 << 6
)

// Elem defines the element of stack
type Elem struct {
	position  int
	allocated bool
}

// Ostack defines the ostack object
type Ostack struct {
	entranced bool
	fenced    []bool
	size      int
	cap_      int
	expcap    []int
	begin     int
	bottom    *Elem
	top       *Elem
	empty     bool
	mapped    bool
	_map      map[int]*Elem
	__map     map[*Elem][]interface{}
	expanded  bool
}

// Openstack defines the ostack interface
type Openstack interface {
	init()
	Size() int
	List() []*Elem
	GetBottom() *Elem
	GetTop() *Elem
	AddElem(e *Elem) bool
	RemoveElem(e *Elem) bool
	IsEmpty() bool
	Check(err error)
	IsExist(e *Elem) bool
	SetMap(e *Elem, v ...interface{}) bool
	GetMap(index int) (*Elem, []interface{})
	Destroy(index int)
	IsExpand(expanded bool) bool
	Expand() []*Ostack
	expand() (newcap int)
	abs_(a, b int) int
}

func (o *Ostack) init() {
	o.entranced = true
	for i := 0; i < o.size; i++ {
		o.fenced[i] = false
	}
	o.size = 0
	o.cap_ = 0

	exp := VA
	for i := 0; i < VB; i++ {
		exp *= (i + 1)
		o.expcap[i] = exp
	}
	o.begin = 0
	o.bottom = nil
	o.top = nil
	o.empty = true
	o.mapped = false
	o.expanded = false
}

func (o *Ostack) Size() int {
	if o.empty {
		return 0
	}
	return o.size
}

// List shows the element of ostack
func (o *Ostack) List() []*Elem {
	// the default value of cap_ is 0, but when we create new
	// slice, the value of cap_ will increased 1.5 times than
	// o.size, this action make sure the expand can be proces
	// -sed correctly.
	store := make([]*Elem, o.size, (3/2)*o.size)
	if !o.IsEmpty() {
		for key, value := range o._map {
			store[key] = value
		}
	}
	return store
}

// GetBottom gets the bottom of ostack
func (o *Ostack) GetBottom() *Elem {
	if o.empty {
		return nil
	}
	return o.bottom
}

// GetTop gets the top of ostack
func (o *Ostack) GetTop() *Elem {
	if o.empty {
		return nil
	}
	return o.top
}

// AddElem adds the element to ostack
func (o *Ostack) AddElem(e *Elem) (added bool) {
	if e.position < 0 || e.position > o.size {
		added = false
		log.Warnf("Cannot add element cause of invalid index.")
	}
	if o.empty {
		o._map[0] = e
	}
	if !e.allocated {
		o._map[o.begin] = e
		o.begin++
	}
	if o.begin >= o.cap_ {
		o.expanded = true
	}
	if o.IsExpand(o.expanded) {
		o.Expand()
		added = true
	}
	return
}

// RemoveElem removes the element from ostack
func (o *Ostack) RemoveElem(e *Elem) (removed bool) {
	if o.empty {
		removed = false
	}

	if o._map[e.position] != nil {
		o.Destroy(e.position)
	}
	removed = true
	return
}

// isEmpty checks if the openstack is empty.
func (o *Ostack) IsEmpty() bool {
	if o.empty {
		return true
		log.Warnf("Openstack is empty, please add some elements.")
	}
	return false
}

// Check checks if err
func (o *Ostack) Check(err error) {
	if err != nil {
		panic(err)
	}
}

// isExist checks if the element e exist.
func (o *Ostack) IsExist(e *Elem) bool {
	if o.__map[e] != nil && e != nil {
		return true
	} else if o.__map[e] == nil && e != nil {
		return true
	}
	return false
}

// When openstack's fence status is true, we turn down the entrance
// and start to mapping elements.
func (o *Ostack) SetMap(e *Elem, v ...interface{}) (mapped []bool) {
	for k := 0; k < o.size; k++ {
		if !o.fenced[k] {
			mapped[k] = false
		}
		o.entranced = true
		mapped[k] = true
	}
	// Here means element e mapped to any value but with slice, for
	// now, this map indicate the related resource of element e.
	if len(v) > 0 {
		o.__map[e] = v[0].([]interface{})
	}
	o.__map[e] = v

	return
}

// getMap gets the value mapped by index and return it's mapping value.
func (o *Ostack) GetMap(index int) (e *Elem, v []interface{}) {
	if index < 0 || index > o.size {
		log.Warnf("Invalid index.")
	}
	e = o._map[index]
	v = o.__map[e]

	return
}

// In common stack, we just delete that element then move forward.
// But in openstack, we just delete that element and remove its 'grid'
// and add into tail. But we have follows cases:
//		1) the element's memory cannot occupied before the occupied one
//		2) the map full of elements, we just delete the index mapped to
func (o *Ostack) Destroy(index int) {
	if !o._map[index].allocated {
		log.Warnf("This position have no element, you needn't to delete it!")
	}
	if !o.IsExist(o._map[index-1]) && o.IsExist(o._map[index]) {
		log.Warnf("Invalid request, impossible!")
	}
	delete(o._map, index)
	delete(o.__map, o._map[index])
}

// IsExpand checks if ostack is expand enable
func (o *Ostack) IsExpand(expanded bool) bool {
	if expanded {
		return true
	}
	return false
}

// Expand expands capacity of openstack, first we create double space for
// openstack then move original elements to new space, after then delete
// original space.
func (o *Ostack) Expand() []*Ostack {
	newstack := make([]*Ostack, 2*o.size)
	newstack = append(newstack, o)
	for i := 0; i < o.size; i++ {
		o.Destroy(i)
	}
	return newstack
}

func (o *Ostack) expand() (newcap int) {
	if o.begin > o.size {
		need := o.begin - o.size
		if need <= 2*o.cap_ {
			if o.cap_ < VA {
				o.cap_ *= 2
			} else {
				o.cap_ *= (5 / 4)
			}
		} else {
			i := 0
			o.cap_ = o.expcap[i]
			if o.cap_ < need {
				i++
				o.cap_ = o.expcap[i]
			}
		}
	}
	newcap = o.cap_
	return
}

func (o *Ostack) abs_(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
