package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("qricxmco574dent")
		if err != nil {
			return err
		}

		// add
		new_published := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "mch74hpz",
			"name": "published",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_published)
		collection.Schema.AddField(new_published)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("qricxmco574dent")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("mch74hpz")

		return dao.SaveCollection(collection)
	})
}
