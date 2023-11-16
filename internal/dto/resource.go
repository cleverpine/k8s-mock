package dto

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Resource map[string]any

func (r *Resource) Validate() error {
	if r.Get("metadata") == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Metadata must exist")
	}

	return nil
}

func (r *Resource) GetSubResource(key string) *Resource {
	return resourceGetT[Resource](r, key)
}

func (r *Resource) GetString(key string) string {
	v := resourceGetT[string](r, key)
	if v == nil {
		return ""
	}
	return *v
}

func (r *Resource) Get(key string) any {
	v, ok := (*r)[key]
	if ok {
		return v
	}

	var sr map[string]any = *r
	subKeys := strings.Split(key, ".")
	for i, k := range subKeys {
		if t, ok := sr[k]; ok {
			if i >= len(subKeys)-1 {
				return t
			}

			if p, ok := t.(map[string]any); ok {
				sr = p
				continue
			}
		}
		return nil
	}

	return nil
}

func (r *Resource) Set(key string, value any) {
	sr := r

	subKeys := strings.Split(key, ".")
	for i, k := range subKeys {
		if i >= len(subKeys) {
			sr.Set(k, value)
			break
		}

		tmpSr := sr.GetSubResource(k)
		if tmpSr == nil {

			tmpSr = &Resource{}
			(*sr)[k] = tmpSr
		}

		sr = tmpSr
	}
}

func resourceGetT[T any](r *Resource, key string) *T {
	val := r.Get(key)
	if val != nil {
		typedVal, ok := val.(T)
		if ok {
			return &typedVal
		}
	}

	return nil
}
