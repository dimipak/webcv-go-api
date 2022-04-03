package models

import (
	db "app/system"
	"errors"
	"fmt"
	"reflect"
	"sort"
)

type skillsMethods interface {
	Update(profile Profile)
}

type Skill struct {
	SkillId     int    `json:"skill_id" gorm:"primarykey"`
	ProfileId   int    `json:"profile_id"`
	Name        string `json:"name"`
	Progress    int    `json:"progress"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Skills []Skill

func (s Skills) GetById(id int) (Skill, error) {
	for i, skill := range s {
		if skill.SkillId == id {
			return s[i], nil
		}
	}

	return Skill{}, errors.New("skill does not exist")
}

func (s *Skill) Update(skill Skill) error {
	return db.GORM().Model(s).Updates(skill).Error
}

func (s *Skill) Delete() error {
	return db.GORM().Delete(s).Error
}

func (s Skills) OrderByOrder() Skills {
	// s.FieldByName("order")
	// s[0].getValueByTagName("name")
	sort.Slice(s, func(i, j int) bool {
		return s[i].Order < s[j].Order
	})
	// fmt.Println(getField("order"))
	return s
}

func (s *Skill) getValueByTagName(tagName string) reflect.Value {
	// t, _ := reflect.TypeOf(&Skill{}).Elem().FieldByName("Name")
	// fmt.Println(t)

	for i := 0; i < reflect.TypeOf(s).Elem().NumField(); i++ {

		field := reflect.TypeOf(s).Elem().Field(i)
		// value := reflect.ValueOf(s).Elem().Field(i)

		v := field.Tag.Get("json")

		if v == tagName {
			fmt.Println(reflect.ValueOf(*s).Field(i))
			return reflect.ValueOf(s).Elem().Field(i)
		}

		// fmt.Println(field)
		// fmt.Println(v)
	}

	return reflect.ValueOf(s).Elem().Field(0)
}
