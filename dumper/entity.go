package dumper

import "cloud.google.com/go/datastore"

type Entity map[string]interface{}

func load(mm Entity, properties []datastore.Property) error {
	var err error
	for _, p := range properties {
		if st, ok := p.Value.(*datastore.Entity); ok {
			nested := Entity{}
			err = load(nested, st.Properties)
			if err != nil {
				return err
			}
			mm[p.Name] = nested
		} else {
			mm[p.Name] = p.Value
		}
	}
	return nil
}

func (e Entity) Load(properties []datastore.Property) error {
	return load(e, properties)
}

func (e *Entity) Save() ([]datastore.Property, error) {
	panic("not implemented")
}
