package openstack

/*
 * OPENSTACK is a simulation to validate this kind data structure work well with special needs,
 * but with unstable and errors. In some part this data structure just a suppose which I cannot
 * validate.
 *
 *
 */

import (
	_ "errors"
	"log"
	_ "os"
	_ "reflect"
)

type Elem struct {
	position  int
	allocated bool
}

type Ostack struct {
	entranced bool
	lenced    []bool
	size      int
	begin     int
	buttom    *Elem
	top       *Elem
	empty     bool
	mapped    bool
	_map      map[int]*Elem
	__map     map[*Elem][]interface{}
}

type Openstack interface {
	init()
	Size() int
	List() []*Elem
	GetButtom() int
	GetTop() int
	AddElem(e *Elem) bool
	RemoveElem(e *Elem) bool
	IsEmpty() bool
	Check(err error)
	IsExist(e *Elem) bool
	SetMap(e *Elem) bool
	GetMap(index int) (*Elem, []interface{})
	Destory(index int)
}

func (o Ostack) init() {
	o.entranced = true
	for i := 0; i < o.size; i++ {
		o.lenced[i] = false
	}
	o.size = 0
	o.begin = 0
	o.empty = true
	o.mapped = false
}

func (o Ostack) Size() int {
	return o.size
}

func (o Ostack) List() []*Elem {
	store := make([]*Elem, o.size)
	if !o.IsEmpty() {
		for key, value := range o._map {
			store[key] = value
		}
	}
	return store
}

func (o Ostack) GetButtom() *Elem {
	return o.buttom
}

func (o Ostack) GetTop() *Elem {
	return o.top
}

func (o Ostack) AddElem(e *Elem) (added bool) {
	if e.position < 0 || e.position > o.size {
		added = false
	}

	if o.empty {
		o._map[0] = e
	}

	if !e.allocated {
		o.begin++
	}
	o._map[o.begin] = e
	added = true
	return
}

func (o Ostack) RemoveElem(e *Elem) (removed bool) {
	if o.empty {
		removed = false
	}

	if o._map[e.position] != nil {
		o.Destory(e.position)
	}
	removed = true
	return
}

// isEmpty checks if the openstack is empty.
func (o Ostack) IsEmpty() bool {
	if o.empty {
		return true
		log.Fatalf("Openstack is empty, please add some elements.")
	}
	return false
}

func (o Ostack) Check(err error) {
	if err != nil {
		panic(err)
	}
}

// isExist checks if the element e exist.
func (o Ostack) IsExist(e *Elem) bool {
	if o.__map[e] != nil && e != nil {
		return true
	} else if o.__map[e] == nil && e != nil {
		return true
	}
	return false
}

// When openstack's lence status is true, we turn down the entrance
// and start to mapping elements, but it just a little complicated,
// I cannot figured out it, so just wait a moment.
func (o Ostack) SetMap(e *Elem) (mapped []bool) {
	for k := 0; k < o.size; k++ {
		if !o.lenced[k] {
			mapped[k] = false
		}
		o.entranced = true
		mapped[k] = true
	}
	// Here means element e mapped to any value but with slice, for
	// now, this map indicate the related resource of element e.
	// But I don't know whether the related resource means some
	// requests or connections or other stuffs which we can processed.
	// So, in here, we use null interface slice, refactor later.
	o.__map[e] = []interface{}{}

	return
}

// getMap gets the value mapped by index and return it's mapping value.
func (o Ostack) GetMap(index int) (e *Elem, v []interface{}) {
	if index < 0 || index > o.size {
		log.Fatalf("Invalid index.")
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
func (o Ostack) Destory(index int) {
	if !o._map[index].allocated {
		log.Fatalf("This position have no element, you needn't to delete it!")
	}
	if !o.IsExist(o._map[index-1]) && o.IsExist(o._map[index]) {
		log.Fatalf("Invalid request, impossible!")
	}
	delete(o._map, index)
	delete(o.__map, o._map[index])
}
