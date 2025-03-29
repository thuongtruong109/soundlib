package base

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

type Identifiable interface {
    GetID() string
}

type Repository[T Identifiable] struct {
    DataPath string
}

func getID[T Identifiable](item T) string {
    switch v := any(item).(type) {
    case *Identifiable:
        return (*v).GetID()
    default:
        return item.GetID()
    }
}

func (br *Repository[T]) GetAll() ([]T, error) {
    data, err := gouse.ReadFileObj[T](br.DataPath)
    if err != nil {
        return nil, err
    }
    // if data == nil {
    //     return nil, nil
    // }

    return data, nil
}

func (br *Repository[T]) GetByID(id string) (*T, error) {
    items, err := br.GetAll()
    if err != nil {
        return nil, err
    }
    for _, item := range items {
        if getID(item) == id {
            return &item, nil
        }
    }
    return nil, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
}

func (br *Repository[T]) Create(item *T) error {
    items, err := br.GetAll()
    if err != nil {
        return err
    }
    items = append(items, *item)
    return gouse.WriteFileObj(br.DataPath, items)
}

func (br *Repository[T]) Update(item *T) error {
    items, err := br.GetAll()
    if err != nil {
        return err
    }
    for i, existingItem := range items {
        if getID(existingItem) == getID(*item) {
            items[i] = *item
            break
        }
    }
    return gouse.WriteFileObj(br.DataPath, items)
}

func (br *Repository[T]) Delete(id string) error {
    items, err := br.GetAll()
    if err != nil {
        return err
    }
    for i, item := range items {
        if getID(item) == id {
            items = append(items[:i], items[i+1:]...)
            break
        }
    }
    return gouse.WriteFileObj(br.DataPath, items)
}